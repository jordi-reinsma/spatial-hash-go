package spatialhash2d

import (
	"testing"
)

func TestNewSpatialHash(t *testing.T) {
	type args struct {
		heigth   Unit
		width    Unit
		cellsize Unit
	}
	type want struct {
		grids int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"10x10 grid of 1-sized cells", args{10, 10, 1}, want{100}},
		{"100x100 grid of 20-sized cells", args{100, 100, 20}, want{25}},
		{"20x20 grid of 10-sized cells", args{20, 20, 10}, want{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.heigth, tt.args.width, tt.args.cellsize)
			if len(got.grids) != tt.want.grids {
				t.Errorf("got %v, want %v", len(got.grids), tt.want.grids)
			}
		})
	}
}

func BenchmarkGetNearby(b *testing.B) {
	dim := Unit(128)
	s := New(dim, dim, 16)
	for i := Unit(0); i < 1024; i++ {
		v := Vector{i % dim, i % dim}
		s.Insert(v)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := Unit(i) % dim
		s.GetNearby(Vector{x, x})
	}
}
