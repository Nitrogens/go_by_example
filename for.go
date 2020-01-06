package main 

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 573; j <= 578; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}