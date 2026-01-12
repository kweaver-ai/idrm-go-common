package demand_management

import "context"

type Driven interface {
	GetShareApply(ctx context.Context, req *GetShareApplyReq) (*GetShareApplyResp, error)
	GetUserShareApplyResource(ctx context.Context, req *UserShareApplyResourceReq) (*UserShareApplyResourceResp, error)
	// 通过分析场景产物ID查询产物及需求名称
	GetNameByAnalOutputItemID(ctx context.Context, analOutputItemID string) (*NameGetResp, error)
}

type GetShareApplyReq struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	IsAll     bool   `json:"is_all"`
}

type GetShareApplyResp struct {
	Entries []struct {
		Id           string `json:"id"`
		Code         string `json:"code"`
		Name         string `json:"name"`
		Status       string `json:"status"`
		AuditStatus  string `json:"audit_status"`
		RejectReason string `json:"reject_reason"`
		ApplyOrgCode string `json:"apply_org_code"`
		ApplyOrgName string `json:"apply_org_name"`
		ApplyOrgPath string `json:"apply_org_path"`
		Applier      string `json:"applier"`
		Phone        string `json:"phone"`
		ViewNum      int    `json:"view_num"`
		ApiNum       int    `json:"api_num"`
		CreatedAt    int64  `json:"created_at"`
		CreatedBy    string `json:"created_by"`
		FinishDate   int64  `json:"finish_date"`
	} `json:"entries"`
	TotalCount int `json:"total_count"`
}

type UserShareApplyResourceReq struct {
	Type          int    `json:"type"` // 1 目录 2 接口服务 0 全部
	UserId        string `json:"user_id"`
	DataCatalogId string `form:"data_catalog_id" binding:"omitempty"`
}

type UserShareApplyResourceResp struct {
	Entries []UserShareApplyResourceItem `json:"entries"`
}

type UserShareApplyResourceItem struct {
	ResId   string `json:"res_id"`   // 资源id
	ResType int    `json:"res_type"` // 资源类型
}

type NameGetResp struct {
	AnalOutputItemID   string `json:"anal_output_item_id"`   // 分析场景产物ID
	AnalOutputItemName string `json:"anal_output_item_name"` // 分析场景产物名称
	DataAnalReqID      string `json:"data_anal_req_id"`      // 数据分析需求ID
	DataAnalReqName    string `json:"data_anal_req_name"`    // 数据分析需求名称
}
