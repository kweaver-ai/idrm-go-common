package audit

import (
	"fmt"

	audit "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
)

const descriptionFormat = "%s%q%s%q"
const descriptionFormatWithoutObject = "%s%q%s"

func GenerateSimplifiedChineseDescription(operator *audit.Operator, operation audit.Operation, object audit.ResourceObject) string {
	return generateSimplifiedChineseDescription(operator, operation, object)
}

// generateSimplifiedChineseDescription generates Simplified Chinese
func generateSimplifiedChineseDescription(operator *audit.Operator, operation audit.Operation, object audit.ResourceObject) string {
	// 操作者的类型、名称
	operatorType, operatorName := generateSimplifiedChineseOperator(operator)

	var operationName string = operation.SimplifiedChineseName()

	if object.GetName() == "" {
		return fmt.Sprintf(descriptionFormatWithoutObject, operatorType, operatorName, operationName)
	}
	return fmt.Sprintf(descriptionFormat, operatorType, operatorName, operationName, object.GetName())
}

const (
	simplifiedChineseOperatorTypeUnknown string = "未知类型操作者"
	simplifiedChineseOperatorNameUnknown string = "未知名称"
)

// generateSimplifiedChineseOperator 返回操作者的简体中文类型和名称，用于生成审
// 计日志的描述
func generateSimplifiedChineseOperator(operator *audit.Operator) (operatorType, operatorName string) {
	if operator == nil {
		return simplifiedChineseOperatorTypeUnknown, simplifiedChineseOperatorNameUnknown
	}

	// 类型
	switch operator.Type {
	case audit.OperatorAPP:
		operatorType = "应用"
	case audit.OperatorAuthenticatedUser:
		operatorType = "用户"
	default:
		operatorType = simplifiedChineseOperatorTypeUnknown
	}

	// 名称
	operatorName = operator.Name
	if operatorName == "" {
		operatorName = simplifiedChineseOperatorNameUnknown
	}

	return
}
