// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/proto/specy/request/TaskResponse.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Result struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	TaskResult           []byte   `protobuf:"bytes,2,opt,name=task_result,json=taskResult,proto3" json:"task_result,omitempty"`
	ErrorInfo            string   `protobuf:"bytes,3,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_58f612313eb57b82, []int{0}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Result) GetTaskResult() []byte {
	if m != nil {
		return m.TaskResult
	}
	return nil
}

func (m *Result) GetErrorInfo() string {
	if m != nil {
		return m.ErrorInfo
	}
	return ""
}

type TaskResponse struct {
	Taskhash             []byte   `protobuf:"bytes,1,opt,name=taskhash,proto3" json:"taskhash,omitempty"`
	Result               *Result  `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	RuleFileHash         []byte   `protobuf:"bytes,3,opt,name=rule_file_hash,json=ruleFileHash,proto3" json:"rule_file_hash,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58f612313eb57b82, []int{1}
}
func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetTaskhash() []byte {
	if m != nil {
		return m.Taskhash
	}
	return nil
}

func (m *TaskResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *TaskResponse) GetRuleFileHash() []byte {
	if m != nil {
		return m.RuleFileHash
	}
	return nil
}

func (m *TaskResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "specy.request.Result")
	proto.RegisterType((*TaskResponse)(nil), "specy.request.TaskResponse")
}

func init() {
	proto.RegisterFile("relayer/proto/specy/request/TaskResponse.proto", fileDescriptor_58f612313eb57b82)
}

var fileDescriptor_58f612313eb57b82 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x89, 0x2b, 0x65, 0x3b, 0x5b, 0x3d, 0x44, 0x56, 0x8a, 0x28, 0x96, 0xc5, 0x43, 0x2f,
	0xa6, 0xa0, 0xff, 0xc0, 0x83, 0xe8, 0x35, 0x78, 0xf2, 0x52, 0xa3, 0x4c, 0xdd, 0xb2, 0xa1, 0xa9,
	0x33, 0xc9, 0xa1, 0x3f, 0xc6, 0xff, 0x2a, 0x4d, 0xab, 0xee, 0x1e, 0xe7, 0xf1, 0xcd, 0x7b, 0x8f,
	0x07, 0x8a, 0xd0, 0x9a, 0x01, 0xa9, 0xea, 0xc9, 0x79, 0x57, 0x71, 0x8f, 0x1f, 0x43, 0x45, 0xf8,
	0x15, 0x90, 0x7d, 0xf5, 0x62, 0x78, 0xa7, 0x91, 0x7b, 0xd7, 0x31, 0xaa, 0x08, 0xc8, 0x93, 0x48,
	0xa8, 0x99, 0xd8, 0xbc, 0x41, 0xa2, 0x91, 0x83, 0xf5, 0xf2, 0x1c, 0x12, 0xf6, 0xc6, 0x07, 0xce,
	0x45, 0x21, 0xca, 0xa5, 0x9e, 0x2f, 0x79, 0x0d, 0x2b, 0x6f, 0x78, 0x57, 0x53, 0xc4, 0xf2, 0xa3,
	0x42, 0x94, 0x99, 0x06, 0x3f, 0x39, 0x8f, 0x8f, 0x57, 0x00, 0x48, 0xe4, 0xa8, 0x6e, 0xbb, 0xc6,
	0xe5, 0x8b, 0x42, 0x94, 0xa9, 0x4e, 0xa3, 0xf2, 0xdc, 0x35, 0x6e, 0xf3, 0x2d, 0x20, 0xdb, 0xef,
	0x21, 0x2f, 0x60, 0x39, 0x7e, 0x6f, 0x0d, 0x6f, 0x63, 0x54, 0xa6, 0xff, 0x6e, 0x79, 0x0b, 0xc9,
	0x5e, 0xce, 0xea, 0x6e, 0xad, 0x0e, 0xea, 0xaa, 0x29, 0x52, 0xcf, 0x90, 0xbc, 0x81, 0x53, 0x0a,
	0x16, 0xeb, 0xa6, 0xb5, 0x58, 0x47, 0xc3, 0x45, 0x34, 0xcc, 0x46, 0xf5, 0xb1, 0xb5, 0xf8, 0x34,
	0x9a, 0x5e, 0x42, 0xca, 0xed, 0x67, 0x67, 0x7c, 0x20, 0xcc, 0x8f, 0x23, 0xf0, 0x2f, 0x3c, 0xac,
	0x5f, 0xcf, 0x7e, 0x27, 0x9c, 0xc6, 0xf3, 0x43, 0x8f, 0xfc, 0x9e, 0xc4, 0xb9, 0xee, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x6d, 0xbc, 0x43, 0xd5, 0x60, 0x01, 0x00, 0x00,
}