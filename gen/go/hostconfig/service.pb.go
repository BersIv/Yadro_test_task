// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/service.proto

package hostconfig

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ChangeHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ChangeHostnameRequest) Reset() {
	*x = ChangeHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHostnameRequest) ProtoMessage() {}

func (x *ChangeHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHostnameRequest.ProtoReflect.Descriptor instead.
func (*ChangeHostnameRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeHostnameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangeHostnameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName,proto3" json:"newName,omitempty"`
}

func (x *ChangeHostnameResponse) Reset() {
	*x = ChangeHostnameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHostnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHostnameResponse) ProtoMessage() {}

func (x *ChangeHostnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHostnameResponse.ProtoReflect.Descriptor instead.
func (*ChangeHostnameResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeHostnameResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ChangeHostnameResponse) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type ListDNSServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDNSServersRequest) Reset() {
	*x = ListDNSServersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDNSServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDNSServersRequest) ProtoMessage() {}

func (x *ListDNSServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDNSServersRequest.ProtoReflect.Descriptor instead.
func (*ListDNSServersRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

type ListDNSServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServers []string `protobuf:"bytes,1,rep,name=dnsServers,proto3" json:"dnsServers,omitempty"`
}

func (x *ListDNSServersResponse) Reset() {
	*x = ListDNSServersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDNSServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDNSServersResponse) ProtoMessage() {}

func (x *ListDNSServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDNSServersResponse.ProtoReflect.Descriptor instead.
func (*ListDNSServersResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListDNSServersResponse) GetDnsServers() []string {
	if x != nil {
		return x.DnsServers
	}
	return nil
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x16,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x38, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x32, 0xe2, 0x01, 0x0a, 0x0a,
	0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6d, 0x0a, 0x0e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x68,
	0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x65, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x4e, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x64, 0x6e, 0x73,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x20, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_service_proto_goTypes = []any{
	(*ChangeHostnameRequest)(nil),  // 0: hostconfig.ChangeHostnameRequest
	(*ChangeHostnameResponse)(nil), // 1: hostconfig.ChangeHostnameResponse
	(*ListDNSServersRequest)(nil),  // 2: hostconfig.ListDNSServersRequest
	(*ListDNSServersResponse)(nil), // 3: hostconfig.ListDNSServersResponse
}
var file_proto_service_proto_depIdxs = []int32{
	0, // 0: hostconfig.HostConfig.ChangeHostname:input_type -> hostconfig.ChangeHostnameRequest
	2, // 1: hostconfig.HostConfig.ListDNSServers:input_type -> hostconfig.ListDNSServersRequest
	1, // 2: hostconfig.HostConfig.ChangeHostname:output_type -> hostconfig.ChangeHostnameResponse
	3, // 3: hostconfig.HostConfig.ListDNSServers:output_type -> hostconfig.ListDNSServersResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeHostnameRequest); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeHostnameResponse); i {
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
		file_proto_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListDNSServersRequest); i {
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
		file_proto_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListDNSServersResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
