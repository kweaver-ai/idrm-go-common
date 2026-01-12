package authorization

import (
	"context"
	"fmt"
	"time"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/samber/lo"
)

type Authorization interface {
	AuthConfig  //获取配置， 配置什么返回什么
	AuthEnforce //授权角色，返回实际权限的最终结果
	AuthExtend  //扩展功能
}

// AuthConfig 策略配置相关，返回的是配置，不是最终的结果
type AuthConfig interface {
	//CreatePolicy  新建策略
	CreatePolicy(ctx context.Context, req []*CreatePolicyReq) (*CreatePolicyResp, error)
	//UpdatePolicy  更新策略
	UpdatePolicy(ctx context.Context, ids string, req []*UpdatePolicyReq) error
	//DeletePolicy  删除策略
	DeletePolicy(ctx context.Context, ids string) error
	//GetAccessorPolicy 获取访问者策略配置
	GetAccessorPolicy(ctx context.Context, req *AccessorPolicyArgs) (*base.PageResult[AccessorPolicy], error)
	//GetResourcePolicy 获取资源实例策略配置
	GetResourcePolicy(ctx context.Context, req *GetResourcePolicyReq) (*base.PageResult[ResourcePolicy], error)
}

// AuthEnforce 策略决策
type AuthEnforce interface {
	//OperationCheck 单个决策, 返回值，true代表允许，false代表拒绝
	OperationCheck(ctx context.Context, req *OperationCheckArgs) (*OperationCheckResult, error)
	//ResourceFilter  资源过滤
	ResourceFilter(ctx context.Context, req *ResourceFilterArgs) ([]*ResourceFilterResp, error)
	//ResourceList 资源列表,可以查询用户最终的资源
	ResourceList(ctx context.Context, req *ResourceListArgs) ([]*ResourceListRespItem, error)
}

type AuthExtend interface {
	MenuResourceEnforce(ctx context.Context, userID string, operation string, keys ...string) (policyEnforceEffect *MenuResourceEnforceResult, err error)
}

// region  GetAccessorPolicy

type AccessorPolicyArgs struct {
	AccessorID   string `query:"accessor_id"`   //访问者唯一标识, 必填
	AccessorType string `query:"accessor_type"` //访问者类型, user,role,app
	ResourceType string `query:"resource_type"` //资源类型, 该参数不传时，返回所有资源类型的策略配置。
	ResourceID   string `query:"resource_id"`   //资源实例唯一标识, 该参数传入时，resource_type必传, 返回指定资源实例的策略配置。
	Offset       int64  `query:"offset"`        //获取数据起始下标
	Limit        int64  `query:"limit"`         //获取数据量
	//返回的额外信息：
	//obligation_types：义务类型
	//obligations：义务
	Include []string `query:"include"`
}

type AccessorPolicy struct {
	ID        string          `json:"id"`         //策略唯一标识
	Resource  *ResourceObject `json:"resource"`   //资源
	Operation *AuthOperation  `json:"operation"`  //操作
	ExpiresAt time.Time       `json:"expires_at"` //到期时间(秒级)，RFC3339格式，UNIX TIME时间纪元(1970-01-01T08:00:00+08:00)表示永久有效
}

func (a *AccessorPolicy) Allowed() []string {
	if a.Operation == nil {
		return []string{}
	}
	return a.Operation.Allowed()
}

type ResourceObject struct {
	ID   string `json:"id"`   //资源唯一标识
	Name string `json:"name"` //资源名称, 仅作结果返回
	Type string `json:"type"` //资源类型
}

type AuthOperation struct {
	Allow []*OperationObject `json:"allow"` //允许操作
	Deny  []*OperationObject `json:"deny"`  //拒绝操作
}

func (o *AuthOperation) Allowed() []string {
	return lo.Times(len(o.Allow), func(i int) string {
		return o.Allow[i].ID
	})
}

//endregion

//region OperationCheck

type OperationCheckArgs struct {
	Method    string         `json:"method"`    //方法,GET,必填
	Accessor  Accessor       `json:"accessor"`  //访问者,必填
	Resource  ResourceObject `json:"resource"`  //资源信息,必填
	Operation []string       `json:"operation"` //检查的操作,必填
	Include   []string       `json:"include"`   //是否包含义务，operation_obligations
}

type OperationCheckResult struct {
	Result  bool                   `json:"result"`
	Include *ObligationIncludeBody `json:"include,omitempty"`
}

type ObligationIncludeBody struct {
	OperationObligations []OperationObligation `json:"operation_obligations"`
}

type OperationObligation struct {
	Operation   string           `json:"operation"`
	Obligations []ObligationItem `json:"obligations"`
}

type ObligationItem struct {
	Value  map[string]any `json:"value"`
	TypeId string         `json:"type_id"`
}

func (o *ObligationIncludeBody) OperationScope() string {
	if len(o.OperationObligations) <= 0 {
		return ""
	}
	if len(o.OperationObligations[0].Obligations) <= 0 {
		return ""
	}
	if len(o.OperationObligations[0].Obligations[0].Value) <= 0 {
		return ""
	}
	v, ok := o.OperationObligations[0].Obligations[0].Value["scope"]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

func (o *OperationCheckResult) OperationScope() string {
	if o.Include == nil {
		return ""
	}
	return o.Include.OperationScope()
}

//endregion

//region CreatePolicy

type CreatePolicyReq struct {
	Accessor  Accessor       `json:"accessor"`
	Resource  ResourceObject `json:"resource"`
	Operation AuthOperation  `json:"operation"`
	ExpiresAt string         `json:"expires_at"`
}

type Accessor struct {
	ID         string `json:"id"`                    //资源唯一标识
	Type       string `json:"type"`                  //资源类型
	Name       string `json:"name,omitempty"`        //资源名称
	Ip         string `json:"ip,omitempty"`          //访问ip
	ClientType string `json:"client_type,omitempty"` //终端类型
}

type CreatePolicyResp struct {
	Ids []string `json:"ids"`
}

//endregion

//region UpdatePolicy

type UpdatePolicyReq struct {
	Operation AuthOperation `json:"operation"`
	ExpiresAt string        `json:"expires_at"`
}

//endregion

//region DeletePolicy

type PathIDReq struct {
	Ids string `json:"ids" uri:"ids"`
}

//endregion

//region GetResourcePolicy

type GetResourcePolicyReq struct {
	ResourceID   string `query:"resource_id"`   //资源唯一标识
	ResourceType string `query:"resource_type"` //资源类型
	Offset       int64  `query:"offset"`        //获取数据起始下标
	Limit        int64  `query:"limit"`         //获取数据量
	//返回的额外信息：
	//obligation_types：义务类型
	//obligations：义务
	Include []string `query:"include"`
}

type ResourcePolicy struct {
	ID        string        `json:"id"`         //策略ID
	Accessor  Accessor      `json:"accessor"`   //访问者
	Operation AuthOperation `json:"operation"`  //操作
	ExpiresAt time.Time     `json:"expires_at"` //到期时间(秒级)
}

//endregion

//region MenuResourceEnforce

type MenuResourceEnforceResult struct {
	Scope  string `json:"scope"`
	Effect bool   `json:"effect"`
}

//endregion

//region  ResourceFilter

type ResourceFilterArgs struct {
	Method         string           `json:"method"`
	Accessor       Accessor         `json:"accessor"`
	Resources      []ResourceObject `json:"resources"`
	Operation      []string         `json:"operation"`
	AllowOperation bool             `json:"allow_operation"`
	Include        []string         `json:"include"`
}

type ResourceFilterResp struct {
	Id             string                 `json:"id"`
	AllowOperation []string               `json:"allow_operation"`
	Include        *ObligationIncludeBody `json:"include,omitempty"`
}

func (r *ResourceFilterResp) OperationScope() string {
	if r.Include == nil {
		return ""
	}
	return r.Include.OperationScope()
}

//endregion

//region  ResourceFilter

type ResourceListArgs struct {
	Method    string         `json:"method"`
	Accessor  Accessor       `json:"accessor"`
	Resource  ResourceObject `json:"resource"`
	Operation []string       `json:"operation"`
	Include   []string       `json:"include"`
}

type ResourceListRespItem struct {
	Id      string            `json:"id"`
	Include IncludeObligation `json:"include"`
}

type IncludeObligation struct {
	OperationObligations []ObligationWithOperation `json:"operation_obligations"`
}

type ObligationWithOperation struct {
	Operation   string           `json:"operation"`
	Obligations []ObligationItem `json:"obligations"`
}

//endregion
