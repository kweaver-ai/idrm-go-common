package v1

// Hollow 代表空接口，无任何实际操作
type Hollow struct{}

// DataPush implements Interface.
func (Hollow) DataPush() DataPushCallbackServiceClient {
	return &HollowDataPushCallbackServiceClient{}
}

var _ Interface = &Hollow{}
