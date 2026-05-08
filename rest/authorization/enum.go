package authorization

import (
	"strings"

	"github.com/kweaver-ai/idrm-go-frame/core/enum"
)

func JoinDisplay[T any](operations []string) string {
	names := make([]string, 0)
	for _, operation := range operations {
		obj := enum.GetObj[T](operation)
		names = append(names, obj.Display)
	}
	return strings.Join(names, ",")
}

type ResourceTypeEnum enum.Object

var (
	ResourceTypeEnumDataView         = enum.New[ResourceTypeEnum](1, "data_view", "数据视图")
	ResourceTypeEnumKnowledgeNetwork = enum.New[ResourceTypeEnum](2, "knowledge_network", "业务知识网络")
)

type ViewOperationEnum enum.Object

var (
	ViewOperationEnumDataQuery     = enum.New[ViewOperationEnum](1, "data_query", "数据查询")
	ViewOperationEnumViewDetail    = enum.New[ViewOperationEnum](2, "view_detail", "查看")
	ViewOperationEnumModify        = enum.New[ViewOperationEnum](3, "modify", "修改")
	ViewOperationEnumDelete        = enum.New[ViewOperationEnum](4, "delete", "删除")
	ViewOperationEnumCreate        = enum.New[ViewOperationEnum](5, "create", "新建")
	ViewOperationEnumAuthorize     = enum.New[ViewOperationEnum](6, "authorize", "权限管理")
	ViewOperationEnumImport        = enum.New[ViewOperationEnum](7, "import", "导入")
	ViewOperationEnumExport        = enum.New[ViewOperationEnum](8, "export", "导出")
	ViewOperationEnumRuleManage    = enum.New[ViewOperationEnum](9, "rule_manage", "行列规则管理")
	ViewOperationEnumRuleAuthorize = enum.New[ViewOperationEnum](10, "rule_authorize", "行列规则授权")
)

type KNOperationEnum enum.Object

var (
	KNOperationEnumViewDetail = enum.New[KNOperationEnum](1, "view_detail", "查看")
	KNOperationEnumCreate     = enum.New[KNOperationEnum](2, "create", "新建")
	KNOperationEnumModify     = enum.New[KNOperationEnum](3, "modify", "编辑")
	KNOperationEnumDelete     = enum.New[KNOperationEnum](4, "delete", "删除")
	KNOperationEnumDataQuery  = enum.New[KNOperationEnum](5, "data_query", "数据查询")
	KNOperationEnumAuthorize  = enum.New[KNOperationEnum](6, "authorize", "权限管理")
	KNOperationEnumImport     = enum.New[KNOperationEnum](7, "import", "导入")
	KNOperationEnumExport     = enum.New[KNOperationEnum](8, "export", "导出")
	KNOperationEnumTaskManage = enum.New[KNOperationEnum](9, "task_manage", "任务管理")
)

type AccessorTypeEnum enum.Object

var (
	AccessorTypeEnumUser       = enum.New[AccessorTypeEnum](1, "user", "用户")
	AccessorTypeEnumDepartment = enum.New[AccessorTypeEnum](2, "department", "部门")
	AccessorTypeEnumGroup      = enum.New[AccessorTypeEnum](3, "group", "用户组")
	AccessorTypeEnumRole       = enum.New[AccessorTypeEnum](4, "role", "角色")
	AccessorTypeEnumApp        = enum.New[AccessorTypeEnum](5, "app", "应用")
)
