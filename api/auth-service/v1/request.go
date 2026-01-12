package v1

// EnforceRequest 定义一个鉴权请求，auth-service 判断这个请求是否被允许。
type EnforceRequest struct {
	// 操作者
	Subject
	// 资源
	Object
	// 动作
	Action Action `json:"action,omitempty"`
}

type EnforceResponse struct {
	EnforceRequest `json:",inline"`

	Effect PolicyEffect `json:"effect,omitempty"`
}

type RulePolicyEnforce struct {
	UserID     string `json:"user_id" form:"user_id"`                                             //用户ID
	ObjectId   string `json:"object_id" form:"object_id" binding:"required,VerifyNameEn,max=128"` //资源id
	ObjectType string `json:"object_type" form:"object_type" binding:"required"`                  //资源类型 domain 主题域 data_view 逻辑视图 api 接口 sub_view 子视图 indicator 指标
	Action     string `json:"action" form:"action" binding:"required"`                            //请求动作 view 查看 read 读取 download 下载
}

type RulePolicyEnforceEffect struct {
	ObjectId   string `json:"object_id" binding:"required,VerifyNameEn,max=128"`                                                       //资源id
	ObjectType string `json:"object_type" binding:"required,oneof=domain data_view api sub_view indicator indicator_dimensional_rule"` //资源类型 domain 主题域 data_view 逻辑视图 api 接口 sub_view 子视图 indicator 指标
	Action     string `json:"action" binding:"required,oneof=view read download"`                                                      //请求动作 view 查看 read 读取 download 下载
	Effect     string `json:"effect"`                                                                                                  //策略结果 allow 允许 deny 拒绝
}
