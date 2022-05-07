package btree

import "fmt"

type BTree struct {
	root *BTreeNode // Pointer to root node
	t    int        // Minimum degree
}

func New(t int) *BTree {
	return &BTree{
		root: nil,
		t:    t,
	}
}

func (bt BTree) Traverse() {
	if bt.root != nil {
		bt.root.Traverse()
		fmt.Sprintln()
	}
}

// Search - function to search a key in this tree
func (bt BTree) Search(k int) *BTreeNode {
	if bt.root != nil {
		return bt.root.Search(k)
	}

	return nil
}
