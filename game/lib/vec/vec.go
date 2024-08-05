package vec

import "math"

type number interface {
	comparable
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
type Vec2[T number] struct{ X, Y T }
type Vec2f = Vec2[float64]
type Vec2i = Vec2[int]

func New[T number](x, y T) Vec2[T] { return Vec2[T]{x, y} }

func (v *Vec2[T]) Add(o Vec2[T]) { v.X += o.X; v.Y += o.Y }
func (v *Vec2[T]) Sub(o Vec2[T]) { v.X -= o.X; v.Y -= o.Y }
func (v *Vec2[T]) Mul(o Vec2[T]) { v.X *= o.X; v.Y *= o.Y }
func (v *Vec2[T]) Div(o Vec2[T]) { v.X /= o.X; v.Y /= o.Y }

func (v Vec2[T]) Dot(o Vec2[T]) T   { return v.X*o.X + v.Y*o.Y }
func (v Vec2[T]) Cross(o Vec2[T]) T { return v.X*o.Y - v.Y*o.X }
func (v Vec2[T]) Len() float64      { return math.Hypot(float64(v.X), float64(v.Y)) }
func (v Vec2[T]) Tuple() (T, T)     { return v.X, v.Y }

func (v Vec2[T]) Normalized() Vec2[T] {
	l := v.Len()
	return Vec2[T]{T(float64(v.X) / l), T(float64(v.Y) / l)}
}

func (v Vec2[T]) Angle() float64 {
	return math.Atan2(float64(v.Y), float64(v.X))
}

func (v Vec2[T]) Rotate(a float64) Vec2[T] {
	s, c := math.Sincos(a)
	return Vec2[T]{T(float64(v.X)*c - float64(v.Y)*s), T(float64(v.X)*s + float64(v.Y)*c)}
}

func (v Vec2[T]) Distance(o Vec2[T]) float64 {
	return math.Hypot(float64(v.X-o.X), float64(v.Y-o.Y))
}
