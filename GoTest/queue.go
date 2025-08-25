package main

import (
	"errors"
	"fmt"
)

// 初始化一个队列结构
type Queue struct {
	elements []interface{} // 切片,存储多种类型的元素
	size     int           // 长度
}

// 创建一个队列 size = 0 表示无容量限制
func NewQueue(size int) *Queue {
	return &Queue{
		make([]interface{}, 0),
		size,
	}
}

// 像队列尾部添加一个元素
// Enqueue 入队，接收一个interface{}类型的参数element
// 它需要依赖一个具体的队列实例（q *Queue）才能操作（比如修改队列的内部数据）
// 因此必须通过接收者（(q *Queue)）绑定到 Queue 结构体。
func (q *Queue) Enqueue(val interface{}) error {
	// q 为结构体的实例 q.elements队列内部用于存储实际元素的数据容器
	// len(q.elements) >= q.size 队列中存储的数量已经达到最大容量
	if q.size > 0 && len(q.elements) >= q.size {
		return errors.New("queue is full")
	}
	q.elements = append(q.elements, val)
	return nil
}

// 从队列头部移除并返回一个元素
func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("queue is empty")
	}
	val := q.elements[0]GoTest
	q.elements = q.elements[1:]
	return val, nil
}

//type Queue_int struct {
//	int []interface{} // 切片,存储多种类型的元素
//	size     int           // 长度
//}

// 遍历队列
func (q *Queue) Traverse() {
	for i := 0; i < q.size; i++ {
		ele := q.elements[i] // 获取当前元素
		fmt.Println(ele)
	}
}
