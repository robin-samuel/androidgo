package pair

import (
	"fmt"
	"hash/fnv"
)

type Pair[A, B comparable] struct {
	First  A
	Second B
}

func (p *Pair[A, B]) Equals(other interface{}) bool {
	if otherPair, ok := other.(*Pair[A, B]); ok {
		return p.First == otherPair.First && p.Second == otherPair.Second
	}
	return false
}

func (p *Pair[A, B]) HashCode() int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v%v", p.First, p.Second)))
	return int(h.Sum32())
}
