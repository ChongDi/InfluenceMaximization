package InfluenceMaximization

import (
	"bufio"
	"encoding/csv"
	"fmt"

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

func InInt64Slice(haystack []int64, needle int64) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}

	return false
}

type PropagationSimulation interface {
	IC_model([]int64, float64) int
	WC_model([]int64) int
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
		if node_s != node_d { // avoid self-edge
			g.SetEdge(g.NewEdge(graph.Node(node_s), graph.Node(node_d)))
		}
	}
}

// Independent Cascade propagation model for undirected graphs
// Parameter:
//			seed: set of the original active nodes
//			p:    activation probability between nodes
// return: the number of activated node in this simulation
func (g *UndirectedGraph) IC_model(seed []int64, p float64) int {
	active := seed
	for {
		var active_i []int64
		for _, node := range active {
			if g.Node(node) != nil {
				neighbors := g.From(node)
				for neighbors.Next() {
					neighbor := neighbors.Node().ID()
					if !InInt64Slice(seed, neighbor) {
						r := rand_01()
						if r <= p {
							active_i = append(active_i, neighbor)
						}
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

// Weighted Cascade propagation model for undirected graphs
// Parameter:
//			seed: set of the original active nodes
// *the activation probability for node i is set to 1/degree(n)
// return: the number of activated node in this simulation
func (g *UndirectedGraph) WC_model(seed []int64) int {
	active := seed
	for {
		var active_i []int64
		for _, node := range active {
			if g.Node(node) != nil {
				neighbors := g.From(node)
				for neighbors.Next() {
					neighbor := neighbors.Node().ID()
					if !InInt64Slice(seed, neighbor) {
						r := rand_01()
						p := 1.0 / float64(g.From(neighbor).Len())
						if r <= p {
							active_i = append(active_i, neighbor)
						}
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

func (g *UndirectedGraph) LT_model() {

}

func IMEntranceUndirected(g *graph.UndirectedGraph) int {
	var g_ *UndirectedGraph = new(UndirectedGraph) //(graph.UndirectedGraph -> UndirectedGraph)
	g_.UndirectedGraph = g
	seed := []int64{0}
	return g_.IC_model(seed, 0.01)
}

func ModelTest(g_ *graph.UndirectedGraph) {
	MCNum := 10000

	var g *UndirectedGraph = new(UndirectedGraph) //(graph.UndirectedGraph -> UndirectedGraph)
	var g_p PropagationSimulation
	g_p = g
	g.UndirectedGraph = g_
	nodes := g.Nodes()
	var max_node int64
	var max_range float64

	f, err := os.Create("test_GrQc_WC.csv") // write result to .csv file
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for nodes.Next() {
		node := nodes.Node().ID()
		seed := append([]int64{}, node)
		iter := make(chan int, MCNum)
		results := make(chan int, MCNum)
		for i := 0; i < 64; i++ { // workers
			go worker(g_p, iter, results, seed, 0.01)
		}
		for i := 0; i < MCNum; i++ {
			iter <- i
		}
		close(iter)
		sum_ := 0
		for i := 0; i < MCNum; i++ {
			res := <-results
			sum_ += res
		}
		avg_ := float64(sum_) / float64(MCNum)
		if avg_ >= max_range {
			max_node = node
			max_range = avg_
		}
		fmt.Println(node, avg_, " | ", max_node, max_range)
		res := []string{strconv.FormatInt(node, 10), strconv.FormatFloat(avg_, 'E', -1, 64)}
		w.Flush()
		err := w.Write(res)
		if err != nil {
			check(err)
		}
	}
	fmt.Println(max_node, max_range)
	res := []string{strconv.FormatInt(max_node, 10), strconv.FormatFloat(max_range, 'E', -1, 64)}
	w.Write(res)
	w.Flush()
}

func worker(g PropagationSimulation, jobs <-chan int, results chan<- int, seed []int64, p float64) {
	for i := range jobs {
		_ = i
		// results <- g.IC_model(seed, p)
		results <- g.WC_model(seed)
	}
}
