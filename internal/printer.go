package internal

type Printer interface {
	Print(fitness float64, fitnessCount uint)
	Close()
}
