package tddbc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_あまおうLの文字列表現を取得する(t *testing.T) {
	sut := Strawberry{
		kind: "あまおう",
		size: "L",
	}
	actual := sut.ToString()

	assert.Equal(t, "あまおう: L", actual)
}
func Test_もういっこMの文字列表現を取得する(t *testing.T) {
	sut := Strawberry{
		kind: "もういっこ",
		size: "M",
	}
	actual := sut.ToString()

	assert.Equal(t, "もういっこ: M", actual)
}
