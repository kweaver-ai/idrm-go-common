package validation

import (
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

type validateSubjectPolicyOptions struct {
	// subject policy 支持的 action
	supportedActions sets.Set[v1.Action]
}

func validateSubjectPolicy(policy *v1.SubjectPolicy, fldPath *field.Path, opts *validateSubjectPolicyOptions) (allErrs field.ErrorList) {
	// 当前仅支持 app, user，不支持 role，domain
	if policy.Subject.Type == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("subject_type"), ""))
	} else if !supportedSubjectTypes.Has(policy.Type) {
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("subject_type"), policy.Subject.Type, sets.List(supportedSubjectTypes)))
	}

	allErrs = append(allErrs, validateRequiredUUID(policy.Subject.ID, fldPath.Child("subject_id"))...)

	seen := sets.New[v1.Action]()
	for i, act := range policy.Actions {
		if !opts.supportedActions.Has(act) {
			allErrs = append(allErrs, field.NotSupported(fldPath.Child("actions").Index(i), act, sets.List(opts.supportedActions)))
			continue
		}
		if seen.Has(act) {
			allErrs = append(allErrs, field.Duplicate(fldPath.Child("actions").Index(i), act))
			continue
		}
		seen.Insert(act)
	}
	if len(seen) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("actions"), ""))
	}

	return
}
