package v1

// 查询目录共享申请状态的请求
//
//	Repository: demand-management
//	Commit:     5016e6deca3f0362c0c631576251c9cd8e0ff044
//	File:       domain/v1/shared_declaration/interface.go
//	Line:       128
type SharedDeclarationStatusReq struct {
	// 目录 id
	CatalogIDs []string `json:"catalog_ids"`
}

// 查询目录共享申请状态的响应
//
//	Repository: demand-management
//	Commit:     5016e6deca3f0362c0c631576251c9cd8e0ff044
//	File:       domain/v1/shared_declaration/interface.go
//	Line:       132
type SharedDeclarationStatusResp struct {
	// 目录id
	CatalogID string `json:"catalog_id"`
	// 共享申请状态
	Status SharedDeclarationStatus `json:"status"`
}

// 共享申请状态
type SharedDeclarationStatus string

// 共享申请状态的枚举值
//
//	Repository: demand-management
//	Commit:     5016e6deca3f0362c0c631576251c9cd8e0ff044
//	File:       domain/v1/shared_declaration/interface.go
//	Line:       134
const (
	// 未申请
	SharedDeclarationNotApplied SharedDeclarationStatus = "not_applied"
	// 已添加到待申请列表
	SharedDeclarationAdded SharedDeclarationStatus = "added"
	// 已申请
	SharedDeclarationApplied SharedDeclarationStatus = "applied"
)
