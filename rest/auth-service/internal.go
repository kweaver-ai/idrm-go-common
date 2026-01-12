package auth_service

type MenuResourceActionsResp struct {
	ResourceID string   `json:"resource_id"` //菜单的key
	Actions    []string `json:"actions"`     //允许的操作
}

func (m *MenuResourceActionsResp) HasManageAction() bool {
	for _, a := range m.Actions {
		if manageActionDict[a] {
			return true
		}
	}
	return false
}
