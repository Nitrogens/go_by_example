package main 
import "fmt"
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for id, num := range nums {
		sum += num
		fmt.Println(id, num, sum)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%v -> %v\n", key, value)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}