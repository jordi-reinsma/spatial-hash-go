package spatialhash

type Unit uint32

type SpatialHash struct {
	grids      [][]Object
	dimensions []Unit
	cellsize   Unit
}

func New(dimensions []Unit, cellsize Unit) *SpatialHash {
	var gridsize Unit = 1
	dims := make([]Unit, len(dimensions))
	for i := range dimensions {
		size := dimensions[i] / cellsize
		dims[i] = size
		gridsize *= size
	}
	grids := make([][]Object, gridsize)
	return &SpatialHash{grids: grids, dimensions: dims, cellsize: cellsize}
}

func (s *SpatialHash) Clear() {
	for i := range s.grids {
		s.grids[i] = nil
	}
}

func (s *SpatialHash) Insert(o Object) {
	grids := o.GetGrids(s.dimensions, s.cellsize)
	for _, grid := range grids {
		s.grids[grid] = append(s.grids[grid], o)
	}
}

func (s *SpatialHash) GetNearby(o Object) (nearby []Object) {
	grids := o.GetGrids(s.dimensions, s.cellsize)
	for _, grid := range grids {
		nearby = append(nearby, s.grids[grid]...)
	}
	return
}
