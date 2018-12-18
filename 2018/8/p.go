package p

type Node struct {
	Metadata []int
	Children []*Node
}

func Parse(p []int) (n *Node, q []int) {
	n = &Node{}
	n.Children, p = make([]*Node, p[0]), p[1:]
	n.Metadata, p = make([]int, p[0]), p[1:]
	for i := range n.Children {
		n.Children[i], p = Parse(p)
	}
	for i := range n.Metadata {
		n.Metadata[i], p = p[0], p[1:]
	}
	q = p
	return
}

func (n *Node) P1() (s int) {
	for _, m := range n.Metadata {
		s += m
	}
	for _, c := range n.Children {
		s += c.P1()
	}
	return
}

func (n *Node) P2() (s int) {
	if len(n.Children) == 0 {
		for _, m := range n.Metadata {
			s += m
		}
		return
	}
	for _, m := range n.Metadata {
		i := m - 1
		if i < len(n.Children) {
			s += n.Children[i].P2()
		}
	}
	return
}
