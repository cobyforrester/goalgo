package data_structures

// Nodes and their neighbors
type UndirectedGraph struct {
	Nodes map[int][]int
}

func (u *UndirectedGraph) AddNode (node int) error{
	if _, ok := u.Nodes[node]; !ok {
		u.Nodes[node] = []int{}
	}
	return nil
}

func (u *UndirectedGraph) RemoveNode (node int) error{
	delete(u.Nodes, node)
	return nil
}
