package v1

import (
	"path"

	"github.com/gin-gonic/gin"

	v1 "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
	"github.com/kweaver-ai/idrm-go-common/audit"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
)

// AuditLogger implements middleware.Middleware.
func (m *Middleware) AuditLogger() gin.HandlerFunc {
	return AuditLogger(m.auditLogger, m.configurationCenterDriven)
}

// AuditLogger 返回配置审计日志日志器的 gin 中间件
func AuditLogger(logger audit.Logger, configurationCenter configuration_center.Driven) gin.HandlerFunc {
	return func(c *gin.Context) {
		if logger.IsZero() {
			return
		}

		// get operator(user or application) from context
		operator := operatorFromContextWithClient(c, configurationCenter)

		// create new logger with operator
		logger := logger.WithOperator(operator)

		// save operator into context
		//  1. gin.Context
		//  2. http.Request.Context
		audit.SetCustomContext(c, logger)
		c.Request = c.Request.WithContext(audit.NewContext(c.Request.Context(), logger))
	}
}

func AuditDepartmentsFromUserParentDeps(userParentDeps []configuration_center.DepartmentPath) (auditDepartments []v1.Department) {
	return auditDepartmentsFromUserParentDeps(userParentDeps)
}

// auditDepartmentsFromUserParentDeps converts
// []configuration_center.DepartmentPath to []audit.Department.
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

// operatorFromContextWithClient 返回请求的操作者及其所属部门。操作者可能是用户或应用。
func operatorFromContextWithClient(c *gin.Context, client configuration_center.Driven) (operator v1.Operator) {
	// agent
	operator.Agent = audit.AgentFromRequest(c.Request)

	// 获取操作者，操作者可能是用户、应用
	if u, err := UserFromContext(c); err != nil {
		operator.Type = v1.OperatorUnknown
	} else {
		operator.ID = u.ID
		operator.Name = u.Name
	}

	// 获取操作者类型。只有用户类型的操作者需要获取所属部门
	if operator.Type = operatorTypeFromContext(c); operator.Type != v1.OperatorAuthenticatedUser {
		return
	}

	// 获取用户所属部门
	user, err := client.GetUser(c, operator.ID, configuration_center.GetUserOptions{
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

// operatorFromContextWithClient 返回请求的操作者类型
func operatorTypeFromContext(c *gin.Context) v1.OperatorType {
	tokenType, ok := c.Get(interception.TokenType)
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
