package main

import (
	"math/rand"
	"spatialhash"
)

func main() {
	s := spatialhash.New([]spatialhash.Unit{128, 128, 128}, 16)
	for i := 0; i < 1000; i++ {
		v := spatialhash.Vector{Position: []spatialhash.Unit{
			spatialhash.Unit(rand.Intn(128)), spatialhash.Unit(rand.Intn(128)), spatialhash.Unit(rand.Intn(128)),
		}}
		s.Insert(v)
	}

	for i := 0; i < 200; i++ {
		x := spatialhash.Unit(i) % 128
		s.GetNearby(spatialhash.Vector{Position: []spatialhash.Unit{x, x, x}})
	}
}
