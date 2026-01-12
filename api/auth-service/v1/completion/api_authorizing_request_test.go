package completion

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
)

func newGenerateIDFuncForStatic(id string) func() string { return func() string { return id } }

func TestCompleteAPIAuthorizingRequestCreate(t *testing.T) {
	type args struct {
		req         *v1.APIAuthorizingRequest
		requesterID string
	}
	tests := []struct {
		name           string
		args           args
		generateIDFunc func() string
		want           *v1.APIAuthorizingRequest
	}{
		{
			name: "empty",
			args: args{
				req: &v1.APIAuthorizingRequest{
					Status: v1.APIAuthorizingRequestStatus{
						Phase: v1.APIAuthorizingRequestAuditing,
					},
				},
			},
			generateIDFunc: newGenerateIDFuncForStatic("00000000-0000-0000-0000-000000000000"),
			want: &v1.APIAuthorizingRequest{
				ID: "00000000-0000-0000-0000-000000000000",
				Status: v1.APIAuthorizingRequestStatus{
					Phase: v1.APIAuthorizingRequestAuditing,
				},
			},
		},
		{
			name: "already exists",
			args: args{
				req: &v1.APIAuthorizingRequest{
					ID: "11111111-1111-1111-1111-111111111111",
					Status: v1.APIAuthorizingRequestStatus{
						Phase: v1.APIAuthorizingRequestAuditing,
					},
				},
			},
			generateIDFunc: newGenerateIDFuncForStatic("00000000-0000-0000-0000-000000000000"),
			want: &v1.APIAuthorizingRequest{
				ID: "11111111-1111-1111-1111-111111111111",
				Status: v1.APIAuthorizingRequestStatus{
					Phase: v1.APIAuthorizingRequestAuditing,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateIDFunc = tt.generateIDFunc
			CompleteAPIAuthorizingRequestCreate(tt.args.req, tt.args.requesterID)
			assert.Equal(t, tt.want, tt.args.req)
		})
	}
}

func TestCompleteAPIAuthorizingRequestSpec(t *testing.T) {
	type args struct {
		spec        *v1.APIAuthorizingRequestSpec
		requesterID string
	}
	tests := []struct {
		name string
		args args
		want *v1.APIAuthorizingRequestSpec
	}{
		{
			name: "empty",
			args: args{
				spec:        &v1.APIAuthorizingRequestSpec{},
				requesterID: "00000000-0000-0000-0000-000000000000",
			},
			want: &v1.APIAuthorizingRequestSpec{
				RequesterID: "00000000-0000-0000-0000-000000000000",
			},
		},
		{
			name: "already exists",
			args: args{
				spec: &v1.APIAuthorizingRequestSpec{
					RequesterID: "11111111-1111-1111-1111-111111111111",
				},
				requesterID: "00000000-0000-0000-0000-000000000000",
			},
			want: &v1.APIAuthorizingRequestSpec{
				RequesterID: "11111111-1111-1111-1111-111111111111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CompleteAPIAuthorizingRequestSpec(tt.args.spec, tt.args.requesterID)
			assert.Equal(t, tt.want, tt.args.spec)
		})
	}
}

func TestCompleteAPIAuthorizingRequestStatus(t *testing.T) {
	type args struct {
		status *v1.APIAuthorizingRequestStatus
	}
	tests := []struct {
		name string
		args args
		want *v1.APIAuthorizingRequestStatus
	}{
		{
			name: "empty",
			args: args{
				status: &v1.APIAuthorizingRequestStatus{},
			},
			want: &v1.APIAuthorizingRequestStatus{Phase: v1.APIAuthorizingRequestAuditing},
		},
		{
			name: "already exists",
			args: args{
				status: &v1.APIAuthorizingRequestStatus{Phase: v1.APIAuthorizingRequestRejected},
			},
			want: &v1.APIAuthorizingRequestStatus{Phase: v1.APIAuthorizingRequestRejected},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CompleteAPIAuthorizingRequestStatus(tt.args.status)
			assert.Equal(t, tt.want, tt.args.status)
		})
	}
}
