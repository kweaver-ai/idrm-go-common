package v1

// APIAuthorizingRequest 定义接口授权申请
type APIAuthorizingRequest struct {
	// 接口授权申请的 ID，由服务端设置，不可修改
	ID string `json:"id,omitempty"`
	// 接口授权申请所期望的指标、用户、权限等
	Spec APIAuthorizingRequestSpec `json:"spec,omitempty"`
	// 接口授权申请的状态，处于当前状态的原因等，由服务端维护和更新，客户端无法
	// 修改
	Status APIAuthorizingRequestStatus `json:"status,omitempty"`
}

// APIAuthorizingRequestWithReferences 定义接口授权申请及其引用的资源
type APIAuthorizingRequestWithReferences struct {
	APIAuthorizingRequest
	// 接口授权申请所引用的资源列表
	References []ReferenceSource `json:"references,omitempty"`
}

// APIAuthorizingRequestSpec 描述了接口授权申请是为哪些用户申请哪个接口，执行哪
// 些动作的授权。
type APIAuthorizingRequestSpec struct {
	// 接口的 ID，格式：UUID
	ID string `json:"id,omitempty"`
	// 用户权限策略列表，定义对于这个接口，为哪些用户，申请哪些动作的权限
	Policies []SubjectPolicy `json:"policies,omitempty"`
	// 申请者的 ID，格式：UUID，由服务端设置
	RequesterID string `json:"requester_id,omitempty"`
	// 申请理由，申请授权必须有理由
	Reason string `json:"reason,omitempty"`
}

// APIAuthorizingRequestStatus 描述了接口授权申请当前的状态，处于当前
// 状态的原因等。
type APIAuthorizingRequestStatus struct {
	// 接口授权申请在其生命周期中所处的阶段
	Phase APIAuthorizingRequestPhase `json:"phase,omitempty"`
	// 处于当前阶段的原因
	Message string `json:"message,omitempty"`
}

// APIAuthorizingRequestPhase 定义 APIAuthorizingRequest 在其生命周
// 期中所处的阶段
type APIAuthorizingRequestPhase string

const (
	// 审批中
	APIAuthorizingRequestAuditing APIAuthorizingRequestPhase = "Auditing"
	// 申请被拒绝
	APIAuthorizingRequestRejected APIAuthorizingRequestPhase = "Rejected"
	// 申请被允许
	APIAuthorizingRequestApproved APIAuthorizingRequestPhase = "Approved"
	// 申请被发起者撤回
	APIAuthorizingRequestUndone APIAuthorizingRequestPhase = "Undone"
	// 失败。创建资源失败，或应用权限策略失败
	APIAuthorizingRequestFailed APIAuthorizingRequestPhase = "Failed"
	// 完成
	APIAuthorizingRequestCompleted APIAuthorizingRequestPhase = "Completed"
)
