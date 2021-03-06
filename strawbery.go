package strawberry

import "errors"

type Strawberry struct {
	//CAUTION: kindとsizeの値が確定したらenumに変更する。
	kind string
	size string
}

// 重さについては、ステークホルダーとも整数の範囲でOKと合意した（という設定）
func New(kind string, weight uint) (*Strawberry, error) {
	if weight == 0 {
		return nil, errors.New("重さは1以上の整数を入力してください")
	}
	return &Strawberry{
		kind,
		"S",
	}, nil
}

func (berry Strawberry) String() string {
	return berry.kind + ": " + berry.size
}
