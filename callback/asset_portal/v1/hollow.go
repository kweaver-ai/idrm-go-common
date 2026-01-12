package v1

// Hollow 代表空接口，无任何实际操作
type Hollow struct{}

// Notification implements Interface.
func (h *Hollow) Notification() NotificationServiceClient {
	return &HollowNotificationServiceClient{}
}

var _ Interface = &Hollow{}
