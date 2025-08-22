package main

// 定义链表节点
type ListNode struct {
	Val  int       // 数据域
	Next *ListNode // 指针域，指向下一个指针
}

// 创建链表 -> 将列表转化为链表
func CreatLinkList(nums []int) *ListNode {
	// 创建哨兵节点，简化链表操作
	head := &ListNode{}
	current := head

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head.Next
}

// 遍历链表
func TraverseList(head *ListNode) []int {
	// 使用切片更加灵活
	// 切片[]int{}  数组[...]int{}
	list := []int{}
	cur := head
	for cur != nil {
		value := cur.Val
		// 判断容量，是否迭代新的数组,动态扩容
		list = append(list, value)
		cur = cur.Next
	}
	return list
}

// 链表头部插入
func Insert_head(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Val: val}
	newHead.Next = head
	return newHead
}

// 在链表的尾部插入
func Insert_tail(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	cur := head
	if head == nil {
		return newNode
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode.Next
	return head
}

// 在指定的位子插入数据 插入的位置小于链表的长度
func Inser_point(head *ListNode, val int, point int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		return newNode
	}
	if point == 0 {
		newNode.Next = head
		return newNode
	}
	cur := head
	for i := 0; i < point; i++ {
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	return head
}

// 删除指定值的节点
func Delete_value(head *ListNode, val int) *ListNode {
	// 设置一个虚拟节点
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			// 以 cur 的视角判断 cur.val
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 删除指定索引，指定位置
