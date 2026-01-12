package data_application_service

import (
	"context"

	"github.com/google/uuid"

	v1 "github.com/kweaver-ai/idrm-go-common/api/data_application_service/v1"
)

type Driven interface {
	// Service 返回指定 ID 的接口服务
	Service(ctx context.Context, id string) (*v1.Service, error)
	QueryDomainServices(ctx context.Context, flag string, isOperator bool, id ...string) (*QueryDomainServicesResp, error)
	QueryDomainApplicationServiceCountMap(ctx context.Context, flag string, isOperator bool, id ...string) (map[string]int64, error)
	// ServiceOwnerID 返回指定接口服务的 OwnerID
	ServiceOwnerID(ctx context.Context, id string) (string, error)
	GetServicesDataView(ctx context.Context, serviceId string) (*GetServicesDataViewRes, error)  //接口关联的视图
	GetDataViewServices(ctx context.Context, dataViewId string) (*GetDataViewServicesRes, error) //视图关联的接口
	InternalGetServiceDetail(ctx context.Context, id string) (*v1.Service, error)
	// 批量发布、上线接口服务
	InternalBatchPublishAndOnline(ctx context.Context, batch *v1.BatchPublishAndOnline) error
	GetSubServiceSimple(ctx context.Context, id string) (*SubService, error)
	UserServiceAuth(ctx context.Context, userID string, serviceID ...string) ([]string, error)
	GeSubServiceByServices(ctx context.Context, ids []string) (map[string][]string, error)
	// 通过接口ID列表获取接口列表
	InternalGetServicesByIDs(ctx context.Context, ids []string) (*ArrayResult[v1.Service], error)
	// 通过接口ID与网关同步
	InternalSyncServicesToGateway(ctx context.Context, id string) (SyncServicesToGatewayRes, error)
}

// region QueryDomainServices

const (
	QueryFlagAll   = "all"
	QueryFlagCount = "count"
	QueryFlagTotal = "total"
)

type QueryDomainServicesArgs struct {
	Flag       string   `json:"flag"`        //如果是all, 返回所有的数量；如果是count, 返回下面数组的数量,  如果是total ，只返回总的数量即可
	IsOperator bool     `json:"is_operator"` //如果为true，表示该用户是数据运营角色或者数据开发角色，这时展示所有的视图数据
	ID         []string `json:"id"`          //业务域，业务对象ID
}

type QueryDomainServicesResp struct {
	Total       int64                   `json:"total"`
	RelationNum []DomainServiceRelation `json:"relation_num"`
}

type DomainServiceRelation struct {
	SubjectDomainID    string `json:"subject_domain_id"` //业务域，业务对象ID
	RelationServiceNum int64  `json:"relation_service_num"`
}

// endregion

type ArrayResult[T any] struct {
	Entries []*T `json:"entries" binding:"omitempty"` // 对象列表
}

//region GetServicesDataView

type GetServicesDataViewRes struct {
	DataViewId string `json:"data_view_id"`
	ArrayResult[ServicesGetByDataViewId]
}

//endregion
//region GetDataViewServices

type GetDataViewServicesRes struct {
	ArrayResult[ServicesGetByDataViewId]
}

type ServicesGetByDataViewId struct {
	ServiceID   string `json:"service_id"`   // 接口ID
	ServiceCode string `json:"service_code"` // 接口编码
	ServiceName string `json:"service_name"` // 接口名称
}

//endregion

// SubService 子视图
type SubService struct {
	// ID
	ID uuid.UUID `json:"id,omitempty" path:"id"`
	// 名称
	Name string `json:"name,omitempty"`
	// 子视图所属逻辑视图的 ID
	ServiceID uuid.UUID `json:"service_id,omitempty"`
	//  授权范围, 可能是视图ID，可能是行列规则
	AuthScopeID uuid.UUID `json:"auth_scope_id,omitempty"`
}

//region SyncServicesToGateway

type SyncServicesToGatewayRes struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

//endregion
