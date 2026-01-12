package v1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative notification.proto

type NotificationGetter interface {
	Notification() NotificationServiceClient
}

// HollowNotificationServiceClient 代表空接口，无任何实际操作
type HollowNotificationServiceClient struct{}

// Create implements NotificationServiceClient.
func (h *HollowNotificationServiceClient) Create(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

var _ NotificationServiceClient = &HollowNotificationServiceClient{}
