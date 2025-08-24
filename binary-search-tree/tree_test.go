package binarysearchtree_test

import (
	"testing"

	binarysearchtree "github.com/chaewonkong/data-structure/binary-search-tree"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("insert, find", func(t *testing.T) {
		tree := new(binarysearchtree.Tree)

		tree.InsertNode(4)
		tree.InsertNode(6)
		tree.InsertNode(2)
		tree.InsertNode(3)
		tree.InsertNode(1)
		tree.InsertNode(7)

		// find 3
		n3 := tree.FindNode(3)
		if n3 == nil || n3.Value != 3 {
			t.Errorf("n3 must be 3")
			return
		}

		if n3.Parent == nil || n3.Parent.Value != 2 {
			t.Errorf("n3 parent should be 3")
			return
		}

		if n3.Parent.Left == nil || n3.Parent.Left.Value != 1 {
			t.Errorf("n3.Parent.Left.Value shoud be 1")
			return
		}

		if n3.Parent.Right == nil || n3.Parent.Right != n3 {
			t.Errorf("n3.parent.right shoud be n3")
		}
	})

	t.Run("remove", func(t *testing.T) {
		tree := new(binarysearchtree.Tree)

		tree.InsertNode(4)
		tree.InsertNode(2)
		tree.InsertNode(7)
		tree.InsertNode(3)
		tree.InsertNode(1)
		tree.InsertNode(5)
		tree.InsertNode(9)
		tree.InsertNode(8)
		tree.InsertNode(11)

		n7 := tree.FindNode(7)
		tree.RemoveNode(n7)

		v := tree.FindNode(7)
		if v != nil {
			t.Errorf("7 should be removed")
			return
		}

		root := tree.FindNode(4)
		n8 := tree.FindNode(8)
		if n8.Parent == nil || root.Right == nil || root.Right != n8 {
			t.Errorf("n8 should be root.Right")
			return
		}

		if n8.Left == nil || n8.Left.Value != 5 {
			t.Errorf("n8..Left should be 5")
			return // node left가 사라진듯
		}

		if n8.Right == nil || n8.Right.Value != 9 {
			t.Errorf("n8..Right should be 9")

			if n8.Right.Right == nil || n8.Right.Right.Value != 11 {
				t.Errorf("n8..Right.Right should be 11")
				return
			}
		}
	})
}
