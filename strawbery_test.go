package tddbc

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
	}

	for _, test := range cases {
		sut := Strawberry{kind: test.kind, size: test.size}

		assert.Equal(t, test.expected, sut.String())
	}
}
