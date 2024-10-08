package internal

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type GA struct {
	GAConfig
	Population
	Selector       Selector
	Printer        Printer
	BestIndividual Individual
	PrintCallBack  func()
}

type GAConfig struct {
	PopulationSize uint
	NGenerations   uint
	CrossoverRate  float64
	MutationRate   float64
	ParallelEval   bool
}

type Population struct {
	Individuals Individuals
	Generations uint
}

func NewGA(gaConfig GAConfig, selector Selector, printer Printer) *GA {
	return &GA{
		GAConfig: gaConfig,
		Selector: selector,
		Printer:  printer,
	}
}

func (ga *GA) initPopulation(g Genome) {
	indis := make(Individuals, ga.PopulationSize)
	for i := range indis {
		indis[i].Chromosome = g.Initialization()
	}
	indis.Evaluate(ga.ParallelEval)
	ga.Population.Generations = 0
	ga.Population.Individuals = indis
	ga.Population.Individuals.SortByFitness()
	ga.BestIndividual = ga.Population.Individuals[0]
}

func (ga *GA) evolve() error {
	ga.Generations++
	rand.Seed(time.Now().UnixNano())
	offSprings := make(Individuals, ga.PopulationSize)
	selected, err := ga.Selector.Select(ga.Population.Individuals)
	if err != nil {
		log.Fatal(err)
	}

	for i := range offSprings {
		if i == len(selected)-1 {
			offSprings[i] = selected[i].Clone()
		} else {
			if rand.Float64() < ga.CrossoverRate {
				offSprings[i].Chromosome = selected[i].Chromosome.Crossover(selected[i+1].Chromosome)
			} else {
				offSprings[i] = selected[i].Clone()
			}
		}
		if rand.Float64() < ga.MutationRate {
			offSprings[i].Chromosome.Mutation()
		}
	}

	offSprings.Evaluate(ga.ParallelEval)
	offSprings.SortByFitness()
	ga.updateBest(offSprings[0])
	ga.Population.Individuals = offSprings.Clone()
	return nil
}

func (ga *GA) updateBest(indi Individual) {
	if ga.BestIndividual.Fitness > indi.Fitness {
		ga.BestIndividual = indi.Clone()
	}
}

func (ga *GA) Minimize(g Genome) error {
	ga.initPopulation(g)

	for i := uint(1); i <= ga.NGenerations; i++ {
		if err := ga.evolve(); err != nil {
			return err
		}
		if ga.PrintCallBack != nil {
			ga.PrintCallBack()
		} else {
			fmt.Printf("Generation %3d: Fitness=%.3f Solution=%.3f\n", i, ga.BestIndividual.Fitness, ga.BestIndividual.Chromosome)
		}
	}
	return nil
}
