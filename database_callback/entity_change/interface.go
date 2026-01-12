package entity_change

const (
	BusinessRelationGraph            = "business-relation-graph"              //业务架构知识图谱名称
	CognitiveSearchDataResourceGraph = "cognitive-search-data-resource-graph" //认知搜索图谱_数据资源版
	CognitiveSearchDataCatalogGraph  = "cognitive-search-data-catalog-graph"  //认知搜索图谱_数据目录版
	SmartRecommendationGraph         = "smart-recommendation-graph"           //AF智能推荐场景图谱
)

// MessageSender 发送消息的，一般是MQ，抽象成接口，以防各个项目不同的MQ实现
// 各个项目的实现应该将topic隐藏其中
type MessageSender interface {
	Send(body []byte) error
	DestGraph(modelName string) []string //给出model名称，返回涉及的图谱名称
}
