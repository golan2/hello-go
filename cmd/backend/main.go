package main

import (
	"fmt"
	"hello-go/pkg/interfaces"
)

func main() {
	sa := interfaces.SimpleIntAggregator{}
	res := Aggregate(&sa, nil)
	fmt.Printf("%v\n", res)
}

func Aggregate(agg interfaces.Aggregator[int], values []int) []int {
	aggregate, _ := agg.Aggregate(values)
	return aggregate
}
