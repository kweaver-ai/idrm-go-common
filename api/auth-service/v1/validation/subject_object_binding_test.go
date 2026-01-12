package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

var _ = fldPathTest

func TestValidateSubjectTypeAndObjectType(t *testing.T) {
	type args struct {
		sub     v1.SubjectType
		obj     v1.ObjectType
		fldPath *field.Path
	}
	tests := []struct {
		st   v1.SubjectType
		ot   v1.ObjectType
		want field.ErrorList
	}{
		{
			st: v1.SubjectUser,
			ot: v1.ObjectAPI,
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "test",
					BadValue: v1.SubjectUser,
					Detail:   "不支持授权资源 api 给访问者 user",
				},
			},
		},
		{
			st: v1.SubjectUser,
			ot: v1.ObjectDataView,
		},
		{
			st: v1.SubjectUser,
			ot: v1.ObjectIndicator,
		},
		{
			st: v1.SubjectUser,
			ot: v1.ObjectIndicatorDimensionalRule,
		},
		{
			st: v1.SubjectUser,
			ot: v1.ObjectSubView,
		},
		{
			st: v1.SubjectAPP,
			ot: v1.ObjectAPI,
		},
		{
			st: v1.SubjectAPP,
			ot: v1.ObjectDataView,
		},
		{
			st: v1.SubjectAPP,
			ot: v1.ObjectIndicator,
		},
		{
			st: v1.SubjectAPP,
			ot: v1.ObjectIndicatorDimensionalRule,
		},
		{
			st: v1.SubjectAPP,
			ot: v1.ObjectSubView,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.st, tt.ot), func(t *testing.T) {
			got := ValidateSubjectTypeAndObjectType(tt.st, tt.ot, fldPathTest)
			assert.Equal(t, tt.want, got)
		})
	}

	testedBindings := sets.New[subjectTypeObjectTypeBinding]()
	for _, tt := range tests {
		testedBindings.Insert(subjectTypeObjectTypeBinding{tt.st, tt.ot})
	}

	supportedBindings := sets.New[subjectTypeObjectTypeBinding]()
	for st := range supportedSubjectTypes {
		for ot := range supportedObjectTypes {
			supportedBindings.Insert(subjectTypeObjectTypeBinding{st, ot})
		}
	}

	for b := range supportedBindings.Difference(testedBindings) {
		t.Errorf("uncovered test case: %s_%s", b.subjectType, b.objectType)
	}
}
