package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/set"
)

func main()  {
	s := set.New()
	s.Insert(1)
	s.Insert(5)
	s.Insert(3)
	s.Insert(4)
	s.Insert(2)

	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		fmt.Printf("%v\n", iter.Value())
	}
}
