package binarytree

import (
	"container/list"
)

type BSTIterator struct {
	stack *list.List
}

func Constructor(root *TreeNode) BSTIterator {
	stack := list.New()
	stack2 := list.New()
	node := root
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			node = stack.Remove(stack.Back()).(*TreeNode)
			stack2.PushBack(node)
			node = node.Right
		}
	}
	return BSTIterator{stack: stack2}
}

func (this *BSTIterator) Next() int {
	if this.stack.Len() > 0 {
		node := this.stack.Remove(this.stack.Front()).(*TreeNode)
		return node.Val
	}
	return 0
}

func (this *BSTIterator) HasNext() bool {
	return this.stack.Len() != 0
}
