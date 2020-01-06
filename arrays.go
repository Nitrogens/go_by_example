package main 
import "fmt"
func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	fmt.Println("cap:", cap(a))

	b := [...]int{5, 7, 3, 5, 7, 8}
	fmt.Println(b)

	c := [5]int{}
	fmt.Println(c)

	d := [6]int{5, 7, 8}
	fmt.Println(d)

	e := [...]int{}
	fmt.Println(e)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	f := []int{1, 2, 3, 4}
	fmt.Println(f)
}