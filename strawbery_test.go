package strawberry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_いちごの文字列表現を取得する(t *testing.T) {
	cases := []struct {
		kind     string
		size     string
		expected string
	}{
		{
			kind:     "あまおう",
			size:     "L",
			expected: "あまおう: L",
		},
		{
			kind:     "あまおう",
			size:     "LL",
			expected: "あまおう: LL",
		},
		{
			kind:     "とちおとめ",
			size:     "L",
			expected: "とちおとめ: L",
		},
		{
			kind:     "もういっこ",
			size:     "M",
			expected: "もういっこ: M",
		},
		{
			kind:     "あまおう",
			size:     "S",
			expected: "あまおう: S",
		},
	}
	for _, test := range cases {
		sut := Strawberry{kind: test.kind, size: test.size}

		assert.Equal(t, test.expected, sut.String())
	}
}

func Test_重さが8gのあまおうの場合の文字列表現を取得(t *testing.T) {
	sut, _ := New("あまおう", 8)

	assert.Equal(t, "あまおう: S", sut.String())
}

func Test_重さに応じたサイズを取得する_25g_LLサイズ(t *testing.T) {
	assert.Equal(t, "LL", WeightToSize(25))
}

func Test_重さが0gの場合はエラーになる(t *testing.T) {
	_, err := New("あまおう", 0)

	assert.NotNil(t, err)
}
