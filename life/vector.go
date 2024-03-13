package life

import (
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func (vec *Vector2) GetLen() float64 {
	len_squared := vec.X*vec.X + vec.Y*vec.Y
	return math.Sqrt(len_squared)
}

func (vec *Vector2) Normalize() {
	len := vec.GetLen()
	vec.X /= len
	vec.Y /= len
}

func GetLenBetweenVectors(v1 *Vector2, v2 *Vector2) float64 {
	v3 := Vector2{v1.X - v2.X, v1.Y - v2.Y}
	return v3.GetLen()
}
