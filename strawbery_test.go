package tddbc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_いちごの文字列表現を取得する(t *testing.T) {
	sut := Strawberry{
		kind: "あまおう",
		size: "L",
	}
	actual := sut.String()
	assert.Equal(t, "あまおう: L", actual)
}
