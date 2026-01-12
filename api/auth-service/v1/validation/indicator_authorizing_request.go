package validation

import (
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

func ValidateIndicatorAuthorizingRequestCreate(req *v1.IndicatorAuthorizingRequest) (allErrs field.ErrorList) {
	allErrs = append(allErrs, validateIndicatorAuthorizingRequestSpecCreate(&req.Spec, field.NewPath("spec"))...)
	return
}

// indicator 所支持的 action
var supportedIndicatorActions = sets.New(v1.ActionRead)

const indicatorAuthorizingRequestSpecReasonMaxLength = 800

func validateIndicatorAuthorizingRequestSpecCreate(spec *v1.IndicatorAuthorizingRequestSpec, fldPath *field.Path) (allErrs field.ErrorList) {
	allErrs = append(allErrs, validateRequiredSonyflakeID(spec.ID, fldPath.Child("id"))...)

	if len(spec.Policies) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("policies"), ""))
	}
	for i, p := range spec.Policies {
		allErrs = append(allErrs, validateSubjectPolicy(&p, fldPath.Child("policies").Index(i), &validateSubjectPolicyOptions{supportedActions: supportedIndicatorActions})...)
		allErrs = append(allErrs, ValidateSubjectTypeAndObjectType(p.Type, v1.ObjectIndicator, fldPath.Child("policies").Index(i).Child("subject_type"))...)
	}

	// requester_id
	allErrs = append(allErrs, validateRequiredUUID(spec.RequesterID, fldPath.Child("requester_id"))...)

	// reason
	allErrs = append(allErrs, validateUTF8EncodedStringWithMaxLength(spec.Reason, indicatorAuthorizingRequestSpecReasonMaxLength, fldPath.Child("reason"))...)

	return
}
