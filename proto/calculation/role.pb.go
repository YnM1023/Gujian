// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/role.proto

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

// List of dps roles
// Next field: 10
type Role int32

const (
	Role_ZHAN_FENG Role = 0
	Role_TIAN_GANG Role = 1
	Role_SHEN_HONG Role = 2
	Role_FEI_XING  Role = 3
	Role_MIAO_FA   Role = 4
	Role_SI_MING   Role = 5
	Role_ZHOU_YIN  Role = 6
	Role_LONG_JUN  Role = 7
	Role_LING_YING Role = 8
	Role_GUAN_RI   Role = 9
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ZHAN_FENG",
		1: "TIAN_GANG",
		2: "SHEN_HONG",
		3: "FEI_XING",
		4: "MIAO_FA",
		5: "SI_MING",
		6: "ZHOU_YIN",
		7: "LONG_JUN",
		8: "LING_YING",
		9: "GUAN_RI",
	}
	Role_value = map[string]int32{
		"ZHAN_FENG": 0,
		"TIAN_GANG": 1,
		"SHEN_HONG": 2,
		"FEI_XING":  3,
		"MIAO_FA":   4,
		"SI_MING":   5,
		"ZHOU_YIN":  6,
		"LONG_JUN":  7,
		"LING_YING": 8,
		"GUAN_RI":   9,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_role_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_role_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{0}
}

// List of dps branches
// Next field: 2
type Branch int32

const (
	Branch_ZF_CHUN_HAO  Branch = 0
	Branch_ZF_CHUN_YING Branch = 1
)

// Enum value maps for Branch.
var (
	Branch_name = map[int32]string{
		0: "ZF_CHUN_HAO",
		1: "ZF_CHUN_YING",
	}
	Branch_value = map[string]int32{
		"ZF_CHUN_HAO":  0,
		"ZF_CHUN_YING": 1,
	}
)

func (x Branch) Enum() *Branch {
	p := new(Branch)
	*p = x
	return p
}

func (x Branch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Branch) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_role_proto_enumTypes[1].Descriptor()
}

func (Branch) Type() protoreflect.EnumType {
	return &file_proto_role_proto_enumTypes[1]
}

func (x Branch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Branch.Descriptor instead.
func (Branch) EnumDescriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{1}
}

var File_proto_role_proto protoreflect.FileDescriptor

var file_proto_role_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x67, 0x75, 0x6a, 0x69, 0x61, 0x6e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x93, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0d, 0x0a, 0x09, 0x5a, 0x48, 0x41, 0x4e, 0x5f, 0x46, 0x45, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x49, 0x41, 0x4e, 0x5f, 0x47, 0x41, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x48, 0x45, 0x4e, 0x5f, 0x48, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x45, 0x49, 0x5f, 0x58, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49,
	0x41, 0x4f, 0x5f, 0x46, 0x41, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49, 0x5f, 0x4d, 0x49,
	0x4e, 0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x5a, 0x48, 0x4f, 0x55, 0x5f, 0x59, 0x49, 0x4e,
	0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x4a, 0x55, 0x4e, 0x10, 0x07,
	0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12,
	0x0b, 0x0a, 0x07, 0x47, 0x55, 0x41, 0x4e, 0x5f, 0x52, 0x49, 0x10, 0x09, 0x2a, 0x2b, 0x0a, 0x06,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0f, 0x0a, 0x0b, 0x5a, 0x46, 0x5f, 0x43, 0x48, 0x55,
	0x4e, 0x5f, 0x48, 0x41, 0x4f, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x5a, 0x46, 0x5f, 0x43, 0x48,
	0x55, 0x4e, 0x5f, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_role_proto_rawDescOnce sync.Once
	file_proto_role_proto_rawDescData = file_proto_role_proto_rawDesc
)

func file_proto_role_proto_rawDescGZIP() []byte {
	file_proto_role_proto_rawDescOnce.Do(func() {
		file_proto_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_role_proto_rawDescData)
	})
	return file_proto_role_proto_rawDescData
}

var file_proto_role_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_role_proto_goTypes = []interface{}{
	(Role)(0),   // 0: gujian.calculation.Role
	(Branch)(0), // 1: gujian.calculation.Branch
}
var file_proto_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_role_proto_init() }
func file_proto_role_proto_init() {
	if File_proto_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_role_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_role_proto_goTypes,
		DependencyIndexes: file_proto_role_proto_depIdxs,
		EnumInfos:         file_proto_role_proto_enumTypes,
	}.Build()
	File_proto_role_proto = out.File
	file_proto_role_proto_rawDesc = nil
	file_proto_role_proto_goTypes = nil
	file_proto_role_proto_depIdxs = nil
}
