package internal

type Selector interface {
	Select(Individuals) (Individuals, error)
}
