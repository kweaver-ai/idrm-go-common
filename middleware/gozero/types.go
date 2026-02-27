package gozero

import "net/http"

// MiddlewareFunc 中间件函数类型定义
// 这是 go-zero rest.Middleware 的类型定义，避免直接引用 go-zero 包导致 IDE 类型解析错误
// 使用函数类型而非结构体，与 go-zero 的 rest.Middleware 兼容
type MiddlewareFunc func(handler http.HandlerFunc) http.HandlerFunc
