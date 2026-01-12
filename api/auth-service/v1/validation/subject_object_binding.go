package validation

import (
	"fmt"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// subjectTypeObjectTypeBinding 定义访问者类型和资源类型的关系
type subjectTypeObjectTypeBinding struct {
	subjectType v1.SubjectType
	objectType  v1.ObjectType
}

// supportedSubjectTypeObjectTypeBindings 定义支持的访问者类型和资源类型的关系
var supportedSubjectTypeObjectTypeBindings = sets.New(
	subjectTypeObjectTypeBinding{v1.SubjectAPP, v1.ObjectAPI},
	subjectTypeObjectTypeBinding{v1.SubjectAPP, v1.ObjectDataView},
	subjectTypeObjectTypeBinding{v1.SubjectAPP, v1.ObjectIndicator},
	subjectTypeObjectTypeBinding{v1.SubjectAPP, v1.ObjectIndicatorDimensionalRule},
	subjectTypeObjectTypeBinding{v1.SubjectAPP, v1.ObjectSubView},
	subjectTypeObjectTypeBinding{v1.SubjectUser, v1.ObjectDataView},
	subjectTypeObjectTypeBinding{v1.SubjectUser, v1.ObjectIndicator},
	subjectTypeObjectTypeBinding{v1.SubjectUser, v1.ObjectIndicatorDimensionalRule},
	subjectTypeObjectTypeBinding{v1.SubjectUser, v1.ObjectSubView},
)

func ValidateSubjectTypeAndObjectType(st v1.SubjectType, ot v1.ObjectType, fldPath *field.Path) (allErrs field.ErrorList) {
	if !supportedSubjectTypeObjectTypeBindings.Has(subjectTypeObjectTypeBinding{st, ot}) {
		allErrs = append(allErrs, field.Invalid(fldPath, st, fmt.Sprintf("不支持授权资源 %s 给访问者 %s", ot, st)))
	}
	return
}
