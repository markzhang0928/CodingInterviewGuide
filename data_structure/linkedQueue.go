package data_structure

import "sync"

type LinkedQueue struct {
	head *LNode
	end  *LNode
	sync.RWMutex
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

// 判断队列是否为空,如果为空返回true，否则返回false
func (p *LinkedQueue) IsEmpty() bool {
	return p.head == nil
}

// 获取栈中元素的个数
func (p *LinkedQueue) Size() int {
	size := 0
	node := p.head
	for node != nil {
		node = node.Next
		size++
	}
	return size
}

// 入队列：把元素e加到队列尾
func (p *LinkedQueue) EnQueue(e interface{}) {
	p.Lock()
	defer p.Unlock()
	node := &LNode{Data: e}
	if p.head == nil {
		p.head = node
		p.end = node
	} else {
		p.end.Next = node
		p.end = node
	}
}

// 出队列，删除队列首元素
func (p *LinkedQueue) DeQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.head == nil {
		return nil
	}
	res := p.head
	p.head = p.head.Next
	if p.head == nil {
		p.end = nil
	}
	return res
}

// 取得队列首元素
func (p *LinkedQueue) GetFront() interface{} {
	if p.head == nil {
		return nil
	}
	return p.head.Data
}

// 取得队列尾元素
func (p *LinkedQueue) GetBack() interface{} {
	if p.end == nil {
		return nil
	}
	return p.end.Data
}
