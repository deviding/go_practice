package multi_div

import (
	"errors"
)

// 掛け算を行う関数
func Calc_multi(a int, b int) int {
	return a * b
}

// 割り算を行う関数
func Calc_div(a int, b int) (float32, error) {
	if b == 0 {
		return 0.0, errors.New("b is zero")
	}
	return float32(float32(a) / float32(b)), nil
}
