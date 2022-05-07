package btree

import "fmt"

type BTreeNode struct {
	keys []int       // An array of keys
	t    int         // Minimum degree (defines the range for number of keys)
	C    []BTreeNode // An array of child pointers
	n    int         // Current number of keys
	leaf bool        // Is true when node is leaf. Otherwise false
}

func NewNode(t int, leaf bool) *BTreeNode {
	return &BTreeNode{
		t:    t,
		leaf: leaf,
		keys: make([]int, 2*t-1),
		C:    []BTreeNode{},
		n:    0,
	}
}

// Traverse - A function to traverse all nodes in a subtree rooted with this node
func (btn BTreeNode) Traverse() {
	// There are n keys and n+1 children, traverse through n keys
	// and first n children
	i := 0
	for i = 0; i < btn.n; i++ {
		if !btn.leaf {
			btn.C[i].Traverse()
		}

		fmt.Printf("%d ", btn.keys[i])
	}

	if btn.leaf == false {
		btn.C[i].Traverse()
	}
}

// Search - A function to search a key in the subtree rooted with this node.
func (btn BTreeNode) Search(k int) *BTreeNode {
	// Find the first key greater than or equal to k
	i := 0
	for i < btn.n && k > btn.keys[i] {
		i++
	}

	// If the found key is equal to k, return this node
	if btn.keys[i] == k {
		return &btn
	}

	// If the key is not found here and this is a leaf node
	if btn.leaf {
		return nil
	}

	return btn.C[i].Search(k)
}
