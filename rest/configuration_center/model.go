package configuration_center

import "time"

type GetInfoSystemByIdsRes struct {
	ID           string `json:"id"`             // 信息系统业务id
	Name         string `json:"name"`           // 信息系统名称
	OnlineStatus string `json:"online_status"`  //信息系统的状态，草稿，建设中，上线，已停用
	Description  string `json:"description"`    // 信息系统描述
	CreatedAt    int64  `json:"created_at"`     // 创建时间
	CreatedByUID string `json:"created_by_uid"` // 创建用户ID
	UpdatedAt    int64  `json:"updated_at"`     // 更新时间
	UpdatedByUID string `json:"updated_by_uid"` // 更新用户ID
}

type DataSourcesPrecision struct {
	DataSourceID uint64 `json:"data_source_id"` // 数据源雪花id
	ID           string `json:"id"`             // 数据源业务id
	InfoSystemID string `json:"info_system_id"` // 信息系统id
	DepartmentID string `json:"department_id"`  // 归属部门ID
	Name         string `json:"name"`           // 数据源名称
	CatalogName  string `json:"catalog_name"`   // 数据源catalog名称
	Type         int32  `json:"type"`           // 数据库类型
	TypeName     string `json:"type_name"`      // 数据库类型名称
	Host         string `json:"host"`           // 连接地址
	Port         int32  `json:"port"`           // 端口
	Username     string `json:"username"`       // 用户名
	DatabaseName string `json:"database_name"`  // 数据库名称
	Schema       string `json:"schema"`         // 数据库模式
	SourceType   int32  `json:"source_type"`    // 数据源类型 1:记录型、2:分析型
	CreatedByUID string `json:"created_by_uid"` // 创建人id
	CreatedAt    int64  `json:"created_at"`     // 创建时间
	UpdatedByUID string `json:"updated_by_uid"` // 更新人id
	UpdatedAt    int64  `json:"updated_at"`     // 更新时间
	HuaAoId      string `json:"hua_ao_id"`      // 第三方数据源id
}
type DataSources struct {
	DataSourceID uint64    `json:"data_source_id"` // 数据源雪花id
	ID           string    `json:"id"`             // 数据源业务id
	InfoSystemID string    `json:"info_system_id"` // 信息系统id
	DepartmentID string    `json:"department_id"`  // 归属部门ID  可能没给值
	Name         string    `json:"name"`           // 数据源名称
	CatalogName  string    `json:"catalog_name"`   // 数据源catalog名称
	Type         int32     `json:"type"`           // 数据库类型
	TypeName     string    `json:"type_name"`      // 数据库类型名称
	Host         string    `json:"host"`           // 连接地址
	Port         int32     `json:"port"`           // 端口
	Username     string    `json:"username"`       // 用户名
	DatabaseName string    `json:"database_name"`  // 数据库名称
	Schema       string    `json:"schema"`         // 数据库模式
	SourceType   int32     `json:"source_type"`    // 数据源类型 1:记录型、2:分析型
	CreatedByUID string    `json:"created_by_uid"` // 创建人id
	CreatedAt    time.Time `json:"created_at"`     // 创建时间
	UpdatedByUID string    `json:"updated_by_uid"` // 更新人id
	UpdatedAt    time.Time `json:"updated_at"`     // 更新时间
}

type DataSourcePage struct {
	ID           string `json:"id"`             // 数据源id
	Name         string `json:"name"`           // 数据源名称
	CatalogName  string `json:"catalog_name"`   // 数据源catalog名称
	Type         string `json:"type"`           // 数据源类型
	SourceType   string `json:"source_type"`    // 数据源来源类型
	DatabaseName string `json:"database_name"`  // 数据库名称
	Schema       string `json:"schema"`         // 数据库模式
	UpdatedByUID string `json:"updated_by_uid"` // 修改人
	UpdatedAt    int64  `json:"updated_at"`     // 修改时间
	// 第三方数据源 ID
	HuaAoID string `json:"hua_ao_id,omitempty"`
}

// region QueryDataGrade

type QueryDataGradeReq struct {
	Ids string `uri:"ids"`
}

type QueryDataGradeResp struct {
	Entries []HierarchyTag
}

type HierarchyTag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	SortWeight  int    `json:"sort_weight"`  //排序值
	Description string `json:"description"`  // 标签描述
	LabelPath   string `json:"name_display"` //  标签路径
}

//endregion

// region ConfigurationCenterDriven

type GetLabelByIdRes struct {
	ID                  string  `json:"id"`                    // 标签id
	Name                string  `json:"name"`                  // 标签名称
	SortWeight          int     `json:"sort_weight"`           //排序值
	Description         string  `json:"description"`           // 标签描述
	LabelIcon           string  `json:"icon"`                  //  标签颜色
	LabelPath           string  `json:"name_display"`          //  标签路径
	SensitiveAttri      *string `json:"sensitive_attri"`       // 敏感属性预设
	SecretAttri         *string `json:"secret_attri"`          // 涉密属性预设
	ShareCondition      *string `json:"share_condition"`       // 共享条件：不共享，有条件共享，无条件共享
	DataProtectionQuery bool    `json:"data_protection_query"` // 数据保护查询开关, true开启
}

type GetLabelByIdsRes struct {
	Entries []*GetLabelByIdRes `json:"entries"` // 标签
}

//endregion

// region GetProcessBindByAuditType

type GetProcessBindByAuditTypeReq struct {
	AuditType string `json:"audit_type" uri:"audit_type" binding:"required,VerifyNameEn"` // 审核类型
}

type GetProcessBindByAuditTypeRes struct {
	ID          string `json:"id"`         // id
	AuditType   string `json:"audit_type"` // 审核类型
	ProcDefKey  string `json:"proc_def_key"`
	ServiceType string `json:"service_type"`
}

//endregion

// region DeleteProcessBindByAuditType

type DeleteProcessBindByAuditTypeReq struct {
	AuditType string `json:"audit_type" uri:"audit_type" binding:"required,VerifyNameEn"` // 审核类型
}

//endregion

// region thirdPart Address

type GetThirdPartyAddressRes struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

//endregion

//region 查询用户的角色

type Role struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	System int    `json:"system"`
}

//endregion

// region 应用授权信息

type Apps struct {
	ID             string      `json:"id" example:"f5600699-b4c8-443e-a37e-39e3fd5d2159"` // 应用ID
	Name           string      `json:"name" example:"name"`                               // 应用名称
	Description    string      `json:"description" example:"description"`                 // 应用描述
	PassID         string      `json:"pass_id" example:"passid"`                          // PassID
	Token          string      `json:"token" example:"token"`                             // Token
	InfoSystem     *SystemInfo `json:"info_systems"`                                      // 信息系统
	IpAddrs        []*IpAddr   `json:"ip_addr"`
	AccountName    string      `json:"account_name" example:"account_name"`                                                                              // 账号名称
	AccountID      string      `json:"account_id" example:"f5600699-b4c8-443e-a37e-39e3fd5d2159"`                                                        // 账号Id
	AuthorityScope []string    `json:"authority_scope" example:"demand_task,business_grooming,standardization,resource_management,configuration_center"` // 权限范围
	CreatedAt      int64       `json:"created_at" example:"1684301771000"`                                                                               // 创建时间
	CreatedName    string      `json:"creator_name" example:"创建人名称"`                                                                                     //创建人
	UpdatedAt      int64       `json:"updated_at" example:"1684301771000"`                                                                               // 更新时间
	UpdatedName    string      `json:"updater_name" example:"更新人名称"`                                                                                     //更新人

}

type AppSimpleInfo struct {
	ID   string `json:"id"`   // 应用ID
	Name string `json:"name"` // 应用名称
}

type SystemInfo struct {
	ID   string `json:"id"`   // 信息系统id，uuid
	Name string `json:"name"` // 信息系统名称
}
type IpAddr struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

// endregion
type AppList struct {
	Entries    []Apps `json:"entries"`
	TotalCount int    `json:"total_count"`
}

func (a *AppList) CheckAppIDExists(id string) bool {
	for _, v := range a.Entries {
		if v.ID == id {
			return true
		}
	}
	return false
}

// 应用，仅包含部分字段，需要时再补充缺少的字段
type Application struct {
	// ID
	ID string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
}

//region Generate 编码列表

type CodeList struct {
	// 编码列表
	Entries []string `json:"entries,omitempty"`
	// 编码的数量
	TotalCount int `json:"total_count,omitempty"`
}

const (
	GenerateCodeIdDataView        = "13daf448-d9c4-11ee-81aa-005056b4b3fc"
	GenerateCodeIdDataCatalog     = "28fa2073-2b5f-4ab5-9630-c73800fed3e5"
	GenerateCodeIdInfoCatalog     = "d6aaf704-91e5-438a-bb6b-c1e1b83d50c9"
	GenerateCodeIdApi             = "15d8b9f8-f87b-11ee-aeae-005056b4b3fc"
	GenerateCodeIdDataRequirement = "cef39a5e-dc4f-11ee-b798-005056b4b3fc"
	GenerateCodeIdFileResource    = "7b83283c-ecff-11ef-a99d-cad97a383659"
)

//endregion

// region   GetGlobalConfig

type GetConfigValueReq struct {
	Key string `query:"key"` // 配置表中对应的key
}

type GetConfigValueRes struct {
	Key   string `json:"key" example:"AISampleDataShow"` // 配置表中的key的值
	Value string `json:"value" example:"YES"`            // 配置表中的value的值
}

// endregion
// User 定义用户结构。仅包含用到的属性，需要其他属性时再补充。
type User struct {
	// 用户 ID
	ID string `json:"id,omitempty"`
	// 用户显示名
	Name string `json:"name,omitempty"`
	// 用户所属的部门。因为用户可以属于多个部门，所以是个 slice。
	ParentDeps []DepartmentPath `json:"parent_deps,omitempty"`
}

// User 定义用户结构。仅包含用到的属性，需要其他属性时再补充。
type UserBase struct {
	// 用户 ID
	ID string `json:"id,omitempty"`
	// 用户显示名
	Name string `json:"name,omitempty"`
}

// DepartmentPath 定义用户部门的层级结构，由高到低。
type DepartmentPath []Department

// Department 定义部门对象。type 属性仅支持 department，可以忽略。
type Department struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetUserOptions 定义获取用户的选项。
type GetUserOptions struct {
	// 需要获取的 User 的字段列表
	Fields []Field
}

// Field 定义对象的字段
type Field string

// 仅定义了需要用到的字段，需要其他字段时再补充。
const (
	// 用户显示名
	FieldName Field = "name"
	// 父部门信息
	FieldParentDeps Field = "parent_deps"
)

type ObjectUpdateFileReq struct {
	FileId   string `json:"file_id" binding:"required,uuid" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"` // 文件ID
	FileName string `json:"file_name" binding:"required" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"`    // 文件名称
}

type BusinessMattersObject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetBusinessMatterPageReq struct {
	Limit int `json:"limit"`
}

type GetBusinessMatterPageRes struct {
	Entries    []GetBusinessMatterPageEntry `json:"entries"`
	TotalCount int                          `json:"total_count"`
}

type GetBusinessMatterPageEntry struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	TypeKey         string `json:"type_key"`
	TypeValue       string `json:"type_value"`
	DepartmentID    string `json:"department_id"`
	DepartmentName  string `json:"department_name"`
	DepartmentPath  string `json:"department_path"`
	MaterialsNumber int    `json:"materials_number"`
}
