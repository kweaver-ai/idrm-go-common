package completion

import "github.com/google/uuid"

// 生成 ID 的函数。为了方便在单元测试。
var generateIDFunc = func() string { return uuid.Must(uuid.NewV7()).String() }
