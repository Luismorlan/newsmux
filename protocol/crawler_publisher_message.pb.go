// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: protocol/crawler_publisher_message.proto

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
		mi := &file_protocol_crawler_publisher_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerMessage) ProtoMessage() {}

func (x *CrawlerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_crawler_publisher_message_proto_msgTypes[0]
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
	return file_protocol_crawler_publisher_message_proto_rawDescGZIP(), []int{0}
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

	// sub_source_id is the sub_source id generated in DB, in table "sub_sources"
	// for a crawler, when it starts crawl, it should read DB first
	// and when crawler sends message, each message should have
	// sub_source_id
	// [Update]: we decided to remove source_id, since it can be inferred from subsource
	// Also, all sources have default sub source "default"
	SubSourceId string `protobuf:"bytes,1,opt,name=sub_source_id,json=subSourceId,proto3" json:"sub_source_id,omitempty"`
	// title is the title of the post
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// content is the plain text content of the post
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// image_urls stores urls to images attached to the post
	ImageUrls []string `protobuf:"bytes,4,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	// files_urls stores urls to files attached to the post
	FilesUrls []string `protobuf:"bytes,5,rep,name=files_urls,json=filesUrls,proto3" json:"files_urls,omitempty"`
	// content_generated_at is the actual time content is generated by 3rd party website
	ContentGeneratedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=content_generated_at,json=contentGeneratedAt,proto3" json:"content_generated_at,omitempty"`
	// origin_url is the url of crawled website
	// example: "http://companies.caixin.com/2021-04-10/101688620.html"
	OriginUrl string `protobuf:"bytes,7,opt,name=origin_url,json=originUrl,proto3" json:"origin_url,omitempty"`
}

func (x *CrawlerMessage_CrawledPost) Reset() {
	*x = CrawlerMessage_CrawledPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_crawler_publisher_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerMessage_CrawledPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerMessage_CrawledPost) ProtoMessage() {}

func (x *CrawlerMessage_CrawledPost) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_crawler_publisher_message_proto_msgTypes[1]
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
	return file_protocol_crawler_publisher_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CrawlerMessage_CrawledPost) GetSubSourceId() string {
	if x != nil {
		return x.SubSourceId
	}
	return ""
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

var File_protocol_crawler_publisher_message_proto protoreflect.FileDescriptor

var file_protocol_crawler_publisher_message_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x03, 0x0a, 0x0e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x49, 0x70, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x8c,
	0x02, 0x0a, 0x0b, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x72, 0x6c, 0x73,
	0x12, 0x4c, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x69, 0x73,
	0x6d, 0x6f, 0x72, 0x6c, 0x61, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x6d, 0x75, 0x78, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_crawler_publisher_message_proto_rawDescOnce sync.Once
	file_protocol_crawler_publisher_message_proto_rawDescData = file_protocol_crawler_publisher_message_proto_rawDesc
)

func file_protocol_crawler_publisher_message_proto_rawDescGZIP() []byte {
	file_protocol_crawler_publisher_message_proto_rawDescOnce.Do(func() {
		file_protocol_crawler_publisher_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_crawler_publisher_message_proto_rawDescData)
	})
	return file_protocol_crawler_publisher_message_proto_rawDescData
}

var file_protocol_crawler_publisher_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_crawler_publisher_message_proto_goTypes = []interface{}{
	(*CrawlerMessage)(nil),             // 0: protocol.CrawlerMessage
	(*CrawlerMessage_CrawledPost)(nil), // 1: protocol.CrawlerMessage.CrawledPost
	(*timestamp.Timestamp)(nil),        // 2: google.protobuf.Timestamp
}
var file_protocol_crawler_publisher_message_proto_depIdxs = []int32{
	1, // 0: protocol.CrawlerMessage.post:type_name -> protocol.CrawlerMessage.CrawledPost
	2, // 1: protocol.CrawlerMessage.crawled_at:type_name -> google.protobuf.Timestamp
	2, // 2: protocol.CrawlerMessage.CrawledPost.content_generated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protocol_crawler_publisher_message_proto_init() }
func file_protocol_crawler_publisher_message_proto_init() {
	if File_protocol_crawler_publisher_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_crawler_publisher_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protocol_crawler_publisher_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_protocol_crawler_publisher_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_crawler_publisher_message_proto_goTypes,
		DependencyIndexes: file_protocol_crawler_publisher_message_proto_depIdxs,
		MessageInfos:      file_protocol_crawler_publisher_message_proto_msgTypes,
	}.Build()
	File_protocol_crawler_publisher_message_proto = out.File
	file_protocol_crawler_publisher_message_proto_rawDesc = nil
	file_protocol_crawler_publisher_message_proto_goTypes = nil
	file_protocol_crawler_publisher_message_proto_depIdxs = nil
}
