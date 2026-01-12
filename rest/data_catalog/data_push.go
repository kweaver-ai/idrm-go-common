package data_catalog

type DataPushCreateReq struct {
	Name                string `json:"name" binding:"required,TrimSpace,max=128"`         // 数据推送模型名称
	Description         string `json:"description" binding:"TrimSpace,omitempty,max=300"` // 描述
	ResponsiblePersonID string `json:"responsible_person_id" binding:"required,uuid"`     // 责任人ID
	Channel             int32  `json:"channel" binding:"required,oneof=1 2 3"`            // 数据提供方法，1web 2share_apply 3catalog_report  不填默认是web

	PushStatus *int32 `json:"push_status" binding:"omitempty,oneof=0 1 2"` // 推送状态，不填默认是2

	SourceCatalogID    string                 `json:"source_catalog_id" binding:"required"`            // 来源视图编目的数据目录ID，冗余字段
	TargetDatasourceID string                 `json:"target_datasource_id" binding:"required,uuid"`    // 数据源（目标端）
	TargetTableID      string                 `json:"target_table_id" binding:"omitempty,uuid"`        // 目标表ID，物理表的ID
	TargetTableExists  *bool                  `json:"target_table_exists" binding:"required"`          // 目标表在本次推送是否存在，0不存在，1存在
	TargetTableName    string                 `json:"target_table_name"  binding:"required,max=255"`   // 目标表名称
	SyncModelFields    []*SyncModelCheckField `json:"sync_model_fields" binding:"required,min=1,dive"` // 同步模型选定的字段
	FilterCondition    string                 `json:"filter_condition"   binding:"omitempty"`          // 过滤表达式，SQL后面的where条件

	TransmitMode       int32  `json:"transmit_mode" binding:"required,oneof=1 2"`            // 传输模式，1 增量 ; 2 全量
	IncrementField     string `json:"increment_field" binding:"TrimSpace,omitempty,max=128"` // 增量字段，当推送类型选择增量时，选一个字段作为增量字段，（技术名称）
	IncrementTimeStamp int64  `json:"increment_timestamp" binding:"omitempty"`               // 增量时间戳值，单位毫秒；当推送类型选择增量时，该字段必填
	PrimaryKey         string `json:"primary_key" binding:"omitempty"`                       // 主键，技术名称，当推送类型选择增量时，该字段必填

	ScheduleType  string `json:"schedule_type" binding:"required,oneof=ONCE PERIOD"` // 调度计划:once一次性,timely定时
	ScheduleTime  string `json:"schedule_time" binding:"omitempty,LocalDateTime"`    // 调度时间，格式 2006-01-02 15:04:05;  空：立即执行；非空：定时执行
	ScheduleStart string `json:"schedule_start" binding:"omitempty,LocalDate"`       // 计划开始日期, 格式 2006-01-02
	ScheduleEnd   string `json:"schedule_end" binding:"omitempty,LocalDate"`         // 计划结束日期, 格式 2006-01-02
	CrontabExpr   string `json:"crontab_expr" binding:"omitempty,cron"`              // linux crontab表达式, 5级
}

type SyncModelCheckField struct {
	SourceTechName string `json:"source_tech_name" binding:"required"` // 来源表字段的技术名称
	BusinessName   string `json:"business_name" binding:"required"`    // 来源字段的列业务名称，和目标字段公用
	TechnicalName  string `json:"technical_name" binding:"required"`   // 列技术名称
	PrimaryKey     bool   `json:"primary_key" `                        // 是不是主键
	DataType       string `json:"data_type"`                           // 数据类型
	DataLength     int32  `json:"data_length"`                         // 数据长度
	Precision      *int   `json:"precision"`                           // 数据精度
	Comment        string `json:"comment"`                             // 字段注释
	IsNullable     string `json:"is_nullable"`                         // 是否为空
}

type UpdateReq struct {
	ID                  string `json:"id"`
	Name                string `json:"name" binding:"required,TrimSpace,max=128"`         // 数据推送模型名称
	Description         string `json:"description" binding:"TrimSpace,omitempty,max=300"` // 描述
	ResponsiblePersonID string `json:"responsible_person_id" binding:"required,uuid"`     // 责任人ID

	PushStatus         *int32                 `json:"push_status" binding:"omitempty,oneof=0 1 2 3 4 5 6"` // 推送状态，不填不修改
	SourceCatalogID    string                 `json:"source_catalog_id" binding:"required"`                // 来源视图编目的数据目录ID，冗余字段
	TargetDatasourceID string                 `json:"target_datasource_id" binding:"required,uuid"`        // 数据源（目标端）
	TargetTableID      string                 `json:"target_table_id" binding:"omitempty,uuid"`            // 目标表ID，视图的ID
	TargetTableExists  *bool                  `json:"target_table_exists" binding:"required"`              // 目标表在本次推送是否存在，0不存在，1存在
	TargetTableName    string                 `json:"target_table_name"  binding:"required,max=255"`       // 目标表名称
	SyncModelFields    []*SyncModelCheckField `json:"sync_model_fields" binding:"required,min=1,dive"`     // 同步模型选定的字段
	FilterCondition    string                 `json:"filter_condition"   binding:"omitempty"`              // 过滤表达式，SQL后面的where条件

	TransmitMode       int32  `json:"transmit_mode" binding:"required,oneof=1 2"`            // 传输模式，1 增量 ; 2 全量
	IncrementField     string `json:"increment_field" binding:"TrimSpace,omitempty,max=128"` // 增量字段，当推送类型选择增量时，选一个字段作为增量字段，（技术名称）
	IncrementTimeStamp int64  `json:"increment_timestamp" binding:"omitempty"`               // 增量时间戳值，单位毫秒；当推送类型选择增量时，该字段必填
	PrimaryKey         string `json:"primary_key" binding:"omitempty,uuid"`                  // 主键，技术名称，当推送类型选择增量时，该字段必填

	ScheduleType  string `json:"schedule_type" binding:"required,oneof=ONCE PERIOD"` // 调度计划:once一次性,timely定时
	ScheduleTime  string `json:"schedule_time" binding:"omitempty,LocalDateTime"`    // 调度时间，格式 2006-01-02 15:04:05;  空：立即执行；非空：定时执行
	ScheduleStart string `json:"schedule_start" binding:"omitempty,LocalDate"`       // 计划开始日期, 格式 2006-01-02
	ScheduleEnd   string `json:"schedule_end" binding:"omitempty,LocalDate"`         // 计划结束日期, 格式 2006-01-02
	CrontabExpr   string `json:"crontab_expr" binding:"omitempty,cron"`              // linux crontab表达式, 5级
}

type TaskExecuteHistoryReq struct {
	Direction       string `query:"direction" binding:"oneof=asc desc" default:"desc"`             // 排序方向，枚举：asc：正序；desc：倒序。默认倒序
	Sort            string `query:"sort" binding:"oneof=start_time end_time" default:"updated_at"` // 排序类型，枚举：updated_at：按更新时间排序。默认按创建时间排序
	Offset          int    `query:"offset" binding:"min=1" default:"1" example:"1"`                // 页码，默认1
	Limit           int    `query:"limit" binding:"min=0,max=2000" default:"10" example:"2"`       // 每页大小，默认10 limit=0不分页
	Step            string `query:"step"  binding:"omitempty"`                                     //执行的,参考采集加工，传的是insert
	Status          string `query:"status"`                                                        //执行状态，SUCCESS FAILURE RUNNING_EXECUTION
	ScheduleExecute string `query:"scheduleExecute"`                                               //执行方式，手动执行还是自动执行
	ModelUUID       string `query:"model_uuid" binding:"required"`                                 //推送模型的ID
}

type TaskLogInfo struct {
	StartTime         string `json:"start_time"`          //请求时间
	EndTime           string `json:"end_time"`            //完成时间
	SyncCount         string `json:"sync_count"`          //推送总数
	SyncTime          string `json:"sync_time"`           //执行时间，单位，秒
	SyncMethod        string `json:"sync_method"`         //执行方式
	Status            string `json:"status"`              //状态
	StepName          string `json:"step_name"`           //步骤名称，uuid
	StepId            string `json:"step_id"`             //步骤ID，序号
	ProcessInstanceId string `json:"process_instance_id"` //处理实例ID
}

type BatchUpdateStatusReq struct {
	ModelID []uint64 `json:"model_id"  binding:"required"`
}
