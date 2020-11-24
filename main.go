package main

import (
	"fmt"

	graph "gonum.org/v1/gonum/graph/simple"

	IM "IM/IMpackage"
)

func main() {
	path := "GraphData/facebook_combined.txt"
	g := graph.NewUndirectedGraph()
	IM.CreatUndirectedGraphFromFile(path, g)
	fmt.Println(IM.IM_Entrance_Undirected(g))

}
