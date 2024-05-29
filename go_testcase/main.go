package main

import (
	"add_sub"
	"fmt"
	"multi_div"
)

func main() {
	// 引数設定
	x := 3
	y := 5

	// 足し算と引き算
	fmt.Printf("%v + %v = %v\n", x, y, add_sub.Calc_add(x, y))
	fmt.Printf("%v - %v = %v\n", x, y, add_sub.Calc_sub(x, y))

	// 掛け算と割り算
	fmt.Printf("%v * %v = %v\n", x, y, multi_div.Calc_multi(x, y))
	res, err := multi_div.Calc_div(x, y)
	if err == nil {
		fmt.Printf("%v / %v = %v\n", x, y, res)
	} else {
		fmt.Println(err.Error())
	}
}
