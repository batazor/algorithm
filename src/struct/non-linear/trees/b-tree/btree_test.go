package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// create a b-tree
	tree := New(3)

	// expect empty b-tree
	var emptytNode *BTreeNode
	assert.Equal(t, emptytNode, tree.Search(1))

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(12)
	tree.Insert(30)
	tree.Insert(7)
	tree.Insert(17)

	tree.Traverse()

	assert.Equal(t, emptytNode, tree.Search(7))
	assert.Equal(t, emptytNode, tree.Search(15))
}
