package tddbc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_あまおうL文字列表現の取得する(t *testing.T) {
	sut := Strawberry{
		kind: "あまおう",
		size: "L",
	}
	actual := sut.ToString()

	assert.Equal(t, "あまおう: L", actual)
}
