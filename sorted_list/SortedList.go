package sorted_list

import (
	"sync"
)

type IntList struct {
	head   *intNode
	length int
	mu     sync.RWMutex
}

type intNode struct {
	val    int
	next   *intNode
	marked bool
}

// newNode 返回一个全新的链表节点
func newNode(val int) *intNode {
	return &intNode{val: val}
}

// NewInt 返回一个全新的有序链表
func NewInt() *IntList {
	return &IntList{head: newNode(0)}
}

// Contains 检查一个元素是否存在，如果存在则返回 true，否则返回 false
func (l *IntList) Contains(value int) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	node := l.head.next
	for node != nil && node.val < value {
		node = node.next
	}
	return node != nil && node.val == value
}

// 插入一个元素，如果此操作成功插入一个元素，则返回 true，否则返回 false
func Insert(value int) bool {
	
}

// 删除一个元素，如果此操作成功删除一个元素，则返回 true，否则返回 false
func Delete(value int) bool

// 遍历此有序链表的所有元素，如果 f 返回 false，则停止遍历
func (l *IntList) Range(f func(value int) bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	node := l.head.next
	for node != nil {
		if !f(node.val) {
			break
		}
		node = node.next
	}
}

// Len 返回有序链表的元素个数
func (l *IntList) Len() int {
	return l.length
}
