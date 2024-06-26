// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: reservista/table/table.proto

package proto_table

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{0}
}

type AddTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfSeats int32  `protobuf:"varint,1,opt,name=NumberOfSeats,proto3" json:"NumberOfSeats,omitempty"`
	IsReserved    bool   `protobuf:"varint,2,opt,name=IsReserved,proto3" json:"IsReserved,omitempty"`
	TableNumber   int32  `protobuf:"varint,3,opt,name=TableNumber,proto3" json:"TableNumber,omitempty"`
	RestaurantID  string `protobuf:"bytes,4,opt,name=restaurantID,proto3" json:"restaurantID,omitempty"`
}

func (x *AddTableRequest) Reset() {
	*x = AddTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTableRequest) ProtoMessage() {}

func (x *AddTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTableRequest.ProtoReflect.Descriptor instead.
func (*AddTableRequest) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{1}
}

func (x *AddTableRequest) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

func (x *AddTableRequest) GetIsReserved() bool {
	if x != nil {
		return x.IsReserved
	}
	return false
}

func (x *AddTableRequest) GetTableNumber() int32 {
	if x != nil {
		return x.TableNumber
	}
	return 0
}

func (x *AddTableRequest) GetRestaurantID() string {
	if x != nil {
		return x.RestaurantID
	}
	return ""
}

type TableObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumberOfSeats int32             `protobuf:"varint,2,opt,name=NumberOfSeats,proto3" json:"NumberOfSeats,omitempty"`
	IsReserved    bool              `protobuf:"varint,3,opt,name=IsReserved,proto3" json:"IsReserved,omitempty"`
	TableNumber   int32             `protobuf:"varint,4,opt,name=TableNumber,proto3" json:"TableNumber,omitempty"`
	Restaurant    *RestaurantObject `protobuf:"bytes,5,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
}

func (x *TableObject) Reset() {
	*x = TableObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableObject) ProtoMessage() {}

func (x *TableObject) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableObject.ProtoReflect.Descriptor instead.
func (*TableObject) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{2}
}

func (x *TableObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TableObject) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

func (x *TableObject) GetIsReserved() bool {
	if x != nil {
		return x.IsReserved
	}
	return false
}

func (x *TableObject) GetTableNumber() int32 {
	if x != nil {
		return x.TableNumber
	}
	return 0
}

func (x *TableObject) GetRestaurant() *RestaurantObject {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

type UpdateTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumberOfSeats int32  `protobuf:"varint,2,opt,name=NumberOfSeats,proto3" json:"NumberOfSeats,omitempty"`
	IsReserved    bool   `protobuf:"varint,3,opt,name=IsReserved,proto3" json:"IsReserved,omitempty"`
	TableNumber   int32  `protobuf:"varint,4,opt,name=TableNumber,proto3" json:"TableNumber,omitempty"`
}

func (x *UpdateTableRequest) Reset() {
	*x = UpdateTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTableRequest) ProtoMessage() {}

func (x *UpdateTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTableRequest.ProtoReflect.Descriptor instead.
func (*UpdateTableRequest) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTableRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTableRequest) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

func (x *UpdateTableRequest) GetIsReserved() bool {
	if x != nil {
		return x.IsReserved
	}
	return false
}

func (x *UpdateTableRequest) GetTableNumber() int32 {
	if x != nil {
		return x.TableNumber
	}
	return 0
}

type RestaurantObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Contact string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *RestaurantObject) Reset() {
	*x = RestaurantObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantObject) ProtoMessage() {}

func (x *RestaurantObject) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantObject.ProtoReflect.Descriptor instead.
func (*RestaurantObject) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{4}
}

func (x *RestaurantObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RestaurantObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestaurantObject) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RestaurantObject) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type IDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDRequest) Reset() {
	*x = IDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDRequest) ProtoMessage() {}

func (x *IDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDRequest.ProtoReflect.Descriptor instead.
func (*IDRequest) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{5}
}

func (x *IDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{6}
}

func (x *StatusResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type TableListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*TableObject `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *TableListResponse) Reset() {
	*x = TableListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservista_table_table_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableListResponse) ProtoMessage() {}

func (x *TableListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservista_table_table_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableListResponse.ProtoReflect.Descriptor instead.
func (*TableListResponse) Descriptor() ([]byte, []int) {
	return file_reservista_table_table_proto_rawDescGZIP(), []int{7}
}

func (x *TableListResponse) GetTables() []*TableObject {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_reservista_table_table_proto protoreflect.FileDescriptor

var file_reservista_table_table_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x73, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x9d,
	0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x65,
	0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x22, 0xbe,
	0x01, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53,
	0x65, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22,
	0x8c, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x49, 0x73, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x6a,
	0x0a, 0x10, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x3f, 0x0a, 0x11, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x32, 0xf1, 0x03, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x61, 0x69, 0x64, 0x6f, 0x73, 0x74,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x73, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservista_table_table_proto_rawDescOnce sync.Once
	file_reservista_table_table_proto_rawDescData = file_reservista_table_table_proto_rawDesc
)

func file_reservista_table_table_proto_rawDescGZIP() []byte {
	file_reservista_table_table_proto_rawDescOnce.Do(func() {
		file_reservista_table_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservista_table_table_proto_rawDescData)
	})
	return file_reservista_table_table_proto_rawDescData
}

var file_reservista_table_table_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_reservista_table_table_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: table.Empty
	(*AddTableRequest)(nil),    // 1: table.AddTableRequest
	(*TableObject)(nil),        // 2: table.TableObject
	(*UpdateTableRequest)(nil), // 3: table.UpdateTableRequest
	(*RestaurantObject)(nil),   // 4: table.RestaurantObject
	(*IDRequest)(nil),          // 5: table.IDRequest
	(*StatusResponse)(nil),     // 6: table.StatusResponse
	(*TableListResponse)(nil),  // 7: table.TableListResponse
}
var file_reservista_table_table_proto_depIdxs = []int32{
	4,  // 0: table.TableObject.restaurant:type_name -> table.RestaurantObject
	2,  // 1: table.TableListResponse.tables:type_name -> table.TableObject
	0,  // 2: table.Table.GetAllTables:input_type -> table.Empty
	5,  // 3: table.Table.GetTablesByRestId:input_type -> table.IDRequest
	5,  // 4: table.Table.GetTable:input_type -> table.IDRequest
	1,  // 5: table.Table.AddTable:input_type -> table.AddTableRequest
	3,  // 6: table.Table.UpdateTableById:input_type -> table.UpdateTableRequest
	5,  // 7: table.Table.DeleteTableById:input_type -> table.IDRequest
	5,  // 8: table.Table.GetAvailableTables:input_type -> table.IDRequest
	5,  // 9: table.Table.GetReservedTables:input_type -> table.IDRequest
	7,  // 10: table.Table.GetAllTables:output_type -> table.TableListResponse
	7,  // 11: table.Table.GetTablesByRestId:output_type -> table.TableListResponse
	2,  // 12: table.Table.GetTable:output_type -> table.TableObject
	6,  // 13: table.Table.AddTable:output_type -> table.StatusResponse
	6,  // 14: table.Table.UpdateTableById:output_type -> table.StatusResponse
	6,  // 15: table.Table.DeleteTableById:output_type -> table.StatusResponse
	7,  // 16: table.Table.GetAvailableTables:output_type -> table.TableListResponse
	7,  // 17: table.Table.GetReservedTables:output_type -> table.TableListResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_reservista_table_table_proto_init() }
func file_reservista_table_table_proto_init() {
	if File_reservista_table_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservista_table_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_reservista_table_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTableRequest); i {
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
		file_reservista_table_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableObject); i {
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
		file_reservista_table_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTableRequest); i {
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
		file_reservista_table_table_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantObject); i {
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
		file_reservista_table_table_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDRequest); i {
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
		file_reservista_table_table_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_reservista_table_table_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableListResponse); i {
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
			RawDescriptor: file_reservista_table_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservista_table_table_proto_goTypes,
		DependencyIndexes: file_reservista_table_table_proto_depIdxs,
		MessageInfos:      file_reservista_table_table_proto_msgTypes,
	}.Build()
	File_reservista_table_table_proto = out.File
	file_reservista_table_table_proto_rawDesc = nil
	file_reservista_table_table_proto_goTypes = nil
	file_reservista_table_table_proto_depIdxs = nil
}
