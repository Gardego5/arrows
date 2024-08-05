package grid

import (
	"math"

	"github.com/Gardego5/arrows/game/lib/vec"
)

type Grid[T any] struct {
	data          []T
	width, height int
	size          float64
}

func New[T any](w, h int, size ...float64) Grid[T] {
	s := 1.0
	if len(size) > 0 {
		s = size[0]
	}
	return Grid[T]{make([]T, w*h), w, h, s}
}

func (g *Grid[T]) idx(x, y int) int { return x*g.height + y }

func (g *Grid[T]) Get(x, y int) T         { return g.data[g.idx(x, y)] }
func (g *Grid[T]) GetPointer(x, y int) *T { return &g.data[g.idx(x, y)] }
func (g *Grid[T]) Set(x, y int, v T)      { g.data[g.idx(x, y)] = v }

func (g Grid[T]) Width() int     { return g.width }
func (g Grid[T]) Height() int    { return g.height }
func (g Grid[T]) Size() float64  { return g.size }
func (g Grid[T]) Vec() vec.Vec2f { return vec.New(float64(g.width)*g.size, float64(g.height)*g.size) }

func (g Grid[T]) CellAt(x, y float64) (int, int) { return int(x / g.size), int(y / g.size) }
func (g Grid[T]) CellCenterAt(x, y float64) (float64, float64) {
	return (math.Floor(x/g.size) + 0.5) * g.size, (math.Floor(y/g.size) + 0.5) * g.size
}

func (g *Grid[T]) Window(x, y, w, h int) Grid[T] {
	if x+w > g.width || y+h > g.height {
		panic("window out of bounds")
	}

	window := New[T](w, h, g.size)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			window.Set(i, j, g.Get(x+i, y+j))
		}
	}

	return window
}

func (g *Grid[T]) Splat(from Grid[T], x, y int) {
	for i := 0; i < from.width; i++ {
		for j := 0; j < from.height; j++ {
			g.Set(x+i, y+j, from.Get(i, j))
		}
	}
}
