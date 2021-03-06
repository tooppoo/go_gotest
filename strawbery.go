package strawberry

type Strawberry struct {
	//CAUTION: kindとsizeの値が確定したらenumに変更する。
	kind string
	size string
}

func New(kind string, weight int) Strawberry {
	return Strawberry{
		kind,
		"S",
	}
}

func (berry Strawberry) String() string {
	return berry.kind + ": " + berry.size
}
