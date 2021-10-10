package panoptic

const (
	// Task emitted by executor and is in pending state.
	TOPIC_PENDING_JOB  = "topic.pending_job"
	TOPIC_EXECUTED_JOB = "topic.executed_job"

	LAMBDA_AWS_ROLE      = "arn:aws:iam::213288384225:role/service-role/test_ddog_logging-role-8qnsddqu"
	DATA_COLLECTOR_IMAGE = "213288384225.dkr.ecr.us-west-1.amazonaws.com/data_collector:latest"
	AWS_REGION           = "us-west-1"

	// Datadog related
	DDOG_TASK_STATE_COUNTER           = "task_state_counter"
	DDOG_TASK_SUCCESS_MESSAGE_COUNTER = "task_crawled_message_counter"
	DDOG_TASK_FAILURE_MESSAGE_COUNTER = "task_failure_message_counter"
)

type LambdaExecutorState int64

const (
	UNINITIALIZED LambdaExecutorState = 0
	RUNNABLE      LambdaExecutorState = 1
	RUNNING       LambdaExecutorState = 2
)

type LambdaFunctionState int64

const (
	// Lambda function is active, can take in new job.
	ACTIVE LambdaFunctionState = 0
	// Lambda function is stale, should not accept new job. But shouldn't be
	// cleaned up because it still has pending tasks.
	STALE LambdaFunctionState = 1
	// Lambda function is both stale and has no pending job, should be removed.
	REMOVABLE LambdaFunctionState = 2
)
