package main

import (
	"container/list"
	"fmt"
)

// 定义二叉树的结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建一个二叉树节点
func CreateTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// 向二叉搜索树中插入节点
func (node *TreeNode) Insert(val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = CreateTreeNode(val)
		} else {
			node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = CreateTreeNode(val)
		} else {
			node.Right.Insert(val)
		}
	}
}

// 前序遍历
func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	node.Left.PreOrder()
	node.Right.PreOrder()
}

// 中序遍历
func (node *TreeNode) MidOrder() {
	if node == nil {
		return
	}
	node.Left.PreOrder()
	fmt.Println(node.Val)
	node.Right.PreOrder()
}

// 后序遍历
func (node *TreeNode) TailOrder() {
	if node == nil {
		return
	}
	node.Left.PreOrder()
	node.Right.PreOrder()
	fmt.Println(node.Val)
}

// 层序遍历
func LevelOrder(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	/*
		PushBack：尾部入队
		Front: 获取头部元素
		Remove: 出队
	*/
	// 本质为创建双向链表
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(*TreeNode)
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return ans
}

// 查找某个指定值的节点是否存在
func (node *TreeNode) Search(val int) bool {
	if node == nil {
		return false
	}
	if node.Val == val {
		return true
	}
	/*
		判断值在左边还是右边
	*/
	if node.Val > val {
		return node.Right.Search(val)
	} else {
		return node.Left.Search(val)
	}
}
