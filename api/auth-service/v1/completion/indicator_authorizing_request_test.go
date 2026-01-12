package completion

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
)

func TestCompleteIndicatorAuthorizingRequestSpec(t *testing.T) {
	const (
		userID0 = "00000000-0000-0000-0000-000000000000"
		userID1 = "00000000-0000-0000-1111-000000000000"
	)
	type args struct {
		spec        *v1.IndicatorAuthorizingRequestSpec
		requesterID string
	}
	tests := []struct {
		name string
		args args
		want *v1.IndicatorAuthorizingRequestSpec
	}{
		{
			name: "补全 RequesterID",
			args: args{
				spec:        &v1.IndicatorAuthorizingRequestSpec{},
				requesterID: userID0,
			},
			want: &v1.IndicatorAuthorizingRequestSpec{RequesterID: userID0},
		},
		{
			name: "不会修改已指定的 RequesterID",
			args: args{
				spec:        &v1.IndicatorAuthorizingRequestSpec{RequesterID: userID0},
				requesterID: userID1,
			},
			want: &v1.IndicatorAuthorizingRequestSpec{RequesterID: userID0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CompleteIndicatorAuthorizingRequestSpec(tt.args.spec, tt.args.requesterID)
			assert.Equal(t, tt.want, tt.args.spec)
		})
	}
}
