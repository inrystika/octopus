// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: server/platform-server/internal/conf/conf.proto

package conf

import (
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

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App     *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app"`
	Server  *Server  `protobuf:"bytes,2,opt,name=server,proto3" json:"server"`
	Data    *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
	Service *Service `protobuf:"bytes,4,opt,name=service,proto3" json:"service"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Version  string `protobuf:"bytes,2,opt,name=version,proto3" json:"version"`
	IsDev    bool   `protobuf:"varint,3,opt,name=isDev,proto3" json:"isDev"`
	LogLevel string `protobuf:"bytes,4,opt,name=logLevel,proto3" json:"logLevel"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetIsDev() bool {
	if x != nil {
		return x.IsDev
	}
	return false
}

func (x *App) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseServerAddr           string      `protobuf:"bytes,1,opt,name=baseServerAddr,proto3" json:"baseServerAddr"`
	Redis                    *Data_Redis `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis"`
	BaseServerRequestTimeout string      `protobuf:"bytes,3,opt,name=baseServerRequestTimeout,proto3" json:"baseServerRequestTimeout"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetBaseServerAddr() string {
	if x != nil {
		return x.BaseServerAddr
	}
	return ""
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data) GetBaseServerRequestTimeout() string {
	if x != nil {
		return x.BaseServerRequestTimeout
	}
	return ""
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenExpirationSec int64 `protobuf:"varint,1,opt,name=tokenExpirationSec,proto3" json:"tokenExpirationSec"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Service) GetTokenExpirationSec() int64 {
	if x != nil {
		return x.TokenExpirationSec
	}
	return 0
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network    string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network"`
	Addr       string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr"`
	Timeout    *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout"`
	JwtSecrect string               `protobuf:"bytes,4,opt,name=jwtSecrect,proto3" json:"jwtSecrect"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_HTTP) GetJwtSecrect() string {
	if x != nil {
		return x.JwtSecrect
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr     string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_server_platform_server_internal_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_server_platform_server_internal_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_server_platform_server_internal_conf_conf_proto protoreflect.FileDescriptor

var file_server_platform_server_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf7, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x33, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61,
	0x70, 0x70, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x65, 0x0a, 0x03, 0x41, 0x70, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x73, 0x44, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x44, 0x65, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x22, 0xd3, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x04, 0x68,
	0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x89, 0x01, 0x0a, 0x04, 0x48,
	0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x77, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x63, 0x74, 0x22, 0xff, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x26, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x3e, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x1a, 0x53, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x39, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x63, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_server_platform_server_internal_conf_conf_proto_rawDescOnce sync.Once
	file_server_platform_server_internal_conf_conf_proto_rawDescData = file_server_platform_server_internal_conf_conf_proto_rawDesc
)

func file_server_platform_server_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_server_platform_server_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_server_platform_server_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_platform_server_internal_conf_conf_proto_rawDescData)
	})
	return file_server_platform_server_internal_conf_conf_proto_rawDescData
}

var file_server_platform_server_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_server_platform_server_internal_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: platformserver.internal.conf.Bootstrap
	(*App)(nil),                 // 1: platformserver.internal.conf.App
	(*Server)(nil),              // 2: platformserver.internal.conf.Server
	(*Data)(nil),                // 3: platformserver.internal.conf.Data
	(*Service)(nil),             // 4: platformserver.internal.conf.Service
	(*Server_HTTP)(nil),         // 5: platformserver.internal.conf.Server.HTTP
	(*Data_Redis)(nil),          // 6: platformserver.internal.conf.Data.Redis
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
}
var file_server_platform_server_internal_conf_conf_proto_depIdxs = []int32{
	1, // 0: platformserver.internal.conf.Bootstrap.app:type_name -> platformserver.internal.conf.App
	2, // 1: platformserver.internal.conf.Bootstrap.server:type_name -> platformserver.internal.conf.Server
	3, // 2: platformserver.internal.conf.Bootstrap.data:type_name -> platformserver.internal.conf.Data
	4, // 3: platformserver.internal.conf.Bootstrap.service:type_name -> platformserver.internal.conf.Service
	5, // 4: platformserver.internal.conf.Server.http:type_name -> platformserver.internal.conf.Server.HTTP
	6, // 5: platformserver.internal.conf.Data.redis:type_name -> platformserver.internal.conf.Data.Redis
	7, // 6: platformserver.internal.conf.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_server_platform_server_internal_conf_conf_proto_init() }
func file_server_platform_server_internal_conf_conf_proto_init() {
	if File_server_platform_server_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_platform_server_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_server_platform_server_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
		file_server_platform_server_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_server_platform_server_internal_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_server_platform_server_internal_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_server_platform_server_internal_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_server_platform_server_internal_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
			RawDescriptor: file_server_platform_server_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_platform_server_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_server_platform_server_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_server_platform_server_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_server_platform_server_internal_conf_conf_proto = out.File
	file_server_platform_server_internal_conf_conf_proto_rawDesc = nil
	file_server_platform_server_internal_conf_conf_proto_goTypes = nil
	file_server_platform_server_internal_conf_conf_proto_depIdxs = nil
}
