package btree

// Tree is a binary tree.
type Tree[T any] struct {
	// cmp compares two T values.
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

// PreOrder traversal is to visit the root first.
// Then traverse the left subtree.
// Finally, traverse the right subtree.
func (t *Tree[T]) PreOrder() []T {
	res := []T{*t.Value()}

	if t.left != nil {
		res = append(res, t.left.PreOrder()...)
	}

	if t.right != nil {
		res = append(res, t.right.PreOrder()...)
	}

	return res
}

// InOrder traversal is to traverse the left subtree first.
// Then visit the root.
// Finally, traverse the right subtree.
func (t *Tree[T]) InOrder() []T {
	var res []T

	if t.left != nil {
		res = append(res, t.left.InOrder()...)
	}

	res = append(res, *t.Value())

	if t.right != nil {
		res = append(res, t.right.InOrder()...)
	}

	return res
}

// PostOrder traversal is to traverse the left subtree first.
// Then traverse the right subtree.
// Finally, visit the root.
func (t *Tree[T]) PostOrder() []T {
	var res []T

	if t.left != nil {
		res = append(res, t.left.PostOrder()...)
	}
	if t.right != nil {
		res = append(res, t.right.PostOrder()...)
	}

	res = append(res, *t.Value())

	return res
}

func (t *Tree[T]) BFS(level int) map[int][]T {
	res := map[int][]T{}

	res[level] = append(res[level], *t.Value())

	if t.left != nil {
		leftRes := t.left.BFS(level + 1)

		for i := range leftRes {
			res[i] = append(res[i], leftRes[i]...)
		}
	}

	if t.right != nil {
		rightRes := t.right.BFS(level + 1)

		for i := range rightRes {
			res[i] = append(res[i], rightRes[i]...)
		}
	}

	return res
}

func (t *Tree[T]) Depth(depth int) int {
	leftDepth := 0
	if t.left != nil {
		leftDepth = t.left.Depth(depth + 1)
	}

	rightDepth := 0
	if t.right != nil {
		rightDepth = t.right.Depth(depth + 1)
	}

	if leftDepth >= rightDepth && leftDepth > depth {
		return leftDepth
	}

	if rightDepth >= leftDepth && rightDepth > depth {
		return rightDepth
	}

	return depth
}
