package data_catalog

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-frame/core/models"
)

type Driven interface {
	Catalog
	DataPush
}

type Catalog interface {
	//GetDataCatalogDetail 获取目录详情
	GetDataCatalogDetail(ctx context.Context, catalogID string) (*GetDataCatalogDetailResp, error)
	//GetDataCatalogColumnList  获取目录的信息项
	GetDataCatalogColumnList(ctx context.Context, catalogID string) (*GetDataCatalogColumnsRes, error)
	//GetDataCatalogMountList  获取目录挂载的列表
	GetDataCatalogMountList(ctx context.Context, catalogID string) (*GetDataCatalogMountListRes, error)
	//GetInfoCatalogByStandardForm  通过业务标准表查询信息资源目录
	GetInfoCatalogByStandardForm(ctx context.Context, formID []string) ([]*GetCatalogByStandardFormItem, error)

	GetTemplateDetail(ctx context.Context, id string) (res *TemplateReq, err error)

	GetColumnMapByIds(ctx context.Context, req *GetColumnListByIdsReq) (map[uint64]*ColumnNameInfo, error)

	GetCatalogColumnByViewID(ctx context.Context, id string) ([]*ColumnInfo, error)
	GetComprehensionDetail(ctx context.Context, catalogId, templateID string) (*ComprehensionDetail, error)
	GetResourceFavoriteByID(ctx context.Context, req *CheckV1Req) (map[uint64]*CheckV1Resp, error)
}

type DataPush interface {
	CreateDataPush(ctx context.Context, req *DataPushCreateReq) (string, error)
	UpdateDataPush(ctx context.Context, req *UpdateReq) (string, error)
	BatchUpdateStatus(ctx context.Context, req *BatchUpdateStatusReq) ([]uint64, error)
	History(ctx context.Context, req *TaskExecuteHistoryReq) (*base.PageResult[TaskLogInfo], error)
	SandboxPushCount(ctx context.Context, ids []string) (map[string]int, error)
}

type QueryCatalogIDReq struct {
	CatalogID  string `uri:"catalog_id"`
	Limit      int    `query:"limit"`
	ReportShow bool   `json:"report_show"` //上报信息展示
}

type QueryCatalogSimpleResp struct {
	CatalogID    string `json:"catalog_id"`
	Title        string `json:"title"`
	Code         string `json:"code"`
	ReportStage  int    `json:"report_stage"`
	OrgCode      string `json:"org_code"`
	ResourceType int32  `json:"resource_type"`
}

// region GetDataCatalogDetail

type GetDataCatalogDetailResp struct {
	ID               string `json:"id"`                                          //数据目录的ID
	Name             string `json:"name" binding:"required,VerifyNameStandard"`  // 数据资源目录名称
	Code             string `json:"code"`                                        // 目录编码
	SourceDepartment string `json:"source_department" binding:"omitempty,uuid"`  // 数据资源来源部门
	ResourceType     int8   `json:"resource_type" binding:"omitempty,oneof=1 2"` // 资源类型 1逻辑视图 2 接口
	ReportInfo
	DepartmentInfo
	InfoSystemInfo
	CategoryInfos      []*CategoryInfo `json:"category_infos"`                                             //自定义类目
	SubjectInfo        []*SubjectInfo  `json:"subject_info"`                                               // 所属主题
	AppSceneClassify   int8            `json:"app_scene_classify"  binding:"omitempty"`                    // 应用场景分类
	DataRelatedMatters string          `json:"data_related_matters"  binding:"required"`                   // 数据所属事项
	DataRange          int32           `json:"data_range" binding:"omitempty,oneof=1 2 3"`                 // 数据范围：字典DM_DATA_SJFW，01全市 02市直 03区县
	UpdateCycle        int32           `json:"update_cycle" binding:"omitempty,min=1,max=9"`               // 更新频率 参考数据字典：GXZQ，1不定时 2实时 3每日 4每周 5每月 6每季度 7每半年 8每年 9其他
	DataClassify       string          `json:"data_classify" binding:"required"`                           // 数据分级
	Description        string          `json:"description" binding:"omitempty,VerifyDescription,max=1000"` // 数据资源目录描述
	SharedOpenInfo                     //共享开放信息
	MoreInfo                           //更多信息
	PublishStatus      string          `json:"publish_status"` // 发布状态
	PublishAt          int64           `json:"publish_at"`     // 发布时间
	OnlineStatus       string          `json:"online_status"`  // 上线状态
	OnlineTime         int64           `json:"online_time"`    // 上线时间
	CreatedAt          int64           `json:"created_at"`     // 创建时间
	UpdatedAt          int64           `json:"updated_at"`     // 编辑时间
}
type ReportInfo struct {
	DataDomain            int32  `json:"data_domain" binding:"omitempty,min=1,max=27"`                                              // 数据所属领域
	DataLevel             int32  `json:"data_level"  binding:"omitempty,min=1,max=4"`                                               // 数据所在层级
	TimeRange             string `json:"time_range" `                                                                               // 数据时间范围
	ProviderChannel       int32  `json:"provider_channel" binding:"omitempty,min=1,max=3"`                                          // 提供渠道
	AdministrativeCode    int32  `json:"administrative_code" binding:"omitempty,ValidateAdministrativeCode"`                        // 行政区划代码
	CentralDepartmentCode int32  `json:"central_department_code" binding:"omitempty,min=2,max=99"`                                  // 中央业务指导部门代码
	ProcessingLevel       string `json:"processing_level" binding:"omitempty,oneof=sjjgcd01 sjjgcd02 sjjgcd03 sjjgcd04 sjjgcd05 0"` // 数据加工程度
	CatalogTag            int32  `json:"catalog_tag"  binding:"omitempty,min=1,max=50"`                                             // 目录 标签
	IsElectronicProof     bool   `json:"is_electronic_proof,default=true" default:"true"`                                           // 是否电子证明编码
}

type SharedOpenInfo struct {
	SharedType      int8   `json:"shared_type" binding:"required,oneof=1 2 3"`                                                        // 共享属性 1 无条件共享 2 有条件共享 3 不予共享
	SharedCondition string `json:"shared_condition" binding:"required_unless=SharedType 1,omitempty,VerifyDescription,min=1,max=255"` // 共享条件
	OpenType        int8   `json:"open_type" binding:"required,oneof=1 2"`                                                            // 开放属性 1 向公众开放 2 不向公众开放
	OpenCondition   string `json:"open_condition" binding:"omitempty,VerifyDescription,max=255"`                                      // 开放条件
	SharedMode      int8   `json:"shared_mode" binding:"required_unless=SharedType 3,omitempty,oneof=1 2 3"`                          // 共享方式 1 共享平台方式 2 邮件方式 3 介质方式
}

type MoreInfo struct {
	PhysicalDeletion *int8  `json:"physical_deletion" binding:"omitempty,oneof=0 1"`       // 挂接实体资源是否存在物理删除(1 是 ; 0 否)
	SyncMechanism    int8   `json:"sync_mechanism" binding:"omitempty,oneof=1 2"`          // 数据归集机制(1 增量 ; 2 全量)
	SyncFrequency    string `json:"sync_frequency" binding:"omitempty,VerifyNameStandard"` // 同步频率
	PublishFlag      *int8  `json:"publish_flag" binding:"required,omitempty,oneof=0 1"`   // 是否发布到超市 (1 是 ; 0 否)
}

type SubjectInfo struct {
	SubjectID   string `json:"subject_id"`   // 所属主题id
	SubjectName string `json:"subject"`      // 所属主题
	SubjectPath string `json:"subject_path"` // 所属主题路径
}
type DepartmentInfo struct {
	DepartmentID   string `json:"department_id"`   // 所属部门id
	Department     string `json:"department"`      // 所属部门
	DepartmentPath string `json:"department_path"` // 所属部门路径
}
type InfoSystemInfo struct {
	InfoSystemID string `json:"info_system_id"` // 信息系统id
	InfoSystem   string `json:"info_system"`    // 关联信息系统
}

type CategoryInfo struct {
	CategoryID     string `json:"category_id"`      // 资源属性分类id
	Category       string `json:"category"`         // 资源属性分类
	CategoryNodeID string `json:"category_node_id"` // 资源属性分类节点id
	CategoryNode   string `json:"category_node"`    // 资源属性分类节点
}

//endregion

type QueryCatalogByPageReq struct {
	Offset    *int    `query:"offset"`    // 页码，默认1
	Limit     *int    `query:"limit"`     // 每页大小，默认10 limit=0不分页
	Direction *string `query:"direction"` // 排序方向，枚举：asc：正序；desc：倒序。默认倒序
	Sort      *string `query:"sort"`      // 排序类型，枚举：created_at：按创建时间排序；updated_at：按更新时间排序。默认按创建时间排序
}

type QueryCatalogByPageResp struct{}

// region  GetDataCatalogColumns

type GetDataCatalogColumnsRes struct {
	Columns []*ColumnInfoRes `json:"columns"` // 关联信息项
}
type ColumnInfoRes struct {
	ColumnInfo
	StandardCode     string `json:"standard_code"`      // 数据标准code
	Standard         string `json:"standard"`           // 数据标准名称
	StandardType     int    `json:"standard_type"`      // 数据标准类型
	StandardTypeName string `json:"standard_type_name"` // 数据标准类型名称
	StandardStatus   string `json:"standard_status"`    // 数据标准状态
	CodeTable        string `json:"code_table"`         // 码表名称
	CodeTableStatus  string `json:"code_table_status"`  // 码表状态
}

type ColumnInfo struct {
	ID              string `json:"id" form:"id" uri:"id"`
	BusinessName    string `json:"business_name"`    // 信息项业务名称
	TechnicalName   string `json:"technical_name"`   // 信息项技术名称
	SourceID        string `json:"source_id"`        // 来源id
	StandardCode    string `json:"standard_code"`    // 关联数据标准code
	CodeTableID     string `json:"code_table_id"`    // 关联码表IDe
	DataType        *int32 `json:"data_type"`        // 字段类型 0:数字型 1:字符型 2:日期型 3:日期时间型 4:时间戳型 5:布尔型 6:二进制
	DataLength      *int32 `json:"data_length"`      // 数据长度
	DataRange       string `json:"data_range" `      // 数据值域 中英文、数字、下划线及中划线，且不能以下划线和中划线开头 128个字符
	SharedType      int8   `json:"shared_type" `     // 共享属性 1 无条件共享 2 有条件共享 3 不予共享
	SharedCondition string `json:"shared_condition"` // 共享条件
	OpenType        int8   `json:"open_type"`        // 开放属性 1 无条件开 2 有条件开 3 不予开
	OpenCondition   string `json:"open_condition"`   // 开放条件
	ClassifiedFlag  *int8  `json:"classified_flag"`  // 是否涉密属性(1 是 ; 0 否)
	SensitiveFlag   *int8  `json:"sensitive_flag"`   // 是否敏感属性(1 是 ; 0 否)
	TimestampFlag   *int8  `json:"timestamp_flag"`   // 是否时间戳(1 是 ; 0 否)
	PrimaryFlag     *int8  `json:"primary_flag"`     // 是否主键(1 是 ; 0 否)

	//上报属性
	SourceTechnicalName string `json:"source_technical_name,omitempty"` // 来源技术名称
	SourceSystemId      string `json:"source_system_id,omitempty"`      // 来源系统id
	SourceSystemSchema  string `json:"source_system_schema,omitempty"`  // 来源系统
	SourceSystemLevel   int32  `json:"source_system_level,omitempty"`   // 来源系统分级 1 自建自用 2 国直(国家部委统一平台) 3省直(省级统一平台) 4市直(市级统一平台) 5县直(县级统一平台)
	InfoItemLevel       string `json:"info_item_level"`                 // 信息项分级
}

//endregion

type GetDataCatalogMountListRes struct {
	MountResource []*MountResourceRes `json:"mount_resource"` // 挂载资源
}

type MountResourceRes struct {
	MountResource
	Name           string `json:"name"`
	Code           string `json:"code"`            // 统一编目编码
	DepartmentId   string `json:"department_id"`   // 所属部门id
	Department     string `json:"department"`      // 所属部门
	DepartmentPath string `json:"department_path"` // 所属部门路径
	PublishAt      int64  `json:"publish_at"`      // 发布时间
	Status         int8   `json:"status"`          // 视图状态,1正常,2删除
}
type MountResource struct {
	ResourceType int8   `json:"resource_type"` // 挂接资源类型 1逻辑视图 2 接口
	ResourceID   string `json:"resource_id"`   // 挂接资源ID
	//上报属性
	SchedulingInfo         // 调度信息
	RequestFormat  string  `json:"request_format"`  // 服务请求报文格式application/json application/xml application/x-www-form-urlencoded multipart/form-data text/plain;charset-uft-8 others
	ResponseFormat string  `json:"response_format"` // 服务响应报文格式application/json application/xml application/x-www-form-urlencoded multipart/form-data text/plain;charset-uft-8 others
	RequestBody    []*Body `json:"request_body"`    // 请求体
	ResponseBody   []*Body `json:"response_body"`   // 响应体
}

type SchedulingInfo struct {
	SchedulingPlan int32  `json:"scheduling_plan" ` // 调度计划 1 一次性、2按分钟、3按天、4按周、5按月
	Interval       int32  `json:"interval"`         // 间隔
	Time           string `json:"time"`             // 时间
}

type Body struct {
	ID         string `json:"id"`
	Name       string `json:"name"`        //参数名
	Type       string `json:"type"`        //int32 int64 float double byte binary date date-time boolean
	IsArray    bool   `json:"is_array"`    //是否数组
	HasContent bool   `json:"has_content"` //是否有内容
}

// [通过业务标准表查询目录]
type GetCatalogByStandardForm struct {
	StandardFormID []string `query:"standard_form_id"` // 业务标准表ID
}

type GetCatalogByStandardFormItem struct {
	ID             string `json:"id"`               // 信息类ID
	Name           string `json:"name"`             // 信息类名称
	Code           string `json:"code"`             // 信息类编码
	BusinessFormID string `json:"business_form_id"` //业务标准表ID
}

type TemplateReq struct {
	Name           string         `json:"name" binding:"required,min=1,max=255" example:"xxxx"`                    //理解模板名称
	Description    string         `json:"description" binding:"TrimSpace,omitempty,lte=300" example:"description"` //理解模板描述
	TemplateConfig TemplateConfig `json:"template_config" binding:"required"`                                      //理解模板配置
}
type TemplateConfig struct {
	BusinessObject *bool `json:"business_object"  binding:"required"` //业务对象

	//时间维度

	TimeRange              *bool `json:"time_range" binding:"required"`               //时间范围
	TimeFieldComprehension *bool `json:"time_field_comprehension" binding:"required"` //时间字段理解

	//空间维度

	SpatialRange              *bool `json:"spatial_range" binding:"required"`               //空间范围
	SpatialFieldComprehension *bool `json:"spatial_field_comprehension" binding:"required"` //空间字段理解

	BusinessSpecialDimension *bool `json:"business_special_dimension" binding:"required"` //业务特殊维度
	CompoundExpression       *bool `json:"compound_expression" binding:"required"`        //复合表达
	ServiceRange             *bool `json:"service_range" binding:"required"`              //服务范围
	ServiceAreas             *bool `json:"service_areas" binding:"required"`              //服务领域
	FrontSupport             *bool `json:"front_support" binding:"required"`              //正面支撑
	NegativeSupport          *bool `json:"negative_support" binding:"required"`           //负面支撑

	//业务规则

	ProtectControl *bool `json:"protect_control" binding:"required"` //保护/控制什么
	PromotePush    *bool `json:"promote_push" binding:"required"`    //促进/推动什么
}

//region GetColumnListByIds

type GetColumnListByIdsReq struct {
	IDs []uint64 `json:"ids"` // 信息项id
}
type FavorIDBase struct {
	ID    uint64 `gorm:"column:id;primaryKey" json:"id,string"` // 收藏项ID
	ResID string `gorm:"column:res_id" json:"res_id"`           // 资源ID
}
type CheckV1Resp struct {
	IsFavored bool   `json:"is_favored"`                // 是否已收藏
	FavorID   uint64 `json:"favor_id,string,omitempty"` // 收藏项ID，仅已收藏时返回该字段
}

type CheckV1Req struct {
	ResType   string `form:"res_type" json:"res_type" binding:"TrimSpace,required,oneof=data-catalog info-catalog elec-licence-catalog data-view interface-svc indicator" example:"data-catalog"` // 收藏资源类型 data-catalog 数据资源目录 info-catalog 信息资源目录 elec-licence-catalog 电子证照目录
	ResID     string `form:"res_id" json:"res_id" binding:"TrimSpace,required,min=1,max=64" example:"544217704094017271"`                                                                         // 收藏资源ID
	CreatedBy string `form:"created_by" json:"created_by" binding:"TrimSpace,required,min=1,max=64" example:"544217704094017271"`
}

type GetColumnListByIdsResp struct {
	Columns []*ColumnNameInfo `json:"columns"` // 信息项
}
type ColumnNameInfo struct {
	ID            uint64 `json:"id" binding:"required" example:"1"`
	BusinessName  string `json:"business_name" binding:"required,min=1,max=255" example:"业务名称"`  // 信息项业务名称
	TechnicalName string `json:"technical_name" binding:"required,min=1,max=255" example:"技术名称"` // 信息项技术名称
}

//endregion

//region GetComprehensionDetail

type ComprehensionDetail struct {
	CatalogID               models.ModelID      `json:"catalog_id"`               //编目ID
	CatalogCode             string              `json:"catalog_code"`             //编目code
	CatalogInfo             *CatalogInfo        `json:"catalog_info"`             //数据编目相关信息
	Note                    string              `json:"note"`                     //数据理解提示语
	Status                  int8                `json:"status"`                   //编目状态
	AuditAdvice             string              `json:"audit_advice"`             //审核意见，仅驳回时有用
	UpdatedAt               int64               `json:"updated_at"`               // 更新时间
	ComprehensionDimensions []*DimensionConfig  `json:"comprehension_dimensions"` //数据理解详情
	ColumnComments          []ColumnComment     `json:"column_comments"`          //字段注释理解
	Choices                 map[string][]Choice `json:"choices"`
	Icons                   map[string]string   `json:"icons"` // icon
}

type CatalogInfo struct {
	ID              models.ModelID `json:"id"`              //编目ID
	DepartmentInfos []*SummaryInfo `json:"department_path"` //部门处室
	Name            string         `json:"name"`            //目录中文名称
	NameEn          string         `json:"name_en"`         //目录英文名称
	BusinessDuties  []string       `json:"business_duties"` //业务职责
	BaseWorks       []string       `json:"base_works"`      //开展工作
	UpdateCycle     int32          `json:"update_cycle"`    //更新周期
	TableName       string         `json:"table_name"`      //挂载的表名
	TableId         string         `json:"table_id"`        //挂载资源id
	TableDesc       string         `json:"table_desc"`      //表含义
	UpdatedAt       int64          `json:"updated_at"`      //理解更新人
	UpdaterUID      string         `json:"updater_uid"`     //理解更新时间
	UpdaterName     string         `json:"updater_name"`    //更新用户名称
}

type SummaryInfo struct {
	ID     string `json:"id" `     // 对象ID
	Name   string `json:"name"`    // 对象名称
	Type   string `json:"type"`    // 对象类型
	Path   string `json:"path"`    // 对象路径
	PathID string `json:"path_id"` // 对象ID路径
}

type DimensionConfig struct {
	CatalogId models.ModelID     `json:"-"`                     //目录ID
	Id        string             `json:"id" binding:"required"` //配置的ID
	Name      string             `json:"name"`                  //数据理解名称
	Category  string             `json:"category"`              //所属分类, business:业务信息;value:价值信息
	Error     string             `json:"error,omitempty"`       //维度错误,详情接口的返回错误
	Children  []*DimensionConfig `json:"children,omitempty"`    //子维度,，叶子节点没有Children配置
	Detail    *DimensionDetail   `json:"detail,omitempty"`      //具体的配置, 非叶子节点没有该配置
}

type DimensionDetail struct {
	CatalogId         models.ModelID `json:"-"`                        //目录ID
	DimensionConfigId string         `json:"-"`                        //理解配置ID
	DimensionName     string         `json:"-"`                        //维度名称
	Required          bool           `json:"required"`                 //是否必填
	IsMulti           bool           `json:"is_multi"`                 //是否有多个值
	MaxMulti          int            `json:"max_multi"`                //最高支持几个值
	ItemLength        int            `json:"item_length"`              //单个理解的最大长度
	Content           any            `json:"content,omitempty"`        //理解的内容
	AIContent         any            `json:"ai_content,omitempty"`     //AI理解的内容
	ContentType       int            `json:"content_type"`             //理解内容的类型
	ContentErrors     ContentError   `json:"content_errors,omitempty"` //包含的错误
	Note              string         `json:"note"`                     //悬浮提示
	Error             string         `json:"error,omitempty"`          //理解维度报错

	ListErr string `json:"-"` //列表错误
}

type ContentError map[string]string

// ColumnComment  信息项详情
type ColumnComment struct {
	ID         models.ModelID `json:"id" example:"1"`             //字段ID
	ColumnName string         `json:"column_name" example:"name"` //字段名称
	NameCN     string         `json:"name_cn" example:"字段中文名称"`   //字段中文名称
	DataFormat int32          `json:"data_format" example:"1"`    //字段类型
	Comment    string         `json:"comment"`                    //字段注释理解
	AIComment  string         `json:"ai_comment"`                 //AI生成的理解
	Error      string         `json:"error,omitempty"`            //检查生成的错误
	Sync       bool           `json:"sync" example:"false"`       //是否同步到数据资产中心的字段理解
}

type Choice struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//endregion
