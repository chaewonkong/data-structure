package binarysearchtree

// Tree 이진탐색트리
type Tree struct {
	Root *TreeNode
}

// FindNode tree에서 노드를 찾는 함수
func (t *Tree) FindNode(target int) *TreeNode {
	if t.Root == nil {
		return nil
	}

	return FindValue(t.Root, target)
}

func (t *Tree) InsertNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Value: val}
		return
	}

	InsertValue(t.Root, val)
}

func (t *Tree) RemoveNode(node *TreeNode) {
	defer func() {
		// clean up
		node.Parent = nil
		node.Right = nil
		node.Left = nil
	}()

	if t == nil || node == nil {
		return
	}

	// leaf node
	if node.Left == nil && node.Right == nil {
		// root node
		if node.Parent == nil {
			t.Root = nil
			return
		}

		if node.Parent.Left == node {
			node.Parent.Left = nil
			return
		}

		node.Parent.Right = nil
		return
	}

	// one child
	if node.Left == nil || node.Right == nil {
		child := node.Left
		if node.Left == nil {
			child = node.Right
		}

		// node is root
		if node.Parent == nil {
			t.Root = child
			return
		}

		child.Parent = node.Parent
		if node.Parent.Left == node {
			node.Parent.Left = child
			return
		}

		node.Parent.Right = child
		return
	}

	// two children
	// 1. successor 찾기: node.Right의 모든 재귀적인 자식 노드들 중 가장 작은 node를 찾고 node의 위치에 바꿔치기
	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	// 2. successor를 결합에서 제거
	t.RemoveNode(successor) // successor의 자식은 0 또는 1개 (2개였으면 재귀탐색 했지)

	// 3. successor를 제거할 노드의 위치에 추가
	// node가 root
	if node.Parent == nil {
		t.Root = successor
	} else if node.Parent.Left == node {
		node.Parent.Left = successor
	} else {
		node.Parent.Right = successor
	}

	successor.Parent = node.Parent

	// node의 왼쪽 자식을 successor에 붙여주기
	successor.Left = node.Left
	node.Left.Parent = successor

	// node의 오른쪽 자식을 successor에 붙여주기
	successor.Right = node.Right
	if node.Right != nil {
		node.Right.Parent = successor
	}
}

// TreeNode 이진탐색트리의 노드
type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func FindValue(cur *TreeNode, target int) *TreeNode {
	if cur == nil {
		return nil
	}

	if cur.Value == target {
		return cur
	}

	if cur.Value > target && cur.Left != nil {
		return FindValue(cur.Left, target)
	}

	if cur.Value < target && cur.Right != nil {
		return FindValue(cur.Right, target)
	}

	return nil
}

func InsertValue(cur *TreeNode, val int) {
	if cur.Value == val {
		return // insert 필요 없음
	}

	if cur.Value > val {
		if cur.Left != nil {
			InsertValue(cur.Left, val)
			return
		}

		cur.Left = &TreeNode{Value: val, Parent: cur}
		return
	} else {
		if cur.Right != nil {
			InsertValue(cur.Right, val)
			return
		}

		cur.Right = &TreeNode{Value: val, Parent: cur}
		return
	}
}
