package main

import (
	IM "IM/IMpackage"
)

func main() {
	path := "GraphData/CA-GrQc.txt"

	g := IM.CreatUndirectedGraphFromFile(path)
	IM.ModelTest(g)
}
