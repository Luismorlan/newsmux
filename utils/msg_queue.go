package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	. "github.com/Luismorlan/newsmux/utils/log"
)

type MessageQueueMessage struct {
	Message       *string
	MessageId     *string
	ReceivedTimes int
	SentTimeStamp int
	ReceiptHandle string
}

type MessageQueueReader interface {
	ReceiveMessages(int64) ([]*MessageQueueMessage, error)
	DeleteMessage(*MessageQueueMessage) error
}

type SQSMessageQueueReader struct {
	readTimeout int64
	queueName   string
	url         string
	client      *sqs.SQS
}

func NewSQSMessageQueueReader(queueName string, readingTimeout int64) (*SQSMessageQueueReader, error) {
	// Initialize a message queue

	if queueName == "" {
		return nil, errors.New("please specify queue name")
	}

	if readingTimeout < 0 || readingTimeout > 20 {
		return nil, errors.New("readingTimeout should be >= 0 and <= 20")
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	// TODO: Move to .env
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := sqs.New(sess)

	url, err := client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			return nil, errors.New(fmt.Sprintf("Unable to find queue %q.", queueName))
		}
		return nil, errors.New(fmt.Sprintf("Unable to queue %q, %v.", queueName, err))
	}

	return &SQSMessageQueueReader{
		queueName:   queueName,
		url:         *url.QueueUrl,
		readTimeout: readingTimeout,
		client:      client,
	}, nil
}

func (reader *SQSMessageQueueReader) DeleteMessage(msg *MessageQueueMessage) error {
	deleteHandler, err := msg.GetIDForDelete()
	if err != nil {
		return err
	}

	_, err = reader.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &reader.url,
		ReceiptHandle: &deleteHandler,
	})

	if err != nil {
		return err
	}

	Log.Info("message Deleted")
	return nil
}

func (reader *SQSMessageQueueReader) ReceiveMessages(maxNumberOfMessages int64) (msgs []*MessageQueueMessage, err error) {
	if maxNumberOfMessages < 1 || maxNumberOfMessages > 10 {
		return nil, errors.New("maxNumberOfMessages should be >= 1 and <= 10")
	}

	Log.Info("waiting for new message")
	result, err := reader.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &reader.url,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
			"ApproximateReceiveCount",
		}),
		MaxNumberOfMessages: aws.Int64(maxNumberOfMessages), // Receive at most 1, polling will close as soon as there is any messages received, whether 1 or many
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: &reader.readTimeout,
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to read: %q, error: %v.", reader.queueName, err))
	}

	res := []*MessageQueueMessage{}

	for _, msg := range result.Messages {
		var (
			count, sentTime int
		)
		if val, ok := msg.Attributes["ApproximateReceiveCount"]; ok {
			count, _ = strconv.Atoi(*val)
		}

		if val, ok := msg.Attributes["SentTimestamp"]; ok {
			sentTime, _ = strconv.Atoi(*val)
		}

		res = append(res, &MessageQueueMessage{
			Message:       msg.Body,
			MessageId:     msg.MessageId,
			ReceivedTimes: count,
			SentTimeStamp: sentTime,
			ReceiptHandle: *msg.ReceiptHandle,
		})
	}

	return res, nil
}

func (msg *MessageQueueMessage) Read() (string, error) {
	return *msg.Message, nil
}

func (msg *MessageQueueMessage) GetIDForDelete() (string, error) {
	return *&msg.ReceiptHandle, nil
}