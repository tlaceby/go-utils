package math

import (
	"math"

	lib "github.com/tlaceby/go-utils"
)

type Vec2D[N lib.Number] struct {
	x N
	y N
}

func New[N lib.Number]() *Vec2D[N] {
	return &Vec2D[N]{}
}

func (v *Vec2D[N]) Zero() {
	v.x = 0
	v.y = 0
}

func (v *Vec2D[N]) Plus(o Vec2D[N]) {
	v.x += o.x
	v.y += o.y
}

func (v *Vec2D[N]) Minus(o Vec2D[N]) {
	v.x -= o.x
	v.y -= o.y
}

func (v *Vec2D[N]) Divide(o Vec2D[N]) {
	v.x /= o.x
	v.y /= o.y
}

func (v *Vec2D[N]) Cross(o Vec2D[N]) {
	v.x *= o.x
	v.y *= o.y
}

func (v *Vec2D[N]) Magnitude() N {
	x := float64(v.x * v.x)
	y := float64(v.y * v.y)
	return N(math.Sqrt(x + y))
}

// Tranforms current vector to a unit vector
func (v *Vec2D[N]) Normalize() {
	mag := v.Magnitude()
	v.x /= mag
	v.y /= mag
}

func (v *Vec2D[N]) Dot(o Vec2D[N]) N {
	return (v.x * o.x) + (v.y * o.y)
}
