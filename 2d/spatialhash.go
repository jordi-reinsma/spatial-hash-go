package spatialhash2d

type Unit uint64

type SpatialHash struct {
	grids    [][]Object
	height   Unit
	width    Unit
	cellsize Unit
}

func New(height, width, cellsize Unit) *SpatialHash {
	h := height / cellsize
	w := width / cellsize
	grids := make([][]Object, h*w)
	return &SpatialHash{grids: grids, height: h, width: w, cellsize: cellsize}
}

func (s *SpatialHash) Clear() {
	for i := range s.grids {
		s.grids[i] = nil
	}
}

func (s *SpatialHash) Insert(o Object) {
	grids := o.GetGrids(s.height, s.width, s.cellsize)
	for _, grid := range grids {
		s.grids[grid] = append(s.grids[grid], o)
	}
}

func (s *SpatialHash) GetNearby(o Object) (nearby []Object) {
	grids := o.GetGrids(s.height, s.width, s.cellsize)
	for _, grid := range grids {
		nearby = append(nearby, s.grids[grid]...)
	}
	return
}
