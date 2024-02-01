package collections

import "fmt"
import "github.com/samber/lo"

//goland:noinspection GoUnusedExportedFunction
func UseLoMapOnMap() {
	a := []int{1, 2, 3}
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("%v\n", m)

	b := lo.Map(a, func(item int, index int) float32 {
		return float32(item) / 2
	})

	fmt.Printf("%v\n", b)

}
