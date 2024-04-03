package main

import (
	"container/list"
)

/*
*
102. 二叉树的递归遍历
*/
func levelOrder1(root *TreeNode) [][]int {
	arr := [][]int{}

	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return arr
}

/*
*
102. 二叉树的层序遍历
*/
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)

	var tmpArr []int

	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			tmpArr = append(tmpArr, node.Val)               //将值加入本层切片中
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}

	return res
}

/*
*
102. 二叉树的层序遍历：使用切片模拟队列，易理解
*/
func levelOrder3(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	curLevel := []*TreeNode{root} // 存放当前层节点
	for len(curLevel) > 0 {
		nextLevel := []*TreeNode{} // 准备通过当前层生成下一层
		vals := []int{}

		for _, node := range curLevel {
			vals = append(vals, node.Val) // 收集当前层的值
			// 收集下一层的节点
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, vals)
		curLevel = nextLevel // 将下一层变成当前层
	}

	return
}
