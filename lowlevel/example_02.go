package main

import "fmt"

func Example02() {
	var (
		a = 0.1
		b = 0.2
		c = 0.3
	)

	fmt.Println(a+b == c)
	fmt.Printf("a:   %.20f\n", a)
	fmt.Printf("b:   %.20f\n", b)
	fmt.Printf("a+b: %.20f\n", a+b)
	fmt.Printf("c:   %.20f\n", c)


	const (
		ac = 0.1
		bc = 0.2
		cc = 0.3
	)

	fmt.Println(ac+bc == cc)
	fmt.Printf("ac:    %.20f\n", ac)
	fmt.Printf("bc:    %.20f\n", bc)
	fmt.Printf("ac+bc: %.20f\n", ac+bc)
	fmt.Printf("cc:    %.20f\n", cc)
}
