package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	graph "gonum.org/v1/gonum/graph/simple"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func creat_graph_from_file(path string, g *graph.UndirectedGraph) int {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		node_s, _ := strconv.Atoi(strings.Split(string(a), " ")[0])
		node_d, _ := strconv.Atoi(strings.Split(string(a), " ")[1])
		g.SetEdge(g.NewEdge(graph.Node(node_s), graph.Node(node_d)))
	}
	fmt.Println(g.Nodes().Len())

	return 0
}

func main() {
	path := "GraphData/facebook_combined.txt"
	g := graph.NewUndirectedGraph()
	creat_graph_from_file(path, g)
}
