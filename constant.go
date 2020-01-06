package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 21474836488
	fmt.Println(n)

	const m = 500000000

	const d = 3e20 / m
	fmt.Println(d)
	fmt.Println(int64(d))

	const q = 1000
	fmt.Println(math.Sin(q))

}