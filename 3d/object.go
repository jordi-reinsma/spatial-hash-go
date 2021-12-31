package spatialhash3d

type Object interface {
	// GetGrids transforms the object spatial position into indexes of a flattened
	// 3-dimensional array of grids, where each grid has sides of `cellsize` length.
	GetGrids(heigth, width, depth, cellsize Unit) []Unit
}

func Flatten(x, y, z, width, depth Unit) Unit {
	return x + (y+z*depth)*width
}

type Vector struct {
	X, Y, Z Unit
}

func (p Vector) GetGrids(heigth, width, depth, cellsize Unit) []Unit {
	x := p.X / cellsize
	y := p.Y / cellsize
	z := p.Z / cellsize
	return []Unit{Flatten(x, y, z, width, depth)}
}
