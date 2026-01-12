package v1

// ProcessingState 代表一次处理的期望状态
type ProcessingState string

const (
	// 处理的期望状态：存在。如果不存在则创建，否则无操作
	ProcessingStatePresent ProcessingState = "Present"
	// 处理的期望状态：不存在。如果存在则删除，否则无操作
	ProcessingStateAbsent ProcessingState = "Absent"
)
