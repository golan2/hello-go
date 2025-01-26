package find_kth_bit_in_nth_binary_string

import (
	"fmt"
	"math"
)

func FindKthBit() {
	for k := 0; k < 16; k++ {
		fmt.Printf("%v", findKthBit(4, k))
	}
}

func findKthBit(n int, k int) byte {
	return recursiveFind(n, k)
}

func recursiveFind(n int, k int) byte {
	if n == 1 {
		return 0
	}
	indexOfMiddleBit := treeSize(n - 1) // size of left subtree is the index of the middle because we start index at 0
	if k < indexOfMiddleBit {
		return recursiveFind(n-1, k)
	} else if k == indexOfMiddleBit {
		return 1
	} else {
		newIndex := treeSize(n) - k - 1
		return not(recursiveFind(n-1, newIndex))
	}
}

func not(b byte) byte {
	if b == 0 {
		return 1
	} else if b == 1 {
		return 0
	} else {
		return 250
	}
}

// number of elements in a tree with the given size
func treeSize(treeHeight int) int {
	return int(math.Pow(2, float64(treeHeight))) - 1
}
