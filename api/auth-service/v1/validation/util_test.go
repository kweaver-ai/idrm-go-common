package validation

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

// 参数校验规则与字段路径无关所以所有用例都是用这个作为字段路径
var fldPathTest = field.NewPath("test")

func Test_validateRequiredUUID(t *testing.T) {
	type args struct {
		id      string
		fldPath *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "valid",
			args: args{id: "01916960-4614-795d-956e-d91540963e14", fldPath: fldPathTest},
		},
		{
			name: "missing",
			args: args{fldPath: fldPathTest},
			want: field.ErrorList{{Type: field.ErrorTypeRequired, BadValue: "", Field: fldPathTest.String()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateRequiredUUID(tt.args.id, fldPathTest)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validateUUID(t *testing.T) {
	type args struct {
		id    string
		field *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "valid",
			args: args{id: "01916960-4614-795d-956e-d91540963e14", field: fldPathTest},
		},
		{
			name: "invalid",
			args: args{id: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", field: fldPathTest},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    fldPathTest.String(),
					BadValue: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					Detail:   "test 必须是一个有效的 uuid",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateUUID(tt.args.id, tt.args.field)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validateUTF8EncodedStringWithMaxLength(t *testing.T) {
	type args struct {
		s         string
		maxLength int
		fldPath   *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "valid",
			args: args{s: "中文", maxLength: 2, fldPath: fldPathTest},
		},
		{
			name: "missing",
			args: args{maxLength: 2, fldPath: fldPathTest},
			want: field.ErrorList{{Type: field.ErrorTypeRequired, Field: fldPathTest.String(), BadValue: ""}},
		},
		{
			name: "too long",
			args: args{s: "啊啊啊", maxLength: 2, fldPath: fldPathTest},
			want: field.ErrorList{{Type: field.ErrorTypeTooLong, Field: fldPathTest.String(), BadValue: "啊啊啊", Detail: "test 长度必须不超过 2"}},
		},
		{
			name: "not uft8 encoded string",
			args: args{s: "\xff\xfe\xfd", maxLength: 2, fldPath: fldPathTest},
			want: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: fldPathTest.String(), BadValue: "\xff\xfe\xfd", Detail: "不是 uf8 编码"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateUTF8EncodedStringWithMaxLength(tt.args.s, tt.args.maxLength, tt.args.fldPath)
			assert.Equal(t, tt.want, got)
		})
	}
}

// 2 ^ 63 - 1
const uint64_2_63 uint64 = 1 << 63

func Test_validateRequiredSonyflakeID(t *testing.T) {
	type args struct {
		id      string
		fldPath *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "valid",
			args: args{id: strconv.FormatUint(uint64_2_63-1, 10), fldPath: fldPathTest},
		},
		{
			name: "missing",
			args: args{fldPath: fldPathTest},
			want: field.ErrorList{{Type: field.ErrorTypeRequired, BadValue: "", Field: fldPathTest.String()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateRequiredSonyflakeID(tt.args.id, fldPathTest)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validateSonyflakeID(t *testing.T) {

	type args struct {
		id      string
		fldPath *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "2^63-1",
			args: args{id: strconv.FormatUint(uint64_2_63-1, 10), fldPath: fldPathTest},
		},
		{
			name: "2^63",
			args: args{id: strconv.FormatUint(uint64_2_63, 10), fldPath: fldPathTest},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    fldPathTest.String(),
					BadValue: strconv.FormatUint(uint64_2_63, 10),
					Detail:   "9223372036854775808 必须是一个有效的 sonyflake id",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateSonyflakeID(tt.args.id, tt.args.fldPath)
			assert.Equal(t, tt.want, got)
		})
	}
}
