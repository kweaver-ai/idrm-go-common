package v1

import (
	"path"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// Object 定义资源
type Object struct {
	// 资源类型
	Type ObjectType `json:"object_type,omitempty"`
	// 资源 ID
	ID string `json:"object_id,omitempty"`
}

func (o *Object) Key() string {
	return path.Join(string(o.Type), o.ID)
}

// ObjectWithPermissions 定义资源及其权限
type ObjectWithPermissions struct {
	// 资源
	Object
	// 权限
	Permissions []Permission `json:"permissions,omitempty"`
	// 过期时间，空值代表永久生效
	ExpiredAt *meta_v1.Time `json:"expired_at,omitempty"`
}

// ObjectType 定义资源类型
type ObjectType string

const (
	// 主题域
	ObjectDomain ObjectType = "domain"
	// 数据目录
	ObjectDataCatalog ObjectType = "data_catalog"
	// 逻辑视图（数据表视图）
	ObjectDataView ObjectType = "data_view"
	// 接口
	ObjectAPI ObjectType = "api"
	// 行列规则（子视图）
	ObjectSubView ObjectType = "sub_view"
	// 指标
	ObjectIndicator ObjectType = "indicator"
	// 指标维度规则
	ObjectIndicatorDimensionalRule ObjectType = "indicator_dimensional_rule"
	ObjectTypePermissionResource              = "permission_resource" // 功能接口
	ObjectSubService               ObjectType = "sub_service"
)
