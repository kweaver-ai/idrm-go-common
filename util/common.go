package util

import (
	"context"
	"unsafe"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"go.uber.org/zap"
)

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// [整数与布尔值的类型转换工具]
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func BoolToInt[T Int]() map[bool]T {
	return map[bool]T{true: 1, false: 0}
}

func IntToBool[T Int]() map[T]bool {
	return map[T]bool{1: true, 0: false}
} // [/]

// 记录错误日志，这里的err可能是recover的返回值所以不用error类型
func RecordErrLog(ctx context.Context, err any) {
	if err != nil {
		log.WithContext(ctx).Error("error: ", zap.Any("error", err), zap.Stack("call stack"))
	}
}

// 处理请求并记录错误日志
func HandleReqWithErrLog[T any](ctx context.Context, handler func(context.Context) (T, error)) (res T, err error) {
	res, err = handler(ctx)
	RecordErrLog(ctx, err)
	return
}

// 处理请求并记录错误日志和进行链路追踪
func HandleReqWithTraceIncludingErrLog[T any](ctx context.Context, handler func(context.Context) (T, error)) (res T, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()
	res, err = HandleReqWithErrLog(ctx, handler)
	return
}

func Contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
