// 実行方法は同階層に移動して以下のコマンドを実行
// １：まとめて実行する場合
// go test -v ./...
// ２：個別に実行する場合
// go test -v add_sub_test.go
package gotest

import (
	"add_sub"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcAdd(t *testing.T) {
	expected := 8
	actual := add_sub.Calc_add(3, 5)
	assert.Equal(t, expected, actual)
}

func TestCalcSub(t *testing.T) {
	expected := -2
	actual := add_sub.Calc_sub(3, 5)
	assert.Equal(t, expected, actual)
}
