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

func (bt *BTree) Traverse() {
	if bt.root != nil {
		bt.root.Traverse()
		fmt.Sprintln()
	}
}

// Search - function to search a key in this tree
func (bt *BTree) Search(k int) *BTreeNode {
	if bt.root != nil {
		return bt.root.Search(k)
	}

	return nil
}

// The main function that inserts a new key in this B-Tree
func (bt *BTree) Insert(k int) {
	if bt.root == nil {
		bt.root = NewNode(bt.t, true)
		bt.root.keys[0] = k
		bt.root.n = 1
		return
	}

	if bt.root.n == 2*bt.t-1 {
		// Allocate memory for new root
		s := NewNode(bt.t, false)

		// Make old root as child of new root
		s.Child[0] = bt.root

		// Split the old root and move 1 key to the new root
		s.splitChild(0, bt.root)

		// New root has two children now.  Decide which of the
		// two children is going to have new key
		i := 0
		if s.keys[0] < k {
			i++
		}

		s.Child[i].insertNonFull(k)

		// Change root
		bt.root = s

		return
	}

	// If root is not full, call insertNonFull for root
	bt.root.insertNonFull(k)
}

// The main function that removes a new key in thie B-Tree
func (bt *BTree) Remove(k int) {

}
