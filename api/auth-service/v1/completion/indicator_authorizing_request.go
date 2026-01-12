package completion

import v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"

// CompleteIndicatorAuthorizingRequestSpec 补全 IndicatorAuthorizingRequestSpec
func CompleteIndicatorAuthorizingRequestSpec(spec *v1.IndicatorAuthorizingRequestSpec, requesterID string) {
	// 补全 RequesterID
	if spec.RequesterID == "" && requesterID != "" {
		spec.RequesterID = requesterID
	}
}
