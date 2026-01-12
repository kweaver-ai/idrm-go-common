package built_in

const (
	NONE                                 int32 = 0
	NormalPlatform                       int32 = 1 << (iota - 1) // 标品平台
	DataResourceManagementBackupPlatform                         // 数据资源管理平台-后台
	DataResourceManagementPortalPlatform                         // 数据资源管理平台-门户
	CognitiveApplicationPlatform                                 // 认知应用平台
	CognitiveDiagnosisPlatform                                   // 业务认知分析平台
)
