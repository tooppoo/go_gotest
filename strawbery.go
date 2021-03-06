package strawberry

type Strawberry struct {
	//CAUTION: kindとsizeの値が確定したらenumに変更する。
	kind string
	size string
}

// 重さについては、ステークホルダーとも整数の範囲でOKと合意した（という設定）
func New(kind string, weight int) Strawberry {
	return Strawberry{
		kind,
		"S",
	}
}

func (berry Strawberry) String() string {
	return berry.kind + ": " + berry.size
}
