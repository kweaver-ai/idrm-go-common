package register

import (
	"context"

	"google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service.proto

type UserServiceGetter interface {
	UserService() UserServiceClient
}

// HollowUserServiceClient 是一个空的实现
type HollowUserServiceClient struct{}

// Create implements UserServiceClient.
func (HollowUserServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Result, error) {
	return &Result{SyncFlag: "success", Msg: "hollow implementation"}, nil
}

// Update implements UserServiceClient.
func (HollowUserServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Result, error) {
	return &Result{SyncFlag: "success", Msg: "hollow implementation"}, nil
}

// StatusUpdate implements UserServiceClient.
func (HollowUserServiceClient) StatusUpdate(ctx context.Context, in *StatusUpdateRequest, opts ...grpc.CallOption) (*Result, error) {
	return &Result{SyncFlag: "success", Msg: "hollow implementation"}, nil
}
