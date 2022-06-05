package contest

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Optimal recursion
func averageOfSubtree(root *TreeNode) int {
	total := 0
	countNodes(root, &total)
	return total
}

func countNodes(root *TreeNode, count *int) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftSum, leftN := countNodes(root.Left, count)
	rightSum, rightN := countNodes(root.Right, count)
	sum := leftSum + rightSum + root.Val
	n := leftN + rightN + 1
	if math.Round(float64(sum/n)) == float64(root.Val) {
		*count++
	}
	return sum, n
}

//Not optimal recursion
//My first submission. Compare and better understand recursion
func averageOfSubtree1(root *TreeNode) int {
	total := 0
	countNodes(root, &total)
	return total
}

func countNodes1(root *TreeNode, count *int) {
	if root != nil {
		sum, n := 0, 0
		nodeN := root.Val
		sumOfNodes(root, &sum, &n)
		sum = int(math.Round(float64(sum / n)))
		if nodeN == sum {
			*count++
		}
		countNodes(root.Left, count)
		countNodes(root.Right, count)
	}
}

func sumOfNodes(root *TreeNode, sum, i *int) {
	if root == nil {
		return
	}
	*i++
	*sum += root.Val
	sumOfNodes(root.Left, sum, i)
	sumOfNodes(root.Right, sum, i)
}
