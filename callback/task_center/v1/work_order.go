package v1

import (
	"context"

	"google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative work_order.proto

type WorkOrderGetter interface {
	WorkOrder() WorkOrderCallbackServiceClient
}

// HollowWorkOrderCallbackServiceClient 代表空接口，无任何实际操作
type HollowWorkOrderCallbackServiceClient struct{}

// OnApproved implements WorkOrderCallbackServiceClient.
func (HollowWorkOrderCallbackServiceClient) OnApproved(ctx context.Context, in *WorkOrderApprovedEvent, opts ...grpc.CallOption) (*CommonResult, error) {
	return &CommonResult{}, nil
}

var _ WorkOrderCallbackServiceClient = &HollowWorkOrderCallbackServiceClient{}
