package main

import (
	"math/rand"

	"github.com/seipan/graph-isomorphism-problem-ga/internal"
)

type Variables struct {
	Individual Individual
	Graph      *Graph
}

type Individual []bool

func (ind Variables) Initialization() internal.Genome {
	numNodes := len(ind.Individual)
	genesind := make(Individual, numNodes)
	for i := 0; i < numNodes; i++ {
		genesind[i] = rand.Intn(2) == 1
	}
	genes := Variables{Individual: genesind, Graph: ind.Graph}
	return genes
}

func (ind Variables) Fitness() float64 {
	if isClique(ind.Individual, ind.Graph) {
		return float64(cliqueSize(ind.Individual))
	} else {
		return 0.0 // またはペナルティとして負の値を返す
	}
}

func (v Variables) Mutation() {
	ind := v.Individual
	mutationRate := 0.05
	for i := range ind {
		if rand.Float64() < mutationRate {
			ind[i] = !ind[i]
		}
	}
}

func (v Variables) Crossover(partner internal.Genome) internal.Genome {
	ind := v.Individual
	parent2 := partner.(Variables)
	childind := make(Individual, len(ind))
	child := Variables{Individual: childind, Graph: v.Graph}
	crossoverPoint := rand.Intn(len(ind))
	for i := 0; i < len(ind); i++ {
		if i < crossoverPoint {
			child.Individual[i] = ind[i]
		} else {
			child.Individual[i] = parent2.Individual[i]
		}
	}
	return child
}
