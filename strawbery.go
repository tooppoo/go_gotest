package tddbc

type Strawberry struct {
	//CAUTION: kindとsizeの値が確定したらenumに変更する。
	kind string
	size string
}

func (berry Strawberry) String() string {
	return berry.kind + ": " + berry.size
}
