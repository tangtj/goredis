package list

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

// NewNode returns a new node with the given value.
func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

func (n *Node) GetValue() interface{} {
	return n.value
}

// GetNext is get node next
func (n *Node) GetNext() *Node {
	return n.next
}
