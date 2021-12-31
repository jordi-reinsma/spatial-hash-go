package spatialhash2d

type Object interface {
	// GetGrids transforms the object spatial position into indexes of a flattened
	// 2-dimensional array of grids, where each grid has sides of `cellsize` length.
	GetGrids(heigth, width, cellsize Unit) []Unit
}

func Flatten(x, y, width Unit) Unit {
	return x + y*width
}

type Vector struct {
	X, Y Unit
}

func (v Vector) GetGrids(heigth, width, cellsize Unit) []Unit {
	x := v.X / cellsize
	y := v.Y / cellsize
	return []Unit{Flatten(x, y, width)}
}

type Circle struct {
	X, Y, Radius Unit
}

func (c Circle) GetGrids(heigth, width, cellsize Unit) []Unit {
	var (
		minX, maxX Unit = 0, heigth
		minY, maxY Unit = 0, width
	)
	if c.X > c.Radius {
		minX = (c.X - c.Radius) / cellsize
	}
	if c.Y > c.Radius {
		minY = (c.Y - c.Radius) / cellsize
	}
	if c.X+c.Radius < maxX*cellsize {
		maxX = (c.X + c.Radius) / cellsize
	}
	if c.Y+c.Radius < maxY*cellsize {
		maxX = (c.X + c.Radius) / cellsize
	}

	grids := make([]Unit, 0, 4)
	// Top left point
	grids = append(grids, Flatten(minX, minY, width))
	// Top right point
	grid := Flatten(maxX, maxY, width)
	if grid != grids[0] {
		grids = append(grids, grid)
	}
	// Bottom right point
	grid = Flatten(maxX, maxY, width)
	if grid != grids[1] {
		grids = append(grids, grid)
	}
	// Bottom left point
	grid = Flatten(minX, maxY, width)
	if grid != grids[0] && grid != grids[len(grids)-1] {
		grids = append(grids, grid)
	}

	return grids
}
