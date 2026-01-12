package validation

import (
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// 支持的操作
var supportedActions = sets.New(
	v1.ActionDownload,
	v1.ActionRead,
)

// 验证操作
func ValidateAction(act v1.Action, fldPath *field.Path) (allErrs field.ErrorList) {
	if act == "" {
		allErrs = append(allErrs, field.Required(fldPath, ""))
	} else if !supportedActions.Has(act) {
		allErrs = append(allErrs, field.NotSupported(fldPath, act, sets.List(supportedActions)))
	}
	return
}
