// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proto/lock/v1alpha1/lock.proto

package v1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Lock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Service   string                 `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Who       string                 `protobuf:"bytes,3,opt,name=who,proto3" json:"who,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Lock) Reset() {
	*x = Lock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lock) ProtoMessage() {}

func (x *Lock) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lock.ProtoReflect.Descriptor instead.
func (*Lock) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{0}
}

func (x *Lock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lock) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Lock) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *Lock) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Who     string `protobuf:"bytes,3,opt,name=who,proto3" json:"who,omitempty"`
}

func (x *CreateLockRequest) Reset() {
	*x = CreateLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLockRequest) ProtoMessage() {}

func (x *CreateLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLockRequest.ProtoReflect.Descriptor instead.
func (*CreateLockRequest) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLockRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *CreateLockRequest) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

type CreateLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lock *Lock `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
}

func (x *CreateLockResponse) Reset() {
	*x = CreateLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLockResponse) ProtoMessage() {}

func (x *CreateLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLockResponse.ProtoReflect.Descriptor instead.
func (*CreateLockResponse) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLockResponse) GetLock() *Lock {
	if x != nil {
		return x.Lock
	}
	return nil
}

type GetLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLockRequest) Reset() {
	*x = GetLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockRequest) ProtoMessage() {}

func (x *GetLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockRequest.ProtoReflect.Descriptor instead.
func (*GetLockRequest) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{3}
}

func (x *GetLockRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lock *Lock `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
}

func (x *GetLockResponse) Reset() {
	*x = GetLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockResponse) ProtoMessage() {}

func (x *GetLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockResponse.ProtoReflect.Descriptor instead.
func (*GetLockResponse) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{4}
}

func (x *GetLockResponse) GetLock() *Lock {
	if x != nil {
		return x.Lock
	}
	return nil
}

type UnLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UnLockRequest) Reset() {
	*x = UnLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnLockRequest) ProtoMessage() {}

func (x *UnLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnLockRequest.ProtoReflect.Descriptor instead.
func (*UnLockRequest) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{5}
}

func (x *UnLockRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Count   int64  `protobuf:"varint,10,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UnLockResponse) Reset() {
	*x = UnLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnLockResponse) ProtoMessage() {}

func (x *UnLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnLockResponse.ProtoReflect.Descriptor instead.
func (*UnLockResponse) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{6}
}

func (x *UnLockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UnLockResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnLockResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListLocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerPage *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Page    *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListLocksRequest) Reset() {
	*x = ListLocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocksRequest) ProtoMessage() {}

func (x *ListLocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocksRequest.ProtoReflect.Descriptor instead.
func (*ListLocksRequest) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{7}
}

func (x *ListLocksRequest) GetPerPage() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PerPage
	}
	return nil
}

func (x *ListLocksRequest) GetPage() *wrapperspb.Int32Value {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListLocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locks      []*Lock `protobuf:"bytes,1,rep,name=locks,proto3" json:"locks,omitempty"`
	TotalCount uint32  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListLocksResponse) Reset() {
	*x = ListLocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocksResponse) ProtoMessage() {}

func (x *ListLocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lock_v1alpha1_lock_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocksResponse.ProtoReflect.Descriptor instead.
func (*ListLocksResponse) Descriptor() ([]byte, []int) {
	return file_proto_lock_v1alpha1_lock_proto_rawDescGZIP(), []int{8}
}

func (x *ListLocksResponse) GetLocks() []*Lock {
	if x != nil {
		return x.Locks
	}
	return nil
}

func (x *ListLocksResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_proto_lock_v1alpha1_lock_proto protoreflect.FileDescriptor

var file_proto_lock_v1alpha1_lock_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x87, 0x01, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x77, 0x68, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x6c, 0x6f, 0x63,
	0x6b, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x6b, 0x52, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x6e, 0x4c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x55, 0x6e, 0x4c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x22, 0x67, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x88, 0x04, 0x0a, 0x0b,
	0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x79,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x78, 0x0a, 0x06, 0x55, 0x6e, 0x4c,
	0x6f, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x4c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x55, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x27, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_lock_v1alpha1_lock_proto_rawDescOnce sync.Once
	file_proto_lock_v1alpha1_lock_proto_rawDescData = file_proto_lock_v1alpha1_lock_proto_rawDesc
)

func file_proto_lock_v1alpha1_lock_proto_rawDescGZIP() []byte {
	file_proto_lock_v1alpha1_lock_proto_rawDescOnce.Do(func() {
		file_proto_lock_v1alpha1_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lock_v1alpha1_lock_proto_rawDescData)
	})
	return file_proto_lock_v1alpha1_lock_proto_rawDescData
}

var file_proto_lock_v1alpha1_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_lock_v1alpha1_lock_proto_goTypes = []interface{}{
	(*Lock)(nil),                   // 0: tracker.lock.v1alpha1.Lock
	(*CreateLockRequest)(nil),      // 1: tracker.lock.v1alpha1.CreateLockRequest
	(*CreateLockResponse)(nil),     // 2: tracker.lock.v1alpha1.CreateLockResponse
	(*GetLockRequest)(nil),         // 3: tracker.lock.v1alpha1.GetLockRequest
	(*GetLockResponse)(nil),        // 4: tracker.lock.v1alpha1.GetLockResponse
	(*UnLockRequest)(nil),          // 5: tracker.lock.v1alpha1.UnLockRequest
	(*UnLockResponse)(nil),         // 6: tracker.lock.v1alpha1.UnLockResponse
	(*ListLocksRequest)(nil),       // 7: tracker.lock.v1alpha1.ListLocksRequest
	(*ListLocksResponse)(nil),      // 8: tracker.lock.v1alpha1.ListLocksResponse
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*wrapperspb.UInt32Value)(nil), // 10: google.protobuf.UInt32Value
	(*wrapperspb.Int32Value)(nil),  // 11: google.protobuf.Int32Value
}
var file_proto_lock_v1alpha1_lock_proto_depIdxs = []int32{
	9,  // 0: tracker.lock.v1alpha1.Lock.created_at:type_name -> google.protobuf.Timestamp
	0,  // 1: tracker.lock.v1alpha1.CreateLockResponse.lock:type_name -> tracker.lock.v1alpha1.Lock
	0,  // 2: tracker.lock.v1alpha1.GetLockResponse.lock:type_name -> tracker.lock.v1alpha1.Lock
	10, // 3: tracker.lock.v1alpha1.ListLocksRequest.per_page:type_name -> google.protobuf.UInt32Value
	11, // 4: tracker.lock.v1alpha1.ListLocksRequest.page:type_name -> google.protobuf.Int32Value
	0,  // 5: tracker.lock.v1alpha1.ListLocksResponse.locks:type_name -> tracker.lock.v1alpha1.Lock
	1,  // 6: tracker.lock.v1alpha1.LockService.CreateLock:input_type -> tracker.lock.v1alpha1.CreateLockRequest
	3,  // 7: tracker.lock.v1alpha1.LockService.GetLock:input_type -> tracker.lock.v1alpha1.GetLockRequest
	5,  // 8: tracker.lock.v1alpha1.LockService.UnLock:input_type -> tracker.lock.v1alpha1.UnLockRequest
	7,  // 9: tracker.lock.v1alpha1.LockService.ListLocks:input_type -> tracker.lock.v1alpha1.ListLocksRequest
	2,  // 10: tracker.lock.v1alpha1.LockService.CreateLock:output_type -> tracker.lock.v1alpha1.CreateLockResponse
	4,  // 11: tracker.lock.v1alpha1.LockService.GetLock:output_type -> tracker.lock.v1alpha1.GetLockResponse
	6,  // 12: tracker.lock.v1alpha1.LockService.UnLock:output_type -> tracker.lock.v1alpha1.UnLockResponse
	8,  // 13: tracker.lock.v1alpha1.LockService.ListLocks:output_type -> tracker.lock.v1alpha1.ListLocksResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_lock_v1alpha1_lock_proto_init() }
func file_proto_lock_v1alpha1_lock_proto_init() {
	if File_proto_lock_v1alpha1_lock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lock_v1alpha1_lock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lock); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLockRequest); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLockResponse); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockRequest); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLockResponse); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnLockRequest); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnLockResponse); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocksRequest); i {
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
		file_proto_lock_v1alpha1_lock_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocksResponse); i {
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
			RawDescriptor: file_proto_lock_v1alpha1_lock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_lock_v1alpha1_lock_proto_goTypes,
		DependencyIndexes: file_proto_lock_v1alpha1_lock_proto_depIdxs,
		MessageInfos:      file_proto_lock_v1alpha1_lock_proto_msgTypes,
	}.Build()
	File_proto_lock_v1alpha1_lock_proto = out.File
	file_proto_lock_v1alpha1_lock_proto_rawDesc = nil
	file_proto_lock_v1alpha1_lock_proto_goTypes = nil
	file_proto_lock_v1alpha1_lock_proto_depIdxs = nil
}
