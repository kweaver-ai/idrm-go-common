package validation

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/sony/sonyflake"

	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// validateRequiredUUID 验证 id 非空且是一个 uuid
func validateRequiredUUID(id string, fldPath *field.Path) (allErrs field.ErrorList) {
	if id == "" {
		allErrs = append(allErrs, field.Required(fldPath, ""))
		return
	}

	allErrs = append(allErrs, validateUUID(id, fldPath)...)
	return
}

// validateUUID 验证 id 是一个 uuid
//
// Deprecated: Use GoCommon/util/validation.ValidateUUID instead.
func validateUUID(id string, fldPath *field.Path) (allErrs field.ErrorList) {
	if _, err := uuid.Parse(id); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, id, fmt.Sprintf("%s 必须是一个有效的 uuid", fldPath)))
	}
	return
}

// Deprecated: Use GoCommon/util/validation.ValidateUTF8EncodedString instead.
func validateUTF8EncodedStringWithMaxLength(s string, maxLength int, fldPath *field.Path) (allErrs field.ErrorList) {
	if len(s) == 0 {
		allErrs = append(allErrs, field.Required(fldPath, ""))
		return
	}

	if !utf8.ValidString(s) {
		allErrs = append(allErrs, field.Invalid(fldPath, s, "不是 uf8 编码"))
		return
	}

	if utf8.RuneCountInString(s) > maxLength {
		allErrs = append(allErrs, field.TooLong(fldPath, s, maxLength))
	}

	return
}

// sonyflake id 的 bit 长度
const BitLenSonyflakeID = sonyflake.BitLenTime + sonyflake.BitLenSequence + sonyflake.BitLenMachineID

// validateRequiredSonyflakeID 验证指定字符串非空且是一个合法的 sonyflake id
func validateRequiredSonyflakeID(id string, fldPath *field.Path) (allErrs field.ErrorList) {
	if id == "" {
		allErrs = append(allErrs, field.Required(fldPath, ""))
		return
	}

	allErrs = append(allErrs, validateSonyflakeID(id, fldPath)...)
	return
}

// validateSonyflakeID 验证指定字符串是一个合法的 sonyflake id
func validateSonyflakeID(id string, fldPath *field.Path) (allErrs field.ErrorList) {
	if _, err := strconv.ParseUint(id, 10, BitLenSonyflakeID); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath, id, fmt.Sprintf("%s 必须是一个有效的 sonyflake id", id)))
	}
	return
}
