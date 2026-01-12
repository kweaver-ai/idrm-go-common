package impl

import (
	"encoding/json"
	"strconv"

	"github.com/kweaver-ai/idrm-go-common/rest/points_management"
	"github.com/kweaver-ai/idrm-go-common/workflow/mq/common"
)

const (
	PointsEventTopic = "af.points.event"
)

type PointsEventPubRepoImpl struct {
	kafkaClient common.MQInterface
}

// 创建积分事件发布者
// 依赖注入: kafkaClient
// kafka客户端 路径: github.com/kweaver-ai/idrm-go-common/workflow/mq/kafka/kafka NewKafka
func NewPointsEventPubRepoImpl(kafkaClient common.MQInterface) points_management.PointsEventPubRepo {
	return &PointsEventPubRepoImpl{
		kafkaClient: kafkaClient,
	}
}

// 发布目录反馈事件
func (p *PointsEventPubRepoImpl) PublishCatalogFeedbackEvent(event *points_management.CatalogFeedbackEvent) {
	catalogRatingMsg := points_management.PointsEventPub{
		Type:             points_management.CatalogRating,
		PointObject:      event.DepartmentID,
		Score:            strconv.Itoa(event.Rating),
		PointsObjectName: event.DepartmentName,
	}
	ccatalogRatingMsgBytes, err := json.Marshal(catalogRatingMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), ccatalogRatingMsgBytes)
	catalogFeedbackMsg := points_management.PointsEventPub{
		Type:             points_management.CatalogFeedback,
		PointObject:      event.DepartmentID,
		PointsObjectName: event.DepartmentName,
	}
	catalogFeedbackMsgBytes, err := json.Marshal(catalogFeedbackMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), catalogFeedbackMsgBytes)
}

// 发布共享申请成效提交反馈事件
func (p *PointsEventPubRepoImpl) PublishShareApplicationFeedbackEvent(event *points_management.ShareApplicationFeedbackEvent) {
	shareApplicationFeedbackMsg := points_management.PointsEventPub{
		Type:             points_management.ShareApplicationFeedback,
		PointObject:      event.FeedbackUserID,
		PointsObjectName: event.FeedbackUserName,
	}
	shareApplicationFeedbackMsgBytes, err := json.Marshal(shareApplicationFeedbackMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), shareApplicationFeedbackMsgBytes)
}

// 发布数据归集任务完成事件
func (p *PointsEventPubRepoImpl) PublishDataAggregationCompleteEvent(event *points_management.DataAggregationCompleteEvent) {
	dataAggregationCompleteMsg := points_management.PointsEventPub{
		Type:             points_management.DataAggregationComplete,
		PointObject:      event.TaskExecutorID,
		PointsObjectName: event.TaskExecutorName,
	}
	dataAggregationCompleteMsgBytes, err := json.Marshal(dataAggregationCompleteMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), dataAggregationCompleteMsgBytes)
}

// 发布数据归集任务发布事件
func (p *PointsEventPubRepoImpl) PublishDataAggregationReleaseEvent(event *points_management.DataAggregationReleaseEvent) {
	dataAggregationReleaseMsg := points_management.PointsEventPub{
		Type:             points_management.DataAggregationRelease,
		PointObject:      event.TaskPublisherID,
		PointsObjectName: event.TaskPublisherName,
	}
	dataAggregationReleaseMsgBytes, err := json.Marshal(dataAggregationReleaseMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), dataAggregationReleaseMsgBytes)
}

// 发布供需申请提交目录事件
func (p *PointsEventPubRepoImpl) PublishSupplyAndDemandApplicationSubmisionDirectoryEvent(event *points_management.SupplyAndDemandApplicationSubmisionDirectoryEvent) {
	supplyAndDemandApplicationSubmisionDirectoryMsg := points_management.PointsEventPub{
		Type:             points_management.SupplyAndDemandApplicationSubmisionDirectory,
		PointObject:      event.DepartmentID,
		PointsObjectName: event.DepartmentName,
	}
	supplyAndDemandApplicationSubmisionDirectoryMsgBytes, err := json.Marshal(supplyAndDemandApplicationSubmisionDirectoryMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), supplyAndDemandApplicationSubmisionDirectoryMsgBytes)
}

// 发布共享申请提供资源事件
func (p *PointsEventPubRepoImpl) PublishShareApplicationSubmissionResourceEvent(event *points_management.ShareApplicationSubmissionResourceEvent) {
	shareApplicationSubmissionResourceMsg := points_management.PointsEventPub{
		Type:             points_management.ShareApplicationSubmissionResource,
		PointObject:      event.ResourceProviderID,
		PointsObjectName: event.ResourceProviderName,
	}
	shareApplicationSubmissionResourceMsgBytes, err := json.Marshal(shareApplicationSubmissionResourceMsg)
	if err != nil {
		return
	}
	p.kafkaClient.Produce(PointsEventTopic, []byte(``), shareApplicationSubmissionResourceMsgBytes)
}
