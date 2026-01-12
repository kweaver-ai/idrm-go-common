package v1

import "encoding/json"

// ResourceObject 定义了被操作的资源对象
type ResourceObject interface {
	// GetName 返回资源的名称，用于生成审计日志的 Description
	GetName() string
	// GetDetail 返回资源的详情，用于在 Web 界面展示。例如：
	//
	//  (f *Foo) GetDetail() json.RawMessage {
	//      b, _ := json.Marshal(f)
	//      return b
	//  }
	//
	// 或
	//
	//  (b *Bar) GetDetail() json.RawMessage { return lo.Must(json.Marshal(b)) }
	GetDetail() json.RawMessage
}
