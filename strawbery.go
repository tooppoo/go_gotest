package tddbc

type Strawberry struct {
	kind string
	size string
}

func (berry Strawberry) ToString() string {
	return "あまおう: L"
}