package data_lineage

import "github.com/kweaver-ai/idrm-go-frame/core/enum"

type LineageEntityType enum.Object

var (
	LineageEntityTypeColumn    = enum.New[LineageEntityType](1, "column", "字段")
	LineageEntityTypeIndicator = enum.New[LineageEntityType](2, "indicator", "指标")
	LineageEntityTypeTable     = enum.New[LineageEntityType](3, "table", "数据表")
)

type LineageDirection enum.Object

var (
	LineageDirectionForward   = enum.New[LineageDirection](1, "forward", "正向")
	LineageDirectionReversely = enum.New[LineageDirection](2, "reversely", "反向")
)

type LineageNodeType enum.Object

var (
	LineageNodeTypeDataTable     = enum.New[LineageNodeType](1, "data_table", "数据表")
	LineageNodeTypeFormView      = enum.New[LineageNodeType](2, "form_view", "元数据视图")
	LineageNodeTypeLogicView     = enum.New[LineageNodeType](3, "logic_view", "逻辑实体视图")
	LineageNodeTypeCustomView    = enum.New[LineageNodeType](4, "custom_view", "自定义视图")
	LineageNodeTypeIndicator     = enum.New[LineageNodeType](5, "indicator", "指标")
	LineageNodeTypeColumn        = enum.New[LineageNodeType](6, "column", "字段")
	LineageNodeTypeExternalTable = enum.New[LineageNodeType](7, "external_table", "第三方表")
)
