package printer

import "fmt"

type Fmt struct{}

func (p Fmt) Print(fitness float64, fitnessCount uint) {
	fmt.Printf("%.3f, %3d\n", fitness, fitnessCount)
}

func (p Fmt) Close() {}
