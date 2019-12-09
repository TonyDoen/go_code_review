package medium

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type CBTInserter struct {
	root             *TreeNode
	LastLastNodeList []*TreeNode
	LastNodeList     []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	this := CBTInserter{}
	this.root = root

	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		l := len(nodeList)
		this.LastLastNodeList = this.LastNodeList
		this.LastNodeList = nodeList
		for i := 0; i < l; i++ {
			nodeTmp := nodeList[0]
			nodeList = nodeList[1:]

			if nodeTmp.Left != nil {
				nodeList = append(nodeList, nodeTmp.Left)
			}
			if nodeTmp.Right != nil {
				nodeList = append(nodeList, nodeTmp.Right)
			}
		}
	}
	return this
}

func (this *CBTInserter) Insert(v int) int {
	if len(this.LastNodeList) >= len(this.LastLastNodeList)*2 {
		this.LastLastNodeList = this.LastNodeList
		this.LastNodeList = []*TreeNode{}
	}
	newNode := &TreeNode{Val: v}
	l := len(this.LastNodeList)
	if l%2 == 0 {
		this.LastLastNodeList[l/2].Left = newNode
	} else {
		this.LastLastNodeList[l/2].Right = newNode
	}
	this.LastNodeList = append(this.LastNodeList, newNode)
	return this.LastLastNodeList[l/2].Val
}

func (this *CBTInserter) GetRoot() *TreeNode {
	return this.root
}
