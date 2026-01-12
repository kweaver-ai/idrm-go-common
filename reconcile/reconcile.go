package reconcile

import (
	"context"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

type Reconciler[T any] interface {
	Reconcile(context.Context, *meta_v1.WatchEvent[T]) error
}
