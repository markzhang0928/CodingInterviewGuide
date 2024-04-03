package main

func postorderRecurTraversal(root *TreeNode) []int {
	var postorder func(*TreeNode)
	var vals = []int{}
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		vals = append(vals, node.Val)
	}
	postorder(root)
	return vals
}

func preorderRecurTraversal(root *TreeNode) []int {
	var preorder func(*TreeNode)
	var vals = []int{}
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return vals
}
