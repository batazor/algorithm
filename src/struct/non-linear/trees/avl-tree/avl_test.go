package avl_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	// create
	tree := New(func(a, b int) int {
		return a - b
	})

	// crud operation
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(9)

	// drop operation
	tree.Delete(5)
	assert.Equal(t, 1, *tree.findMin().Value())
}
