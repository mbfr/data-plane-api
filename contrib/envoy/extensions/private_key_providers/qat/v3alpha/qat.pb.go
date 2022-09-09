// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: contrib/envoy/extensions/private_key_providers/qat/v3alpha/qat.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QatPrivateKeyMethodConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Private key to use in the private key provider. If set to inline_bytes or
	// inline_string, the value needs to be the private key in PEM format.
	PrivateKey *v3.DataSource `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// How long to wait before polling the hardware accelerator after a
	// request has been submitted there. Having a small value leads to
	// quicker answers from the hardware but causes more polling loop
	// spins, leading to potentially larger CPU usage. The duration needs
	// to be set to a value greater than or equal to 1 millisecond.
	PollDelay *durationpb.Duration `protobuf:"bytes,2,opt,name=poll_delay,json=pollDelay,proto3" json:"poll_delay,omitempty"`
}

func (x *QatPrivateKeyMethodConfig) Reset() {
	*x = QatPrivateKeyMethodConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QatPrivateKeyMethodConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QatPrivateKeyMethodConfig) ProtoMessage() {}

func (x *QatPrivateKeyMethodConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QatPrivateKeyMethodConfig.ProtoReflect.Descriptor instead.
func (*QatPrivateKeyMethodConfig) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescGZIP(), []int{0}
}

func (x *QatPrivateKeyMethodConfig) GetPrivateKey() *v3.DataSource {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *QatPrivateKeyMethodConfig) GetPollDelay() *durationpb.Duration {
	if x != nil {
		return x.PollDelay
	}
	return nil
}

var File_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDesc = []byte{
	0x0a, 0x44, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x2f, 0x71, 0x61, 0x74, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x71, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x71,
	0x61, 0x74, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x19, 0x51, 0x61, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4,
	0x02, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x48,
	0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0xfa,
	0x42, 0x0b, 0xaa, 0x01, 0x08, 0x08, 0x01, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x09, 0x70,
	0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x42, 0xb1, 0x01, 0x0a, 0x40, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x71, 0x61, 0x74, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x08, 0x51,
	0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x71, 0x61, 0x74, 0x2f, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescData = file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDesc
)

func file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDescData
}

var file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_goTypes = []interface{}{
	(*QatPrivateKeyMethodConfig)(nil), // 0: envoy.extensions.private_key_providers.qat.v3alpha.QatPrivateKeyMethodConfig
	(*v3.DataSource)(nil),             // 1: envoy.config.core.v3.DataSource
	(*durationpb.Duration)(nil),       // 2: google.protobuf.Duration
}
var file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.private_key_providers.qat.v3alpha.QatPrivateKeyMethodConfig.private_key:type_name -> envoy.config.core.v3.DataSource
	2, // 1: envoy.extensions.private_key_providers.qat.v3alpha.QatPrivateKeyMethodConfig.poll_delay:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_init() }
func file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_init() {
	if File_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QatPrivateKeyMethodConfig); i {
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
			RawDescriptor: file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto = out.File
	file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_rawDesc = nil
	file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_goTypes = nil
	file_contrib_envoy_extensions_private_key_providers_qat_v3alpha_qat_proto_depIdxs = nil
}
