package v1

type Permission struct {
	Action Action       `json:"action,omitempty"`
	Effect PolicyEffect `json:"effect,omitempty"`
}
