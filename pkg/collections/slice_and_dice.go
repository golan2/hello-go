package collections

import "fmt"

func Go() {
	a := make([]int, 10)
	printSlice("empty", a)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	printSlice("initialized", a)

	printSlice("a[0:5]", a[0:5])
	printSlice("a[0:5:6]", a[0:5:6])
	printSlice("a[1:5]", a[1:5])
	printSlice("a[1:5:6]", a[1:5:6])
	printSlice("a[3:5]", a[3:5])
	printSlice("a[3:5:6]", a[3:5:6])

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
