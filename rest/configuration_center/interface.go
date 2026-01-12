package configuration_center

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"

	"github.com/kweaver-ai/idrm-go-common/access_control"
)

type Driven interface {
	AccessControlService //权限空值接口
	DataSourceService    // 数据源接口
	DepartmentService    // 部门接口
	LabelService         // 分类分级标签服务
	UserAndRoleService   // 用户和角色
	ApplicationService   //应用服务
	GlobalConfig         //全局通用配置
	DataDict             //字典校验方法
	GetInfoSystemsPrecision(ctx context.Context, ids []string, names []string) ([]*GetInfoSystemByIdsRes, error)
	GetThirdPartyAddr(ctx context.Context, name string) ([]*GetThirdPartyAddressRes, error)
	GetUsingType(ctx context.Context) (int, error)
	Generate(ctx context.Context, id string, count int) (*CodeList, error)
	GetBusinessMatters(ctx context.Context, ids []string) (res []*BusinessMattersObject, err error)
	GetBusinessMatterPage(ctx context.Context, req *GetBusinessMatterPageReq) (res *GetBusinessMatterPageRes, err error)
}

type ApplicationService interface {
	HasAccessPermissionApps(ctx context.Context) (*AppList, error)
	GetApplication(ctx context.Context, id string) (*Apps, error)                        // 获取应用
	GetAppSimpleInfo(ctx context.Context, ids []string) ([]*AppSimpleInfo, error)        //　批量查询app的简单信息，当前仅国开分支有该接口
	GetApplicationInternal(ctx context.Context, id string) (*Apps, error)                // 根据应用名称获取应用
	ListApplicationsByDeveloperID(ctx context.Context, id string) ([]Application, error) //获取应用列表，指定应用开发者
	GetAppsByAccountId(ctx context.Context, id string) (*Apps, error)
	GetAppsByAccountIDs(ctx context.Context, ids []string) ([]*Apps, error)
}

type AccessControlService interface {
	HasAccessPermission(ctx context.Context, uid string, accessType access_control.AccessType, resource access_control.Resource) (bool, error)
	GetProcessBindByAuditType(ctx context.Context, auditType *GetProcessBindByAuditTypeReq) (res *GetProcessBindByAuditTypeRes, err error)
	GetProcessBindByResourceId(ctx context.Context, id string) (res *GetProcessBindByAuditTypeRes, err error)
	DeleteProcessBindByAuditType(ctx context.Context, auditType *DeleteProcessBindByAuditTypeReq) (err error)
	HasAppsAccessPermission(ctx context.Context, uid string, resource string) (bool, error)
}

type DataSourceService interface {
	GetDataSourcePrecision(ctx context.Context, ids []string) ([]*DataSourcesPrecision, error)
	GetDataSourcesByType(ctx context.Context, dataBaseType string) ([]*DataSourcePage, error)
	GetDatasourcesByHuaAoID(ctx context.Context, huaAoID string) (*DataSourcePage, error)
	GetAllDataSources(ctx context.Context) ([]*DataSources, error)
}

type UserAndRoleService interface {
	GetUser(ctx context.Context, id string, opts GetUserOptions) (*User, error)
	GetUsers(ctx context.Context, ids []string) ([]*User, error)
	GetBaseUserByIds(ctx context.Context, ids []string) ([]*UserBase, error)
	GetUserInfo(ctx context.Context, userID string) (*UserRespItem, error)
	GetUserInfoSlice(ctx context.Context, userID ...string) (*base.PageResult[UserRespItem], error)
	GetUsersByDeptRoleID(ctx context.Context, deptID, roleID string) ([]*UserRespItem, error)
}

type LabelService interface {
	QueryDataGrade(ctx context.Context, id ...string) (map[string]*HierarchyTag, error)
	GetLabelById(ctx context.Context, id string) (*GetLabelByIdRes, error)
	GetLabelByIds(ctx context.Context, ids string) (*GetLabelByIdsRes, error)
	GetLabelByName(ctx context.Context, name string) (*GetLabelByIdRes, error)
}

type GlobalConfig interface {
	// GetGlobalConfig  获取配置中心全局配置通用接口
	GetGlobalConfig(ctx context.Context, key string) (string, error)
	GetGlobalSwitch(ctx context.Context, key string) (bool, error)
	// PutGlobalConfig  修改配置中心全局配置通用接口
	PutGlobalConfig(ctx context.Context, key, value string) error
	GetCssjjSwitch(ctx context.Context) (bool, error) // 获取长沙数据局开关配置
	GetThirdPartySwitch(ctx context.Context) (bool, error)
}

// DataDict  数据字典接口，2.0.9发布
type DataDict interface {
	//GetDictItemType 批量获取枚举值
	//使用方法：
	//values, err :=cc.GetDictItemType(ctx, DataDictQueryTypeSSZD，type1, type2)
	//...错误处理
	//values.Filter(type1,values1...) 过滤下有哪些是合法的值
	//values.Check(type2,values2...) 判断下values2里面的值是否都是type2的值
	GetDictItemType(ctx context.Context, queryType string, dictType ...string) (*GetDictItemTypeResp, error)
	BatchCheckNotExistTypeKey(ctx context.Context, req CheckDictTypeKeyReq) (*[]string, error)

	//GetDictItemPage 查询字典项分页
	GetDictItemPage(ctx context.Context, req *GetDictItemPageReq) (*GetDictItemPageRes, error)
	GetGradeLabel(ctx context.Context, req *GetGradeLabelReq) (*GetGradeLabelRes, error) //查询数据分级
}
