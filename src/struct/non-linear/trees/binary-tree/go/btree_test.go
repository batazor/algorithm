package btree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree(t *testing.T) {
	tree := New(func(a, b int) int {
		return a - b
	})

	expected := new(*int)
	assert.Equal(t, *expected, tree.Find(1).Value())

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(9)

	assert.Equal(t, 4, *tree.Find(4).Value())
	assert.Equal(t, 1, *tree.Min().Value())
	assert.Equal(t, 9, *tree.Max().Value())

	tree.Delete(5)
	assert.Equal(t, 9, *tree.Max().Value())
	tree.Delete(9)
	tree.Delete(1)
	tree.Insert(10)
	assert.Equal(t, 10, *tree.Max().Value())
}
