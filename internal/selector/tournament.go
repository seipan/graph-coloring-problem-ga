package selector

import (
	"errors"
	"math/rand"
	"time"

	ga "github.com/seipan/graph-isomorphism-problem-ga/internal"
)

type Tournament struct {
	NContestants int
}

func (t Tournament) Select(indi ga.Individuals) (ga.Individuals, error) {
	if len(indi) < t.NContestants {
		return nil, errors.New("invalid NSelection: Too large NSelection")
	}
	selected := make(ga.Individuals, len(indi))
	rand.Seed(time.Now().UnixNano())
	for i := range selected {
		winner := indi[rand.Intn(len(indi))]
		for j := 0; j < t.NContestants; j++ {
			next := indi[rand.Intn(len(indi))]
			if winner.Fitness > next.Fitness {
				winner = next
			}
		}
		selected[i] = winner
	}

	return selected, nil
}
