package data_view

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/kweaver-ai/idrm-go-common/rest/virtual_engine"
)

type Driven interface {
	DeleteRelated(ctx context.Context, req *DeleteRelatedReq) error
	QueryViewCount(ctx context.Context, flag string, isOperator bool, id ...string) (*QueryViewDetailBySubjectIDResp, error)
	QueryViewCountInMap(ctx context.Context, flag string, isOperator bool, id ...string) (map[string]int64, error)
	QueryViewFieldInfo(ctx context.Context, isOperator bool, id ...string) ([]*SubjectFormViewInfo, error)
	BatchQueryViewFieldInfo(ctx context.Context, ids ...string) ([]*GetViewFieldsResp, error)
	GetDataViewDetails(ctx context.Context, id string) (*GetFormViewDetailsRes, error)
	GetDataViewField(ctx context.Context, id string) (*GetFieldsRes, error)
	GetDataViewFieldByInternal(ctx context.Context, id string) (*GetFieldsRes, error)
	GetSubViewIDs(ctx context.Context, opts *GetSubViewIDsOptions) ([]string, error)                                 // 获取子视图（行列规则）ID 列表
	GetSubViewLogicViewID(ctx context.Context, id string) (string, error)                                            // 获取子视图（行列规则）所属逻辑视图的 ID
	GeSubViewByViews(ctx context.Context, ids []string) (map[string][]string, error)                                 // 批量获取视图的子视图
	GetSubView(ctx context.Context, id string) (*SubView, error)                                                     // 查询子视图
	GetLogicViewReportInfo(ctx context.Context, req *GetLogicViewReportInfoBody) (*GetLogicViewReportInfoRes, error) // 获取逻辑视图上报信息
	GetDataViewBasic(ctx context.Context, ids []string) ([]*ViewBasicInfo, error)
	GetDataPreview(ctx context.Context, req *DataPreviewReq) (*DataPreviewResp, error)
	GetViewByName(ctx context.Context, name string, datasourceID string) (*GetViewFieldsResp, error)
	GetViewInfo(ctx context.Context, req *GetViewInfoReq) (*GetViewInfoResp, error)
	GetSyntheticDataCatalog(ctx context.Context, id string) (*virtual_engine.FetchDataRes, error)
	GetSampleData(ctx context.Context, id string) (*GetSampleDataRes, error)
	FormViewExplore
	InternalFormView
	GetTableCount(ctx context.Context, departmentID string) (int64, error)
	GetByAuditStatus(ctx context.Context, req *GetByAuditStatusReq) (*GetByAuditStatusResp, error)
	CreateWorkOrderTask(ctx context.Context, req *CreateWorkOrderTaskReq) (*CreateWorkOrderTaskResp, error)
	CreateExploreTask(ctx context.Context, req *CreateExploreTaskReq) (*CreateExploreTaskResp, error)
}

// FormViewExplore 原数据视图探查相关
type FormViewExplore interface {
	GetExploreReport(ctx context.Context, id string, thirdParty bool) (*ExploreReportResp, error)
	GetViewBusinessUpdateTime(ctx context.Context, id string) (*GetBusinessUpdateTimeResp, error)
	QueryViewExplore(ctx context.Context, formViewID string, thirdParty bool) (*ViewExploreDetail, error)
	GetExploreRule(ctx context.Context, req *GetRuleListReq) ([]*GetRuleResp, error)
	BatchGetExploreReport(ctx context.Context, req *BatchGetExploreReportReq) (*BatchGetExploreReportResp, error)
	GetExploreTaskList(ctx context.Context, workOrderId string) (*ListExploreTaskResp, error)
	GetViewExploreReport(ctx context.Context, id string, version *int32) (*ExploreReportResp, error)
}

type InternalFormView interface {
	GetViewBasicInfoByName(ctx context.Context, req *GetViewListByTechnicalNameInMultiDatasourceReq) (*GetViewListByTechnicalNameInMultiDatasourceRes, error)
	GetViewField(ctx context.Context, id string) (*GetFieldsRes, error)
	UserViewAuth(ctx context.Context, userID string, viewID ...string) ([]string, error)
}

// region DeleteRelated

type DeleteRelatedReq struct {
	SubjectDomainIDs []string      `json:"subject_domain_ids"`
	LogicEntityIDs   []string      `json:"logic_entity_ids"`
	MoveDeletes      []*MoveDelete `json:"move_deletes"`
}

type MoveDelete struct {
	SubjectDomainID string `json:"subject_domain_id"`
	LogicEntityID   string `json:"logic_entity_id"`
}

func (d *DeleteRelatedReq) Empty() bool {
	return len(d.LogicEntityIDs) <= 0 && len(d.SubjectDomainIDs) <= 0 && len(d.MoveDeletes) <= 0
}

func (d *DeleteRelatedReq) Reader() io.Reader {
	bts, _ := json.Marshal(d)
	return bytes.NewReader(bts)
}

//endregion

//region QueryViewDetailBySubjectIDReq

const (
	QueryFlagAll   = "all"
	QueryFlagCount = "count"
	QueryFlagTotal = "total"
)

type QueryViewDetailBySubjectIDReq struct {
	Flag       string   `json:"flag"`        //如果是all, 返回所有的数量；如果是count, 返回下面数组的数量,  如果是total ，只返回总的数量即可
	IsOperator bool     `json:"is_operator"` //如果为true，表示该用户是数据运营角色或者数据开发角色，这时展示所有的视图数据
	ID         []string `json:"id"`          //业务域，业务对象ID
}

type QueryViewDetailBySubjectIDResp struct {
	Total       int64                `json:"total"`
	RelationNum []DomainViewRelation `json:"relation_num"`
}

type DomainViewRelation struct {
	SubjectDomainID string `json:"subject_domain_id"` //业务域，业务对象ID
	RelationViewNum int64  `json:"relation_view_num"`
}

//endregion

// region QueryViewFieldInfo

type GetRelatedFieldInfoReq struct {
	IsOperator bool   `query:"is_operator"` //如果为true，表示该用户是数据运营角色或者数据开发角色，这时展示所有的视图数据
	IDs        string `query:"ids"`
}

type GetRelatedFieldInfoResp struct {
	Data []*SubjectFormViewInfo `json:"data"`
}

type SubjectFormViewInfo struct {
	FormViewID    string              `json:"form_view_id"`
	CatalogName   string              `json:"catalog_name"`
	Schema        string              `json:"schema"`
	BusinessName  string              `json:"business_name"`
	TechnicalName string              `json:"technical_name"`
	Fields        []*SubjectViewField `json:"fields"`
}

type SubjectViewField struct {
	ID            string       `json:"id"`
	BusinessName  string       `json:"business_name"`
	TechnicalName string       `json:"technical_name"`
	DataType      string       `json:"data_type"`
	Property      *SubjectProp `json:"property"`
	SubjectID     string       `json:"subject_id"`
	IsPrimary     bool         `json:"is_primary"`
}

type SubjectProp struct {
	ID       string `json:"id"`        //属性ID
	Name     string `json:"name"`      //属性的名称
	PathID   string `json:"path_id"`   //ID的路径
	PathName string `json:"path_name"` //属性的名称路径
}

type GetViewFieldsResp struct {
	FormViewID    string             `json:"form_view_id"`
	BusinessName  string             `json:"business_name"`
	TechnicalName string             `json:"technical_name"`
	Fields        []*SimpleViewField `json:"fields"`
}

type SimpleViewField struct {
	ID               string `json:"id"`                 // 视图ID
	BusinessName     string `json:"business_name"`      // 业务名称
	TechnicalName    string `json:"technical_name"`     // 技术名称
	PrimaryKey       bool   `json:"primary_key"`        // 是否主键
	Comment          string `json:"comment"`            // 列注释
	DataType         string `json:"data_type"`          // 数据类型
	DataLength       int32  `json:"data_length"`        // 数据长度
	OriginalDataType string `json:"original_data_type"` // 原始数据类型
	DataAccuracy     int32  `json:"data_accuracy"`      // 数据精度（仅DECIMAL类型）
	IsNullable       string `json:"is_nullable"`        // 是否为空 (YES/NO)
	StandardCode     string `json:"standard_code"`      // 数据标准code
	StandardName     string `json:"standard_name"`      // 数据标准名称
	CodeTableID      string `json:"code_table_id"`      // 码表ID
	Index            int    `json:"index"`              // 字段顺序
}

func (s *SimpleViewField) IsPrimaryKey() int32 {
	if s.PrimaryKey {
		return 1
	}
	return 0
}

func (s *SimpleViewField) IsRequired() int32 {
	if s.IsNullable == "YES" {
		return 0
	}
	return 1
}

type GetFormViewDetailsRes struct {
	TechnicalName         string          `json:"technical_name"`       // 表技术名称
	BusinessName          string          `json:"business_name"`        // 表业务名称
	UniformCatalogCode    string          `json:"uniform_catalog_code"` // 逻辑视图编码
	DatasourceName        string          `json:"datasource_name"`      // 数据源名称 （不可编辑）
	DatasourceID          string          `json:"datasource_id"`        //　数据源ID
	Schema                string          `json:"schema"`               // 库名称  （不可编辑）
	InfoSystem            string          `json:"info_system"`          // 关联信息系统  （不可编辑）
	Description           string          `json:"description"`          // 描述
	SubjectID             string          `json:"subject_id"`           // 所属主题id
	SubjectPathID         string          `json:"subject_path_id"`      // 所属主题path id
	Subject               string          `json:"subject"`              // 所属主题
	DepartmentID          string          `json:"department_id"`        // 所属部门id
	Department            string          `json:"department"`           // 所属部门
	OwnerID               string          `json:"owner_id"`             // 数据Owner id
	Owner                 string          `json:"owner"`                // 数据Owner
	Owners                []DataViewOwner `json:"owners,omitempty"`
	SceneAnalysisId       string          `json:"scene_analysis_id"`        // 场景分析画布id
	ViewSourceCatalogName string          `json:"view_source_catalog_name"` // 视图源
	PublishAt             int64           `json:"publish_at"`               // 发布时间
	CreatedAt             int64           `json:"created_at"`               // 创建时间
	CreatedByUser         string          `json:"created_by"`               // 创建人
	UpdatedAt             int64           `json:"updated_at"`               // 编辑时间
	UpdatedByUser         string          `json:"updated_by"`               // 编辑人

	Sheet            string `json:"sheet"`               // sheet页,逗号分隔
	StartCell        string `json:"start_cell"`          // 起始单元格
	EndCell          string `json:"end_cell"`            // 结束单元格
	HasHeaders       bool   `json:"has_headers"`         // 是否首行作为列名
	SheetAsNewColumn bool   `json:"sheet_as_new_column"` // 是否将sheet作为新列
	ExcelFileName    string `json:"excel_file_name"`     // excel文件名
}

// DataViewOwner 逻辑视图的 Owner
type DataViewOwner struct {
	// 数据Owner id
	OwnerID string `json:"owner_id,omitempty"`
	// 数据Owner
	Owner string `json:"owner,omitempty"`
}

type GetFieldsRes struct {
	FieldsRes             []*FieldsRes `json:"fields"`
	ID                    string       `json:"id"`                       // 逻辑视图ID
	LastPublishTime       int64        `json:"last_publish_time"`        // 最新发布时间（已发布）
	UniformCatalogCode    string       `json:"uniform_catalog_code"`     // 逻辑视图编码
	TechnicalName         string       `json:"technical_name"`           // 表技术名称
	BusinessName          string       `json:"business_name"`            // 表业务名称
	Status                string       `json:"status"`                   // 表状态
	EditStatus            string       `json:"edit_status"`              // 编辑状态
	DatasourceId          string       `json:"datasource_id"`            // 数据源id
	DatasourceType        string       `json:"datasource_type"`          // 数据源类型
	ViewSourceCatalogName string       `json:"view_source_catalog_name"` // 视图源
	Type                  string       `json:"type"`                     // 视图类型
	ExploreJobId          string       `json:"explore_job_id"`           // 探查作业ID
	ExploreJobVer         int          `json:"explore_job_version"`      // 探查作业版本
}

type FieldsRes struct {
	ID                  string  `json:"id"`                     // 列uuid
	TechnicalName       string  `json:"technical_name"`         // 列技术名称
	OriginalName        string  `json:"original_name"`          // 原始字段名称
	BusinessName        string  `json:"business_name"`          // 列业务名称
	Comment             string  `json:"comment"`                // 列注释
	Status              string  `json:"status"`                 // 列视图状态(扫描结果) 0：无变化、1：新增、2：删除
	PrimaryKey          bool    `json:"primary_key"`            // 是否主键
	DataType            string  `json:"data_type"`              // 数据类型
	DataLength          int32   `json:"data_length"`            // 数据长度
	DataAccuracy        int32   `json:"data_accuracy"`          // 数据精度（仅DECIMAL类型）
	OriginalDataType    string  `json:"original_data_type"`     // 原始数据类型
	IsNullable          string  `json:"is_nullable"`            // 是否为空 (YES/NO)
	BusinessTimestamp   bool    `json:"business_timestamp"`     // 是否业务时间字段
	StandardCode        string  `json:"standard_code"`          // 数据标准code
	Standard            string  `json:"standard"`               // 数据标准名称
	StandardType        string  `json:"standard_type"`          // 数据标准类型
	StandardTypeName    string  `json:"standard_type_name"`     // 数据标准类型名称
	StandardStatus      string  `json:"standard_status"`        // 数据标准状态
	CodeTableID         string  `json:"code_table_id"`          // 码表ID
	CodeTable           string  `json:"code_table"`             // 码表名称
	CodeTableStatus     string  `json:"code_table_status"`      // 码表状态
	IsReadable          bool    `json:"is_readable"`            // 当前用户是否有此字段的读取权限
	IsDownloadable      bool    `json:"is_downloadable"`        // 当前用户是否有此字段的下载权限
	AttributeID         *string `json:"attribute_id"`           // L5属性ID
	AttributeName       string  `json:"attribute_name"`         // L5属性名称
	AttributePath       string  `json:"attribute_path"`         // 路径id
	LabelID             string  `json:"label_id"`               // 标签ID
	LabelName           string  `json:"label_name"`             // 标签名称
	LabelIcon           string  `json:"label_icon"`             // 标签颜色
	LabelPath           string  `json:"label_path"`             //标签路径
	LabelIsProtected    bool    `json:"label_is_protected"`     // 标签是否受数据查询保护
	ClassfityType       *int    `json:"classfity_type"`         // 分类类型(1自动2人工)
	GradeType           *int    `json:"grade_type"`             // 分级类型(1自动2人工)
	EnableRules         int     `json:"enable_rules"`           // 已开启字段级规则数
	TotalRules          int     `json:"total_rules"`            // 字段级规则总数
	ResetBeforeDataType string  `json:"reset_before_data_type"` // 重置前数据类型
	ResetConvertRules   string  `json:"reset_convert_rules"`    // 重置转换规则 （仅日期类型）
	ResetDataLength     int32   `json:"reset_data_length" `     // 重置数据长度（仅DECIMAL类型）
	ResetDataAccuracy   int32   `json:"reset_data_accuracy"`    // 重置数据精度（仅DECIMAL类型）
	SimpleType          string  `json:"simple_type"`            // 数据大类型
	Index               int     `json:"index"`                  // 字段顺序
}

//endregion

type GetSubViewIDsOptions struct {
	// 逻辑视图 ID，非空时返回属于这个逻辑视图的子视图（行列规则）ID 列表
	LogicViewID string
}

// region GetLogicViewReportInfo

type GetLogicViewReportInfoReq struct {
	GetLogicViewReportInfoBody `param_type:"body"`
}
type GetLogicViewReportInfoBody struct {
	FieldIds []string `json:"field_id" form:"field_id" binding:"required,dive,uuid"`
}
type ReportInfo struct {
	FieldTechnicalName string `json:"field_technical_name"`
	DatasourceSchema   string `json:"datasource_schema"`
	DatasourceId       string `json:"datasource_id"`
}

type GetLogicViewReportInfoRes struct {
	ReportInfos map[string]*ReportInfo `json:"report_infos"`
}

//endregion

//region basicInfo

type ViewBasicInfo struct {
	Id                 string      `json:"id"`
	UniformCatalogCode string      `json:"uniform_catalog_code"`
	TechnicalName      string      `json:"technical_name"`
	BusinessName       string      `json:"business_name"`
	Type               int         `json:"type"`
	DatasourceId       string      `json:"datasource_id"`
	PublishAt          interface{} `json:"publish_at"`
	EditStatus         int         `json:"edit_status"`
}

//endregion

//region GetViewInfoReq

type GetViewInfoReq struct {
	IDs []string `json:"ids" form:"ids" binding:"required"` // 视图id
}

type GetViewInfoResp struct {
	Entries []*ViewInfo `json:"entries" binding:"omitempty"` // 逻辑视图列表
}

type ViewInfo struct {
	ID                    string `json:"id"`                       // ID
	UniformCatalogCode    string `json:"uniform_catalog_code"`     // 逻辑视图编码
	TechnicalName         string `json:"technical_name"`           // 表技术名称
	BusinessName          string `json:"business_name"`            // 表业务名称
	Type                  string `json:"type"`                     // 视图类型
	DatasourceName        string `json:"datasource_name"`          // 数据源名称
	Description           string `json:"description"`              // 描述
	DepartmentID          string `json:"department_id"`            // 所属部门id
	Department            string `json:"department"`               // 所属部门
	DepartmentPath        string `json:"department_path"`          // 所属部门路径
	IsAuditRuleConfigured bool   `json:"is_audit_rule_configured"` // 是否已配置稽核规则
	Status                string `json:"status"`                   // 逻辑视图状态\扫描结果
}

//endregion

// region pre-view

type DataPreviewReq struct {
	FormViewId  string    `json:"form_view_id" form:"form_view_id" binding:"required,uuid" example:"88f78432-ee4e-43df-804c-4ccc4ff17f15"` // 逻辑视图id
	Fields      []string  `json:"fields" form:"fields" binding:"required,dive,uuid"`                                                       // 输出字段
	Direction   string    `json:"direction" form:"direction,default=desc" binding:"omitempty,oneof=asc desc" default:"desc"`               // 排序方向，枚举：asc：正序；desc：倒序。默认倒序
	SortFieldId string    `json:"sort_field_id" form:"sort_field_id" binding:"omitempty,uuid"`                                             // 排序字段id
	Filters     []*Member `json:"filters" form:"filters" binding:"omitempty,dive"`                                                         // 过滤规则
	Offset      *int      `json:"offset" form:"offset,default=1" binding:"omitempty,min=1" default:"1"`                                    // 页码，默认1
	Limit       *int      `json:"limit" form:"limit,default=10" binding:"omitempty,min=1,max=1000" default:"10"`                           // 每页大小，默认10
	Configs     string    `json:"configs" form:"configs" binding:"omitempty" default:""`
	IfCount     int       `json:"if_count" form:"if_count" binding:"omitempty" default:"0"`
}

type Member struct {
	FieldObjV1        // 字段对象
	Operator   string `json:"operator" form:"field_id" binding:"required"` // 限定条件
	Value      string `json:"value" form:"value"`                          // 限定比较值
}

type FieldObjV1 struct {
	ID       string `json:"id" form:"id" binding:"omitempty,uuid" example:"0130dc92-2660-44dd-8de8-171d1ef125aa"` // 字段ID
	Name     string `json:"name" form:"name" binding:"omitempty,VerifyName255NoSpace"`                            // 字段名称
	NameEn   string `json:"name_en" form:"name_en" binding:"omitempty,VerifyName255NoSpace"`                      // 原字段名称
	DataType string `json:"data_type" form:"data_type" binding:"omitempty"`                                       // 字段类型
}

type DataPreviewResp struct {
	FetchDataRes
}

type FetchDataRes struct {
	TotalCount int       `json:"total_count"`
	Columns    []*Column `json:"columns"`
	Data       [][]any   `json:"data"`
}
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// endregion

// region GetViewListByTechnicalNameInMultiDatasource

var GetViewListByTechnicalNameInMultiDatasourceUrl = "/api/internal/data-view/v1/datasource/form-view/te-name"

type GetViewListByTechnicalNameInMultiDatasourceReq struct {
	Datasource []*GetViewListDatasource `json:"datasource"`
}
type GetViewListDatasource struct {
	DatasourceID  string   `json:"datasource_id"`
	TechnicalName []string `json:"technical_name"`
	OriginalName  []string `json:"original_name"` // 原始表名称
}
type GetViewListByTechnicalNameInMultiDatasourceRes struct {
	FormViews []*FormView `json:"form_views"`
}

type FormView struct {
	ID                    string `json:"id"`                       // 逻辑视图uuid
	UniformCatalogCode    string `json:"uniform_catalog_code"`     // 逻辑视图编码
	TechnicalName         string `json:"technical_name"`           // 表技术名称
	OriginalName          string `json:"original_name"`            // 原始表名称
	BusinessName          string `json:"business_name"`            // 表业务名称
	Type                  string `json:"type"`                     // 逻辑视图来源
	DatasourceId          string `json:"datasource_id"`            // 数据源id
	Datasource            string `json:"datasource"`               // 数据源
	DatasourceType        string `json:"datasource_type"`          // 数据源类型
	DatasourceCatalogName string `json:"datasource_catalog_name"`  // 数据源catalog
	Status                string `json:"status"`                   // 逻辑视图状态\扫描结果
	PublishAt             int64  `json:"publish_at"`               // 发布时间
	OnlineTime            int64  `json:"online_time"`              // 上线时间
	OnlineStatus          string `json:"online_status"`            // 上线状态
	AuditAdvice           string `json:"audit_advice"`             // 审核意见，仅驳回时有用
	EditStatus            string `json:"edit_status"`              // 内容状态
	ViewSourceCatalogName string `json:"view_source_catalog_name"` // 视图源
}

//endregion

type SubView struct {
	ID string `json:"id,omitempty" example:"0194078b-5413-7022-a7a8-75a820dbf994"`

	SubViewSpec `json:",inline"`
}

type SubViewSpec struct {
	// 行列规则（子视图）名称
	Name string `json:"name,omitempty" example:"北区数据"`
	//授权范围
	AuthScopeID string `json:"auth_scope_id"`
	// 行列规则（子视图）所属的逻辑视图 ID
	LogicViewID string `json:"logic_view_id,omitempty" example:"0194077d-2290-7387-b505-ac3208b20087"`
	// 行列规则（子视图）的详细定义，JSON 字符串
	Detail string `json:"detail,omitempty" example:"{\"fields\":[{\"id\":\"84d26012-e586-4559-93c3-42e1caa49707\",\"name_en\":\"a1611\",\"name\":\"a1611\",\"data_type\":\"int\"}],\"row_filters\":{\"where\":[],\"where_relation\":\"and\"}}"`
}

//region GetByAuditStatusReq

type GetByAuditStatusReq struct {
	Offset         *int     `json:"offset" form:"offset" binding:"omitempty"`
	Limit          *int     `json:"limit" form:"limit" binding:"omitempty"`
	Keyword        string   `json:"keyword" form:"keyword" binding:"KeywordTrimSpace,omitempty,min=1,max=255"` // 关键字查询，字符无限制
	DatasourceType string   `json:"datasource_type" form:"datasource_type" binding:"omitempty"`                // 数据源类型
	DatasourceIds  []string `json:"-" `
	DatasourceId   string   `json:"datasource_id" form:"datasource_id" binding:"omitempty,uuid"`                        // 数据源id
	PublishStatus  string   `json:"publish_status" form:"publish_status" binding:"omitempty,oneof=publish unpublished"` // 发布状态
	IsAudited      *bool    `json:"is_audited"  form:"is_audited" binding:"omitempty"`                                  // 是否已稽核
}

type GetByAuditStatusResp struct {
	Entries    []*FormViewInfo `json:"entries" binding:"required"`                       // 对象列表
	TotalCount int64           `json:"total_count" binding:"required,gte=0" example:"3"` // 当前筛选条件下的对象数量
}

type FormViewInfo struct {
	ID                 string `json:"id"`                   // 逻辑视图uuid
	UniformCatalogCode string `json:"uniform_catalog_code"` // 逻辑视图编码
	TechnicalName      string `json:"technical_name"`       // 表技术名称
	BusinessName       string `json:"business_name"`        // 表业务名称
	DepartmentID       string `json:"department_id"`        // 所属部门id
	DepartmentPath     string `json:"department_path"`      // 所属部门路径
}

//endregion

type BatchGetExploreReportReq struct {
	IDs              []string `json:"ids" binding:"required,min=1,dive,uuid"` // 逻辑视图ID列表（至少1个，支持单个ID）
	Version          *int32   `json:"version" binding:"omitempty"`            // 报告版本（可选，不传默认nil，获取最新版本）
	ThirdParty       bool     `json:"third_party" binding:"omitempty"`        // 第三方报告（可选，不传默认false）
	HasQualityReport bool     `json:"has_quality_report" binding:"omitempty"` // 是否有质量报告标识（可选，不传默认false）
}

type BatchGetExploreReportResp struct {
	Reports []*BatchExploreReportItem `json:"reports"` // 报告列表
}

type BatchExploreReportItem struct {
	FormViewID       string             `json:"form_view_id"`       // 逻辑视图ID
	Success          bool               `json:"success"`            // 是否成功获取报告
	HasQualityReport bool               `json:"has_quality_report"` // 是否存在质量报告（可用性标识）
	Error            string             `json:"error,omitempty"`    // 错误信息（如果失败）
	Report           *ExploreReportResp `json:"report,omitempty"`   // 报告内容（如果成功）
}

//region CreateWorkOrderTaskReq

type CreateWorkOrderTaskReq struct {
	WorkOrderID  string   `json:"work_order_id" binding:"required,uuid"` // 工单id
	FormViewIDs  []string `json:"form_view_ids" binding:"required"`      // 视图id
	CreatedByUID string   `json:"created_by_uid" binding:"required"`     // 创建人id
	TotalSample  int64    `json:"total_sample" binding:"omitempty"`      // 采样数据量,0为全量数据
}

type CreateWorkOrderTaskResp struct {
	Result []Result `json:"result"`
}

type Result struct {
	FormViewId string
	TaskId     string
	Error      error
}

//endregion

//region GetExploreTskList

type ListExploreTaskResp struct {
	Entries    []*ExploreTaskInfo `json:"entries"`     // 对象列表
	TotalCount int64              `json:"total_count"` // 当前筛选条件下的对象数量
}

type ExploreTaskInfo struct {
	TaskID         string `json:"task_id"`         // 任务id
	Type           string `json:"type"`            // 任务类型
	DatasourceID   string `json:"datasource_id"`   // 数据源id
	DatasourceName string `json:"datasource_name"` // 数据源名称
	DatasourceType string `json:"datasource_type"` // 数据源类型
	FormViewID     string `json:"form_view_id"`    // 视图id
	FormViewName   string `json:"form_view_name"`  // 视图名称
	FormViewType   string `json:"form_view_type"`  // 视图类型
	Status         string `json:"status"`          // 任务状态
	Config         string `json:"config"`          // 探查配置
	CreatedAt      int64  `json:"created_at"`      // 开始时间
	CreatedBy      string `json:"created_by"`      // 发起人
	FinishedAt     int64  `json:"finished_at"`     // 结束时间
	Remark         string `json:"remark"`          // 异常原因
}

//endregion

//region CreateExploreTaskReq

type CreateExploreTaskReq struct {
	WorkOrderID  string `json:"work_order_id" binding:"required,uuid"` // 工单id
	FormViewID   string `json:"form_view_id" binding:"required"`       // 视图id
	CreatedByUID string `json:"created_by_uid" binding:"required"`     // 创建人id
	TotalSample  int64  `json:"total_sample" binding:"omitempty"`      // 采样数据量,0为全量数据
	Version      int32  `json:"version" binding:"omitempty"`           // 版本
}

type CreateExploreTaskResp struct {
	FormViewId string `json:"form_view_id"`
	TaskId     string `json:"task_id"`
}

//endregion

//region GetSampleData

type GetSampleDataRes struct {
	Type string `json:"type"` //合成/样例
	*virtual_engine.FetchDataRes
}

//endregion
