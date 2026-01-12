package validation

import (
	auth_service_v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// 验证 Policy
func ValidatePolicy(p *auth_service_v1.Policy, fldPath *field.Path) (allErrs field.ErrorList) {
	// 验证 Subject
	allErrs = append(allErrs, ValidateSubject(&p.Subject, fldPath)...)
	// 验证 Object
	allErrs = append(allErrs, ValidateObject(&p.Object, fldPath)...)
	// 验证 Action
	allErrs = append(allErrs, ValidateAction(p.Action, fldPath.Child("action"))...)
	// 验证是否可以把 Object 授权给 Subject
	allErrs = append(allErrs, ValidateSubjectTypeAndObjectType(p.Subject.Type, p.Object.Type, fldPath)...)
	return
}

// 验证 Policies
func ValidatePolicies(polices []auth_service_v1.Policy, fldPath *field.Path) (allErrs field.ErrorList) {
	for i, p := range polices {
		allErrs = append(allErrs, ValidatePolicy(&p, fldPath.Index(i))...)
	}
	return
}
