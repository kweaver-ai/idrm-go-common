package data_view

// region 视图探查结果

type ExploreReportResp struct {
	Code                   string                `json:"code" `                    // 数据探查报告编号
	TaskId                 string                `json:"task_id" `                 // 任务ID
	Version                int32                 `json:"version" `                 // 任务版本
	ExploreTime            int64                 `json:"explore_time,omitempty"`   // 探查时间
	Overview               *ReportOverview       `json:"overview,omitempty"`       // 总览信息
	TotalSample            int64                 `json:"total_sample"`             // 采样条数
	ExploreMetadataDetails *ExploreDetails       `json:"explore_metadata_details"` // 元数据级探查结果详情
	ExploreFieldDetails    []*ExploreFieldDetail `json:"explore_field_details"`    // 字段级探查结果详情
	ExploreRowDetails      *ExploreDetails       `json:"explore_row_details"`      // 行级探查结果详情
	ExploreViewDetails     []*RuleResult         `json:"explore_view_details"`     // 视图级探查结果详情
}

type ReportOverview struct {
	ScoreTrends []*ScoreTrend      `json:"score_trend,omitempty"` // 六性评分历史趋势数据
	Fields      *ExploreFieldsInfo `json:"fields,omitempty"`      // 表字段信息
	DimensionScores
}

type ExploreDetails struct {
	ExploreDetails []*RuleResult `json:"explore_details"` // 探查结果详情
	DimensionScores
}

type ScoreTrend struct {
	TaskId               string   `json:"task_id" `              // 任务ID
	Version              int      `json:"version" `              // 任务版本
	ExploreTime          int64    `json:"explore_time"`          // 探查时间
	CompletenessScore    *float64 `json:"completeness_score"`    // 完整性维度评分，缺省为NULL
	UniquenessScore      *float64 `json:"uniqueness_score"`      // 唯一性维度评分，缺省为NULL
	StandardizationScore *float64 `json:"standardization_score"` // 规范性维度评分，缺省为NULL
	AccuracyScore        *float64 `json:"accuracy_score"`        // 准确性维度评分，缺省为NULL
	ConsistencyScore     *float64 `json:"consistency_score"`     // 一致性维度评分，缺省为NULL
}

type ExploreFieldsInfo struct {
	TotalCount   int `json:"total_count"`   // 总字段数
	ExploreCount int `json:"explore_count"` // 探查字段数
}

type RuleResult struct {
	RuleId          string `json:"rule_id"`          // 规则ID
	RuleName        string `json:"rule_name"`        // 规则名称
	RuleDescription string `json:"rule_description"` // 规则描述
	Dimension       string `json:"dimension"`        // 维度属性 0准确性,1及时性,2完整性,3唯一性，4一致性,5规范性
	Result          string `json:"result"`           // 规则输出结果 []any规则输出列级结果
	InspectedCount  int64  `json:"inspected_count"`  // 检测数据量
	IssueCount      int64  `json:"issue_count"`      // 问题数据量
	DimensionScores
}

type DimensionScores struct {
	CompletenessScore    *float64 `json:"completeness_score"`    // 完整性维度评分，缺省为NULL
	UniquenessScore      *float64 `json:"uniqueness_score"`      // 唯一性维度评分，缺省为NULL
	StandardizationScore *float64 `json:"standardization_score"` // 规范性维度评分，缺省为NULL
	AccuracyScore        *float64 `json:"accuracy_score"`        // 准确性维度评分，缺省为NULL
	ConsistencyScore     *float64 `json:"consistency_score"`     // 一致性维度评分，缺省为NULL
}

type RuleConfig struct {
	UpdatePeriod *string `json:"update_period" form:"update_period" binding:"omitempty,oneof=day week month quarter half_a_year year"`
}

type GetBusinessUpdateTimeResp struct {
	FieldID            string `json:"field_id"`             // 业务更新字段id
	FieldBusinessName  string `json:"field_business_name"`  // 业务更新字业务名称
	BusinessUpdateTime string `json:"business_update_time"` // 业务更新时间
}

type ViewExploreDetail struct {
	UniquenessScore      *float64 `json:"uniqueness_score"`      // 唯一性分数
	CompletenessScore    *float64 `json:"completeness_score"`    // 完整性分数
	AccuracyScore        *float64 `json:"accuracy_score"`        // 准确性分数
	StandardizationScore *float64 `json:"standardization_score"` // 规范性维度评分，缺省为NULL
	ConsistencyScore     *float64 `json:"consistency_score"`     // 一致性分数
	TimelinessScore      *float64 `json:"timeliness_score"`      // 及时性分数
}

type ExploreFieldDetail struct {
	FieldId  string        `json:"field_id"`  // 字段id
	CodeInfo string        `json:"code_info"` // 码表信息
	Details  []*RuleResult `json:"details"`   // 规则结果明细（仅返回部分需要呈现的字段规则输出结果）
	DimensionScores
}

type GetRuleListReq struct {
	FormViewId string `json:"form_view_id" form:"form_view_id" binding:"omitempty,uuid" example:"13b8a80b-1914-4896-99d8-51559dba26c4"`                                     // 视图id
	RuleLevel  string `json:"rule_level" form:"rule_level" binding:"omitempty,oneof=metadata field row view"`                                                               // 规则级别，元数据级 字段级 行级 视图级
	Dimension  string `json:"dimension" form:"dimension" binding:"omitempty,oneof=completeness standardization uniqueness accuracy consistency timeliness data_statistics"` // 维度，完整性 规范性 唯一性 准确性 一致性 及时性 数据统计
	FieldId    string `json:"field_id" form:"field_id" binding:"omitempty,uuid" example:"962b749f-32e6-41d1-bd79-33ce839a8598"`                                             // 字段id
	Keyword    string `json:"keyword" form:"keyword" binding:"KeywordTrimSpace,omitempty,min=1,max=128"`                                                                    // 关键字查询
	Enable     bool   `json:"enable" form:"enable" binding:"omitempty"`                                                                                                     // 启用状态，true为已启用，false为未启用，不传该参数则不跟据启用状态筛选
}

type GetRuleResp struct {
	RuleId          string  `json:"rule_id"`          // 规则id
	RuleName        string  `json:"rule_name"`        // 规则名称
	RuleDescription string  `json:"rule_description"` // 规则描述
	RuleLevel       string  `json:"rule_level"`       // 规则级别，元数据级 字段级 行级 视图级
	FieldId         string  `json:"field_id"`         // 字段id
	Dimension       string  `json:"dimension"`        // 维度，完整性 规范性 唯一性 准确性 一致性 及时性
	DimensionType   string  `json:"dimension_type"`   // 维度类型
	RuleConfig      *string `json:"rule_config"`      // 规则配置
	Enable          *bool   `json:"enable"`           // 是否启用
	TemplateId      string  `json:"template_id"`      // 模板id
}
