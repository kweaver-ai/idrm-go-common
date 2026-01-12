package completion

import v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"

// CompleteAPIAuthorizingRequestCreate 用于
func CompleteAPIAuthorizingRequestCreate(req *v1.APIAuthorizingRequest, requesterID string) {
	// 补全 ID
	if req.ID == "" {
		req.ID = generateIDFunc()
	}
	// 补全 Spec
	CompleteAPIAuthorizingRequestSpec(&req.Spec, requesterID)
	// 补全 Status
	CompleteAPIAuthorizingRequestStatus(&req.Status)
}

// CompleteAPIAuthorizingRequestSpec 补全 APIAuthorizingRequestSpec
func CompleteAPIAuthorizingRequestSpec(spec *v1.APIAuthorizingRequestSpec, requesterID string) {
	// 补全 RequesterID
	if spec.RequesterID == "" && requesterID != "" {
		spec.RequesterID = requesterID
	}
}

// CompleteAPIAuthorizingRequestStatus 补全 APIAuthorizingRequestStatus
func CompleteAPIAuthorizingRequestStatus(status *v1.APIAuthorizingRequestStatus) {
	if status.Phase == "" {
		status.Phase = v1.APIAuthorizingRequestAuditing
	}
}
