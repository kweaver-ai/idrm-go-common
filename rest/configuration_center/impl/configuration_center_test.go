package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-frame/core/logx/zapx"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func TestMain(m *testing.M) {
	// 初始化日志器避免调用时 panic
	log.InitLogger(zapx.LogConfigs{}, &telemetry.Config{})

	m.Run()
}

type fakeRoundTripper struct {
	// got
	req *http.Request
	// want
	resp *http.Response
	err  error
}

// RoundTrip implements http.RoundTripper.
func (rt *fakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.req = req
	rt.resp.Request = req
	return rt.resp, rt.err
}

var _ http.RoundTripper = &fakeRoundTripper{}

func TestConfigurationCenterDriven_UsersRoles(t *testing.T) {
	// testdata
	var (
		role0 = configuration_center.Role{ID: "0000", Name: "ROLE_0000"}
		role1 = configuration_center.Role{ID: "1111", Name: "ROLE_1111"}
		role2 = configuration_center.Role{ID: "2222", Name: "ROLE_2222"}
	)
	tests := []struct {
		name        string
		baseURL     string
		rt          *fakeRoundTripper
		bearerToken string
		want        []configuration_center.Role
		wantErr     error
	}{
		{
			name:        "ok",
			baseURL:     "https://configuration-center.example.org",
			rt:          &fakeRoundTripper{resp: newHTTPResponseForJSON(t, http.StatusOK, []configuration_center.Role{role0, role1, role2})},
			bearerToken: "BEARER_TOKEN_0000",
			want:        []configuration_center.Role{role0, role1, role2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConfigurationCenterDriven{baseURL: tt.baseURL, client: &http.Client{Transport: tt.rt}}

			got, err := c.UsersRoles(interception.NewContextWithBearerToken(context.Background(), tt.bearerToken))
			// assert underlying request
			assert.Equal(t, http.MethodGet, tt.rt.req.Method, "Request Method")
			assert.Equal(t, fmt.Sprintf("%s/api/configuration-center/v1/users/roles", tt.baseURL), tt.rt.req.URL.String(), "Request URL")
			assert.Equal(t, fmt.Sprintf("Bearer %s", tt.bearerToken), tt.rt.req.Header.Get("Authorization"), "Request Header: Authorization")
			// assert got
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestConfigurationCenterDriven_GetRoleIDs(t *testing.T) {
	// testdata
	var (
		baseURL = loadRequiredEnv(t, "TEST_BASE_URL")
		userID  = loadRequiredEnv(t, "TEST_USER_ID")
	)
	driven := &ConfigurationCenterDriven{baseURL: baseURL, client: http.DefaultClient}
	roleIDs, err := driven.GetRoleIDs(context.Background(), userID)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range roleIDs {
		t.Logf("role id: %s", id)
	}
	t.Logf("round %d role id(s)", len(roleIDs))
}

func TestConfigurationCenterDriven_ListApplicationsByDeveloperID(t *testing.T) {
	// testdata
	var (
		baseURL = loadRequiredEnv(t, "TEST_BASE_URL")
		userID  = loadRequiredEnv(t, "TEST_USER_ID")
	)
	driven := &ConfigurationCenterDriven{baseURL: baseURL, client: http.DefaultClient}
	applications, err := driven.ListApplicationsByDeveloperID(context.Background(), userID)
	if err != nil {
		t.Fatal(err)
	}
	for _, app := range applications {
		t.Logf("application: id=%s, name=%q", app.ID, app.Name)
	}
	t.Logf("round %d application(s)", len(applications))
}

func newHTTPResponseForJSON(t *testing.T, code int, data any) *http.Response {
	var body io.Reader
	switch v := data.(type) {
	case io.Reader:
		body = v
	case []byte:
		body = bytes.NewReader(v)
	case string:
		body = bytes.NewBufferString(v)
	default:
		b, err := json.Marshal(data)
		if err != nil {
			t.Fatalf("encode data as json fail: %v", err)
		}
		body = bytes.NewReader(b)
	}

	return &http.Response{Status: http.StatusText(code), StatusCode: code, Body: io.NopCloser(body)}
}
