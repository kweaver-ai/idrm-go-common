package callback

import "gorm.io/gorm"

const (
	LineageTransportName = "lineage"
)

const (
	LineageEntityTypeTable     string = "table"     //血缘的数据表实体类
	LineageEntityTypeField     string = "column"    //血缘的数据表字段
	LineageEntityTypeIndicator string = "indicator" //血缘的指标
	LineageEntityTypeDolphin   string = "dolphin"   //血缘的加工数据
)

const (
	BusinessRelationGraph            = "business-relation-graph"              //业务架构知识图谱名称
	CognitiveSearchDataResourceGraph = "cognitive-search-data-resource-graph" //认知搜索图谱_数据资源版
	CognitiveSearchDataCatalogGraph  = "cognitive-search-data-catalog-graph"  //认知搜索图谱_数据目录版
	SmartRecommendationGraph         = "smart-recommendation-graph"           //AF智能推荐场景图谱
)

var defaultDatabaseCallback *DatabaseCallback

func Init(db *gorm.DB) {
	defaultDatabaseCallback = NewDatabaseCallback(db)
	defaultDatabaseCallback.RegisterCallback()
	go defaultDatabaseCallback.Run()
}

func Register(callbackModel CallbackModel, transportName string, transports Transport) {
	defaultDatabaseCallback.transporter.Register(callbackModel,
		TransportModule{
			Name:      transportName,
			Transport: transports,
		})
}

func RegisterByTransport(transports Transport, transportName string, callbackModels ...CallbackModel) {
	for _, callbackModel := range callbackModels {
		Register(callbackModel, transportName, transports)
	}
}
