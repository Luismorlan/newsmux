// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: crawler_publisher_message.proto

package protocol

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CrawledSubSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ExternalId string `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	SourceId   string `protobuf:"bytes,4,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	AvatarUrl  string `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	OriginUrl  string `protobuf:"bytes,6,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"`
}

func (x *CrawledSubSource) Reset() {
	*x = CrawledSubSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_publisher_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawledSubSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawledSubSource) ProtoMessage() {}

func (x *CrawledSubSource) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_publisher_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawledSubSource.ProtoReflect.Descriptor instead.
func (*CrawledSubSource) Descriptor() ([]byte, []int) {
	return file_crawler_publisher_message_proto_rawDescGZIP(), []int{0}
}

func (x *CrawledSubSource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CrawledSubSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CrawledSubSource) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *CrawledSubSource) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *CrawledSubSource) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *CrawledSubSource) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

// A CrawlerMessage represent a post generated by crawler
// Used to communicate between crawler and publisher
type CrawlerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The post details that this message sends
	Post *CrawlerMessage_CrawledPost `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	// crawled_at is the time stamp when crawler generate the post
	// note this is different from the actual time content is generated by 3rd party website
	CrawledAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=crawled_at,json=crawledAt,proto3" json:"crawled_at,omitempty"`
	// crawler_ip is the crawler's ip to identify which host/lambda generated the post
	CrawlerIp string `protobuf:"bytes,3,opt,name=crawler_ip,json=crawlerIp,proto3" json:"crawler_ip,omitempty"`
	// crawler_version is the crawler's version, this is place holder in future if we want to do AB test
	CrawlerVersion string `protobuf:"bytes,4,opt,name=crawler_version,json=crawlerVersion,proto3" json:"crawler_version,omitempty"`
	// is_test is to mark if the post is for end-to-end test purpose
	IsTest bool `protobuf:"varint,5,opt,name=is_test,json=isTest,proto3" json:"is_test,omitempty"`
}

func (x *CrawlerMessage) Reset() {
	*x = CrawlerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_publisher_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerMessage) ProtoMessage() {}

func (x *CrawlerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_publisher_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlerMessage.ProtoReflect.Descriptor instead.
func (*CrawlerMessage) Descriptor() ([]byte, []int) {
	return file_crawler_publisher_message_proto_rawDescGZIP(), []int{1}
}

func (x *CrawlerMessage) GetPost() *CrawlerMessage_CrawledPost {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *CrawlerMessage) GetCrawledAt() *timestamp.Timestamp {
	if x != nil {
		return x.CrawledAt
	}
	return nil
}

func (x *CrawlerMessage) GetCrawlerIp() string {
	if x != nil {
		return x.CrawlerIp
	}
	return ""
}

func (x *CrawlerMessage) GetCrawlerVersion() string {
	if x != nil {
		return x.CrawlerVersion
	}
	return ""
}

func (x *CrawlerMessage) GetIsTest() bool {
	if x != nil {
		return x.IsTest
	}
	return false
}

type CrawlerMessage_CrawledPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Original post id that associated with this post in the crawled website.
	// This is for deduplication purpose both in SNS and in Publisher, to make
	// sure we're not inserting the same post twice.
	DeduplicateId string `protobuf:"bytes,1,opt,name=deduplicate_id,json=deduplicateId,proto3" json:"deduplicate_id,omitempty"`
	// sub_source is the sub_source information
	// publisher will use this to upsert table "sub_sources"
	// for new user in retweet chain, subsource will not carry sub_source_id
	// for a crawler, when it starts crawl, it should read DB first
	// and when crawler sends message, each message should have
	// sub_source_id
	// [Update]: we decided to remove source_id, since it can be inferred from subsource
	// Also, all sources have default sub source "default"
	SubSource *CrawledSubSource `protobuf:"bytes,2,opt,name=sub_source,json=subSource,proto3" json:"sub_source,omitempty"`
	// title is the title of the post
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// content is the plain text content of the post
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// image_urls stores urls to images attached to the post
	ImageUrls []string `protobuf:"bytes,5,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	// files_urls stores urls to files attached to the post
	FilesUrls []string `protobuf:"bytes,6,rep,name=files_urls,json=filesUrls,proto3" json:"files_urls,omitempty"`
	// content_generated_at is the actual time content is generated by 3rd party website
	ContentGeneratedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=content_generated_at,json=contentGeneratedAt,proto3" json:"content_generated_at,omitempty"`
	// origin_url is the url of crawled website
	// example: "http://companies.caixin.com/2021-04-10/101688620.html"
	OriginUrl string `protobuf:"bytes,8,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"`
	// post which current post is shared from. For weibo share or retweet. For
	// twitter, this also means quote and retweet.
	SharedFromCrawledPost *CrawlerMessage_CrawledPost `protobuf:"bytes,9,opt,name=shared_from_crawled_post,json=sharedFromCrawledPost,proto3" json:"shared_from_crawled_post,omitempty"`
	// content tags associated with a post
	Tags []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	// Reply to is a relationship expressed in Twitter, where a thread is
	// created by multiple posts replying to each other, connected as a thread.
	ReplyTo *CrawlerMessage_CrawledPost `protobuf:"bytes,11,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
}

func (x *CrawlerMessage_CrawledPost) Reset() {
	*x = CrawlerMessage_CrawledPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_publisher_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerMessage_CrawledPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerMessage_CrawledPost) ProtoMessage() {}

func (x *CrawlerMessage_CrawledPost) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_publisher_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlerMessage_CrawledPost.ProtoReflect.Descriptor instead.
func (*CrawlerMessage_CrawledPost) Descriptor() ([]byte, []int) {
	return file_crawler_publisher_message_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CrawlerMessage_CrawledPost) GetDeduplicateId() string {
	if x != nil {
		return x.DeduplicateId
	}
	return ""
}

func (x *CrawlerMessage_CrawledPost) GetSubSource() *CrawledSubSource {
	if x != nil {
		return x.SubSource
	}
	return nil
}

func (x *CrawlerMessage_CrawledPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CrawlerMessage_CrawledPost) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CrawlerMessage_CrawledPost) GetImageUrls() []string {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *CrawlerMessage_CrawledPost) GetFilesUrls() []string {
	if x != nil {
		return x.FilesUrls
	}
	return nil
}

func (x *CrawlerMessage_CrawledPost) GetContentGeneratedAt() *timestamp.Timestamp {
	if x != nil {
		return x.ContentGeneratedAt
	}
	return nil
}

func (x *CrawlerMessage_CrawledPost) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

func (x *CrawlerMessage_CrawledPost) GetSharedFromCrawledPost() *CrawlerMessage_CrawledPost {
	if x != nil {
		return x.SharedFromCrawledPost
	}
	return nil
}

func (x *CrawlerMessage_CrawledPost) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CrawlerMessage_CrawledPost) GetReplyTo() *CrawlerMessage_CrawledPost {
	if x != nil {
		return x.ReplyTo
	}
	return nil
}

var File_crawler_publisher_message_proto protoreflect.FileDescriptor

var file_crawler_publisher_message_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a,
	0x10, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x53, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72,
	0x6c, 0x22, 0xe7, 0x05, 0x0a, 0x0e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x49, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x1a, 0xfe, 0x03, 0x0a, 0x0b, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65,
	0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x53, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x09, 0x73, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x4c, 0x0a, 0x14, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x5d, 0x0a, 0x18, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x15, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x69, 0x73, 0x6d, 0x6f,
	0x72, 0x6c, 0x61, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6d, 0x75, 0x78, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crawler_publisher_message_proto_rawDescOnce sync.Once
	file_crawler_publisher_message_proto_rawDescData = file_crawler_publisher_message_proto_rawDesc
)

func file_crawler_publisher_message_proto_rawDescGZIP() []byte {
	file_crawler_publisher_message_proto_rawDescOnce.Do(func() {
		file_crawler_publisher_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_crawler_publisher_message_proto_rawDescData)
	})
	return file_crawler_publisher_message_proto_rawDescData
}

var file_crawler_publisher_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crawler_publisher_message_proto_goTypes = []interface{}{
	(*CrawledSubSource)(nil),           // 0: protocol.CrawledSubSource
	(*CrawlerMessage)(nil),             // 1: protocol.CrawlerMessage
	(*CrawlerMessage_CrawledPost)(nil), // 2: protocol.CrawlerMessage.CrawledPost
	(*timestamp.Timestamp)(nil),        // 3: google.protobuf.Timestamp
}
var file_crawler_publisher_message_proto_depIdxs = []int32{
	2, // 0: protocol.CrawlerMessage.post:type_name -> protocol.CrawlerMessage.CrawledPost
	3, // 1: protocol.CrawlerMessage.crawled_at:type_name -> google.protobuf.Timestamp
	0, // 2: protocol.CrawlerMessage.CrawledPost.sub_source:type_name -> protocol.CrawledSubSource
	3, // 3: protocol.CrawlerMessage.CrawledPost.content_generated_at:type_name -> google.protobuf.Timestamp
	2, // 4: protocol.CrawlerMessage.CrawledPost.shared_from_crawled_post:type_name -> protocol.CrawlerMessage.CrawledPost
	2, // 5: protocol.CrawlerMessage.CrawledPost.reply_to:type_name -> protocol.CrawlerMessage.CrawledPost
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_crawler_publisher_message_proto_init() }
func file_crawler_publisher_message_proto_init() {
	if File_crawler_publisher_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crawler_publisher_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawledSubSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawler_publisher_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlerMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawler_publisher_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlerMessage_CrawledPost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crawler_publisher_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crawler_publisher_message_proto_goTypes,
		DependencyIndexes: file_crawler_publisher_message_proto_depIdxs,
		MessageInfos:      file_crawler_publisher_message_proto_msgTypes,
	}.Build()
	File_crawler_publisher_message_proto = out.File
	file_crawler_publisher_message_proto_rawDesc = nil
	file_crawler_publisher_message_proto_goTypes = nil
	file_crawler_publisher_message_proto_depIdxs = nil
}
