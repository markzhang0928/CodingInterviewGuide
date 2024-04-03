package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) != 0 {
		newQueue := []*TreeNode{}
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
			if i == len(queue)-1 {
				res = append(res, queue[i].Val)
			}
		}
		queue = newQueue
	}
	return res
}
