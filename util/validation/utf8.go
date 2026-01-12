package validation

import (
	"unicode/utf8"

	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

type UTF8EncodedStringValidationOptions struct {
	// 允许为空
	AllowEmpty bool
	// 长度限制，0 代表无限制
	LengthLimit int
}

// ValidateUTF8EncodedString 验证 UTF8 编码的字符串
func ValidateUTF8EncodedString(s string, fldPath *field.Path, opts *UTF8EncodedStringValidationOptions) (allErrs field.ErrorList) {
	if !utf8.ValidString(s) {
		allErrs = append(allErrs, field.Invalid(fldPath, s, "不是 uf8 编码"))
		return
	}

	if !opts.AllowEmpty && len(s) == 0 {
		allErrs = append(allErrs, field.Required(fldPath, ""))
		return
	}

	if opts.LengthLimit != 0 && opts.LengthLimit < utf8.RuneCountInString(s) {
		allErrs = append(allErrs, field.TooLong(fldPath, s, opts.LengthLimit))
	}

	return
}
