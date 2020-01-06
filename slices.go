package main 
import "fmt"
func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, 5)
	d := make([]string, 6)
	e := make([]string, 7)
	copy(c, s)
	copy(d, s)
	copy(e, s)
	fmt.Println("c:", c, "cap:", cap(c))
	fmt.Println("d:", d, "cap:", cap(d))
	fmt.Println("e:", e, "cap:", cap(e))

	f := append(c, "k")
	g := append(e, "k")
	fmt.Println("f:", f, "cap:", cap(f))
	fmt.Println("g:", g, "cap:", cap(g))
	
	h := append(f, "i")
	fmt.Println("f:", f)
	fmt.Println("h:", h)
	h[0] = "b"
	fmt.Println("f:", f)
	fmt.Println("h:", h)

	fmt.Println("e:", e, "cap:", cap(e))
	fmt.Println("g:", g, "cap:", cap(g))
	g[0] = "b"
	fmt.Println("e:", e, "cap:", cap(e))
	fmt.Println("g:", g, "cap:", cap(g))
}