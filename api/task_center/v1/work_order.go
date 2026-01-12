package v1

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// 工单
type WorkOrder struct {
	// ID
	ID string `json:"id,omitempty"`
	// 工单 ID，与 ID 相同
	WorkOrderID string `json:"work_order_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 工单类型：数据理解、数据归集
	Type WorkOrderType `json:"type,omitempty"`
	// 工单状态
	Status WorkOrderStatus `json:"status,omitempty"`
	// 责任人
	ResponsibleUID string `json:"responsible_uid,omitempty"`
	// 优先级
	Priority WorkOrderPriority `json:"priority,omitempty"`
	// 截止日期，空代表工单无截止日期
	FinishedAt *meta_v1.Time `json:"finished_at,omitempty"`
	// 关联数据资源目录
	CatalogIDs []string `json:"catalog_ids,omitempty"`
	// 工单说明
	Description string `json:"description,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 来源类型
	SourceType WorkOrderSourceType `json:"source_type,omitempty"`
	// 来源id
	SourceID string `json:"source_id,omitempty"`
	// 来源 id 列表。如果同时指定 SourceId 和 SourceIDs, 要求 SourceID 与
	// SourceIDs 的第一项相同
	SourceIDs []string `json:"source_ids,omitempty"`
	// 所属项目的运营流程节点 ID，仅当工单来源类型是项目时有值。
	NodeID string `json:"node_id,omitempty"`
	// 所属项目的运营流程阶段 ID，仅当工单来源类型是项目时有值。
	StageID string `json:"stage_id,omitempty"`
	// 创建时间
	CreatedAt meta_v1.TimestampUnixMilli `json:"created_at,omitempty"`
	// 归集工单关联的归集清单 ID
	DataAggregationInventoryID string `json:"data_aggregation_inventory_id,omitempty"`
	// 归集信息。借用归集列表借的结构，有机会再重构
	DataAggregationInventory *DataAggregationInventory `json:"data_aggregation_inventory,omitempty"`
}

// 工单类型
type WorkOrderType string

const (
	// 数据理解
	WorkOrderDataComprehension WorkOrderType = "data_comprehension"
	// 数据归集
	WorkOrderDataAggregation WorkOrderType = "data_aggregation"
	// 数据质量
	WorkOrderDataQuality WorkOrderType = "data_quality"
	// 数据质量检测
	WorkOrderDataQualityAudit WorkOrderType = "data_quality_audit"
	// 数据标准化
	WorkOrderDataStandardization WorkOrderType = "data_standardization"
	// 数据融合
	WorkOrderDataFusion WorkOrderType = "data_fusion"
)

// 工单状态
type WorkOrderStatus string

const (
	// 待签收
	WorkOrderStatusPendingSignature WorkOrderStatus = "pending_signature"
	// 已签收
	WorkOrderStatusSignedFor WorkOrderStatus = "signed_for"
	// 进行中
	WorkOrderStatusOngoing WorkOrderStatus = "ongoing"
	// 已完成
	WorkOrderStatusFinished WorkOrderStatus = "finished"
)

// common emergent urgent
// 工单优先级
type WorkOrderPriority string

const (
	// 普通
	WorkOrderCommon WorkOrderPriority = "common"
	// 紧急
	WorkOrderEmergent WorkOrderPriority = "emergent"
	// 非常紧急
	WorkOrderUrgent WorkOrderPriority = "urgent"
)

// 工单来源类型
type WorkOrderSourceType string

const (
	// 无来源、独立
	WorkOrderStandalone WorkOrderSourceType = "standalone"
	// 计划
	WorkOrderPlan WorkOrderSourceType = "plan"
	// 业务表
	WorkOrderBusinessForm WorkOrderSourceType = "business_form"
	// 数据分析
	WorkOrderDataAnalysis WorkOrderSourceType = "data_analysis"
	// 逻辑视图
	WorkOrderFormView WorkOrderSourceType = "form_view"
	// 归集工单
	WorkOrderAggregationWorkOrder WorkOrderSourceType = "aggregation_work_order"
	// 供需申请
	WorkOrderSupplyAndDemand WorkOrderSourceType = "supply_and_demand"
	// 项目
	WorkOrderProject WorkOrderSourceType = "project"
)
