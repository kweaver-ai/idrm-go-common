package validation

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// ValidateUUID 验证 UUID
func ValidateUUID(id string, fldPath *field.Path) (allErrs field.ErrorList) {
	if _, err := uuid.Parse(id); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, id, fmt.Sprintf("%s 必须是一个有效的 uuid", fldPath)))
	}
	return
}
