package label

import (
	"context"
)

type Driven interface {
	GetLabelByIds(ctx context.Context, ids []string) (*LabelListResp, error)
	GetRangeTypeByIds(ctx context.Context, rangeTypeKey string, ids []string) (*LabelListResp, error)
}

type IdsReq struct {
	Ids []string `json:"id" form:"id"  binding:"gte=1,lte=20,required,dive,ValidateSnowflakeID"` // ids集合，最小数组长度1，最大数组长度20
}
type LabelListResp struct {
	LabelResp []*LabelResp `json:"label_resp" binding:"required"` //标签列表
}
type LabelResp struct {
	ID   string `json:"id" form:"id" binding:"required,max=20" example:"1"`            // 唯一id
	Name string `json:"name" form:"name" binding:"required,max=30" example:"1"`        // 标签名称
	Path string `json:"path" form:"path" binding:"omitempty,max=2000" example:"a/b/c"` // 标签路径
}
type RangeTypeIdsReq struct {
	Ids          []string `json:"id" form:"id"  binding:"gte=1,lte=20,required,dive,ValidateSnowflakeID"` // ids集合，最小数组长度1，最大数组长度20
	RangeTypeKey string   `json:"range_type" form:"range_type" binding:"required,max=64" example:"3"`     // 应用范围类型值（字典项码值）
}
