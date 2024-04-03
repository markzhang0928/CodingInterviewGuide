package main

import "container/list"

// 递归
func getdepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getdepth(root.Left)      // 左
	rightDepth := getdepth(root.Right)    // 右
	return max(leftDepth, rightDepth) + 1 // 中
}

// 遍历
func maxdepth(root *TreeNode) int {
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}

func maxDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans++
	}
	return ans
}
