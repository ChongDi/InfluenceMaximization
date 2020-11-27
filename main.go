package main

import (
	IM "IM/IMpackage"
)

func main() {
	path := "GraphData/CA-GrQc.txt"

	g := IM.CreateUndirectedGraphFromFile(path) // return an undirected graph generated from 'path'.
	IM.ModelTest(g)
}
