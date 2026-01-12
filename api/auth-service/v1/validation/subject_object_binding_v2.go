package validation

import (
	"fmt"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// supportedSubjectTypeObjectTypeBindingsV2 定义支持的访问者类型和资源类型的关系
var supportedSubjectTypeObjectTypeBindingsV2 = sets.New[subjectTypeObjectTypeBindingV2]()

// subjectTypeObjectTypeBindingV2 定义访问者类型和资源类型的关系
type subjectTypeObjectTypeBindingV2 struct {
	subjectType v1.SubjectType
	objectType  v1.ObjectType
	Action      v1.Action
}

func binding(subjectType v1.SubjectType, objectType v1.ObjectType, actions ...v1.Action) {
	for _, action := range actions {
		supportedSubjectTypeObjectTypeBindingsV2.Insert(subjectTypeObjectTypeBindingV2{
			subjectType: subjectType,
			objectType:  objectType,
			Action:      action,
		})
	}
}

var totalActions = []v1.Action{
	v1.ActionRead,
	v1.ActionDownload,
	v1.ActionAuth,
	v1.ActionAllocate,
}

func init() {
	//访问是接口，只能给读取权限
	binding(v1.SubjectAPP, v1.ObjectAPI, v1.ActionRead)
	binding(v1.SubjectAPP, v1.ObjectDataView, v1.ActionRead)
	binding(v1.SubjectAPP, v1.ObjectIndicator, v1.ActionRead)
	binding(v1.SubjectAPP, v1.ObjectIndicatorDimensionalRule, v1.ActionRead)
	binding(v1.SubjectAPP, v1.ObjectSubView, v1.ActionRead)
	binding(v1.SubjectAPP, v1.ObjectSubService, v1.ActionRead)
	//访问者是用户，用户可以被授予全部权限
	binding(v1.SubjectUser, v1.ObjectAPI, totalActions...)
	binding(v1.SubjectUser, v1.ObjectDataView, totalActions...)
	binding(v1.SubjectUser, v1.ObjectIndicator, totalActions...)
	binding(v1.SubjectUser, v1.ObjectIndicatorDimensionalRule, totalActions...)
	binding(v1.SubjectUser, v1.ObjectSubView, totalActions...)
	binding(v1.SubjectUser, v1.ObjectSubService, totalActions...)
}

func ValidateSubjectTypeAndObjectTypeV2(st v1.SubjectType, ot v1.ObjectType, action v1.Action, fldPath *field.Path) (allErrs field.ErrorList) {
	if !supportedSubjectTypeObjectTypeBindingsV2.Has(subjectTypeObjectTypeBindingV2{st, ot, action}) {
		allErrs = append(allErrs, field.Invalid(fldPath, st, fmt.Sprintf("不支持授权资源 %s 给访问者 %s", ot, st)))
	}
	return
}
