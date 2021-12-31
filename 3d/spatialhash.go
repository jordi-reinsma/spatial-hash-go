package spatialhash3d

type Unit uint64

type SpatialHash struct {
	grids    [][]Object
	height   Unit
	width    Unit
	depth    Unit
	cellsize Unit
}

func New(height, width, depth, cellsize Unit) *SpatialHash {
	h := height / cellsize
	w := width / cellsize
	d := depth / cellsize
	grids := make([][]Object, h*w*d)
	return &SpatialHash{grids: grids, height: h, width: w, depth: d, cellsize: cellsize}
}

func (s *SpatialHash) Clear() {
	for i := range s.grids {
		s.grids[i] = nil
	}
}

func (s *SpatialHash) Insert(o Object) {
	grids := o.GetGrids(s.height, s.width, s.depth, s.cellsize)
	for _, grid := range grids {
		s.grids[grid] = append(s.grids[grid], o)
	}
}

func (s *SpatialHash) GetNearby(o Object) (nearby []Object) {
	grids := o.GetGrids(s.height, s.width, s.depth, s.cellsize)
	for _, grid := range grids {
		nearby = append(nearby, s.grids[grid]...)
	}
	return
}
