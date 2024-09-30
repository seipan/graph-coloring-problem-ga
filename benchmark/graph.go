package main

type Graph struct {
	Nodes int
	Edges [][]int
}

func NewGraph(nodes int) *Graph {
	edges := make([][]int, nodes)
	return &Graph{Nodes: nodes, Edges: edges}
}

func (g *Graph) AddEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
	g.Edges[v] = append(g.Edges[v], u)
}

func isClique(ind Individual, graph *Graph) bool {
	nodes := []int{}
	for i, included := range ind {
		if included {
			nodes = append(nodes, i)
		}
	}

	// ノード集合がクリークであるかを判定
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			u, v := nodes[i], nodes[j]
			if !isConnected(u, v, graph) {
				return false
			}
		}
	}
	return true
}

func isConnected(u, v int, graph *Graph) bool {
	for _, neighbor := range graph.Edges[u] {
		if neighbor == v {
			return true
		}
	}
	return false
}

func cliqueSize(ind Individual) int {
	size := 0
	for _, included := range ind {
		if included {
			size++
		}
	}
	return size
}
