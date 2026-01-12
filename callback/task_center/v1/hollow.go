package v1

// Hollow 代表空接口，无任何实际操作
type Hollow struct{}

// WorkOrder implements Interface.
func (Hollow) WorkOrder() WorkOrderCallbackServiceClient {
	return &HollowWorkOrderCallbackServiceClient{}
}

var _ Interface = &Hollow{}
