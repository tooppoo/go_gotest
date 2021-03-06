package tddbc

type Strawberry struct {
	kind string
	size string
}

func (berry Strawberry) String() string {
	return berry.kind + ": " + berry.size
}
