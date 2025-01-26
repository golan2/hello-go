package randi

import (
	"fmt"
	"math/rand"
)

//goland:noinspection GoUnusedExportedFunction
func PlayWithRandom() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	r := rand.New(rand.NewSource(0))
	r.Shuffle(10, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	fmt.Printf("%v\n", arr)
}
