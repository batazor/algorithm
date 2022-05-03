package thunk

import (
	"github.com/batazor/algorithm/src/struct/thunk/options"
)

type Thunk[T any] struct {
	doer func() T           // action being thunked
	o    *options.Option[T] // cache for complete thunk data
}

func New[T any](doer func() T) *Thunk[T] {
	return &Thunk[T]{
		doer: doer,
		o:    options.New[T](),
	}
}

func (t *Thunk[T]) Force() T {
	if t.o.IsSome() {
		return t.o.Apply()
	}

	t.o.Set(t.doer())
	return t.o.Apply()
}
