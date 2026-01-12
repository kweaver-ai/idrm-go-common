package indicator_management

import (
	"context"
)

type Driven interface {
	// 获取指定 ID 的指标
	GetIndicator(ctx context.Context, id string) (*Indicator, error)
	QueryDomainIndicators(ctx context.Context, flag string, id ...string) (*QueryDomainIndicatorsResp, error)
	QueryDomainIndicatorCountMap(ctx context.Context, flag string, id ...string) (map[string]int64, error)
	UserIndicatorAuth(ctx context.Context, userID string, indicatorID ...string) ([]string, error)
}

// region GetIndicator

// 指标
//
// TODO: 补充缺少的字段
type Indicator struct {
	// ID
	ID string `json:"id,omitempty"`
	// 编码
	Code string `json:"code,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 指标类型
	IndicatorType IndicatorType `json:"indicator_type,omitempty"`
	// Deprecated: use Owners.OwnerID
	OwnerId string `json:"owner_id"`
	// Deprecated: use Owners.OwnerName``
	OwnerName  string `json:"owner_name"`
	Department struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"department"`
	Owners []DataApplicationServiceOwner `json:"owners,omitempty"`
}

type DataApplicationServiceOwner struct {
	OwnerID   string `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

// 指标类型
type IndicatorType string

const (
	// 原子指标
	IndicatorAtomic IndicatorType = "atomic"
	// 衍生指标
	IndicatorDerived IndicatorType = "derived"
	// 复合指标
	IndicatorComposite IndicatorType = "composite"
)

type GetIndicatorArgs struct {
	// 指标 ID，获取这个 ID 的指标
	ID string `uri:"id"`
}

// endregion

// region QueryDomainIndicators

const (
	QueryFlagAll   = "all"
	QueryFlagCount = "count"
	QueryFlagTotal = "total"
)

type QueryDomainIndicatorsArgs struct {
	Flag string   `json:"flag"` //如果是all, 返回所有的数量；如果是count, 返回下面数组的数量,  如果是total ，只返回总的数量即可
	ID   []string `json:"id"`   //业务域，业务对象ID
}

type QueryDomainIndicatorsResp struct {
	Total       int64                     `json:"total"`
	RelationNum []DomainIndicatorRelation `json:"relation_num"`
}

type DomainIndicatorRelation struct {
	SubjectDomainID string `json:"subject_domain_id"` //业务域，业务对象ID
	RelationCount   int64  `json:"relation_count"`    //关联的服务数量
}

// endregion
