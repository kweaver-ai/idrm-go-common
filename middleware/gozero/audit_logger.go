package gozero

import (
	"context"
	"net/http"
	"path"

	v1 "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
	"github.com/kweaver-ai/idrm-go-common/audit"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
)

// AuditLogger 审计日志中间件，记录操作者（用户或应用）信息到审计上下文
func (m *Middleware) AuditLogger() MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if m.auditLogger.IsZero() {
				next(w, r)
				return
			}

			// 获取操作者（用户或应用）
			operator := operatorFromContextWithClient(r, m.configurationCenterDriven)

			// 创建带有操作者的 logger
			logger := m.auditLogger.WithOperator(operator)

			// 保存到 context
			ctx := audit.NewContext(r.Context(), logger)
			next(w, r.WithContext(ctx))
		}
	}
}

// AuditDepartmentsFromUserParentDeps 将用户父部门路径转换为审计部门格式
func AuditDepartmentsFromUserParentDeps(userParentDeps []configuration_center.DepartmentPath) (auditDepartments []v1.Department) {
	return auditDepartmentsFromUserParentDeps(userParentDeps)
}

// auditDepartmentsFromUserParentDeps 转换 DepartmentPath 为 audit.Department
func auditDepartmentsFromUserParentDeps(userParentDeps []configuration_center.DepartmentPath) (auditDepartments []v1.Department) {
	for _, userParentDep := range userParentDeps {
		var ids, names []string
		for _, d := range userParentDep {
			ids, names = append(ids, d.ID), append(names, d.Name)
		}

		auditDepartment := v1.Department{
			ID:   path.Join(ids...),
			Name: path.Join(names...),
		}

		auditDepartments = append(auditDepartments, auditDepartment)
	}
	return
}

// operatorFromContextWithClient 返回请求的操作者及其所属部门
func operatorFromContextWithClient(r *http.Request, client configuration_center.Driven) (operator v1.Operator) {
	// agent
	operator.Agent = audit.AgentFromRequest(r)

	// 获取操作者类型
	operator.Type = operatorTypeFromContext(r.Context())
	if operator.Type == v1.OperatorUnknown {
		return
	}

	// 获取操作者 ID 和 Name
	if user, ok := r.Context().Value(interception.InfoName).(*middleware.User); ok {
		operator.ID = user.ID
		operator.Name = user.Name
	}

	// 只有用户类型的操作者需要获取所属部门
	if operator.Type != v1.OperatorAuthenticatedUser {
		return
	}

	// 获取用户所属部门
	user, err := client.GetUser(r.Context(), operator.ID, configuration_center.GetUserOptions{
		Fields: []configuration_center.Field{
			configuration_center.FieldName,
			configuration_center.FieldParentDeps,
		},
	})
	if err != nil {
		return
	}
	operator.Department = auditDepartmentsFromUserParentDeps(user.ParentDeps)

	return
}

// operatorTypeFromContext 返回请求的操作者类型
func operatorTypeFromContext(ctx context.Context) v1.OperatorType {
	tokenType, ok := ctx.Value(interception.TokenType).(int)
	if !ok {
		return v1.OperatorUnknown
	}

	switch tokenType {
	case interception.TokenTypeClient:
		return v1.OperatorAPP
	case interception.TokenTypeUser:
		return v1.OperatorAuthenticatedUser
	default:
		return v1.OperatorUnknown
	}
}
