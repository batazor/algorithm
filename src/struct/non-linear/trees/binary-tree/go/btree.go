package btree

// Tree is a binary tree.
type Tree[T any] struct {
	cmp   func(T, T) int
	left  *Tree[T]
	right *Tree[T]
	val   *T
}

func New[T any](cmp func(T, T) int) *Tree[T] {
	return &Tree[T]{
		cmp: cmp,
	}
}

func (t *Tree[T]) Find(key T) *Tree[T] {
	if t.val == nil {
		return t
	}

	switch cmp := t.cmp(key, *t.val); {
	case cmp < 0:
		return t.left.Find(key)
	case cmp > 0:
		return t.right.Find(key)
	default:
		return t
	}
}

func (t *Tree[T]) Insert(key T) *Tree[T] {
	if t.val == nil {
		t.val = &key
		return t
	}

	switch cmp := t.cmp(key, *t.val); {
	case cmp < 0:
		if t.left == nil {
			t.left = New(t.cmp)
			t.left.val = &key
			return t.left
		}
		return t.left.Insert(key)
	case cmp > 0:
		if t.right == nil {
			t.right = New(t.cmp)
			t.right.val = &key
			return t.right
		}
		return t.right.Insert(key)
	default:
		return t
	}
}

func (t *Tree[T]) Delete(key T) *Tree[T] {
	if t == nil {
		return t
	}

	switch cmp := t.cmp(key, *t.Value()); {
	case cmp < 0:
		t.left = t.left.Delete(key)
	case cmp > 0:
		t.right = t.right.Delete(key)
	default:
		if t.left == nil && t.right == nil {
			t = nil
		} else if t.left == nil {
			*t = *t.right
		} else if t.right == nil {
			*t = *t.left
		} else {
			t.val = t.right.Min().Value()
			t.right = t.right.Delete(*t.Value())
		}
	}

	return t
}

func (t *Tree[T]) Min() *Tree[T] {
	if t.left == nil {
		return t
	}

	return t.left.Min()
}

func (t *Tree[T]) Max() *Tree[T] {
	if t.right == nil {
		return t
	}

	return t.right.Max()
}

func (t *Tree[T]) Value() *T {
	return t.val
}
