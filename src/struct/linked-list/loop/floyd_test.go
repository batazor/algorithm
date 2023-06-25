package loop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLoop_Floyd(t *testing.T) {
	loopNode := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}
	loopNode.Next.Next.Next = loopNode.Next

	noLoopNode := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}

	testCases := []struct {
		in        *Node
		nodeValue int
		isLoop    bool
	}{
		{
			in:        loopNode,
			nodeValue: 2,
			isLoop:    true,
		},
		{
			in:        noLoopNode,
			nodeValue: 0,
			isLoop:    false,
		},
	}

	for _, c := range testCases {
		gotNodeValue, gotIsLoop := IsLoop_Floyd(c.in)
		if gotNodeValue != c.nodeValue || gotIsLoop != c.isLoop {
			require.Equal(t, c.nodeValue, gotNodeValue)
			require.Equal(t, c.isLoop, gotIsLoop)
		}
	}
}
