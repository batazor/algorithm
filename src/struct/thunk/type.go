package thunk

import (
	"github.com/batazor/algorithm/src/struct/thunk/options"
)

type Thunk[T any] struct {
	doer func() T           // action being thunked
	o    *options.Option[T] // cache for complete thunk data
}
