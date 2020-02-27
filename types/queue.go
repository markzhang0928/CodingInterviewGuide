package types

import "sync"

type SliceQueue struct {
	arr []interface{}
	sync.RWMutex
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{arr: make([]interface{}, 0)}
}

// 判空
func (p *SliceQueue) IsEmpty() bool {
	return p.Size() == 0
}

// 返回队列大小
func (p *SliceQueue) Size() int {
	return len(p.arr)
}

// 返回队列首元素
func (p *SliceQueue) GetFront() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.arr[0]
}

// 返回队列尾元素
func (p *SliceQueue) GetBack() interface{} {
	if p.IsEmpty() {
		return nil
	}
	return p.arr[p.Size()-1]
}

// 返回并移除队尾元素
func (p *SliceQueue) PopBack() interface{} {
	p.Lock()
	defer p.Unlock()
	if p.IsEmpty() {
		return nil
	}
	ret := p.arr[p.Size()-1]
	p.arr = p.arr[]
}
