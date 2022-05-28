package btree

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTree(t *testing.T) {
	// create
	tree := New(func(a, b int) int {
		return a - b
	})

	// crud operation
	expected := new(*int)
	assert.Equal(t, *expected, tree.Find(1).Value())

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(9)

	// aggregation operation
	assert.Equal(t, 4, *tree.Find(4).Value())
	assert.Equal(t, 1, *tree.Min().Value())
	assert.Equal(t, 9, *tree.Max().Value())

	// drop operation
	tree.Delete(5)
	assert.Equal(t, 9, *tree.Max().Value())
	tree.Delete(9)
	tree.Delete(1)
	tree.Insert(10)
	tree.Insert(11)
	assert.Equal(t, 11, *tree.Max().Value())

	// aggregation operation
	assert.Equal(t, 3, tree.Depth(0))
}

func TestTreeOrder(t *testing.T) {
	// create
	tree := New(func(a, b string) int {
		return strings.Compare(a, b)
	})

	tree.Insert("F")
	tree.Insert("B")
	tree.Insert("A")
	tree.Insert("D")
	tree.Insert("C")
	tree.Insert("E")
	tree.Insert("G")
	tree.Insert("I")
	tree.Insert("H")

	// order operation
	preOrder := []string{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	assert.Equal(t, preOrder, tree.PreOrder())

	inOrder := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	assert.Equal(t, inOrder, tree.InOrder())

	postOrder := []string{"A", "C", "E", "D", "B", "H", "I", "G", "F"}
	assert.Equal(t, postOrder, tree.PostOrder())

	bfs := map[int][]string{
		0: {"F"},
		1: {"B", "G"},
		2: {"A", "D", "I"},
		3: {"C", "E", "H"},
	}
	assert.Equal(t, bfs, tree.BFS(0))

	// aggregation operation
	assert.Equal(t, 3, tree.Depth(0))
}
