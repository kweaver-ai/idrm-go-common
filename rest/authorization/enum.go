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
