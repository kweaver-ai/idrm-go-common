package validation

import (
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

func ValidateAPIAuthorizingRequestCreate(req *v1.APIAuthorizingRequest) (allErrs field.ErrorList) {
	allErrs = append(allErrs, validateAPIAuthorizingRequestSpecCreate(&req.Spec, field.NewPath("spec"))...)
	return
}

// api 所支持的 action
var supportedAPIActions = sets.New(v1.ActionRead)

const apiAuthorizingRequestSpecReasonMaxLength = 800

func validateAPIAuthorizingRequestSpecCreate(spec *v1.APIAuthorizingRequestSpec, fldPath *field.Path) (allErrs field.ErrorList) {
	allErrs = append(allErrs, validateRequiredUUID(spec.ID, fldPath.Child("id"))...)

	if len(spec.Policies) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("policies"), ""))
	}
	for i, p := range spec.Policies {
		allErrs = append(allErrs, validateSubjectPolicy(&p, fldPath.Child("policies").Index(i), &validateSubjectPolicyOptions{supportedActions: supportedAPIActions})...)
		allErrs = append(allErrs, ValidateSubjectTypeAndObjectType(p.Type, v1.ObjectAPI, fldPath.Child("policies").Index(i).Child("subject_type"))...)
	}

	// requester_id
	allErrs = append(allErrs, validateRequiredUUID(spec.RequesterID, fldPath.Child("requester_id"))...)

	// reason
	allErrs = append(allErrs, validateUTF8EncodedStringWithMaxLength(spec.Reason, apiAuthorizingRequestSpecReasonMaxLength, fldPath.Child("reason"))...)

	return
}
