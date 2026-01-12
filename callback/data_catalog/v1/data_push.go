package v1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative data_push.proto

type DataPushGetter interface {
	DataPush() DataPushCallbackServiceClient
}

// HollowDataPushCallbackServiceClient 是一个空的实现
type HollowDataPushCallbackServiceClient struct{}

// CompleteTask implements DataPushCallbackServiceClient.
func (HollowDataPushCallbackServiceClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
