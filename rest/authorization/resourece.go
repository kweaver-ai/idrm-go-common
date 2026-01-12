package authorization

import "context"

type Resource interface {
	//SetResource 设置资源类型，id是资源的唯一ID，支持自定义
	SetResource(ctx context.Context, id string, content *ResourceConfig) error
	//GetResource 获取资源详情，id是资源的唯一ID
	GetResource(ctx context.Context, id string) (*ResourceConfig, error)
	//DeleteResource 删除资源
	DeleteResource(ctx context.Context, id string) error
	//GetResourceTypeOperations 获取资源类型操作
	GetResourceTypeOperations(ctx context.Context, req *GetResourceTypeOperationsArgs) ([]*ResourceOperations, error)
	//GetResourceOperations 获取资源操作
	GetResourceOperations(ctx context.Context, req *GetResourceOperationsArgs) ([]*ResourceOperations, error)
}

type ResourceConfig struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	InstanceUrl string              `json:"instance_url"`
	DataStruct  string              `json:"data_struct"`
	Operation   []ResourceOperation `json:"operation"`
}

type ResourceOperation struct {
	ID          string                  `json:"id"`
	Name        []ResourceOperationName `json:"name"`
	Description string                  `json:"description"`
	Scope       []string                `json:"scope"`
}

type ResourceOperationName struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// ResourceTypeScope 资源信息
type ResourceTypeScope struct {
	Unlimited bool            `json:"unlimited"` //是否无限制
	Types     []*ResourceType `json:"types"`     //资源类型范围
}

type ResourceType struct {
	ID          string    `json:"id"`           //资源类型唯一标识
	Name        string    `json:"name"`         //资源类型名称
	InstanceUrl string    `json:"instance_url"` //资源实例URL
	DataStruct  string    `json:"data_struct"`  //资源类型数据结构
	Operation   Operation `json:"operation"`    //资源操作
}

type Operation struct {
	Type     []*OperationObject `json:"type"`     //资源类型操作
	Instance []*OperationObject `json:"instance"` //资源实例操作
}

type OperationObject struct {
	ID          string           `json:"id"`                   //操作唯一ID
	Name        string           `json:"name"`                 //操作名称
	Description string           `json:"description"`          //操作描述
	Obligations []ObligationItem `json:"obligation,omitempty"` //操作义务
}

// region GetResourceOperations

type GetResourceOperationsArgs struct {
	Method    string           `json:"method"`    //方法  GET，必填
	Accessor  Accessor         `json:"accessor"`  //访问者
	Resources []ResourceObject `json:"resources"` //资源列表, 必填
}

type ResourceOperations struct {
	ID        string   `json:"id"`        //资源唯一标识
	Operation []string `json:"operation"` //操作
}

//endregion

//region ResourceTypeOperations

type GetResourceTypeOperationsArgs struct {
	Method        string   `json:"method"`         //方法, GET
	ResourceTypes []string `json:"resource_types"` //资源类型数组
}

//endregion
