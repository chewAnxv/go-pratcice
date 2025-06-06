package main

import (
	"fmt"
	"math"
)

// 全局常量，可以被函数调用
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(d))
}
