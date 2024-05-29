// 実行方法は同階層に移動して以下のコマンドを実行
// １：まとめて実行する場合
// go test -v ./...
// ２：個別に実行する場合
// go test -v multi_div_test.go
package gotest

import (
	"errors"
	"multi_div"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcMulti(t *testing.T) {
	expected := 15
	actual := multi_div.Calc_multi(3, 5)
	assert.Equal(t, expected, actual)
}

func TestCalcDiv(t *testing.T) {
	type testData struct {
		name string
		a    int
		b    int
		res  float32
		err  error
	}

	// テストデータ作成
	test_data := []testData{
		{name: "正常系1", a: 10, b: 5, res: 2.0, err: nil},
		{name: "正常系2", a: 3, b: 5, res: 0.6, err: nil},
		{name: "異常系3", a: 3, b: 0, res: 0.0, err: errors.New("b is zero")},
	}

	// テスト実施
	for _, data := range test_data {
		// それぞれのテストデータをサブテストとして実施
		t.Run(data.name, func(t *testing.T) {
			actual, err := multi_div.Calc_div(data.a, data.b)
			if data.err == nil {
				assert.Equal(t, data.res, actual)
			} else {
				assert.Equal(t, data.err.Error(), err.Error())
			}
		})
	}
}
