package validation

import (
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// 支持的资源类型
var supportedObjectTypes = sets.New(
	v1.ObjectAPI,
	v1.ObjectDataView,
	v1.ObjectIndicator,
	v1.ObjectIndicatorDimensionalRule,
	v1.ObjectSubView,
)

// 验证资源
func ValidateObject(obj *v1.Object, fldPath *field.Path) (allErrs field.ErrorList) {
	// 验证访问者类型
	if obj.Type == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("object_type"), ""))
	} else if !supportedObjectTypes.Has(obj.Type) {
		allErrs = append(allErrs, field.NotSupported(fldPath.Child("object_type"), obj.Type, sets.List(supportedObjectTypes)))
	}
	// 验证访问者 ID
	if obj.ID != "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("object_id"), ""))
	}
	return
}
