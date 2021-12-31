package spatialhash

import (
	"reflect"
	"testing"
)

func TestNewSpatialHash(t *testing.T) {
	type args struct {
		dimensions []Unit
		cellsize   Unit
	}
	type want struct {
		dimensions []Unit
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"10x10 grid of 1-sized cells", args{[]Unit{10, 10}, 1}, want{[]Unit{10, 10}}},
		{"100x100 grid of 10-sized cells", args{[]Unit{100, 100}, 10}, want{[]Unit{10, 10}}},
		{"20x20x20 grid of 10-sized cells", args{[]Unit{20, 20, 20}, 10}, want{[]Unit{2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.dimensions, tt.args.cellsize)
			wrongDimension := !reflect.DeepEqual(got.dimensions, tt.want.dimensions)
			if wrongDimension {
				t.Errorf("got %v, want %v", got.dimensions, tt.want.dimensions)
			}
		})
	}
}

func BenchmarkGetNearby(b *testing.B) {
	dim := Unit(128)
	s := New([]Unit{dim, dim, dim}, 16)
	for i := Unit(0); i < 1024; i++ {
		v := Vector{[]Unit{i % dim, i % dim, i % dim}}
		s.Insert(v)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := Unit(i) % dim
		s.GetNearby(Vector{[]Unit{x, x, x}})
	}
}
