// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/calculate.proto

package calculation

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Next field: 7
type AttributeCompareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role         Role       `protobuf:"varint,1,opt,name=role,proto3,enum=gujian.calculation.Role" json:"role,omitempty"`
	Branch       Branch     `protobuf:"varint,2,opt,name=branch,proto3,enum=gujian.calculation.Branch" json:"branch,omitempty"`
	OriginalAttr *Attribute `protobuf:"bytes,3,opt,name=original_attr,json=originalAttr,proto3" json:"original_attr,omitempty"`
	NewAttr      *Attribute `protobuf:"bytes,4,opt,name=new_attr,json=newAttr,proto3" json:"new_attr,omitempty"`
	ChangeOnAttr *Attribute `protobuf:"bytes,5,opt,name=change_on_attr,json=changeOnAttr,proto3" json:"change_on_attr,omitempty"`
	Dps          int32      `protobuf:"varint,6,opt,name=dps,proto3" json:"dps,omitempty"`
}

func (x *AttributeCompareRequest) Reset() {
	*x = AttributeCompareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeCompareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeCompareRequest) ProtoMessage() {}

func (x *AttributeCompareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeCompareRequest.ProtoReflect.Descriptor instead.
func (*AttributeCompareRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeCompareRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ZHAN_FENG
}

func (x *AttributeCompareRequest) GetBranch() Branch {
	if x != nil {
		return x.Branch
	}
	return Branch_ZF_CHUN_HAO
}

func (x *AttributeCompareRequest) GetOriginalAttr() *Attribute {
	if x != nil {
		return x.OriginalAttr
	}
	return nil
}

func (x *AttributeCompareRequest) GetNewAttr() *Attribute {
	if x != nil {
		return x.NewAttr
	}
	return nil
}

func (x *AttributeCompareRequest) GetChangeOnAttr() *Attribute {
	if x != nil {
		return x.ChangeOnAttr
	}
	return nil
}

func (x *AttributeCompareRequest) GetDps() int32 {
	if x != nil {
		return x.Dps
	}
	return 0
}

// Next field: 3
type AttributeCompareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewDps            int32       `protobuf:"varint,1,opt,name=new_dps,json=newDps,proto3" json:"new_dps,omitempty"`
	AttrContributions *AttrEffect `protobuf:"bytes,2,opt,name=attr_contributions,json=attrContributions,proto3" json:"attr_contributions,omitempty"`
}

func (x *AttributeCompareResponse) Reset() {
	*x = AttributeCompareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeCompareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeCompareResponse) ProtoMessage() {}

func (x *AttributeCompareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeCompareResponse.ProtoReflect.Descriptor instead.
func (*AttributeCompareResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *AttributeCompareResponse) GetNewDps() int32 {
	if x != nil {
		return x.NewDps
	}
	return 0
}

func (x *AttributeCompareResponse) GetAttrContributions() *AttrEffect {
	if x != nil {
		return x.AttrContributions
	}
	return nil
}

var File_proto_calculate_proto protoreflect.FileDescriptor

var file_proto_calculate_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x75, 0x6a, 0x69, 0x61, 0x6e, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x17, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x67, 0x75, 0x6a, 0x69, 0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x32,
	0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x67, 0x75, 0x6a, 0x69, 0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x12, 0x42, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x75, 0x6a, 0x69,
	0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x75, 0x6a, 0x69, 0x61,
	0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x41, 0x74, 0x74, 0x72,
	0x12, 0x43, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x75, 0x6a, 0x69, 0x61,
	0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f,
	0x6e, 0x41, 0x74, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x64, 0x70, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x18, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x44, 0x70, 0x73, 0x12, 0x4d, 0x0a,
	0x12, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x75, 0x6a, 0x69,
	0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x11, 0x61, 0x74, 0x74, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x7a, 0x0a, 0x09,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x12, 0x2b, 0x2e,
	0x67, 0x75, 0x6a, 0x69, 0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x75, 0x6a,
	0x69, 0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_calculate_proto_rawDescOnce sync.Once
	file_proto_calculate_proto_rawDescData = file_proto_calculate_proto_rawDesc
)

func file_proto_calculate_proto_rawDescGZIP() []byte {
	file_proto_calculate_proto_rawDescOnce.Do(func() {
		file_proto_calculate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_calculate_proto_rawDescData)
	})
	return file_proto_calculate_proto_rawDescData
}

var file_proto_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_calculate_proto_goTypes = []interface{}{
	(*AttributeCompareRequest)(nil),  // 0: gujian.calculation.AttributeCompareRequest
	(*AttributeCompareResponse)(nil), // 1: gujian.calculation.AttributeCompareResponse
	(Role)(0),                        // 2: gujian.calculation.Role
	(Branch)(0),                      // 3: gujian.calculation.Branch
	(*Attribute)(nil),                // 4: gujian.calculation.Attribute
	(*AttrEffect)(nil),               // 5: gujian.calculation.AttrEffect
}
var file_proto_calculate_proto_depIdxs = []int32{
	2, // 0: gujian.calculation.AttributeCompareRequest.role:type_name -> gujian.calculation.Role
	3, // 1: gujian.calculation.AttributeCompareRequest.branch:type_name -> gujian.calculation.Branch
	4, // 2: gujian.calculation.AttributeCompareRequest.original_attr:type_name -> gujian.calculation.Attribute
	4, // 3: gujian.calculation.AttributeCompareRequest.new_attr:type_name -> gujian.calculation.Attribute
	4, // 4: gujian.calculation.AttributeCompareRequest.change_on_attr:type_name -> gujian.calculation.Attribute
	5, // 5: gujian.calculation.AttributeCompareResponse.attr_contributions:type_name -> gujian.calculation.AttrEffect
	0, // 6: gujian.calculation.Calculate.AttributeCompare:input_type -> gujian.calculation.AttributeCompareRequest
	1, // 7: gujian.calculation.Calculate.AttributeCompare:output_type -> gujian.calculation.AttributeCompareResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_calculate_proto_init() }
func file_proto_calculate_proto_init() {
	if File_proto_calculate_proto != nil {
		return
	}
	file_proto_role_proto_init()
	file_proto_attribute_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_calculate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeCompareRequest); i {
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
		file_proto_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeCompareResponse); i {
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
			RawDescriptor: file_proto_calculate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculate_proto_goTypes,
		DependencyIndexes: file_proto_calculate_proto_depIdxs,
		MessageInfos:      file_proto_calculate_proto_msgTypes,
	}.Build()
	File_proto_calculate_proto = out.File
	file_proto_calculate_proto_rawDesc = nil
	file_proto_calculate_proto_goTypes = nil
	file_proto_calculate_proto_depIdxs = nil
}
