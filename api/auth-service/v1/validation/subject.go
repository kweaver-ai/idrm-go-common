package validation

import (
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// 支持的访问者类型
var supportedSubjectTypes = sets.New(
	v1.SubjectAPP,
	v1.SubjectUser,
)

// 验证访问者
func ValidateSubject(sub *v1.Subject, fldPath *field.Path) (allErrs field.ErrorList) {
	// 验证访问者类型
	if sub.Type == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("subject_type"), ""))
	} else if !supportedSubjectTypes.Has(sub.Type) {
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("subject_type"), sub.Type, sets.List(supportedSubjectTypes)))
	}
	// 验证访问者 ID
	if sub.ID != "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("subject_id"), ""))
	}
	return
}
