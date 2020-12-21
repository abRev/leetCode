package rbtree

const (
	red   = false
	black = true
)

type Node struct {
	Key    int
	Color  bool
	Parent *Node
	Left   *Node
	Right  *Node
}

func initNode(key int) *Node {
	node := &Node{
		Key:    key,
		Color:  red,
		Parent: nil,
		Left:   nil,
		Right:  nil,
	}
	return node
}

type RBTree struct {
	header *Node
}

func (this *RBTree) root() *Node {
	return this.header.Left
}

func (this *RBTree) ratateLeft(x *Node) {
	right := x.Right
	x.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = x
	}
	if x == this.root() {
		this.header.Left = right
	} else if x == x.Parent.Left {
		x.Parent.Left = right
	} else {
		x.Parent.Right = right
	}

	right.Left = x
	x.Parent = right

}

func (this *RBTree) ratateRight(x *Node) {
	left := x.Left
	x.Left = left.Right

	if left.Right != nil {
		left.Right.Parent = x
	}

	if x == this.root() {
		this.header.Left = left
	} else if x == x.Parent.Left {
		x.Parent.Left = left
	} else {
		x.Parent.Right = left
	}

	left.Left = x
	x.Parent = left
}

func (this *RBTree) destroy(node *Node) {
	if node == nil {
		return
	}
	this.destroy(node.Left)
	this.destroy(node.Right)
	node = nil
}

func (this *RBTree) insertRebalance(node *Node) {

}

func (this *RBTree) eraseRebalance(node *Node) {

}

func (this *RBTree) Insert(key int) {

}
func (this *RBTree) Find(key int) *Node {
	z := this.root()
	for z != nil {
		if z.Key > key {
			z = z.Left
		} else if z.Key < key {
			z = z.Right
		} else {
			return z
		}
	}
	return z
}
func (this *RBTree) Erase(node *Node) {

}
