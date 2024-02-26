package collections

type SetOfStrings interface {
	Put(value string)
	Contains(value string) bool
}

func NewSetOfStrings(size int) SetOfStrings {
	return &setOfStrings{
		data: make(map[string]struct{}, size),
	}
}

type setOfStrings struct {
	data map[string]struct{}
}

func (s setOfStrings) Put(value string) {
	s.data[value] = struct{}{}
}

func (s setOfStrings) Contains(value string) (exists bool) {
	_, exists = s.data[value]
	return
}
