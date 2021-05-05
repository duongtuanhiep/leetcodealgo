package bfs

// Idea: Do BFS, saves everything to array, keep track of layer
// Get all elem at the last layer, /2 root until its not possible no more
// func subtreeWithAllDeepest(root *TreeNode) *TreeNode {

// 	// Creates queue, arr node
// 	var queue []*TreeNode
// 	arrNode := []*TreeNode{nil}

// 	// BFS, creates arr
// 	var level = 0
// 	arrNode = append(arrNode, root)
// 	queue = append(queue, root)
// 	for len(queue) != 0 {

// 		for i := range queue {
// 			if queue[i] != nil {
// 				break
// 			}
// 			if i == len(queue)-1 {
// 				goto here
// 			}
// 		}

// 		length := len(queue)
// 		level++
// 		for i := 0; i < length; i++ {
// 			cur := queue[0]
// 			if cur == nil { // Adding 2 abitrary nodes here anyways
// 				queue, arrNode = append(queue, nil, nil), append(arrNode, nil, nil)
// 				queue = queue[1:]
// 			} else { // appends as normal
// 				arrNode = append(arrNode, cur.Left)
// 				arrNode = append(arrNode, cur.Right)
// 				queue = append(queue, cur.Left)
// 				queue = append(queue, cur.Right)
// 				queue = queue[1:]
// 			}
// 		}
// 	}

// here:
// 	// Get all last layer nodes pos
// 	var nodePos []int
// 	for i := int(math.Pow(2, float64(level-1))); i < len(arrNode); i++ {
// 		if arrNode[i] != nil {
// 			nodePos = append(nodePos, i)
// 		}
// 	}

// 	lca := nodePos[0]
// 	for i := 0; i < len(nodePos); i++ {
// 		lca = nodePos[0]
// 		if nodePos[i] != lca {
// 			for j := range nodePos {
// 				nodePos[j] = nodePos[j] / 2
// 			}
// 			i = 0
// 		}
// 	}

// 	return arrNode[lca]
// }

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, maxDepth := search(root)
	return maxDepth
}

func search(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	ldep, lnode := search(root.Left)
	rdep, rnode := search(root.Right)
	max := Max(ldep, rdep) + 1
	if ldep == rdep {
		return max, root
	}
	if ldep > rdep {
		return max, lnode
	}
	return max, rnode
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
