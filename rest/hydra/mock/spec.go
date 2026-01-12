package mock

import (
	"go.uber.org/mock/gomock"

	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
)

type MockHydraSpec struct {
	Introspect        MockHydra_IntrospectSpec
	GetClientNameById MockHydra_GetClientNameByIdSpec
}

func SetMockHydraMockRecorder(r *MockHydraMockRecorder, spec *MockHydraSpec) {
	SetMockHydraIntrospectCall(r, &spec.Introspect)
	SetMocGetClientNameByIdCall(r, &spec.GetClientNameById)
}

type MockHydra_IntrospectSpec struct {
	Times int
	// args
	Token string
	// return
	Info TokenIntrospectInfoSpec
	Err  error
}

type MockHydra_GetClientNameByIdSpec struct {
	Times int
	// args
	ID string
	// return
	Name string
	Err  error
}

func SetMockHydraIntrospectCall(r *MockHydraMockRecorder, spec *MockHydra_IntrospectSpec) {
	var info hydra.TokenIntrospectInfo
	setTokenIntrospectInfo(&info, &spec.Info)

	call := r.Introspect(gomock.Not(nil), spec.Token)
	call.Times(spec.Times)
	call.Return(info, spec.Err)
}

type TokenIntrospectInfoSpec struct {
	Active      bool
	VisitorID   string
	ClientID    string
	VisitorType MockVisitorType
}

func setTokenIntrospectInfo(info *hydra.TokenIntrospectInfo, spec *TokenIntrospectInfoSpec) {
	info.Active = spec.Active
	info.VisitorID = spec.VisitorID
	info.ClientID = spec.ClientID
	info.VisitorTyp = spec.VisitorType.VisitorType()
}

type MockVisitorType string

func (t MockVisitorType) VisitorType() hydra.VisitorType {
	switch t {
	case VisitorAPP:
		return hydra.App
	case VisitorUser:
		return hydra.RealName
	default:
		return 0
	}
}

const (
	VisitorAPP  MockVisitorType = "app"
	VisitorUser MockVisitorType = "user"
)

func SetMocGetClientNameByIdCall(r *MockHydraMockRecorder, spec *MockHydra_GetClientNameByIdSpec) {
	call := r.GetClientNameById(gomock.Not(nil), spec.ID)
	call.Times(spec.Times)
	call.Return(spec.Name, spec.Err)
}
