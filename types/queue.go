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
	p.arr = p.arr[:p.Size()-1]
	return ret
}

// 删除队列头元素
func (p *SliceQueue) DeQueue() interface{} {
	p.Lock()
	defer p.Unlock()
	if len(p.arr) != 0 {
		first := p.arr[0]
		p.arr = p.arr[1:]
		return first
	} else {
		return nil
	}
}

//把新元素加入队列尾
func (p *SliceQueue) EnQueue(item interface{}) {
	p.Lock()
	defer p.Unlock()
	p.arr = append(p.arr, item)
}

// 把新元素加入队列首
func (p *SliceQueue) EnQueueFirst(item interface{}) {
	p.Lock()
	defer p.Unlock()
	newQueue := []interface{}{item}
	p.arr = append(newQueue, p.arr[:]...)
}

// 简单实现一个remove
func (p *SliceQueue) Remove(item interface{}) {
	p.Lock()
	defer p.Unlock()
	for k, v := range p.arr {
		if v == item {
			p.arr = append(p.arr[:k], p.arr[k+1:]...)
		}
	}
}

func (p *SliceQueue) List() []interface{} {
	return p.arr
}
