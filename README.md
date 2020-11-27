### Influence Maximization with Go

This project provides models for influence maximization. Including: 

1. Independent cascade model.

2. Weighted cascade model.

3. Linear threshold model. (To be finished)

### Usage
```
import (
        IM .../IMpackage
)
```
1. func CreateUndirectedGraphFromFile
```
func CreateUndirectedGraphFromFile(path string) *UndirectedGraph
```
CreateUndirectedGraphFromFile returns an undirected graph from file, the file format follows datasets from ["Stanford Large Network Dataset Collection"](https://snap.stanford.edu/data/index.html). 

2. func CreateDirectedGraphFromFile
```
func CreateDirectedGraphFromFile(path string) *DirectedGraph
```
CreateUndirectedGraphFromFile returns a directed graph from file, the file format follows datasets from ["Stanford Large Network Dataset Collection"](https://snap.stanford.edu/data/index.html). 

3. func (g *UndirectedGraph) IC_model
```
func (g *UndirectedGraph) IC_model(seed []int64, p float64) int
```
IC_model returns the number of activated node in an undirected graph under independent cascade propagation model.

4. func (g *DirectedGraph) IC_model
```
func (g *DirectedGraph) IC_model(seed []int64, p float64) int
```
IC_model returns the number of activated node in a directed graph under independent cascade propagation model.

5. func (g *UndirectedGraph) WC_model
```
func (g *UndirectedGraph) WC_model(seed []int64) int
```
WC_model returns the number of activated node in an undirected graph under weighted cascade propagation model.

6. func (g *DirectedGraph) WC_model
```
func (g *DirectedGraph) WC_model(seed []int64) int
```
WC_model returns the number of activated node in a directed graph under weighted cascade propagation model.


**LICENSE**

The MIT License (MIT)
