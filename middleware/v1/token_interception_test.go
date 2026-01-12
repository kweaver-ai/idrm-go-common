package v1

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	configuration_center "github.com/kweaver-ai/idrm-go-common/rest/configuration_center/mock"
	hydra "github.com/kweaver-ai/idrm-go-common/rest/hydra/mock"
)

func TestMiddleware_ShouldTokenInterception(t *testing.T) {
	// testdata
	const (
		uuid_0 = "00000000-0000-0000-0000-000000000000"
		uuid_1 = "11111111-1111-1111-1111-111111111111"

		appName_0 = "APP_NAME_0"

		bearerToken_0 string = "BEARER_TOKEN_0"
	)

	if !assert.NotEqual(t, middleware.VirtualEngineApp, appName_0, "appName_0 不应该与虚拟化引擎使用的名称相同") {
		return
	}

	type want struct {
		auth        string
		bearerToken string
		subject     *v1.Subject
	}
	tests := []struct {
		name                string
		auth                string
		hydra               hydra.MockHydraSpec
		configurationCenter configuration_center.MockDrivenSpec
		want                want
	}{
		{
			name: "anonymous",
		},
		{
			name: "non bearer authorization",
			auth: "NON BEARER AUTHORIZATION",
			want: want{
				auth: "NON BEARER AUTHORIZATION",
			},
		},
		{
			name: "inactive",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info:  hydra.TokenIntrospectInfoSpec{},
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
		{
			name: "user",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						Active:      true,
						VisitorID:   uuid_0,
						ClientID:    uuid_1,
						VisitorType: hydra.VisitorUser,
					},
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
				subject: &v1.Subject{
					Type: v1.SubjectUser,
					ID:   uuid_0,
				},
			},
		},
		{
			name: "user inactive",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						VisitorID:   uuid_0,
						ClientID:    uuid_1,
						VisitorType: hydra.VisitorUser,
					},
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
		{
			name: "app",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						Active:      true,
						VisitorID:   uuid_0,
						ClientID:    uuid_0,
						VisitorType: hydra.VisitorAPP,
					},
				},
				GetClientNameById: hydra.MockHydra_GetClientNameByIdSpec{
					Times: 1,
					ID:    uuid_0,
					Name:  appName_0,
				},
			},
			configurationCenter: configuration_center.MockDrivenSpec{
				GetAppsByAccountId: configuration_center.MockDriven_GetAppsByAccountIdSpec{
					Times: 1,
					ID:    uuid_0,
					APP: configuration_center.MockAPPSpec{
						ID: uuid_1,
					},
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
				subject: &v1.Subject{
					Type: v1.SubjectAPP,
					ID:   uuid_1,
				},
			},
		},
		{
			name: "app inactive",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						VisitorID:   uuid_0,
						ClientID:    uuid_0,
						VisitorType: hydra.VisitorAPP,
					},
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
		{
			name: "virtual engine",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						Active:      true,
						VisitorID:   uuid_0,
						ClientID:    uuid_0,
						VisitorType: hydra.VisitorAPP,
					},
				},
				GetClientNameById: hydra.MockHydra_GetClientNameByIdSpec{
					Times: 1,
					ID:    uuid_0,
					Name:  middleware.VirtualEngineApp,
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
		{
			name: "introspect hydra token fail",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Err:   assert.AnError,
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
		{
			name: "get hydra client name fail",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						Active:      true,
						VisitorID:   uuid_0,
						ClientID:    uuid_0,
						VisitorType: hydra.VisitorAPP,
					},
				},
				GetClientNameById: hydra.MockHydra_GetClientNameByIdSpec{
					Times: 1,
					ID:    uuid_0,
					Err:   assert.AnError,
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
		{
			name: "get app fail",
			auth: "Bearer " + bearerToken_0,
			hydra: hydra.MockHydraSpec{
				Introspect: hydra.MockHydra_IntrospectSpec{
					Times: 1,
					Token: bearerToken_0,
					Info: hydra.TokenIntrospectInfoSpec{
						Active:      true,
						VisitorID:   uuid_0,
						ClientID:    uuid_0,
						VisitorType: hydra.VisitorAPP,
					},
				},
				GetClientNameById: hydra.MockHydra_GetClientNameByIdSpec{
					Times: 1,
					ID:    uuid_0,
					Name:  appName_0,
				},
			},
			configurationCenter: configuration_center.MockDrivenSpec{
				GetAppsByAccountId: configuration_center.MockDriven_GetAppsByAccountIdSpec{
					Times: 1,
					ID:    uuid_0,
					APP: configuration_center.MockAPPSpec{
						ID: uuid_1,
					},
					Err: assert.AnError,
				},
			},
			want: want{
				auth:        "Bearer " + bearerToken_0,
				bearerToken: bearerToken_0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if tt.name != "app inactive" {
			// 	t.Skip()
			// }

			ctrl := gomock.NewController(t)
			defer ctrl.Satisfied()

			var (
				mockHydra                     = hydra.NewMockHydra(ctrl)
				mockConfigurationCenterDriven = configuration_center.NewMockDriven(ctrl)
			)
			m := &Middleware{
				hydra:                     mockHydra,
				configurationCenterDriven: mockConfigurationCenterDriven,
			}

			header := make(http.Header)
			header.Set("Authorization", tt.auth)
			request := &http.Request{Header: header}

			c := &gin.Context{Request: request}

			// mock hydra
			hydra.SetMockHydraMockRecorder(mockHydra.EXPECT(), &tt.hydra)
			// mock configuration center
			configuration_center.SetMockDriven(mockConfigurationCenterDriven.EXPECT(), &tt.configurationCenter)

			m.ShouldTokenInterception()(c)

			// assertion: HTTP Request Header: Authorization
			assertValueFromContextFunc(t, interception.AuthFromContext, c, tt.want.auth, "AuthFromContext")
			// assertion: Bearer Token
			assertValueFromContextFunc(t, interception.BearerTokenFromContext, c, tt.want.bearerToken, "BearerTokenFromContext")
			// assertion: auth-service Subject
			assertValueFromContextFunc(t, interception.AuthServiceSubjectFromContext, c, tt.want.subject, "BearerTokenFromContext")
		})
	}
}

func assertValueFromContextFunc[T any](t *testing.T, f func(context.Context) (T, error), ctx context.Context, want T, msgAndArgs ...any) bool {
	t.Helper()

	var wantError error
	if reflect.ValueOf(want).IsZero() {
		wantError = interception.ErrNotExist
	}

	got, err := f(ctx)

	var conditions []bool
	conditions = append(conditions, assert.Equal(t, want, got, msgAndArgs...))
	conditions = append(conditions, assert.ErrorIs(t, err, wantError, msgAndArgs...))

	for _, c := range conditions {
		if c {
			continue
		}
		return false
	}
	return true
}
