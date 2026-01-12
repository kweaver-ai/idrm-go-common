package standardization

import "context"

type Driven interface {
	GetDataElementDetailByCode(ctx context.Context, value ...string) ([]*DataResp, error)
	GetStandardMapByCode(ctx context.Context, value ...string) (map[string]*DataResp, error)
	GetDataElementDetailByID(ctx context.Context, value ...string) ([]*DataResp, error)
	GetStandardMapByID(ctx context.Context, value ...string) (map[string]*DataResp, error)
	GetStandardDict(ctx context.Context, ids []string) (data map[string]DictResp, err error) //码表ids批量获取码表信息
	// 删除标准文件
	DeleteStandFile(ctx context.Context, standFileID string) error
	GetStandardFiles(ctx context.Context, ids ...string) (data map[string]RuleResp, err error)
	GetStandardRule(ctx context.Context, ids []string) (data map[string]RuleResp, err error) //编码规则ids批量获取编码规则信息

	GetStandardList(ctx context.Context, req GetListReq) (*GetStandardListRes, error)
	GetCodeTableList(ctx context.Context, req GetListReq) (*GetCodeTableListRes, error)
}

const (
	IDType   = 1
	CodeType = 2
)

type GetDataElementDetailReq struct {
	IDS   string `query:"ids"`
	Codes string `query:"codes"`
}

type DataResp struct {
	ID            string `json:"id"`             // 标准id
	Code          string `json:"code"`           // 标准code
	NameCn        string `json:"name_cn"`        // 标准中文名
	NameEn        string `json:"name_en"`        // 标准英文名
	DataType      int    `json:"data_type"`      // 数据类型
	DataTypeName  string `json:"data_type_name"` // 数据类型名称
	DataLength    int    `json:"data_length"`    // 数据长度
	DataPrecision *int   `json:"data_precision"` // 数据精度
	DataRange     string `json:"data_range"`     // 值域
	StdType       int    `json:"std_type"`       // 制定依据
	StdTypeName   string `json:"std_type_name"`
	State         string `json:"state"`
	Deleted       bool   `json:"deleted"`
	DictID        string `json:"dict_id"`      // 码表id
	DictNameCN    string `json:"dict_name_cn"` // 码表中文名称
	DictNameEN    string `json:"dict_name_en"` // 码表英文名称
	DictState     string `json:"dict_state"`   // 码表状态
	DictDeleted   bool   `json:"dict_deleted"` // 码表状态
	RuleID        string `json:"rule_id"`      // 编码规则id
	RuleName      string `json:"rule_name"`    // 编码规则中文名称

}

type StandardDictResp struct {
	Code        string     `json:"code"`
	Description string     `json:"description"`
	Data        []DictResp `json:"data"`
}

type DictResp struct {
	ID      string `json:"id"`
	NameZh  string `json:"ch_name"`
	NameEN  string `json:"en_name"`
	State   string `json:"state"`
	Deleted bool   `json:"deleted"`
}

type StandardRuleResp struct {
	Code        string     `json:"code"`
	Description string     `json:"description"`
	Data        []RuleResp `json:"data"`
}

type RuleResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Deleted bool   `json:"deleted"`
}
type GetListReq struct {
	CatalogID string `json:"catalog_id"`
	Keyword   string `json:"keyword"`
}
type GetStandardListRes struct {
	Code        string     `json:"code"`
	Description string     `json:"description"`
	TotalCount  int        `json:"total_count"`
	Solution    string     `json:"solution"`
	Data        []Standard `json:"data"`
}
type Standard struct {
	ID            string `json:"id"`             // 唯一标识、雪花算法
	Code          string `json:"code"`           // 关联标识、雪花算法
	ChName        string `json:"ch_name"`        // 中文名称
	EnName        string `json:"en_name"`        // 英文名称
	CatalogID     string `json:"catalog_id"`     // 目录关联标识，默认目录ID为11
	OrgType       int    `json:"org_type"`       // 组织类型
	State         string `json:"state"`          // 启用停用状态：enable-启用，disable-停用；枚举值：DISABLE, ENABLE, Unknown
	Version       int    `json:"version"`        // 版本号
	DisableReason string `json:"disable_reason"` // 停用理由
	Deleted       bool   `json:"deleted"`        // 是否删除标记：0-未删除，其他值-已删除
	UsedFlag      bool   `json:"used_flag"`      // 使用标记：是否被使用
	CreateTime    string `json:"create_time"`    // 创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreateUser    string `json:"create_user"`    // 创建用户（ID）
	UpdateTime    string `json:"update_time"`    // 修改时间，格式：yyyy-MM-dd HH:mm:ss
	UpdateUser    string `json:"update_user"`    // 修改用户（ID）
}
type GetCodeTableListRes struct {
	Code        string     `json:"code"`
	Description string     `json:"description"`
	TotalCount  int        `json:"total_count"` // 返回总条数
	Solution    string     `json:"solution"`    // 解决策略
	Data        []DataItem `json:"data"`        // 数据对象
}

type DataItem struct {
	ID            string `json:"id"`             // 唯一标识（雪花算法）
	Code          string `json:"code"`           // 关联标识（雪花算法）
	ChName        string `json:"ch_name"`        // 中文名称
	EnName        string `json:"en_name"`        // 英文名称
	CatalogID     string `json:"catalog_id"`     // 目录ID
	OrgType       int    `json:"org_type"`       // 组织类型
	State         string `json:"state"`          // 状态（enable/disable）
	Version       int    `json:"version"`        // 版本号
	DisableReason string `json:"disable_reason"` // 停用理由
	Deleted       bool   `json:"deleted"`        // 是否已删除
	UsedFlag      bool   `json:"used_flag"`      // 是否被使用
	CreateTime    string `json:"create_time"`    // 创建时间
	CreateUser    string `json:"create_user"`    // 创建用户
	UpdateTime    string `json:"update_time"`    // 更新时间
	UpdateUser    string `json:"update_user"`    // 更新用户
}
