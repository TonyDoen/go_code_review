package medium

/**
  url: http://www.cnblogs.com/grandyang/p/6228252.html
  url: https://leetcode.com/problems/delete-node-in-a-bst/discuss/93296/recursive-easy-to-understand-java-solution

  Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
  Basically, the deletion can be divided into two stages:
  Search for a node to remove.
  If the node is found, delete the node.
  Note: Time complexity should be O(height of tree).

  Example:
  root = [5,3,6,2,4,null,7]
  key = 3

      5
     / \
    3   6
   / \   \
  2   4   7

  Given key to delete is 3. So we find the node with value 3 and delete it.
  One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

      5
     / \
    4   6
   /     \
  2       7

  Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
*/

/**
  题意：删除二叉搜索树中的节点
  这道题让我们删除二叉搜索树中的一个节点，这道题的难点在于删除完节点并补上那个节点的位置后还应该是一棵二叉搜索树。被删除掉的节点位置，不一定是由其的左右子节点补上，比如下面这棵树：
        7                                                              7
      /   \                                                          /  \
     4     8             如果我们要删除节点4，                        5     8
   /   \                 那么应该将节点5补到4的位置，                /  \
  2     6                这样才能保证还是BST                      2     6
   \    /               ------------------------->               \
    3  5                                                          3

  先来看一种递归的解法，首先判断根节点是否为空。由于BST的左<根<右的性质，使得我们可以快速定位到要删除的节点，我们对于当前节点值不等于key的情况，根据大小关系对其左右子节点分别调用递归函数。若当前节点就是要删除的节点，我们首先判断是否有一个子节点不存在，那么我们就将root指向另一个节点，如果左右子节点都不存在，那么root就赋值为空了，也正确。难点就在于处理左右子节点都存在的情况，我们需要在右子树找到最小值，即右子树中最左下方的节点，然后将该最小值赋值给root，然后再在右子树中调用递归函数来删除这个值最小的节点
*/

func DeleteNode3(root *Node, key int) *Node { // error
	if nil == root {
		return nil
	}
	if key == root.data {
		if nil == root.right {
			return root.left
		} else {
			cur := root.right
			for nil != cur.left {
				cur = cur.left
			}
			// swap
			val := root.data
			root.data = cur.data
			cur.data = val
		}
	}
	root.left = DeleteNode3(root.left, key)
	root.right = DeleteNode3(root.right, key)
	return root
}
