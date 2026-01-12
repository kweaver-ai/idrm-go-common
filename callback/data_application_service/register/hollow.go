package register

// Hollow 代表空接口，无任何实际操作
type Hollow struct{}

// UserService implements Interface.
func (Hollow) UserService() UserServiceClient {
	return &HollowUserServiceClient{}
}

var _ Interface = &Hollow{}
