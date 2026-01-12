package validation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

func TestValidateIndicatorAuthorizingRequestCreate(t *testing.T) {
	type args struct {
		req *v1.IndicatorAuthorizingRequest
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "valid",
			args: args{
				req: &v1.IndicatorAuthorizingRequest{
					Spec: v1.IndicatorAuthorizingRequestSpec{
						ID: "525948083410872251",
						Policies: []v1.SubjectPolicy{
							{
								Subject: v1.Subject{
									Type: v1.SubjectUser,
									ID:   "00000000-0000-1111-0000-000000000000",
								},
								Actions: []v1.Action{
									v1.ActionRead,
								},
							},
						},
						RequesterID: "00000000-0000-1111-1111-000000000000",
						Reason:      "REASON",
					},
				},
			},
		},
		{
			name: "invalid",
			args: args{
				req: &v1.IndicatorAuthorizingRequest{
					Spec: v1.IndicatorAuthorizingRequestSpec{
						Policies: []v1.SubjectPolicy{
							{
								Subject: v1.Subject{
									Type: v1.SubjectUser,
									ID:   "00000000-0000-1111-0000-000000000000",
								},
								Actions: []v1.Action{
									v1.ActionRead,
								},
							},
						},
						RequesterID: "00000000-0000-1111-1111-000000000000",
						Reason:      "REASON",
					},
				},
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeRequired,
					Field:    "spec.id",
					BadValue: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateIndicatorAuthorizingRequestCreate(tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validateIndicatorAuthorizingRequestSpecCreate(t *testing.T) {
	const (
		indicatorID = "525948083410872251"

		userID0 = "00000000-0000-1111-0000-000000000000"
		userID1 = "00000000-0000-1111-1111-000000000000"
		userID2 = "00000000-0000-1111-2222-000000000000"
	)
	type args struct {
		spec    *v1.IndicatorAuthorizingRequestSpec
		fldPath *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "valid",
			args: args{
				spec: &v1.IndicatorAuthorizingRequestSpec{
					ID: indicatorID,
					Policies: []v1.SubjectPolicy{
						{
							Subject: v1.Subject{
								Type: v1.SubjectUser,
								ID:   userID0,
							},
							Actions: []v1.Action{
								v1.ActionRead,
							},
						},
						{
							Subject: v1.Subject{
								Type: v1.SubjectUser,
								ID:   userID1,
							},
							Actions: []v1.Action{
								v1.ActionRead,
							},
						},
						{
							Subject: v1.Subject{
								Type: v1.SubjectUser,
								ID:   userID2,
							},
							Actions: []v1.Action{
								v1.ActionRead,
							},
						},
					},
					RequesterID: userID0,
					Reason:      strings.Repeat("正", indicatorAuthorizingRequestSpecReasonMaxLength),
				},
				fldPath: fldPathTest,
			},
		},
		{
			name: "missing policies",
			args: args{
				spec: &v1.IndicatorAuthorizingRequestSpec{
					ID:          indicatorID,
					RequesterID: userID0,
					Reason:      "REASON",
				},
				fldPath: fldPathTest,
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeRequired,
					Field:    "test.policies",
					BadValue: "",
				},
			},
		},
		{
			name: "invalid action",
			args: args{
				spec: &v1.IndicatorAuthorizingRequestSpec{
					ID: indicatorID,
					Policies: []v1.SubjectPolicy{
						{
							Subject: v1.Subject{
								Type: v1.SubjectUser,
								ID:   userID0,
							},
							Actions: []v1.Action{
								v1.ActionRead,
								v1.ActionDownload,
							},
						},
					},
					RequesterID: userID0,
					Reason:      strings.Repeat("正", indicatorAuthorizingRequestSpecReasonMaxLength),
				},
				fldPath: fldPathTest,
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeNotSupported,
					Field:    "test.policies[0].actions[1]",
					BadValue: v1.ActionDownload,
					Detail:   "test.policies[0].actions[1] 必须是 read",
				},
			},
		},
		{
			name: "invalid",
			args: args{
				spec: &v1.IndicatorAuthorizingRequestSpec{
					Policies: []v1.SubjectPolicy{
						{
							Subject: v1.Subject{
								Type: v1.SubjectRole,
								ID:   userID0,
							},
							Actions: []v1.Action{
								v1.ActionRead,
							},
						},
					},
				},
				fldPath: fldPathTest,
			},
			want: field.ErrorList{
				{
					Type:     field.ErrorTypeRequired,
					Field:    "test.id",
					BadValue: "",
				},
				{
					Type:     field.ErrorTypeNotSupported,
					Field:    "test.policies[0].subject_type",
					BadValue: v1.SubjectRole,
					Detail:   "test.policies[0].subject_type 必须是 app, user 中的一个",
				},
				{
					Type:     field.ErrorTypeInvalid,
					Field:    "test.policies[0].subject_type",
					BadValue: v1.SubjectRole,
					Detail:   "不支持授权资源 indicator 给访问者 role",
				},
				{
					Type:     field.ErrorTypeRequired,
					Field:    "test.requester_id",
					BadValue: "",
				},
				{
					Type:     field.ErrorTypeRequired,
					Field:    "test.reason",
					BadValue: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateIndicatorAuthorizingRequestSpecCreate(tt.args.spec, tt.args.fldPath)
			assert.Equal(t, tt.want, got)
		})
	}
}
