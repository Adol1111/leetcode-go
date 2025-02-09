package solutions

func Connect(root *TreeNode) *TreeNode {
	startNode := root
	for startNode != nil {
		// nextStart 为下一层的起始节点, 即最左侧节点
		// lastNode 为该层上一次处理的最后一个节点，即当前节点的左侧节点
		var nextStart, lastNode *TreeNode
		handle := func(current *TreeNode) {
			if current == nil {
				return
			}
			if nextStart == nil {
				nextStart = current
			}
			if lastNode != nil {
				lastNode.Next = current
			}
			lastNode = current
		}
		// 每次遍历是给下一层建立Next连接，再遍历下一层时，可以顺着上一层建立的next横向遍历该层
		for p := startNode; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		startNode = nextStart
	}
	return root
}

//       1
//     /   \
//    2 --> 3
//   /     / \
//  4 --> 6-->7

// 1. startNode = 1, 2和3 建立连接, nextStartNode = 2
// 2. startNode = 2
//  2.1 先建立2的子树，由于2没有右子树，所以lastNode = 4, nextStartNode = 4
//  2.2 再建立3的子树，4 -> 6 -> 7
// 3. startNode = 4
//  3.1 先建立4的子树，由于4没有子树, lastNode=nil, nextStartNode = nil
//  3.2 再建立6的子树，由于6没有子树, lastNode=nil, nextStartNode = nil
//  3.3 再建立7的子树，由于7没有子树, lastNode=nil, nextStartNode = nil
