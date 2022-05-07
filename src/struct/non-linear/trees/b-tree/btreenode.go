package btree

import "fmt"

type BTreeNode struct {
	keys  []int        // An array of keys
	t     int          // Minimum degree (defines the range for number of keys)
	Child []*BTreeNode // An array of child pointers
	n     int          // Current number of keys
	leaf  bool         // Is true when node is leaf. Otherwise false
}

func NewNode(t int, leaf bool) *BTreeNode {
	return &BTreeNode{
		t:     t, // Copy the given minimum degree and leaf property
		leaf:  leaf,
		keys:  make([]int, 2*t-1), // Allocate memory for maximum number of possible keys and child pointers
		Child: make([]*BTreeNode, 2*t),
		n:     0, // Initialize the number of keys as 0
	}
}

// Traverse - A function to traverse all nodes in a subtree rooted with this node
func (btn *BTreeNode) Traverse() {
	// There are n keys and n+1 children, traverse through n keys
	// and first n children
	i := 0
	for i = 0; i < btn.n; i++ {
		if !btn.leaf {
			btn.Child[i].Traverse()
		}

		fmt.Printf("%d ", btn.keys[i])
	}

	if !btn.leaf {
		btn.Child[i].Traverse()
	}
}

// Search - A function to search a key in the subtree rooted with this node.
func (btn *BTreeNode) Search(k int) *BTreeNode {
	// Find the first key greater than or equal to k
	i := 0
	for i < btn.n && k > btn.keys[i] {
		i++
	}

	// If the found key is equal to k, return this node
	if btn.keys[i] == k {
		return btn
	}

	// If the key is not found here and this is a leaf node
	if btn.leaf {
		return nil
	}

	return btn.Child[i].Search(k)
}

// A utility function to insert a new key in the subtree rooted with
// this node. The assumption is, the node must be non-full when this
// function is called
func (btn *BTreeNode) insertNonFull(k int) {
	// Initialize index as index of rightmost element
	i := btn.n - 1

	// If this is a leaf node
	if btn.leaf {
		// The following loop does two things
		// a) Finds the location of new key to be inserted
		// b) Moves all greater keys to one place ahead
		for i >= 0 && btn.keys[i] > k {
			btn.keys[i+1] = btn.keys[i]
			i--
		}

		// Insert the new key at found location
		btn.keys[i+1] = k
		btn.n++

		return
	}

	// If this node is not leaf
	// Find the child which is going to have the new key
	for i >= 0 && btn.keys[i] > k {
		i--
	}

	// See if the found child is full
	if btn.Child[i+1].n == 2*btn.t-1 {
		// If the child is full, then split it
		btn.splitChild(i+1, btn.Child[i+1])

		// After split, the middle key of C[i] goes up and
		// C[i] is splitted into two.  See which of the two
		// is going to have the new key
		if btn.keys[i+1] < k {
			i++
		}

		btn.Child[i+1].insertNonFull(k)
	}
}

// A utility function to split the child y of this node. i is index of y in
// child array C[].  The Child y must be full when this function is called
func (btn *BTreeNode) splitChild(i int, y *BTreeNode) {
	// Create a new node which is going to store (t-1) keys of y
	z := NewNode(y.t, y.leaf)
	z.n = btn.t - 1

	// Copy the last (t-1) keys of y to z
	for j := 0; j < btn.t-1; j++ {
		z.keys[j] = y.keys[j+btn.t]
	}

	// Copy the last t children of y to z
	if y.leaf == false {
		for j := 0; j < btn.t; j++ {
			z.Child[j] = y.Child[j+btn.t]
		}
	}

	// Reduce the number of keys in y
	y.n = btn.t - 1

	// Since this node is going to have a new child,
	// create space of new child
	for j := btn.n; j >= i+1; j-- {
		btn.Child[j+1] = btn.Child[j]
	}

	// Link the new child to this node
	btn.Child[i+1] = z

	// A key of y will move to this node. Find the location of
	// new key and move all greater keys one space ahead
	for j := btn.n - 1; j >= i; j-- {
		btn.keys[j+1] = btn.keys[j]
	}

	// Copy the middle key of y to this node
	btn.keys[i] = y.keys[btn.t-1]

	// Increment count of keys in this node
	btn.n++
}

// A function that returns the index of the first key that is greater or equal to k
func (btn *BTreeNode) findKey(k int) int {
	idx := 0
	for idx < btn.n && btn.keys[idx] < k {
		idx++
	}
	return idx
}
