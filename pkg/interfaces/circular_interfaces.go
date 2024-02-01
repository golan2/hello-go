package interfaces

type Aggregator[T any] interface {
	Aggregate(rawData []T) ([]int, error)
}

type IntAggregator interface {
	Aggregate(rawData []int) ([]int, error)
}

type SimpleIntAggregator struct {
}

func (s *SimpleIntAggregator) Aggregate(rawData []int) ([]int, error) {
	return rawData, nil
}
