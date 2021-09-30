/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var head *TreeNode
	head = root
	preorder(root, head)
	root = head
}

func preorder(root *TreeNode, mainroot *TreeNode) {

	if root == nil {
		return
	} else {
		mainroot.Right = root
		mainroot.Left = new(TreeNode)
		preorder(root.Left, mainroot)
		preorder(root.Right, mainroot)
	}
}

