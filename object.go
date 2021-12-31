package spatialhash

type Object interface {
	// GetGrids transforms the object spatial position into indexes of a flattened N-dimensional
	// array of grids, where each grid has `len(dimensions)` sides of `cellsize` length.
	GetGrids(dimensions []Unit, cellsize Unit) []Unit
}

func Flatten(position, dimensions []Unit, cellsize Unit) Unit {
	var res, size Unit = 0, 1
	for i := range position {
		if i == 0 {
			res = position[0] / cellsize
			continue
		}
		size *= dimensions[i]
		res += (position[i] / cellsize) * size
	}
	return res
}

// Vector implements Object interface
type Vector struct {
	Position []Unit
}

func (v Vector) GetGrids(dimensions []Unit, cellsize Unit) []Unit {
	return []Unit{Flatten(v.Position, dimensions, cellsize)}
}
