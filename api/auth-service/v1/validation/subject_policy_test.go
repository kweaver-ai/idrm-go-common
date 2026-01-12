package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

func Test_validateSubjectPolicy(t *testing.T) {
	const (
		userID0 = "00000000-0000-0000-0000-000000000000"
	)
	type args struct {
		policy  *v1.SubjectPolicy
		fldPath *field.Path
		opts    *validateSubjectPolicyOptions
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "supported subject type user",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectUser,
						ID:   userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
						v1.ActionDownload,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
						v1.ActionDownload,
					),
				},
			},
		},
		{
			name: "supported subject type app",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectAPP,
						ID:   userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
						v1.ActionDownload,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
						v1.ActionDownload,
					),
				},
			},
		},
		{
			name: "missing subject type",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						ID: userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
					),
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeRequired,
					Field:    "test.subject_type",
					BadValue: "",
				},
			},
		},
		{
			name: "unsupported subject type role",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectRole,
						ID:   userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
						v1.ActionDownload,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
						v1.ActionDownload,
					),
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeNotSupported,
					Field:    "test.subject_type",
					BadValue: v1.SubjectRole,
					Detail:   "test.subject_type 必须是 app, user 中的一个",
				},
			},
		},
		{
			name: "unsupported subject type department",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectDepartment,
						ID:   userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
						v1.ActionDownload,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
						v1.ActionDownload,
					),
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeNotSupported,
					Field:    "test.subject_type",
					BadValue: v1.SubjectDepartment,
					Detail:   "test.subject_type 必须是 app, user 中的一个",
				},
			},
		},
		{
			name: "missing actions",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectUser,
						ID:   userID0,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
					),
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeRequired,
					Field:    "test.actions",
					BadValue: "",
				},
			},
		},
		{
			name: "unsupported action",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectUser,
						ID:   userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
						v1.ActionDownload,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
					),
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeNotSupported,
					Field:    "test.actions[1]",
					BadValue: v1.ActionDownload,
					Detail:   "test.actions[1] 必须是 read",
				},
			},
		},
		{
			name: "duplicated actions",
			args: args{
				policy: &v1.SubjectPolicy{
					Subject: v1.Subject{
						Type: v1.SubjectUser,
						ID:   userID0,
					},
					Actions: []v1.Action{
						v1.ActionRead,
						v1.ActionDownload,
						v1.ActionRead,
						v1.ActionDownload,
						v1.ActionRead,
					},
				},
				fldPath: fldPathTest,
				opts: &validateSubjectPolicyOptions{
					supportedActions: sets.New(
						v1.ActionRead,
						v1.ActionDownload,
					),
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeDuplicate,
					Field:    "test.actions[2]",
					BadValue: v1.ActionRead,
				},
				{
					Type:     field.ErrorTypeDuplicate,
					Field:    "test.actions[3]",
					BadValue: v1.ActionDownload,
				},
				{
					Type:     field.ErrorTypeDuplicate,
					Field:    "test.actions[4]",
					BadValue: v1.ActionRead,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateSubjectPolicy(tt.args.policy, tt.args.fldPath, tt.args.opts)
			assert.Equal(t, tt.want, got)
		})
	}
}
