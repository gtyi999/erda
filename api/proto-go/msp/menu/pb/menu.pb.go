// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: menu.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetMenuRequest) Reset() {
	*x = GetMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuRequest) ProtoMessage() {}

func (x *GetMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuRequest.ProtoReflect.Descriptor instead.
func (*GetMenuRequest) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

func (x *GetMenuRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *GetMenuRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*MenuItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMenuResponse) Reset() {
	*x = GetMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuResponse) ProtoMessage() {}

func (x *GetMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuResponse.ProtoReflect.Descriptor instead.
func (*GetMenuResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{1}
}

func (x *GetMenuResponse) GetData() []*MenuItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type MenuItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string            `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	ClusterType string            `protobuf:"bytes,2,opt,name=clusterType,proto3" json:"clusterType,omitempty"`
	Key         string            `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	CnName      string            `protobuf:"bytes,4,opt,name=cnName,proto3" json:"cnName,omitempty"`
	EnName      string            `protobuf:"bytes,5,opt,name=enName,proto3" json:"enName,omitempty"`
	Href        string            `protobuf:"bytes,6,opt,name=href,proto3" json:"href,omitempty"`
	Params      map[string]string `protobuf:"bytes,7,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Children    []*MenuItem       `protobuf:"bytes,8,rep,name=children,proto3" json:"children,omitempty"`
	// 前端用于判断菜单是否显示，默认引导页为true，功能页为false，当tenant存在时进行反转
	Exists bool `protobuf:"varint,9,opt,name=exists,proto3" json:"exists,omitempty"`
	// 内部字段: 强制显示
	MustExists bool `protobuf:"varint,10,opt,name=mustExists,proto3" json:"mustExists,omitempty"`
	// 内部字段: 只在K8S集群显示
	OnlyK8S bool `protobuf:"varint,11,opt,name=onlyK8S,proto3" json:"onlyK8S,omitempty"`
	// 内部字段: 只在非K8S集群显示
	OnlyNotK8S bool `protobuf:"varint,12,opt,name=onlyNotK8S,proto3" json:"onlyNotK8S,omitempty"`
	IsK8S      bool `protobuf:"varint,13,opt,name=isK8S,proto3" json:"isK8S,omitempty"`
	IsEdas     bool `protobuf:"varint,14,opt,name=isEdas,proto3" json:"isEdas,omitempty"`
}

func (x *MenuItem) Reset() {
	*x = MenuItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuItem) ProtoMessage() {}

func (x *MenuItem) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuItem.ProtoReflect.Descriptor instead.
func (*MenuItem) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{2}
}

func (x *MenuItem) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *MenuItem) GetClusterType() string {
	if x != nil {
		return x.ClusterType
	}
	return ""
}

func (x *MenuItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MenuItem) GetCnName() string {
	if x != nil {
		return x.CnName
	}
	return ""
}

func (x *MenuItem) GetEnName() string {
	if x != nil {
		return x.EnName
	}
	return ""
}

func (x *MenuItem) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *MenuItem) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *MenuItem) GetChildren() []*MenuItem {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *MenuItem) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *MenuItem) GetMustExists() bool {
	if x != nil {
		return x.MustExists
	}
	return false
}

func (x *MenuItem) GetOnlyK8S() bool {
	if x != nil {
		return x.OnlyK8S
	}
	return false
}

func (x *MenuItem) GetOnlyNotK8S() bool {
	if x != nil {
		return x.OnlyNotK8S
	}
	return false
}

func (x *MenuItem) GetIsK8S() bool {
	if x != nil {
		return x.IsK8S
	}
	return false
}

func (x *MenuItem) GetIsEdas() bool {
	if x != nil {
		return x.IsEdas
	}
	return false
}

type GetSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantGroup string `protobuf:"bytes,1,opt,name=tenantGroup,proto3" json:"tenantGroup,omitempty"`
	TenantId    string `protobuf:"bytes,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
}

func (x *GetSettingRequest) Reset() {
	*x = GetSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingRequest) ProtoMessage() {}

func (x *GetSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingRequest.ProtoReflect.Descriptor instead.
func (*GetSettingRequest) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{3}
}

func (x *GetSettingRequest) GetTenantGroup() string {
	if x != nil {
		return x.TenantGroup
	}
	return ""
}

func (x *GetSettingRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type GetSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*EngineSetting `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSettingResponse) Reset() {
	*x = GetSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingResponse) ProtoMessage() {}

func (x *GetSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingResponse.ProtoReflect.Descriptor instead.
func (*GetSettingResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{4}
}

func (x *GetSettingResponse) GetData() []*EngineSetting {
	if x != nil {
		return x.Data
	}
	return nil
}

type EngineSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddonName string            `protobuf:"bytes,1,opt,name=addonName,proto3" json:"addonName,omitempty"`
	CnName    string            `protobuf:"bytes,2,opt,name=cnName,proto3" json:"cnName,omitempty"`
	EnName    string            `protobuf:"bytes,3,opt,name=enName,proto3" json:"enName,omitempty"`
	Config    map[string]string `protobuf:"bytes,4,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EngineSetting) Reset() {
	*x = EngineSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EngineSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EngineSetting) ProtoMessage() {}

func (x *EngineSetting) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EngineSetting.ProtoReflect.Descriptor instead.
func (*EngineSetting) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{5}
}

func (x *EngineSetting) GetAddonName() string {
	if x != nil {
		return x.AddonName
	}
	return ""
}

func (x *EngineSetting) GetCnName() string {
	if x != nil {
		return x.CnName
	}
	return ""
}

func (x *EngineSetting) GetEnName() string {
	if x != nil {
		return x.EnName
	}
	return ""
}

func (x *EngineSetting) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x1a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf1, 0x03, 0x0a, 0x08, 0x4d, 0x65,
	0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66,
	0x12, 0x3b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x33, 0x0a,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x75,
	0x73, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x6d, 0x75, 0x73, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6e,
	0x6c, 0x79, 0x4b, 0x38, 0x53, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6e, 0x6c,
	0x79, 0x4b, 0x38, 0x53, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x4b,
	0x38, 0x53, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x4e, 0x6f,
	0x74, 0x4b, 0x38, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x4b, 0x38, 0x53, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4b, 0x38, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x45, 0x64, 0x61, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x45, 0x64,
	0x61, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52,
	0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xda, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xda, 0x02,
	0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1d, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0xfa, 0x81, 0xf9, 0x1b, 0x16, 0x0a, 0x14, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x6e,
	0x75, 0x12, 0xb4, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x20, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x65,
	0x6e, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0xfa, 0x81, 0xf9,
	0x1b, 0x36, 0x0a, 0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x1a, 0x10, 0xc2, 0xc4, 0xcb, 0x1c, 0x0b, 0x22,
	0x03, 0x6d, 0x73, 0x70, 0x32, 0x04, 0x10, 0x01, 0x20, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_menu_proto_goTypes = []interface{}{
	(*GetMenuRequest)(nil),     // 0: erda.msp.menu.GetMenuRequest
	(*GetMenuResponse)(nil),    // 1: erda.msp.menu.GetMenuResponse
	(*MenuItem)(nil),           // 2: erda.msp.menu.MenuItem
	(*GetSettingRequest)(nil),  // 3: erda.msp.menu.GetSettingRequest
	(*GetSettingResponse)(nil), // 4: erda.msp.menu.GetSettingResponse
	(*EngineSetting)(nil),      // 5: erda.msp.menu.EngineSetting
	nil,                        // 6: erda.msp.menu.MenuItem.ParamsEntry
	nil,                        // 7: erda.msp.menu.EngineSetting.ConfigEntry
}
var file_menu_proto_depIdxs = []int32{
	2, // 0: erda.msp.menu.GetMenuResponse.data:type_name -> erda.msp.menu.MenuItem
	6, // 1: erda.msp.menu.MenuItem.params:type_name -> erda.msp.menu.MenuItem.ParamsEntry
	2, // 2: erda.msp.menu.MenuItem.children:type_name -> erda.msp.menu.MenuItem
	5, // 3: erda.msp.menu.GetSettingResponse.data:type_name -> erda.msp.menu.EngineSetting
	7, // 4: erda.msp.menu.EngineSetting.config:type_name -> erda.msp.menu.EngineSetting.ConfigEntry
	0, // 5: erda.msp.menu.MenuService.GetMenu:input_type -> erda.msp.menu.GetMenuRequest
	3, // 6: erda.msp.menu.MenuService.GetSetting:input_type -> erda.msp.menu.GetSettingRequest
	1, // 7: erda.msp.menu.MenuService.GetMenu:output_type -> erda.msp.menu.GetMenuResponse
	4, // 8: erda.msp.menu.MenuService.GetSetting:output_type -> erda.msp.menu.GetSettingResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuRequest); i {
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
		file_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuResponse); i {
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
		file_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuItem); i {
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
		file_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingRequest); i {
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
		file_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingResponse); i {
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
		file_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EngineSetting); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}
