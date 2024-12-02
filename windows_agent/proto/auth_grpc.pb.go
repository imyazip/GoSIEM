// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_ValidateJWTForSensor_FullMethodName = "/auth.AuthService/ValidateJWTForSensor"
	AuthService_GenerateJWTForSensor_FullMethodName = "/auth.AuthService/GenerateJWTForSensor"
	AuthService_Login_FullMethodName                = "/auth.AuthService/Login"
	AuthService_ValidateJWTUser_FullMethodName      = "/auth.AuthService/ValidateJWTUser"
	AuthService_CreateUser_FullMethodName           = "/auth.AuthService/CreateUser"
	AuthService_DeleteUser_FullMethodName           = "/auth.AuthService/DeleteUser"
	AuthService_GetUserRole_FullMethodName          = "/auth.AuthService/GetUserRole"
	AuthService_SetUserRole_FullMethodName          = "/auth.AuthService/SetUserRole"
	AuthService_GenerateAPIKey_FullMethodName       = "/auth.AuthService/GenerateAPIKey"
	AuthService_RevokeAPIKey_FullMethodName         = "/auth.AuthService/RevokeAPIKey"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис авторизации
type AuthServiceClient interface {
	// API для сенсоров
	ValidateJWTForSensor(ctx context.Context, in *ValidateJWTForSensorRequest, opts ...grpc.CallOption) (*ValidateJWTForSensorResponse, error)
	GenerateJWTForSensor(ctx context.Context, in *GenerateJWTForSensorRequest, opts ...grpc.CallOption) (*GenerateJWTForSensorResponse, error)
	// API для пользователей
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ValidateJWTUser(ctx context.Context, in *ValidateJWTUserRequest, opts ...grpc.CallOption) (*ValidateJWTUserResponse, error)
	// Управление пользователями
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// Управление ролями
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
	SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error)
	// Управление API-ключами
	GenerateAPIKey(ctx context.Context, in *GenerateAPIKeyRequest, opts ...grpc.CallOption) (*GenerateAPIKeyResponse, error)
	RevokeAPIKey(ctx context.Context, in *RevokeAPIKeyRequest, opts ...grpc.CallOption) (*RevokeAPIKeyResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) ValidateJWTForSensor(ctx context.Context, in *ValidateJWTForSensorRequest, opts ...grpc.CallOption) (*ValidateJWTForSensorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateJWTForSensorResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidateJWTForSensor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateJWTForSensor(ctx context.Context, in *GenerateJWTForSensorRequest, opts ...grpc.CallOption) (*GenerateJWTForSensorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateJWTForSensorResponse)
	err := c.cc.Invoke(ctx, AuthService_GenerateJWTForSensor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateJWTUser(ctx context.Context, in *ValidateJWTUserRequest, opts ...grpc.CallOption) (*ValidateJWTUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateJWTUserResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidateJWTUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, AuthService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserRoleResponse)
	err := c.cc.Invoke(ctx, AuthService_GetUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetUserRoleResponse)
	err := c.cc.Invoke(ctx, AuthService_SetUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateAPIKey(ctx context.Context, in *GenerateAPIKeyRequest, opts ...grpc.CallOption) (*GenerateAPIKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAPIKeyResponse)
	err := c.cc.Invoke(ctx, AuthService_GenerateAPIKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RevokeAPIKey(ctx context.Context, in *RevokeAPIKeyRequest, opts ...grpc.CallOption) (*RevokeAPIKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeAPIKeyResponse)
	err := c.cc.Invoke(ctx, AuthService_RevokeAPIKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
//
// Сервис авторизации
type AuthServiceServer interface {
	// API для сенсоров
	ValidateJWTForSensor(context.Context, *ValidateJWTForSensorRequest) (*ValidateJWTForSensorResponse, error)
	GenerateJWTForSensor(context.Context, *GenerateJWTForSensorRequest) (*GenerateJWTForSensorResponse, error)
	// API для пользователей
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ValidateJWTUser(context.Context, *ValidateJWTUserRequest) (*ValidateJWTUserResponse, error)
	// Управление пользователями
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// Управление ролями
	GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error)
	SetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error)
	// Управление API-ключами
	GenerateAPIKey(context.Context, *GenerateAPIKeyRequest) (*GenerateAPIKeyResponse, error)
	RevokeAPIKey(context.Context, *RevokeAPIKeyRequest) (*RevokeAPIKeyResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) ValidateJWTForSensor(context.Context, *ValidateJWTForSensorRequest) (*ValidateJWTForSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateJWTForSensor not implemented")
}
func (UnimplementedAuthServiceServer) GenerateJWTForSensor(context.Context, *GenerateJWTForSensorRequest) (*GenerateJWTForSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateJWTForSensor not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) ValidateJWTUser(context.Context, *ValidateJWTUserRequest) (*ValidateJWTUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateJWTUser not implemented")
}
func (UnimplementedAuthServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRole not implemented")
}
func (UnimplementedAuthServiceServer) SetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserRole not implemented")
}
func (UnimplementedAuthServiceServer) GenerateAPIKey(context.Context, *GenerateAPIKeyRequest) (*GenerateAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAPIKey not implemented")
}
func (UnimplementedAuthServiceServer) RevokeAPIKey(context.Context, *RevokeAPIKeyRequest) (*RevokeAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAPIKey not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_ValidateJWTForSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateJWTForSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateJWTForSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidateJWTForSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateJWTForSensor(ctx, req.(*ValidateJWTForSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateJWTForSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateJWTForSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateJWTForSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GenerateJWTForSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateJWTForSensor(ctx, req.(*GenerateJWTForSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateJWTUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateJWTUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateJWTUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidateJWTUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateJWTUser(ctx, req.(*ValidateJWTUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SetUserRole(ctx, req.(*SetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GenerateAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateAPIKey(ctx, req.(*GenerateAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RevokeAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RevokeAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RevokeAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RevokeAPIKey(ctx, req.(*RevokeAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateJWTForSensor",
			Handler:    _AuthService_ValidateJWTForSensor_Handler,
		},
		{
			MethodName: "GenerateJWTForSensor",
			Handler:    _AuthService_GenerateJWTForSensor_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "ValidateJWTUser",
			Handler:    _AuthService_ValidateJWTUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _AuthService_GetUserRole_Handler,
		},
		{
			MethodName: "SetUserRole",
			Handler:    _AuthService_SetUserRole_Handler,
		},
		{
			MethodName: "GenerateAPIKey",
			Handler:    _AuthService_GenerateAPIKey_Handler,
		},
		{
			MethodName: "RevokeAPIKey",
			Handler:    _AuthService_RevokeAPIKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

const (
	LogStorageService_TransferRawStringLog_FullMethodName = "/auth.LogStorageService/TransferRawStringLog"
	LogStorageService_TranserSerializedLog_FullMethodName = "/auth.LogStorageService/TranserSerializedLog"
	LogStorageService_GetNewLogs_FullMethodName           = "/auth.LogStorageService/GetNewLogs"
	LogStorageService_AddSecurityEvent_FullMethodName     = "/auth.LogStorageService/AddSecurityEvent"
)

// LogStorageServiceClient is the client API for LogStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogStorageServiceClient interface {
	TransferRawStringLog(ctx context.Context, in *TransferRawStringLogRequest, opts ...grpc.CallOption) (*TransferRawStringLogResponse, error)
	TranserSerializedLog(ctx context.Context, in *TranserSerializedLogRequest, opts ...grpc.CallOption) (*TranserSerializedLogResponse, error)
	GetNewLogs(ctx context.Context, in *GetNewLogsRequest, opts ...grpc.CallOption) (*GetNewLogsResponse, error)
	AddSecurityEvent(ctx context.Context, in *AddSecurityEventRequest, opts ...grpc.CallOption) (*AddSecurityEventResponse, error)
}

type logStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogStorageServiceClient(cc grpc.ClientConnInterface) LogStorageServiceClient {
	return &logStorageServiceClient{cc}
}

func (c *logStorageServiceClient) TransferRawStringLog(ctx context.Context, in *TransferRawStringLogRequest, opts ...grpc.CallOption) (*TransferRawStringLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferRawStringLogResponse)
	err := c.cc.Invoke(ctx, LogStorageService_TransferRawStringLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logStorageServiceClient) TranserSerializedLog(ctx context.Context, in *TranserSerializedLogRequest, opts ...grpc.CallOption) (*TranserSerializedLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranserSerializedLogResponse)
	err := c.cc.Invoke(ctx, LogStorageService_TranserSerializedLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logStorageServiceClient) GetNewLogs(ctx context.Context, in *GetNewLogsRequest, opts ...grpc.CallOption) (*GetNewLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewLogsResponse)
	err := c.cc.Invoke(ctx, LogStorageService_GetNewLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logStorageServiceClient) AddSecurityEvent(ctx context.Context, in *AddSecurityEventRequest, opts ...grpc.CallOption) (*AddSecurityEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSecurityEventResponse)
	err := c.cc.Invoke(ctx, LogStorageService_AddSecurityEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogStorageServiceServer is the server API for LogStorageService service.
// All implementations must embed UnimplementedLogStorageServiceServer
// for forward compatibility.
type LogStorageServiceServer interface {
	TransferRawStringLog(context.Context, *TransferRawStringLogRequest) (*TransferRawStringLogResponse, error)
	TranserSerializedLog(context.Context, *TranserSerializedLogRequest) (*TranserSerializedLogResponse, error)
	GetNewLogs(context.Context, *GetNewLogsRequest) (*GetNewLogsResponse, error)
	AddSecurityEvent(context.Context, *AddSecurityEventRequest) (*AddSecurityEventResponse, error)
	mustEmbedUnimplementedLogStorageServiceServer()
}

// UnimplementedLogStorageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogStorageServiceServer struct{}

func (UnimplementedLogStorageServiceServer) TransferRawStringLog(context.Context, *TransferRawStringLogRequest) (*TransferRawStringLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferRawStringLog not implemented")
}
func (UnimplementedLogStorageServiceServer) TranserSerializedLog(context.Context, *TranserSerializedLogRequest) (*TranserSerializedLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranserSerializedLog not implemented")
}
func (UnimplementedLogStorageServiceServer) GetNewLogs(context.Context, *GetNewLogsRequest) (*GetNewLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewLogs not implemented")
}
func (UnimplementedLogStorageServiceServer) AddSecurityEvent(context.Context, *AddSecurityEventRequest) (*AddSecurityEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSecurityEvent not implemented")
}
func (UnimplementedLogStorageServiceServer) mustEmbedUnimplementedLogStorageServiceServer() {}
func (UnimplementedLogStorageServiceServer) testEmbeddedByValue()                           {}

// UnsafeLogStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogStorageServiceServer will
// result in compilation errors.
type UnsafeLogStorageServiceServer interface {
	mustEmbedUnimplementedLogStorageServiceServer()
}

func RegisterLogStorageServiceServer(s grpc.ServiceRegistrar, srv LogStorageServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogStorageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogStorageService_ServiceDesc, srv)
}

func _LogStorageService_TransferRawStringLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRawStringLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStorageServiceServer).TransferRawStringLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogStorageService_TransferRawStringLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStorageServiceServer).TransferRawStringLog(ctx, req.(*TransferRawStringLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogStorageService_TranserSerializedLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranserSerializedLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStorageServiceServer).TranserSerializedLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogStorageService_TranserSerializedLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStorageServiceServer).TranserSerializedLog(ctx, req.(*TranserSerializedLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogStorageService_GetNewLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStorageServiceServer).GetNewLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogStorageService_GetNewLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStorageServiceServer).GetNewLogs(ctx, req.(*GetNewLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogStorageService_AddSecurityEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSecurityEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogStorageServiceServer).AddSecurityEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogStorageService_AddSecurityEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogStorageServiceServer).AddSecurityEvent(ctx, req.(*AddSecurityEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogStorageService_ServiceDesc is the grpc.ServiceDesc for LogStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.LogStorageService",
	HandlerType: (*LogStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferRawStringLog",
			Handler:    _LogStorageService_TransferRawStringLog_Handler,
		},
		{
			MethodName: "TranserSerializedLog",
			Handler:    _LogStorageService_TranserSerializedLog_Handler,
		},
		{
			MethodName: "GetNewLogs",
			Handler:    _LogStorageService_GetNewLogs_Handler,
		},
		{
			MethodName: "AddSecurityEvent",
			Handler:    _LogStorageService_AddSecurityEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
