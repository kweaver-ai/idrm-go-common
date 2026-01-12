package mock

import (
	"go.uber.org/mock/gomock"

	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
)

type MockDrivenSpec struct {
	GetAppsByAccountId MockDriven_GetAppsByAccountIdSpec

	GetUser MockDriven_GetUserSpec
}

type MockDriven_GetAppsByAccountIdSpec struct {
	Times int
	// args
	ID string
	// return
	APP MockAPPSpec
	Err error
}

type MockDriven_GetUserSpec struct {
	Times int
	// args
	ID   string
	Opts configuration_center.GetUserOptions
	// return
	User *configuration_center.User
	Err  error
}

type MockAPPSpec struct {
	ID string
}

func SetMockDriven(r *MockDrivenMockRecorder, spec *MockDrivenSpec) {
	SetMockDriven_GetAppsByAccountId(r, &spec.GetAppsByAccountId)
	SetMockDriven_GetUser(r, &spec.GetUser)
}

func SetMockDriven_GetAppsByAccountId(r *MockDrivenMockRecorder, spec *MockDriven_GetAppsByAccountIdSpec) {
	var app configuration_center.Apps
	SetMockAPP(&app, &spec.APP)

	call := r.GetAppsByAccountId(gomock.Not(nil), spec.ID)
	call.Times(spec.Times)
	call.Return(&app, spec.Err)
}

func SetMockDriven_GetUser(r *MockDrivenMockRecorder, spec *MockDriven_GetUserSpec) {
	call := r.GetUser(gomock.Not(nil), spec.ID, spec.Opts)
	call.Times(spec.Times)
	call.Return(spec.User, spec.Err)
}

func SetMockAPP(app *configuration_center.Apps, spec *MockAPPSpec) {
	app.ID = spec.ID
}
