package InfluenceMaximization

import (
	"bufio"

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

func CreatGraphFromFile(path string, g *graph.UndirectedGraph) {
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
}
