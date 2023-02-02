// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roles.proto

package roles

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	// 系統管理員
	Role_ADMINISTRATOR Role = 1
	// 主管
	Role_LEADER Role = 2
	// 排產人員
	Role_PLANNER Role = 3
	// 派工人員
	Role_SCHEDULER Role = 4
	// 快檢人員
	Role_INSPECTOR Role = 5
	// 品管人員
	Role_QUALITY_CONTROLLER Role = 6
	// 作業人員
	Role_OPERATOR Role = 7
	// 搬運人員
	Role_BEARER Role = 8
	// 生技
	Role_BIOTECHNOLOGIST Role = 9
)

var Role_name = map[int32]string{
	0: "ROLE_UNSPECIFIED",
	1: "ADMINISTRATOR",
	2: "LEADER",
	3: "PLANNER",
	4: "SCHEDULER",
	5: "INSPECTOR",
	6: "QUALITY_CONTROLLER",
	7: "OPERATOR",
	8: "BEARER",
	9: "BIOTECHNOLOGIST",
}

var Role_value = map[string]int32{
	"ROLE_UNSPECIFIED":   0,
	"ADMINISTRATOR":      1,
	"LEADER":             2,
	"PLANNER":            3,
	"SCHEDULER":          4,
	"INSPECTOR":          5,
	"QUALITY_CONTROLLER": 6,
	"OPERATOR":           7,
	"BEARER":             8,
	"BIOTECHNOLOGIST":    9,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{0}
}

func init() {
	proto.RegisterEnum("roles.Role", Role_name, Role_value)
}

func init() { proto.RegisterFile("roles.proto", fileDescriptor_b96358c61fe6d5ae) }

var fileDescriptor_b96358c61fe6d5ae = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xce, 0x41, 0x4e, 0x03, 0x31,
	0x0c, 0x85, 0x61, 0x0a, 0xed, 0xb4, 0x75, 0xa9, 0x30, 0x06, 0x71, 0x08, 0x16, 0x6c, 0x38, 0x41,
	0x9a, 0x18, 0x6a, 0x29, 0xc4, 0x83, 0x93, 0x59, 0xb0, 0xaa, 0x84, 0xd4, 0xdd, 0x48, 0x41, 0x03,
	0xd7, 0xe2, 0x8e, 0x28, 0xc3, 0xf2, 0x97, 0xec, 0x4f, 0x0f, 0x76, 0x53, 0x1d, 0xcf, 0xdf, 0x4f,
	0x5f, 0x53, 0xfd, 0xa9, 0xb4, 0x9a, 0xe3, 0xf1, 0x77, 0x01, 0x4b, 0xab, 0xe3, 0x99, 0xee, 0x01,
	0x4d, 0x23, 0x9f, 0x86, 0x94, 0x7b, 0xf6, 0xf2, 0x22, 0x1c, 0xf0, 0x82, 0x6e, 0x61, 0xef, 0xc2,
	0x9b, 0x24, 0xc9, 0xc5, 0x5c, 0x51, 0xc3, 0x05, 0x01, 0x74, 0x91, 0x5d, 0x60, 0xc3, 0x4b, 0xda,
	0xc1, 0xba, 0x8f, 0x2e, 0x25, 0x36, 0xbc, 0xa2, 0x3d, 0x6c, 0xb3, 0x3f, 0x72, 0x18, 0x22, 0x1b,
	0x2e, 0x5b, 0xca, 0x6c, 0xb5, 0xb7, 0x15, 0x3d, 0x00, 0xbd, 0x0f, 0x2e, 0x4a, 0xf9, 0x38, 0x79,
	0x4d, 0xc5, 0x34, 0xb6, 0xb3, 0x8e, 0xae, 0x61, 0xa3, 0x3d, 0xff, 0xe3, 0xeb, 0x86, 0x1f, 0xd8,
	0x19, 0x1b, 0x6e, 0xe8, 0x0e, 0x6e, 0x0e, 0xa2, 0x85, 0xfd, 0x31, 0x69, 0xd4, 0x57, 0xc9, 0x05,
	0xb7, 0x9f, 0xdd, 0xbc, 0xfe, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xe8, 0xdb, 0x03, 0xcc,
	0x00, 0x00, 0x00,
}