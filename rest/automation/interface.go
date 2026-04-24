package automation

import "context"

// Driven 对 flow-automation 服务的 REST 包装。
type Driven interface {
	DagByName(ctx context.Context, name string) (*DagMeta, error)
	// DagList 获取流程元数据列表。
	DagList(ctx context.Context, req *DagListArgs) (*DagListResp, error)
	// DagDetail 获取指定流程定义详情。
	DagDetail(ctx context.Context, id string) (*DagDetailResp, error)
	// RunInstanceForm 以表单触发运行流程。
	RunInstanceForm(ctx context.Context, id string, body map[string]any) error
}

type DagListArgs struct {
	Keyword string `query:"keyword"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	SortBy  string `query:"sortby"`
	Order   string `query:"order"`
}

type DagListResp struct {
	Dags  []*DagMeta `json:"dags"`
	Limit int        `json:"limit"`
	Page  int        `json:"page"`
	Total int        `json:"total"`
}

type DagMeta struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Actions     []string `json:"actions"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
	Status      string   `json:"status"`
	UserID      string   `json:"userid"`
	Creator     string   `json:"creator"`
	Trigger     string   `json:"trigger"`
	VersionID   string   `json:"version_id"`
}

type DagDetailArgs struct {
	ID string `uri:"id"`
}

type DagDetailResp struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Status        string         `json:"status"`
	Steps         []*DagStep     `json:"steps"`
	CreatedAt     int64          `json:"created_at"`
	UpdatedAt     int64          `json:"updated_at"`
	Shortcuts     any            `json:"shortcuts"`
	Accessors     []*DagAccessor `json:"accessors"`
	Cron          string         `json:"cron"`
	Published     bool           `json:"published"`
	TriggerConfig map[string]any `json:"trigger_config"`
	UserID        string         `json:"userid"`
}

type DagStep struct {
	ID         string           `json:"id"`
	Title      string           `json:"title"`
	Operator   string           `json:"operator"`
	Parameters DagStepParameter `json:"parameters"`
}

type DagStepParameter struct {
	Fields []*DagFormField `json:"fields"`
}

type DagFormField struct {
	Description   DagFieldDescription `json:"description"`
	Key           string              `json:"key"`
	Name          string              `json:"name"`
	Required      bool                `json:"required"`
	Type          string              `json:"type"`
	Default       any                 `json:"default"`
	AllowedValues any                 `json:"allowed_values"`
}

type DagFieldDescription struct {
	Text string `json:"text"`
}

type DagAccessor struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type RunInstanceFormArgs struct {
	Data map[string]any `json:"data" param_type:"json"`
}
