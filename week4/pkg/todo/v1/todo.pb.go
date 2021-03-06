// Defines the message for todo service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: todo.proto

package todo

import (
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

// Status defines status of todo item
type Status int32

const (
	Status_PENDING Status = 0
	Status_DONE    Status = 10
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "PENDING",
		10: "DONE",
	}
	Status_value = map[string]int32{
		"PENDING": 0,
		"DONE":    10,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_todo_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

// TodoItemRequest is used for create/update a TODO item. It contains the field can be set via API
type TodoItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc   string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=vrealzhou.todo.v1.Status" json:"status,omitempty"`
}

func (x *TodoItemRequest) Reset() {
	*x = TodoItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItemRequest) ProtoMessage() {}

func (x *TodoItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItemRequest.ProtoReflect.Descriptor instead.
func (*TodoItemRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoItemRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *TodoItemRequest) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

// TodoItemResponse is used for present single TODO item. It contains all fields can be viewd.
type TodoItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc      string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Status    Status `protobuf:"varint,3,opt,name=status,proto3,enum=vrealzhou.todo.v1.Status" json:"status,omitempty"`
	Created   int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TodoItemResponse) Reset() {
	*x = TodoItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItemResponse) ProtoMessage() {}

func (x *TodoItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItemResponse.ProtoReflect.Descriptor instead.
func (*TodoItemResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoItemResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoItemResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *TodoItemResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

func (x *TodoItemResponse) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *TodoItemResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ListTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListTodoRequest) Reset() {
	*x = ListTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoRequest) ProtoMessage() {}

func (x *ListTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoRequest.ProtoReflect.Descriptor instead.
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *ListTodoRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTodoRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// TodoListResponse is used for present multiple TODO items. It also contains pagination info.
type TodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Offset int32               `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32               `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Items  []*TodoItemResponse `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TodoListResponse) Reset() {
	*x = TodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoListResponse) ProtoMessage() {}

func (x *TodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoListResponse.ProtoReflect.Descriptor instead.
func (*TodoListResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TodoListResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *TodoListResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TodoListResponse) GetItems() []*TodoItemResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x76, 0x72,
	0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x22,
	0x68, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68,
	0x6f, 0x75, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x10, 0x54, 0x6f,
	0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3f, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x91,
	0x01, 0x0a, 0x10, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68,
	0x6f, 0x75, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x2a, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e,
	0x45, 0x10, 0x0a, 0x32, 0xbf, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x22,
	0x2e, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x76, 0x72, 0x65,
	0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x72, 0x65, 0x61, 0x6c, 0x7a, 0x68, 0x6f, 0x75, 0x2f, 0x67, 0x65,
	0x65, 0x6b, 0x62, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x77, 0x65, 0x65, 0x6b, 0x34, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_todo_proto_goTypes = []interface{}{
	(Status)(0),              // 0: vrealzhou.todo.v1.Status
	(*TodoItemRequest)(nil),  // 1: vrealzhou.todo.v1.TodoItemRequest
	(*TodoItemResponse)(nil), // 2: vrealzhou.todo.v1.TodoItemResponse
	(*ListTodoRequest)(nil),  // 3: vrealzhou.todo.v1.ListTodoRequest
	(*TodoListResponse)(nil), // 4: vrealzhou.todo.v1.TodoListResponse
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: vrealzhou.todo.v1.TodoItemRequest.status:type_name -> vrealzhou.todo.v1.Status
	0, // 1: vrealzhou.todo.v1.TodoItemResponse.status:type_name -> vrealzhou.todo.v1.Status
	2, // 2: vrealzhou.todo.v1.TodoListResponse.items:type_name -> vrealzhou.todo.v1.TodoItemResponse
	1, // 3: vrealzhou.todo.v1.TodoService.SetTodo:input_type -> vrealzhou.todo.v1.TodoItemRequest
	3, // 4: vrealzhou.todo.v1.TodoService.ListTodoItems:input_type -> vrealzhou.todo.v1.ListTodoRequest
	2, // 5: vrealzhou.todo.v1.TodoService.SetTodo:output_type -> vrealzhou.todo.v1.TodoItemResponse
	4, // 6: vrealzhou.todo.v1.TodoService.ListTodoItems:output_type -> vrealzhou.todo.v1.TodoListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItemRequest); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItemResponse); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodoRequest); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoListResponse); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		EnumInfos:         file_todo_proto_enumTypes,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
