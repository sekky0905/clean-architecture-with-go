package main

import "fmt"

func main() {
	result := Sum(2,3)
	fmt.Println(result)
}

// Sum は、引数で受け取った2数を合計して返す。
func Sum(a, b int)int {
	return a + b
}