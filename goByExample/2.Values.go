package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7/3 =", 7/3)         // 整数除法
	fmt.Println("7.0/3.0 =", 7.0/3.0) // 浮点数除法

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
