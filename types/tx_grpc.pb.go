// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	Count(ctx context.Context, in *MsgCount, opts ...grpc.CallOption) (*MsgCountResponse, error)
	RenewLeasesAll(ctx context.Context, in *MsgRenewLeasesAll, opts ...grpc.CallOption) (*MsgRenewLeasesAllResponse, error)
	RenewLease(ctx context.Context, in *MsgRenewLease, opts ...grpc.CallOption) (*MsgRenewLeaseResponse, error)
	GetNShortestLeases(ctx context.Context, in *MsgGetNShortestLeases, opts ...grpc.CallOption) (*MsgGetNShortestLeasesResponse, error)
	Keys(ctx context.Context, in *MsgKeys, opts ...grpc.CallOption) (*MsgKeysResponse, error)
	Rename(ctx context.Context, in *MsgRename, opts ...grpc.CallOption) (*MsgRenameResponse, error)
	MultiUpdate(ctx context.Context, in *MsgMultiUpdate, opts ...grpc.CallOption) (*MsgMultiUpdateResponse, error)
	DeleteAll(ctx context.Context, in *MsgDeleteAll, opts ...grpc.CallOption) (*MsgDeleteAllResponse, error)
	KeyValues(ctx context.Context, in *MsgKeyValues, opts ...grpc.CallOption) (*MsgKeyValuesResponse, error)
	Has(ctx context.Context, in *MsgHas, opts ...grpc.CallOption) (*MsgHasResponse, error)
	GetLease(ctx context.Context, in *MsgGetLease, opts ...grpc.CallOption) (*MsgGetLeaseResponse, error)
	Read(ctx context.Context, in *MsgRead, opts ...grpc.CallOption) (*MsgReadResponse, error)
	Upsert(ctx context.Context, in *MsgUpsert, opts ...grpc.CallOption) (*MsgUpsertResponse, error)
	Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error)
	Update(ctx context.Context, in *MsgUpdate, opts ...grpc.CallOption) (*MsgUpdateResponse, error)
	Delete(ctx context.Context, in *MsgDelete, opts ...grpc.CallOption) (*MsgDeleteResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Count(ctx context.Context, in *MsgCount, opts ...grpc.CallOption) (*MsgCountResponse, error) {
	out := new(MsgCountResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RenewLeasesAll(ctx context.Context, in *MsgRenewLeasesAll, opts ...grpc.CallOption) (*MsgRenewLeasesAllResponse, error) {
	out := new(MsgRenewLeasesAllResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/RenewLeasesAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RenewLease(ctx context.Context, in *MsgRenewLease, opts ...grpc.CallOption) (*MsgRenewLeaseResponse, error) {
	out := new(MsgRenewLeaseResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/RenewLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetNShortestLeases(ctx context.Context, in *MsgGetNShortestLeases, opts ...grpc.CallOption) (*MsgGetNShortestLeasesResponse, error) {
	out := new(MsgGetNShortestLeasesResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/GetNShortestLeases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Keys(ctx context.Context, in *MsgKeys, opts ...grpc.CallOption) (*MsgKeysResponse, error) {
	out := new(MsgKeysResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Keys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Rename(ctx context.Context, in *MsgRename, opts ...grpc.CallOption) (*MsgRenameResponse, error) {
	out := new(MsgRenameResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MultiUpdate(ctx context.Context, in *MsgMultiUpdate, opts ...grpc.CallOption) (*MsgMultiUpdateResponse, error) {
	out := new(MsgMultiUpdateResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/MultiUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAll(ctx context.Context, in *MsgDeleteAll, opts ...grpc.CallOption) (*MsgDeleteAllResponse, error) {
	out := new(MsgDeleteAllResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/DeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) KeyValues(ctx context.Context, in *MsgKeyValues, opts ...grpc.CallOption) (*MsgKeyValuesResponse, error) {
	out := new(MsgKeyValuesResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/KeyValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Has(ctx context.Context, in *MsgHas, opts ...grpc.CallOption) (*MsgHasResponse, error) {
	out := new(MsgHasResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Has", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetLease(ctx context.Context, in *MsgGetLease, opts ...grpc.CallOption) (*MsgGetLeaseResponse, error) {
	out := new(MsgGetLeaseResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/GetLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Read(ctx context.Context, in *MsgRead, opts ...grpc.CallOption) (*MsgReadResponse, error) {
	out := new(MsgReadResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Upsert(ctx context.Context, in *MsgUpsert, opts ...grpc.CallOption) (*MsgUpsertResponse, error) {
	out := new(MsgUpsertResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error) {
	out := new(MsgCreateResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Update(ctx context.Context, in *MsgUpdate, opts ...grpc.CallOption) (*MsgUpdateResponse, error) {
	out := new(MsgUpdateResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Delete(ctx context.Context, in *MsgDelete, opts ...grpc.CallOption) (*MsgDeleteResponse, error) {
	out := new(MsgDeleteResponse)
	err := c.cc.Invoke(ctx, "/bluzelle.curium.crud.Msg/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	Count(context.Context, *MsgCount) (*MsgCountResponse, error)
	RenewLeasesAll(context.Context, *MsgRenewLeasesAll) (*MsgRenewLeasesAllResponse, error)
	RenewLease(context.Context, *MsgRenewLease) (*MsgRenewLeaseResponse, error)
	GetNShortestLeases(context.Context, *MsgGetNShortestLeases) (*MsgGetNShortestLeasesResponse, error)
	Keys(context.Context, *MsgKeys) (*MsgKeysResponse, error)
	Rename(context.Context, *MsgRename) (*MsgRenameResponse, error)
	MultiUpdate(context.Context, *MsgMultiUpdate) (*MsgMultiUpdateResponse, error)
	DeleteAll(context.Context, *MsgDeleteAll) (*MsgDeleteAllResponse, error)
	KeyValues(context.Context, *MsgKeyValues) (*MsgKeyValuesResponse, error)
	Has(context.Context, *MsgHas) (*MsgHasResponse, error)
	GetLease(context.Context, *MsgGetLease) (*MsgGetLeaseResponse, error)
	Read(context.Context, *MsgRead) (*MsgReadResponse, error)
	Upsert(context.Context, *MsgUpsert) (*MsgUpsertResponse, error)
	Create(context.Context, *MsgCreate) (*MsgCreateResponse, error)
	Update(context.Context, *MsgUpdate) (*MsgUpdateResponse, error)
	Delete(context.Context, *MsgDelete) (*MsgDeleteResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Count(context.Context, *MsgCount) (*MsgCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedMsgServer) RenewLeasesAll(context.Context, *MsgRenewLeasesAll) (*MsgRenewLeasesAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewLeasesAll not implemented")
}
func (UnimplementedMsgServer) RenewLease(context.Context, *MsgRenewLease) (*MsgRenewLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewLease not implemented")
}
func (UnimplementedMsgServer) GetNShortestLeases(context.Context, *MsgGetNShortestLeases) (*MsgGetNShortestLeasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNShortestLeases not implemented")
}
func (UnimplementedMsgServer) Keys(context.Context, *MsgKeys) (*MsgKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keys not implemented")
}
func (UnimplementedMsgServer) Rename(context.Context, *MsgRename) (*MsgRenameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedMsgServer) MultiUpdate(context.Context, *MsgMultiUpdate) (*MsgMultiUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiUpdate not implemented")
}
func (UnimplementedMsgServer) DeleteAll(context.Context, *MsgDeleteAll) (*MsgDeleteAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAll not implemented")
}
func (UnimplementedMsgServer) KeyValues(context.Context, *MsgKeyValues) (*MsgKeyValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyValues not implemented")
}
func (UnimplementedMsgServer) Has(context.Context, *MsgHas) (*MsgHasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Has not implemented")
}
func (UnimplementedMsgServer) GetLease(context.Context, *MsgGetLease) (*MsgGetLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLease not implemented")
}
func (UnimplementedMsgServer) Read(context.Context, *MsgRead) (*MsgReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedMsgServer) Upsert(context.Context, *MsgUpsert) (*MsgUpsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedMsgServer) Create(context.Context, *MsgCreate) (*MsgCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMsgServer) Update(context.Context, *MsgUpdate) (*MsgUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMsgServer) Delete(context.Context, *MsgDelete) (*MsgDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Count(ctx, req.(*MsgCount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RenewLeasesAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRenewLeasesAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RenewLeasesAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/RenewLeasesAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RenewLeasesAll(ctx, req.(*MsgRenewLeasesAll))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RenewLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRenewLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RenewLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/RenewLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RenewLease(ctx, req.(*MsgRenewLease))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetNShortestLeases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetNShortestLeases)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetNShortestLeases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/GetNShortestLeases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetNShortestLeases(ctx, req.(*MsgGetNShortestLeases))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Keys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgKeys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Keys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Keys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Keys(ctx, req.(*MsgKeys))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRename)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Rename(ctx, req.(*MsgRename))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MultiUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMultiUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MultiUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/MultiUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MultiUpdate(ctx, req.(*MsgMultiUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/DeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAll(ctx, req.(*MsgDeleteAll))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_KeyValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgKeyValues)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).KeyValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/KeyValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).KeyValues(ctx, req.(*MsgKeyValues))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Has_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgHas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Has(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Has",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Has(ctx, req.(*MsgHas))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/GetLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetLease(ctx, req.(*MsgGetLease))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRead)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Read(ctx, req.(*MsgRead))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpsert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Upsert(ctx, req.(*MsgUpsert))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Create(ctx, req.(*MsgCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Update(ctx, req.(*MsgUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bluzelle.curium.crud.Msg/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delete(ctx, req.(*MsgDelete))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bluzelle.curium.crud.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _Msg_Count_Handler,
		},
		{
			MethodName: "RenewLeasesAll",
			Handler:    _Msg_RenewLeasesAll_Handler,
		},
		{
			MethodName: "RenewLease",
			Handler:    _Msg_RenewLease_Handler,
		},
		{
			MethodName: "GetNShortestLeases",
			Handler:    _Msg_GetNShortestLeases_Handler,
		},
		{
			MethodName: "Keys",
			Handler:    _Msg_Keys_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _Msg_Rename_Handler,
		},
		{
			MethodName: "MultiUpdate",
			Handler:    _Msg_MultiUpdate_Handler,
		},
		{
			MethodName: "DeleteAll",
			Handler:    _Msg_DeleteAll_Handler,
		},
		{
			MethodName: "KeyValues",
			Handler:    _Msg_KeyValues_Handler,
		},
		{
			MethodName: "Has",
			Handler:    _Msg_Has_Handler,
		},
		{
			MethodName: "GetLease",
			Handler:    _Msg_GetLease_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Msg_Read_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _Msg_Upsert_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Msg_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Msg_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Msg_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud/tx.proto",
}
