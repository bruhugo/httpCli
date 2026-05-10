type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	_, m := traverse(root)
	return m
}

func traverse(node *TreeNode) (path, m int) {
	if node == nil {
		return 0, 0
	}

	pathl, maxl := traverse(node.Left)
	pathr, maxr := traverse(node.Right)

	path = max(pathl, pathr) + node.Val
	m = max(maxl, maxr, path, node.Val, pathl+pathr+node.Val)

	return
}