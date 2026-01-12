package configuration_center

import "github.com/kweaver-ai/idrm-go-frame/core/models"

//region GetDictItemType

type GetDictItemTypeReq struct {
	DictType  []string `query:"dict_type"`
	QueryType string   `query:"query_type"`
}

type CheckDictTypeKeyReq struct {
	DictTypeKey []*DictTypeKey `json:"dict_type_key" binding:"required,dive"`
}

type DictTypeKey struct {
	DictType string `json:"dict_type" form:"dict_type" binding:"required"`
	DictKey  string `json:"dict_key" form:"dict_key" binding:"required"`
	Field    string `json:"field" form:"field" binding:"required"`
}

type GetDictItemTypeResp struct {
	DictSlice []*DictItemType `json:"dicts"`
}

type DictItemType struct {
	DictType     string              `json:"dict_type"`
	DictItemResp []*DictItemRespUnit `json:"dict_item_resp"`
}

type DictItemRespUnit struct {
	Id          string `json:"id"`
	DictKey     string `json:"dict_key"`
	DictValue   string `json:"dict_value"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

//region GetDictItemPage

type GetDictItemPageReq struct {
	Offset    int    `json:"offset" form:"offset,default=1" binding:"omitempty,min=1" default:"1"`                      // 页码，默认1
	Limit     int    `json:"limit" form:"limit,default=20" binding:"omitempty,min=1,max=2000" default:"10"`             // 每页大小，默认10
	Direction string `json:"direction" form:"direction,default=desc" binding:"omitempty,oneof=asc desc" default:"desc"` // 排序方向，枚举：asc：正序；desc：倒序。默认倒序
	Sort      string `json:"sort" form:"sort,default=id" binding:"omitempty,oneof=updated_at name id" default:"id"`     // 排序类型，枚举：created_at：按创建时间排序；updated_at：按更新时间排序。默认按创建时间排序
	Name      string `json:"name" form:"name" binding:"VerifyXssString,omitempty,max=128"`                              // 关键字查询，字符无限制
	DictId    string `json:"dict_id" form:"dict_id" binding:"required,ValidateSnowflakeID"`                             //字典ID
}

type GetDictItemPageRes struct {
	DictItemResp []*DictItemRespUnit `json:"entries" binding:"required"`
	TotalCount   int64               `json:"total_count" binding:"required,ge=0" example:"3"`
}

//endregion

//region GetGradeLabel

type GetGradeLabelReq struct {
	Keyword     string `json:"keyword" form:"keyword" binding:"TrimSpace,omitempty,min=1,max=128"` // 关键字查询，字符无限制
	IsShowLabel bool   `json:"is_show_label" form:"is_show_label,default=true" default:"true"`     // 是否展示标签
}

type GetGradeLabelRes struct {
	GradeLabel []*GradeLabel `json:"entries" binding:"required"`
	TotalCount int64         `json:"total_count" binding:"required,ge=0" example:"3"`
}
type GradeLabel struct {
	ID                  string         `json:"id" binding:"required" example:"1"`                                      // 对象ID
	Name                string         `json:"name" binding:"required,min=1,max=128" example:"catalog_class_name"`     // 目录类别名称
	ParentID            models.ModelID `json:"parent_id,omitempty" binding:"required,VerifyModelID" example:"0"`       // 目录类别父节点ID
	Description         string         `json:"description"`                                                            // 描述
	SortWeight          uint64         `json:"sort_weight"`                                                            // 排序权重
	NodeType            int            `json:"node_type"`                                                              // 节点类型
	Icon                string         `json:"icon"`                                                                   // 目录类别描述
	SensitiveAttri      *string        `json:"sensitive_attri"`                                                        // 敏感属性预设
	SecretAttri         *string        `json:"secret_attri"`                                                           // 涉密属性预设
	ShareCondition      *string        `json:"share_condition"`                                                        // 共享条件：不共享，有条件共享，无条件共享
	DataProtectionQuery bool           `json:"data_protection_query"`                                                  // 数据保护查询开关
	CreatedAt           int64          `json:"created_at" binding:"required,gt=0"`                                     // 创建时间，时间戳
	UpdatedAt           int64          `json:"updated_at" binding:"required,gt=0"`                                     // 最终修改时间，时间戳
	RawName             string         `json:"raw_name" binding:"required,min=1,max=128" example:"catalog_class_name"` // 目录类别原始的名称
	Children            []*GradeLabel  `json:"children,omitempty" binding:"omitempty"`                                 // 当前TreeNode的子Node列表
}

//endregion
