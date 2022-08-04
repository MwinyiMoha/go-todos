// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/framework/rpcserver/proto/todo.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Todo struct {
	Id                   string               `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type GetTodosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodosRequest) Reset()         { *m = GetTodosRequest{} }
func (m *GetTodosRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodosRequest) ProtoMessage()    {}
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{1}
}

func (m *GetTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodosRequest.Unmarshal(m, b)
}
func (m *GetTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodosRequest.Marshal(b, m, deterministic)
}
func (m *GetTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodosRequest.Merge(m, src)
}
func (m *GetTodosRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodosRequest.Size(m)
}
func (m *GetTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodosRequest proto.InternalMessageInfo

type GetTodosResponse struct {
	TodoItems            []*Todo  `protobuf:"bytes,1,rep,name=TodoItems,proto3" json:"TodoItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodosResponse) Reset()         { *m = GetTodosResponse{} }
func (m *GetTodosResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodosResponse) ProtoMessage()    {}
func (*GetTodosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{2}
}

func (m *GetTodosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodosResponse.Unmarshal(m, b)
}
func (m *GetTodosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodosResponse.Marshal(b, m, deterministic)
}
func (m *GetTodosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodosResponse.Merge(m, src)
}
func (m *GetTodosResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodosResponse.Size(m)
}
func (m *GetTodosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodosResponse proto.InternalMessageInfo

func (m *GetTodosResponse) GetTodoItems() []*Todo {
	if m != nil {
		return m.TodoItems
	}
	return nil
}

type GetTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{3}
}

func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(m, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoResponse struct {
	TodoItem             *Todo    `protobuf:"bytes,1,opt,name=TodoItem,proto3" json:"TodoItem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{4}
}

func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(m, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetTodoItem() *Todo {
	if m != nil {
		return m.TodoItem
	}
	return nil
}

type CreateTodoRequest struct {
	Description          string   `protobuf:"bytes,1,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{5}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateTodoResponse struct {
	TodoItem             *Todo    `protobuf:"bytes,1,opt,name=TodoItem,proto3" json:"TodoItem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{6}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetTodoItem() *Todo {
	if m != nil {
		return m.TodoItem
	}
	return nil
}

type UpdateTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRequest) Reset()         { *m = UpdateTodoRequest{} }
func (m *UpdateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRequest) ProtoMessage()    {}
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{7}
}

func (m *UpdateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRequest.Unmarshal(m, b)
}
func (m *UpdateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRequest.Merge(m, src)
}
func (m *UpdateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRequest.Size(m)
}
func (m *UpdateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRequest proto.InternalMessageInfo

func (m *UpdateTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateTodoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateTodoResponse struct {
	TodoItem             *Todo    `protobuf:"bytes,1,opt,name=TodoItem,proto3" json:"TodoItem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoResponse) Reset()         { *m = UpdateTodoResponse{} }
func (m *UpdateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoResponse) ProtoMessage()    {}
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{8}
}

func (m *UpdateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoResponse.Unmarshal(m, b)
}
func (m *UpdateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoResponse.Merge(m, src)
}
func (m *UpdateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoResponse.Size(m)
}
func (m *UpdateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoResponse proto.InternalMessageInfo

func (m *UpdateTodoResponse) GetTodoItem() *Todo {
	if m != nil {
		return m.TodoItem
	}
	return nil
}

type DeleteTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{9}
}

func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(m, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3cbe5c5393df32a, []int{10}
}

func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(m, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Todo)(nil), "pb.Todo")
	proto.RegisterType((*GetTodosRequest)(nil), "pb.GetTodosRequest")
	proto.RegisterType((*GetTodosResponse)(nil), "pb.GetTodosResponse")
	proto.RegisterType((*GetTodoRequest)(nil), "pb.GetTodoRequest")
	proto.RegisterType((*GetTodoResponse)(nil), "pb.GetTodoResponse")
	proto.RegisterType((*CreateTodoRequest)(nil), "pb.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "pb.CreateTodoResponse")
	proto.RegisterType((*UpdateTodoRequest)(nil), "pb.UpdateTodoRequest")
	proto.RegisterType((*UpdateTodoResponse)(nil), "pb.UpdateTodoResponse")
	proto.RegisterType((*DeleteTodoRequest)(nil), "pb.DeleteTodoRequest")
	proto.RegisterType((*DeleteTodoResponse)(nil), "pb.DeleteTodoResponse")
}

func init() {
	proto.RegisterFile("internal/framework/rpcserver/proto/todo.proto", fileDescriptor_e3cbe5c5393df32a)
}

var fileDescriptor_e3cbe5c5393df32a = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0xaa, 0x9b, 0x40,
	0x14, 0x46, 0x53, 0xda, 0xe4, 0x04, 0xd2, 0x3a, 0x49, 0x8b, 0xb8, 0xa9, 0x4c, 0x4b, 0xc9, 0xa6,
	0x0a, 0x96, 0xd0, 0x92, 0xac, 0x4a, 0x52, 0x4a, 0x76, 0xc5, 0xa6, 0x0f, 0xa0, 0xf1, 0x24, 0x48,
	0xa3, 0x33, 0x9d, 0x99, 0xa4, 0xef, 0x75, 0x9f, 0xf0, 0xa2, 0xa3, 0xd1, 0xe8, 0xe5, 0x42, 0x76,
	0xf2, 0x9d, 0xef, 0xe7, 0xf8, 0x1d, 0x06, 0x3e, 0xa7, 0xb9, 0x42, 0x91, 0x47, 0x27, 0xff, 0x20,
	0xa2, 0x0c, 0xff, 0x33, 0xf1, 0xd7, 0x17, 0x7c, 0x2f, 0x51, 0x5c, 0x50, 0xf8, 0x5c, 0x30, 0xc5,
	0x7c, 0xc5, 0x12, 0xe6, 0x95, 0x9f, 0xc4, 0xe4, 0xb1, 0xf3, 0xfe, 0xc8, 0xd8, 0xf1, 0x84, 0x7a,
	0x18, 0x9f, 0x0f, 0xbe, 0x4a, 0x33, 0x94, 0x2a, 0xca, 0xb8, 0x26, 0x51, 0x01, 0x2f, 0x76, 0x2c,
	0x61, 0x64, 0x02, 0xe6, 0x36, 0xb1, 0x0d, 0xd7, 0x98, 0x8f, 0x42, 0x73, 0x9b, 0x10, 0x17, 0xc6,
	0x1b, 0x94, 0x7b, 0x91, 0x72, 0x95, 0xb2, 0xdc, 0x36, 0xcb, 0x41, 0x1b, 0x22, 0xdf, 0x60, 0xb4,
	0x16, 0x18, 0x29, 0x4c, 0xbe, 0x2b, 0x7b, 0xe0, 0x1a, 0xf3, 0x71, 0xe0, 0x78, 0x3a, 0xce, 0xab,
	0xe3, 0xbc, 0x5d, 0x1d, 0x17, 0x36, 0x64, 0x6a, 0xc1, 0xeb, 0x9f, 0xa8, 0x8a, 0x58, 0x19, 0xe2,
	0xbf, 0x33, 0x4a, 0x45, 0x97, 0xf0, 0xa6, 0x81, 0x24, 0x67, 0xb9, 0x44, 0xf2, 0x09, 0x46, 0x05,
	0xb0, 0x55, 0x98, 0x49, 0xdb, 0x70, 0x07, 0xf3, 0x71, 0x30, 0xf4, 0x78, 0xec, 0x15, 0x60, 0xd8,
	0x8c, 0xa8, 0x0b, 0x93, 0x4a, 0x5b, 0xb9, 0x75, 0x7f, 0x86, 0x7e, 0xbd, 0x06, 0x5e, 0xcd, 0x3f,
	0xc2, 0xb0, 0x76, 0x28, 0x89, 0x6d, 0xef, 0xeb, 0x84, 0x2e, 0xc0, 0xd2, 0x6b, 0xb7, 0xdd, 0x3b,
	0xd5, 0x18, 0xbd, 0x6a, 0xe8, 0x12, 0x48, 0x5b, 0x76, 0x57, 0xe4, 0x0f, 0xb0, 0xfe, 0xf0, 0xa4,
	0x13, 0x79, 0xf7, 0x75, 0x8a, 0x15, 0xda, 0x36, 0x77, 0xad, 0xf0, 0x01, 0xac, 0x0d, 0x9e, 0xf0,
	0xd9, 0x15, 0xe8, 0x0c, 0x48, 0x9b, 0xa4, 0x03, 0x82, 0x07, 0x13, 0x26, 0x25, 0xf0, 0x6b, 0xfd,
	0x1b, 0xc5, 0x25, 0xdd, 0x23, 0x59, 0xc0, 0xb0, 0x3e, 0x2d, 0x99, 0x16, 0x69, 0x9d, 0xdb, 0x3b,
	0xb3, 0x5b, 0xb0, 0x5a, 0x35, 0x80, 0x57, 0x15, 0x46, 0x48, 0x8b, 0x50, 0x8b, 0xa6, 0x37, 0x58,
	0xa5, 0x59, 0x01, 0x34, 0xbd, 0x93, 0xb7, 0x05, 0xa5, 0x77, 0x3e, 0xe7, 0x5d, 0x17, 0x6e, 0xc4,
	0x4d, 0x63, 0x5a, 0xdc, 0x3b, 0x84, 0x16, 0x3f, 0x51, 0xec, 0x0a, 0xa0, 0x69, 0x43, 0x8b, 0x7b,
	0x15, 0x6a, 0x71, 0xbf, 0xb4, 0xf8, 0x65, 0xf9, 0x5c, 0xbe, 0x3c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0xde, 0x5c, 0x7a, 0xe0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoRPCServiceClient is the client API for TodoRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoRPCServiceClient interface {
	GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error)
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
}

type todoRPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoRPCServiceClient(cc *grpc.ClientConn) TodoRPCServiceClient {
	return &todoRPCServiceClient{cc}
}

func (c *todoRPCServiceClient) GetTodos(ctx context.Context, in *GetTodosRequest, opts ...grpc.CallOption) (*GetTodosResponse, error) {
	out := new(GetTodosResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRPCService/GetTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRPCServiceClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRPCService/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRPCServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRPCService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRPCServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error) {
	out := new(UpdateTodoResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRPCService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRPCServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRPCService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoRPCServiceServer is the server API for TodoRPCService service.
type TodoRPCServiceServer interface {
	GetTodos(context.Context, *GetTodosRequest) (*GetTodosResponse, error)
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error)
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
}

// UnimplementedTodoRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoRPCServiceServer struct {
}

func (*UnimplementedTodoRPCServiceServer) GetTodos(ctx context.Context, req *GetTodosRequest) (*GetTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}
func (*UnimplementedTodoRPCServiceServer) GetTodo(ctx context.Context, req *GetTodoRequest) (*GetTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (*UnimplementedTodoRPCServiceServer) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoRPCServiceServer) UpdateTodo(ctx context.Context, req *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoRPCServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}

func RegisterTodoRPCServiceServer(s *grpc.Server, srv TodoRPCServiceServer) {
	s.RegisterService(&_TodoRPCService_serviceDesc, srv)
}

func _TodoRPCService_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRPCServiceServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRPCService/GetTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRPCServiceServer).GetTodos(ctx, req.(*GetTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRPCService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRPCServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRPCService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRPCServiceServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRPCService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRPCServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRPCService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRPCServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRPCService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRPCServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRPCService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRPCServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRPCService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRPCServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRPCService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRPCServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoRPCService",
	HandlerType: (*TodoRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodos",
			Handler:    _TodoRPCService_GetTodos_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoRPCService_GetTodo_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodoRPCService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoRPCService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoRPCService_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/framework/rpcserver/proto/todo.proto",
}
