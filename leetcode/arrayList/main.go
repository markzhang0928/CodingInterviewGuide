package main

import "fmt"

type ArrayList struct {
	arr []int
	len int
	cap int
}

func (list *ArrayList) GetList(index int) int {
	if list.len == 0 || index > len(list.arr) {
		// panic or return nil
		return -1
	}
	return list.arr[index]
}

// 添加元素到列表
func (list *ArrayList) PushList(val int) *ArrayList {

	curlen := list.len + 1
	if curlen <= list.cap {
		// 添加元素到列表
		list.arr[list.len] = val
		list.len++
	} else if curlen > list.cap {
		// 扩容数组列表
		newList := make([]int, list.len*2)
		for i, value := range list.arr {
			newList[i] = value
		}
		newList[curlen-1] = val
		list.arr = newList
		list.len = curlen
		list.cap = len(newList)
	}
	return list
}

func (list *ArrayList) DeleteList(index int) *ArrayList {
	if list.len == 0 || index > len(list.arr) {
		// panic or return nil
		return list
	}
	nextList := list.arr[index+1:]
	for i, j := index, 0; j < len(list.arr[index+1:]); j++ {
		list.arr[i] = nextList[j]
		i++
	}
	list.len--
	return list
}

func main() {
	list := ArrayList{arr: make([]int, 10), len: 0, cap: 10}
	// push 1... 12
	for i := 0; i < 15; i++ {
		list.PushList(i)
	}
	fmt.Println(list)

	// get index =1
	fmt.Println(list.GetList(5))
	fmt.Println(list.DeleteList(2))
	fmt.Println(list.DeleteList(5))
	fmt.Println(list.DeleteList(10))
}
