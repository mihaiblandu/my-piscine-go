package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 1
	node := l
	for ; node != nil; count++ {
		if count == pos {
			break
		}
		node = node.Next
	}
	return node.Next
}
