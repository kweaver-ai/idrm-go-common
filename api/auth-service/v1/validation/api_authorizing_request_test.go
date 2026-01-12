package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
	"github.com/kweaver-ai/idrm-go-common/util/validation/field"
)

func TestValidateAPIAuthorizingRequestCreate(t *testing.T) {
	req := &v1.APIAuthorizingRequest{
		Spec: v1.APIAuthorizingRequestSpec{
			ID: "00000000-0000-0000-0000-000000000000",
			Policies: []v1.SubjectPolicy{
				{
					Subject: v1.Subject{
						Type: v1.SubjectAPP,
						ID:   "00000000-0000-0000-1111-000000000000",
					},
					Actions: []v1.Action{
						v1.ActionRead,
					},
				},
				{
					Subject: v1.Subject{
						Type: v1.SubjectAPP,
						ID:   "00000000-0000-0000-1111-111111111111",
					},
					Actions: []v1.Action{
						v1.ActionRead,
					},
				},
				{
					Subject: v1.Subject{
						Type: v1.SubjectAPP,
						ID:   "00000000-0000-0000-1111-222222222222",
					},
					Actions: []v1.Action{
						v1.ActionRead,
					},
				},
			},
			RequesterID: "00000000-0000-0000-2222-000000000000",
			Reason:      "REASON",
		},
	}
	got := ValidateAPIAuthorizingRequestCreate(req)
	assert.Empty(t, got)
}

func Test_validateAPIAuthorizingRequestSpecCreate(t *testing.T) {
	// testdata
	var (
		apiID = "00000000-0000-0000-0000-000000000000"

		appID0 = "00000000-0000-0000-1111-000000000000"
		appID1 = "00000000-0000-0000-1111-111111111111"
		appID2 = "00000000-0000-0000-1111-222222222222"

		subjectAPP0 = v1.Subject{Type: v1.SubjectAPP, ID: appID0}
		subjectAPP1 = v1.Subject{Type: v1.SubjectAPP, ID: appID1}
		subjectAPP2 = v1.Subject{Type: v1.SubjectAPP, ID: appID2}

		subjectPolicyAPP0Read = v1.SubjectPolicy{Subject: subjectAPP0, Actions: []v1.Action{v1.ActionRead}}
		subjectPolicyAPP1Read = v1.SubjectPolicy{Subject: subjectAPP1, Actions: []v1.Action{v1.ActionRead}}
		subjectPolicyAPP2Read = v1.SubjectPolicy{Subject: subjectAPP2, Actions: []v1.Action{v1.ActionRead}}

		policies = []v1.SubjectPolicy{subjectPolicyAPP0Read, subjectPolicyAPP1Read, subjectPolicyAPP2Read}

		requesterID = "00000000-0000-0000-2222-000000000000"

		reason = "REASON"

		spec = v1.APIAuthorizingRequestSpec{ID: apiID, Policies: policies, RequesterID: requesterID, Reason: reason}
	)
	type args struct {
		spec    *v1.APIAuthorizingRequestSpec
		fldPath *field.Path
	}
	tests := []struct {
		name string
		args args
		want field.ErrorList
	}{
		{
			name: "ok",
			args: args{spec: &spec, fldPath: field.NewPath("ok")},
		},
		{
			name: "invalid api id",
			args: args{
				spec: &v1.APIAuthorizingRequestSpec{
					ID:          "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					Policies:    policies,
					RequesterID: requesterID,
					Reason:      reason,
				},
				fldPath: field.NewPath("invalid_api_id"),
			},
			want: field.ErrorList{field.Invalid(field.NewPath("invalid_api_id").Child("id"), "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "invalid_api_id.id 必须是一个有效的 uuid")},
		},
		{
			name: "empty policies",
			args: args{
				spec:    &v1.APIAuthorizingRequestSpec{ID: apiID, RequesterID: requesterID, Reason: reason},
				fldPath: field.NewPath("empty_policies"),
			},
			want: field.ErrorList{field.Required(field.NewPath("empty_policies").Child("policies"), "")},
		},
		{
			name: "unsupported subject user",
			args: args{
				spec: &v1.APIAuthorizingRequestSpec{
					ID: apiID,
					Policies: []v1.SubjectPolicy{
						subjectPolicyAPP0Read,
						subjectPolicyAPP1Read,
						{Subject: v1.Subject{Type: v1.SubjectUser, ID: "33333333-3333-3333-3333-333333333333"}, Actions: []v1.Action{v1.ActionRead}},
					},
					RequesterID: requesterID,
					Reason:      reason,
				},
				fldPath: field.NewPath("unsupported_subject_user"),
			},
			want: field.ErrorList{field.Invalid(field.NewPath("unsupported_subject_user").Child("policies").Index(2).Child("subject_type"), v1.SubjectUser, "不支持授权资源 api 给访问者 user")},
		},
		{
			name: "unsupported action download",
			args: args{
				spec: &v1.APIAuthorizingRequestSpec{
					ID: apiID,
					Policies: []v1.SubjectPolicy{
						subjectPolicyAPP0Read,
						subjectPolicyAPP1Read,
						{Subject: subjectAPP2, Actions: []v1.Action{v1.ActionRead, v1.ActionDownload}},
					},
					RequesterID: requesterID,
					Reason:      reason,
				},
				fldPath: field.NewPath("unsupported_action_download"),
			},
			want: field.ErrorList{field.NotSupported(field.NewPath("unsupported_action_download").Child("policies").Index(2).Child("actions").Index(1), v1.ActionDownload, sets.List(supportedAPIActions))},
		},
		{
			name: "empty requester id",
			args: args{
				spec:    &v1.APIAuthorizingRequestSpec{ID: apiID, Policies: policies, Reason: reason},
				fldPath: field.NewPath("empty_requester_id"),
			},
			want: field.ErrorList{field.Required(field.NewPath("empty_requester_id").Child("requester_id"), "")},
		},
		{
			name: "empty reason",
			args: args{
				spec:    &v1.APIAuthorizingRequestSpec{ID: apiID, Policies: policies, RequesterID: requesterID},
				fldPath: field.NewPath("empty_reason"),
			},
			want: field.ErrorList{field.Required(field.NewPath("empty_reason").Child("reason"), "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateAPIAuthorizingRequestSpecCreate(tt.args.spec, tt.args.fldPath)
			assert.Equal(t, tt.want, got)
		})
	}
}
