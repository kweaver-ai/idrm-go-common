package v2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	"github.com/kweaver-ai/idrm-go-common/rest/user_management"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHydra 模拟 Hydra 客户端
type MockHydra struct {
	mock.Mock
}

func (m *MockHydra) Introspect(ctx context.Context, token string) (hydra.TokenIntrospectInfo, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(hydra.TokenIntrospectInfo), args.Error(1)
}

func (m *MockHydra) GetClientNameById(ctx context.Context, clientID string) (string, error) {
	args := m.Called(ctx, clientID)
	return args.String(0), args.Error(1)
}

func (m *MockHydra) GetClientCredentialToken(ctx context.Context, clientID, clientSecret string) (string, int64, error) {
	args := m.Called(ctx, clientID, clientSecret)
	return args.String(0), args.Get(1).(int64), args.Error(2)
}

func (m *MockHydra) RegistCredentialClient(ctx context.Context, clientID, clientSecret string) (string, error) {
	args := m.Called(ctx, clientID, clientSecret)
	return args.String(0), args.Error(1)
}

// MockUserMgm 模拟用户管理服务
type MockUserMgm struct {
	mock.Mock
}

func (m *MockUserMgm) GetUserNameByUserID(ctx context.Context, userID string) (string, bool, []*user_management.DepInfo, error) {
	args := m.Called(ctx, userID)
	return args.String(0), args.Bool(1), args.Get(2).([]*user_management.DepInfo), args.Error(3)
}

func (m *MockUserMgm) GetAppInfo(ctx context.Context, appID string) (user_management.AppInfo, error) {
	args := m.Called(ctx, appID)
	return args.Get(0).(user_management.AppInfo), args.Error(1)
}

func (m *MockUserMgm) GetUserRolesByUserID(ctx context.Context, userID string) ([]user_management.RoleType, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]user_management.RoleType), args.Error(1)
}

func (m *MockUserMgm) GetDepAllUsers(ctx context.Context, depID string) ([]string, error) {
	args := m.Called(ctx, depID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserMgm) GetDepAllUserInfos(ctx context.Context, depID string) ([]user_management.UserInfo, error) {
	args := m.Called(ctx, depID)
	return args.Get(0).([]user_management.UserInfo), args.Error(1)
}

func (m *MockUserMgm) GetDirectDepAllUserInfos(ctx context.Context, depId string) ([]string, error) {
	args := m.Called(ctx, depId)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserMgm) GetGroupMembers(ctx context.Context, groupID string) ([]string, []string, error) {
	args := m.Called(ctx, groupID)
	return args.Get(0).([]string), args.Get(1).([]string), args.Error(2)
}

func (m *MockUserMgm) GetNameByAccessorIDs(ctx context.Context, accessorIDs map[string]user_management.AccessorType) (map[string]string, error) {
	args := m.Called(ctx, accessorIDs)
	return args.Get(0).(map[string]string), args.Error(1)
}

func (m *MockUserMgm) GetUserManagementAppList(ctx context.Context, args *user_management.AppListArgs) (*base.PageResult[user_management.AppEntry], error) {
	return nil, nil
}

func (m *MockUserMgm) GetDepIDsByUserID(ctx context.Context, userID string) ([]string, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserMgm) BatchGetUserInfoByID(ctx context.Context, userIDs []string) (map[string]user_management.UserInfo, error) {
	args := m.Called(ctx, userIDs)
	return args.Get(0).(map[string]user_management.UserInfo), args.Error(1)
}

func (m *MockUserMgm) GetAccessorIDsByUserID(ctx context.Context, userID string) ([]string, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserMgm) GetAccessorIDsByDepartID(ctx context.Context, depID string) ([]string, error) {
	args := m.Called(ctx, depID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockUserMgm) GetUserParentDepartments(ctx context.Context, userID string) ([][]user_management.Department, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([][]user_management.Department), args.Error(1)
}

func (m *MockUserMgm) GetUserInfoByID(ctx context.Context, userID string) (user_management.UserInfo, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(user_management.UserInfo), args.Error(1)
}

func (m *MockUserMgm) GetDepartments(ctx context.Context, level int) ([]*user_management.DepartmentInfo, error) {
	args := m.Called(ctx, level)
	return args.Get(0).([]*user_management.DepartmentInfo), args.Error(1)
}

func (m *MockUserMgm) GetDepartmentParentInfo(ctx context.Context, ids, fields string) ([]*user_management.DepartmentParentInfo, error) {
	args := m.Called(ctx, ids, fields)
	return args.Get(0).([]*user_management.DepartmentParentInfo), args.Error(1)
}

func (m *MockUserMgm) GetDepartmentInfo(ctx context.Context, departmentIds []string, fields string) ([]*user_management.DepartmentInfo, error) {
	args := m.Called(ctx, departmentIds, fields)
	return args.Get(0).([]*user_management.DepartmentInfo), args.Error(1)
}

func (m *MockUserMgm) GetUserManagementApp(ctx context.Context, id string) (*user_management.AppEntry, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*user_management.AppEntry), args.Error(1)
}

func TestTokenInterception_NoToken(t *testing.T) {
	mockHydra := new(MockHydra)
	mockUserMgm := new(MockUserMgm)
	middleware := &Middleware{
		hydra:   mockHydra,
		userMgm: mockUserMgm,
	}

	handler := middleware.TokenInterception()(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestTokenInterception_InvalidToken(t *testing.T) {
	mockHydra := new(MockHydra)
	mockUserMgm := new(MockUserMgm)
	middleware := &Middleware{
		hydra:   mockHydra,
		userMgm: mockUserMgm,
	}

	mockHydra.On("Introspect", mock.Anything, "test-token").Return(hydra.TokenIntrospectInfo{
		Active: false,
	}, nil)

	handler := middleware.TokenInterception()(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestTokenInterception_UserToken_Success(t *testing.T) {
	mockHydra := new(MockHydra)
	mockUserMgm := new(MockUserMgm)
	middleware := &Middleware{
		hydra:   mockHydra,
		userMgm: mockUserMgm,
	}

	token := "user-token"
	mockHydra.On("Introspect", mock.Anything, token).Return(hydra.TokenIntrospectInfo{
		Active:     true,
		VisitorID:  "user-123",
		ClientID:   "client-456",
		VisitorTyp: hydra.RealName,
	}, nil)

	mockUserMgm.On("GetUserNameByUserID", mock.Anything, "user-123").Return("TestUser", true, []*user_management.DepInfo{}, nil)

	var capturedCtx context.Context
	handler := middleware.TokenInterception()(func(w http.ResponseWriter, r *http.Request) {
		capturedCtx = r.Context()
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, capturedCtx)

	// 验证 Context 中的值
	tokenVal := capturedCtx.Value(interception.Token)
	assert.NotNil(t, tokenVal)
	assert.Equal(t, "Bearer "+token, tokenVal)

	infoName := capturedCtx.Value(interception.InfoName)
	assert.NotNil(t, infoName)
}

func TestTokenInterception_AppToken_Success(t *testing.T) {
	mockHydra := new(MockHydra)
	mockUserMgm := new(MockUserMgm)
	middleware := &Middleware{
		hydra:   mockHydra,
		userMgm: mockUserMgm,
	}

	token := "app-token"
	mockHydra.On("Introspect", mock.Anything, token).Return(hydra.TokenIntrospectInfo{
		Active:     true,
		VisitorID:  "app-123",
		ClientID:   "app-123",
		VisitorTyp: hydra.App,
	}, nil)

	mockUserMgm.On("GetAppInfo", mock.Anything, "app-123").Return(user_management.AppInfo{
		ID:   "app-123",
		Name: "TestApp",
	}, nil)

	var capturedCtx context.Context
	handler := middleware.TokenInterception()(func(w http.ResponseWriter, r *http.Request) {
		capturedCtx = r.Context()
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	handler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, capturedCtx)

	tokenVal := capturedCtx.Value(interception.Token)
	assert.NotNil(t, tokenVal)

	tokenType := capturedCtx.Value(interception.TokenType)
	assert.NotNil(t, tokenType)
}
