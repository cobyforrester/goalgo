package main

import (
	"fmt"
	"testing"

	"github.com/cobyforrester/goalgo/data_structures"
)


func TestUndirectedGraph(t *testing.T) {
	g := &data_structures.UndirectedGraph{
		Nodes: make(map[int][]int),
	}
	g.AddNode(1)
	fmt.Print(g)
}