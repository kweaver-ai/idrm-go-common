package af_sailor

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type Driven interface {
	QueryRecLabelList(ctx context.Context, req *QueryRecommendLabelReq) (*SailorResp, error) // 获取标签推荐列表
	StandardConsistency(ctx context.Context, tables []*StandardConsistencyTableInfo) ([][]*StandardConsistencyRespRec, error)
	IndicatorConsistency(ctx context.Context, indicators []*IndicatorConsistency) ([][]*IndicatorConsistencyRespRec, error)
	QueryBusinessSubjectRec(ctx context.Context, recReq []*SailorBusinessSubjectRecReq) (*SailorSubjectRecResp, error)
}

type QueryRecommendLabelReq struct {
	RecommendLabelReq []*RecommendLabelReq `json:"query_items"  binding:"gte=0,lte=20,required,dive"` //批量推荐,最少一个数组，最大20个数组
}

type RecommendLabelReq struct {
	Name         string `json:"name" form:"name" binding:"required,VerifyXssString,max=128" example:"A表"`  // 名称
	Description  string `json:"desc" form:"desc" binding:"VerifyXssString,max=300" example:"用于信息系统建设的业务表"` // 描述
	CategoryId   string `json:"category_id" form:"category_id" binding:"omitempty,max=200" example:"1,2"`  //标签分类ID,同一个名称和描述的多个分类id用","分割
	RangeTypeKey string `json:"range_type" form:"range_type" binding:"omitempty,max=64" example:"3"`       // 应用范围类型值（字典项码值）
}

type QueryBatchRecommendLabelResp struct {
	RecommendLabelResp []*SailorDataResp `json:"items"` //批量标签列表
}

//type RecommendLabelResp struct {
//	ID           string `json:"id" form:"id" binding:"required,min=1,max=20" example:"1"`                          //标签ID
//	Name         string `json:"name" form:"name" binding:"required,min=1,max=50" example:"信息建设"`                   // 标签名称
//	CategoryName string `json:"category_name" form:"category_name" binding:"required,min=1,max=50" example:"信息系统"` // 标签分类名称
//	//Description string           `json:"description" form:"description" binding:"max=300" example:"用于信息系统建设的业务表的标签"` // 标签描述
//	RangeTypeKey string `json:"range_type"  example:"1"` // 应用范围类型值（字典项码值）
//}

type SailorResp struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data []*SailorDataResp `json:"data"`
}

type SailorDataResp struct {
	Name         string                `json:"name" form:"name" binding:"required,min=1,max=50" example:"信息建设"` // 标签名称
	RangeTypeKey string                `json:"range_type" form:"range_type"  example:"1"`                       // 应用范围类型值（字典项码值）
	Rec          []*SailorDataListResp `json:"rec"`                                                             //推荐结果列表
}

type SailorDataListResp struct {
	ID           string `json:"id" form:"id" binding:"required,min=1,max=20" example:"1"`                          //标签ID
	Name         string `json:"name" form:"name" binding:"required,min=1,max=50" example:"信息建设"`                   // 标签名称
	CategoryName string `json:"category_name" form:"category_name" binding:"required,min=1,max=50" example:"信息系统"` // 标签分类名称
	RangeTypeKey string `json:"range_type" form:"range_type" example:"1"`                                          // 应用范围类型值（字典项码值）
	//Description string   `json:"description" form:"description" binding:"max=300" example:"用于信息系统建设的业务表的标签"` // 标签描述
}

//一致性诊断

type CommonConsistencyReq struct {
	Query any `json:"query"  binding:"required"`
}

type SailorCommonResp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// SailorCommonRecBody 标准一致性诊断结果
type SailorCommonRecBody[T any] struct {
	Rate   string `json:"rate"`
	Reason string `json:"reason"`
	Rec    []T    `json:"rec"`
}

// StandardConsistencyTableInfo  标准一致性请求参数
type StandardConsistencyTableInfo struct {
	TableId     string                         `json:"table_id"`
	TableName   string                         `json:"table_name"`
	TableDesc   string                         `json:"table_desc"`
	ProcessInfo ProcessInfo                    `json:"-"`
	Fields      []StandardConsistencyFieldInfo `json:"fields"`
}

type ProcessInfo struct {
	TableId           string `json:"table_id"`
	TableName         string `json:"table_name"`
	ProcessID         string `json:"process_id"`
	ProcessName       string `json:"process_name"`
	BusinessModelID   string `json:"business_model_id"`
	BusinessModelName string `json:"business_model_name"`
}

type StandardConsistencyFieldInfo struct {
	FieldId      string `json:"field_id"`
	FieldName    string `json:"field_name"`
	FieldDesc    string `json:"field_desc"`
	StandardId   string `json:"standard_id"`
	StandardName string `json:"standard_name"`
	StandardType string `json:"standard_type"`
}

// StandardConsistencyRespRec  标准一致性的请求结果
type StandardConsistencyRespRec struct {
	Correct      bool              `json:"correct"`
	Group        []base.IDNameResp `json:"group"`
	StandardName string            `json:"standard_name"`
	StandardType string            `json:"standard_type"`
}

// IndicatorConsistency 指标一致性
type IndicatorConsistency struct {
	IndicatorId        string `json:"indicator_id"`
	IndicatorName      string `json:"indicator_name"`
	IndicatorDesc      string `json:"indicator_desc"`
	IndicatorFormula   string `json:"indicator_formula"`
	IndicatorUnit      string `json:"indicator_unit"`
	IndicatorCycle     string `json:"indicator_cycle"`
	IndicatorCaliber   string `json:"indicator_caliber"`
	BusinessDomainID   string `json:"business_domain_id"`
	BusinessDomainName string `json:"business_domain_name"`
}

type IndicatorConsistencyRespRec struct {
	Correct            bool              `json:"correct"`
	Group              []base.IDNameResp `json:"group"`
	BusinessDomainID   string            `json:"business_domain_id"`
	BusinessDomainName string            `json:"business_domain_name"`
}

// 业务对象识别

type SailorSubjectRecReq struct {
	Query []*SailorBusinessSubjectRecReq `json:"query"  binding:"required,gte=0,dive"` //业务对象识别请求参数
}
type SailorBusinessSubjectRecReq struct {
	TableId                    string                       `json:"table_id"  binding:"omitempty" `         // 业务表ID
	TableName                  string                       `json:"table_name"  binding:"omitempty"`        //业务表名称
	TableDesc                  string                       `json:"field_desc"  binding:"omitempty"`        //业务表描述
	SailorTableFieldsRecReq    []SailorTableFieldsRecReq    `json:"fields" binding:"required,gte=0,dive"`   //字段列表
	SailorSubjectsFieldsRecReq []SailorSubjectsFieldsRecReq `json:"subjects" binding:"required,gte=0,dive"` //逻辑实体属性列表
}

type SailorTableFieldsRecReq struct {
	FieldId    string `json:"field_id" binding:"required" example:"1"`        // 字段ID
	FieldName  string `json:"field_name" binding:"required" example:"字段名称"`   //字段名称
	FieldDesc  string `json:"field_desc"  binding:"omitempty" example:"字段描述"` //字段描述
	StandardId string `json:"standard_id"  binding:"omitempty" example:"1"`   //标准ID
}

type SailorSubjectsFieldsRecReq struct {
	SubjectId   string `json:"subject_id" binding:"required" example:"1"`           // 属性ID
	SubjectName string `json:"subject_name" binding:"required" example:"属性名称"`      //属性名称
	SubjectPath string `json:"subject_path"  binding:"omitempty" example:"主题域层级信息"` //主题域层级信息
}

type SailorBusinessSubjectFieldsRecResp struct {
	FieldId     string `json:"field_id" binding:"required" example:"1"`            // 字段ID
	FieldName   string `json:"field_name" binding:"required" example:"字段名称"`       //字段名称
	SubjectId   string `json:"subject_id" binding:"required" example:"1"`          //逻辑实体属性ID
	SubjectName string `json:"subject_name" binding:"required" example:"逻辑实体属性名称"` //逻辑实体属性名称
	//RecScore    string `json:"score" binding:"omitempty" example:"80"`             //召回+排序阶段返回的检索得分
	//RecReason   string `json:"reason" binding:"omitempty" example:"大模型生成的推荐理由"`    //大模型生成的推荐理由，当前版本都是空
}

type SailorBusinessSubjectRecResp struct {
	TableName string                                `json:"table_name"  binding:"required" example:"业务表名称"` //业务标准表名称
	Rec       []*SailorBusinessSubjectFieldsRecResp `json:"rec"`                                            //字段与逻辑实体属性的对齐列表
}

type SailorSubjectRecResp struct {
	SailorBusinessSubjectRecResp []*SailorBusinessSubjectRecResp `json:"items"` //字段与逻辑实体属性的对齐列表
}
