package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tree := New(2)

	var expectNode *BTreeNode
	assert.Equal(t, expectNode, tree.Search(1))
}
