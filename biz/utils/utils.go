package utils

import (
	"context"
	"runtime/debug"

	"github.com/RanFeng/ilog"
)

// PanicHandler 用于处理panic的函数，该变量可被重赋值
var PanicHandler = func(err interface{}) {
	ilog.EventInfo(context.Background(), "go_routine_panic", "panic", err, "stack", string(debug.Stack()))
}

// WithRecover 用于对函数进行recover包装执行
func WithRecover(fn func()) {
	defer func() {
		handler := PanicHandler
		if handler != nil {
			if err := recover(); err != nil {
				handler(err)
			}
		}
	}()

	fn()
}

func IF[T any](ok bool, a, b T) T {
	if ok {
		return a
	}
	return b
}
