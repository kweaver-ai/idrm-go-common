package points_management

type PointsEventPubRepo interface {
	// 发布目录反馈事件
	PublishCatalogFeedbackEvent(event *CatalogFeedbackEvent)
	// 发布共享申请成效提交反馈事件
	PublishShareApplicationFeedbackEvent(event *ShareApplicationFeedbackEvent)
	// 发布数据归集任务完成事件
	PublishDataAggregationCompleteEvent(event *DataAggregationCompleteEvent)
	// 发布数据归集任务发布事件
	PublishDataAggregationReleaseEvent(event *DataAggregationReleaseEvent)
	// 发布供需申请提交目录事件
	PublishSupplyAndDemandApplicationSubmisionDirectoryEvent(event *SupplyAndDemandApplicationSubmisionDirectoryEvent)
	// 发布共享申请提供资源事件
	PublishShareApplicationSubmissionResourceEvent(event *ShareApplicationSubmissionResourceEvent)
}

// 积分事件code
const (
	// 目录反馈获取评分: catalog_rating
	CatalogRating string = "catalog_rating"
	// 目录反馈提交反馈: catalog_feedback
	CatalogFeedback string = "catalog_feedback"
	// 共享申请成效提交反馈: share_application_feedback
	ShareApplicationFeedback string = "share_application_feedback"
	// 数据归集任务完成: data_aggregation_complete
	DataAggregationComplete string = "data_aggregation_complete"
	// 数据归集任务发布: data_aggregation_release
	DataAggregationRelease string = "data_aggregation_release"
	// 供需申请提交目录: supply_and_demand_application_submission_directory
	SupplyAndDemandApplicationSubmisionDirectory string = "supply_and_demand_application_submission_directory"
	// 共享申请提供资源: share_application_submission_resource
	ShareApplicationSubmissionResource string = "share_application_submission_resource"
)

// 目录反馈事件
type CatalogFeedbackEvent struct {
	// 目录所属部门ID
	DepartmentID string `json:"department_id"`
	// 目录所属部门name
	DepartmentName string `json:"department_name"`
	// 目录评分
	Rating int `json:"rating"`
	// 反馈人ID
	FeedbackUserID string `json:"feedback_user_id"`
	// 反馈人name
	FeedbackUserName string `json:"feedback_user_name"`
}

// 共享申请成效提交反馈事件
type ShareApplicationFeedbackEvent struct {
	// 反馈人ID
	FeedbackUserID string `json:"feedback_user_id"`
	// 反馈人name
	FeedbackUserName string `json:"feedback_user_name"`
}

// 数据归集任务完成事件
type DataAggregationCompleteEvent struct {
	// 任务执行人ID
	TaskExecutorID string `json:"task_executor_id"`
	// 任务执行人name
	TaskExecutorName string `json:"task_executor_name"`
}

// 数据归集任务发布事件
type DataAggregationReleaseEvent struct {
	// 任务发布人ID
	TaskPublisherID string `json:"task_publisher_id"`
	// 任务发布人name
	TaskPublisherName string `json:"task_publisher_name"`
}

// 供需申请提交目录事件
type SupplyAndDemandApplicationSubmisionDirectoryEvent struct {
	// 目录所属部门ID
	DepartmentID string `json:"department_id"`
	// 目录所属部门name
	DepartmentName string `json:"department_name"`
}

// 共享申请提供资源事件
type ShareApplicationSubmissionResourceEvent struct {
	// 资源提供人ID
	ResourceProviderID string `json:"resource_provider_id"`
	// 资源提供人name
	ResourceProviderName string `json:"resource_provider_name"`
}

type PointsEventPub struct {
	Type             string `json:"type"`
	PointObject      string `json:"point_object"`
	Score            string `json:"score"`
	PointsObjectName string `json:"points_object_name"`
}
