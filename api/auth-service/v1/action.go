package v1

// Action 定义动作
type Action string

const (
	ActionRead     Action = "read"     // 读取
	ActionDownload Action = "download" // 下载
	ActionManage   Action = "manage"   // 管理
	ActionAuth     Action = "auth"     // 授权
	ActionAllocate Action = "allocate" // 授权仅分配
)

func (a Action) Str() string {
	return string(a)
}
