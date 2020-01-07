package main
import (
	"errors"
	"fmt"
)
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot work with 42")
	}
	return arg + 3, nil
}
type argError struct {
	arg  int
	prob string
}
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot work with it"}
	}
	return arg + 3, nil
}
func main() {
	for _, i := range []int{7, 42} {
		if res, err := f1(i); err != nil {
			fmt.Println("f1 failed:", err)
		} else {
			fmt.Println("f1 worked:", res)
		}
	}
	for _, i := range []int{7, 42} {
		if res, err := f2(i); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", res)
		}
	}
	_, err := f2(42)
	if val, ok := err.(*argError); ok {
		fmt.Println(val.arg)
		fmt.Println(val.prob)
	}
}