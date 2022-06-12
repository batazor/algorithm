package avl_tree

// Tree is an AVL tree.
type Tree[T any] struct {
	// cmp compares two T values.
	cmp    func(T, T) int
	val    *T
	height int
	left   *Tree[T]
	right  *Tree[T]
}

func New[T any](cmp func(T, T) int) *Tree[T] {
	return &Tree[T]{
		cmp:    cmp,
		height: 0,
	}
}

// Height returns the height of the tree.
func (t *Tree[T]) Height() int {
	return t.height
}

// Value - returns the value of the node.
func (t *Tree[T]) Value() *T {
	return t.val
}

// balance returns the balance of the tree.
func (t *Tree[T]) getBalance() int {
	if t.left == nil || t.right == nil {
		return 0
	}

	return t.right.Height() - t.left.Height()
}

// fixheight updates the height of a node.
func (t *Tree[T]) fixHeight() {
	if t.left == nil || t.right == nil {
		return
	}

	if t.left.Height() > t.right.Height() {
		t.height = t.left.Height() + 1
	} else {
		t.height = t.right.Height() + 1
	}
}

// rotateRight rotates the tree to the right.
func (t *Tree[T]) rotateRight(p *Tree[T]) *Tree[T] {
	q := p.left
	p.left = q.right
	q.right = p
	p.fixHeight()
	q.fixHeight()
	return q
}

// rotateLeft rotates the tree to the left.
func (t *Tree[T]) rotateLeft(p *Tree[T]) *Tree[T] {
	q := p.right
	p.right = q.left
	q.left = p
	p.fixHeight()
	q.fixHeight()
	return q
}

// Balance - makes sure the tree is balanced.
func (t *Tree[T]) Balance(p *Tree[T]) *Tree[T] {
	p.fixHeight()

	if p.getBalance() == 2 {
		if p.right.getBalance() < 0 {
			p.right = p.left.rotateRight(p.right)
		}
		return p.rotateLeft(p)
	}

	if p.getBalance() == -2 {
		if p.left.getBalance() > 0 {
			p.left = p.right.rotateLeft(p.left)
		}
		return p.rotateRight(p)
	}

	return p
}

func (t *Tree[T]) Insert(key T) *Tree[T] {
	if t.val == nil {
		t.val = &key
		return t.Balance(t)
	}

	if t.cmp(key, *t.val) < 0 {
		if t.left == nil {
			t.left = New(t.cmp)
			t.left.val = &key
		} else {
			t.left = t.left.Insert(key)
		}
	} else {
		if t.right == nil {
			t.right = New(t.cmp)
			t.right.val = &key
		} else {
			t.right = t.right.Insert(key)
		}
	}

	return t.Balance(t)
}

// findMin - finds the minimum value in the tree.
func (t *Tree[T]) findMin() *Tree[T] {
	if t.left == nil {
		return t
	}
	return t.left.findMin()
}

// removeMin - removes the minimum value in the tree.
func (t *Tree[T]) removeMin() *Tree[T] {
	if t.left == nil {
		return t.right
	}
	t.left = t.left.removeMin()
	return t.Balance(t)
}

func (t *Tree[T]) Delete(key T) *Tree[T] {
	if t.cmp(key, *t.val) < 0 {
		t.left = t.left.Delete(key)
	} else if t.cmp(key, *t.val) > 0 {
		t.right = t.right.Delete(key)
	} else {
		q := t.left
		r := t.right
		t = nil
		if r == nil {
			return q
		}
		min := r.findMin()
		min.right = r.removeMin()
		min.left = q
		return min.Balance(min)
	}

	return t.Balance(t)
}
