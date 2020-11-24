package InfluenceMaximization

import (
	"bufio"

	"io"
	"os"
	"strconv"
	"strings"

	graph "gonum.org/v1/gonum/graph/simple"
	distuv "gonum.org/v1/gonum/stat/distuv"
)

type UndirectedGraph struct {
	*graph.UndirectedGraph
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PropagationSimulation interface {
	IC_model([]int64, float64) int
}

// random number following uniform distribution U~(0,1)
func rand_01() float64 {
	var uniform *distuv.Uniform = new(distuv.Uniform)
	uniform.Max = 1
	uniform.Min = 0
	return uniform.Rand()
}

// remove duplicate elements from an []int64 array/slice
func remove_dup_int64(d []int64) []int64 {
	check := make(map[int64]int)
	res := make([]int64, 0)
	for _, val := range d {
		check[val] = 1
	}
	for letter := range check {
		res = append(res, letter)
	}
	return res
}

// Creat an undirected graph from file,
// the file format follows datasets from "Stanford Large Network Dataset Collection":
// https://snap.stanford.edu/data/index.html
func CreatUndirectedGraphFromFile(path string, g *graph.UndirectedGraph) {
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

// Independent Cascade propagation model for undirected graphs
// Parameter:
//			seed: set of the original active nodes
//			p:    probability of activation between nodes
// return: the number of activated node in this simulation
func (g *UndirectedGraph) IC_model(seed []int64, p float64) int {
	active := seed
	for {
		var active_i []int64
		for _, node := range active {
			if g.Node(node) != nil {
				neighbors := g.From(node)
				for neighbors.Next() {
					r := rand_01()
					if r <= p {
						active_i = append(active_i, neighbors.Node().ID())
					}
				}
			} else {
				panic("seed does not in the graph!")
			}
		}
		if active_i == nil {
			break
		}
		active = remove_dup_int64(active_i)
		seed = append(seed, active_i...)
	}
	seed = remove_dup_int64(seed)
	return len(seed)
}

func (g *UndirectedGraph) WC_model() {

}

func (g *UndirectedGraph) LT_model() {

}

func IM_Entrance_Undirected(g *graph.UndirectedGraph) int {
	var g_ *UndirectedGraph = new(UndirectedGraph) //(graph.UndirectedGraph -> UndirectedGraph)
	g_.UndirectedGraph = g
	seed := []int64{0, 1, 2, 3, 4, 5}
	return g_.IC_model(seed, 0.01)
}
