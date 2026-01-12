package v1

// PolicyEffect 定义 Policy 的效果
type PolicyEffect string

func (p PolicyEffect) Str() string {
	return string(p)
}

const (
	// 允许
	PolicyAllow PolicyEffect = "allow"
	// 拒绝
	PolicyDeny PolicyEffect = "deny"
)
