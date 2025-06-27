package vectors

import (
	"math"
)

type IVector2 interface {
	Add(vec Vector2)
	Sub(vec Vector2)
	Mul(vec Vector2)
	Div(vec Vector2)
	Normalize() (vec Vector2)
	AngleRadians() float64
	AngleDegrees() float64
	IsZero() bool
	Magnitude() float64
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector3() Vector2
}

type Vector2 struct {
	IVector2

	X float64
	Y float64
}

func (v *Vector2) Add(vec Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

func (v *Vector2) Sub(vec Vector2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

func (v *Vector2) Mul(vec Vector2) {
	v.X *= vec.X
	v.Y *= vec.Y
}

func (v *Vector2) Div(vec Vector2) {
	v.X /= vec.X
	v.Y /= vec.Y
}

func (v *Vector2) Normalize() (vec Vector2) {
	vec = *v
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)

	if magnitude != 0 {
		vec.X /= magnitude
		vec.Y /= magnitude
	}

	return vec
}

func (v Vector2) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vector2) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	return math.Mod(angle+360, 360)
}

func (v *Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector2) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v *Vector2) ClampMagnitude(maxValue float64) {
	magnitude := v.Magnitude()

	if magnitude == 0 || magnitude <= maxValue {
		return
	}

	scale := maxValue / magnitude

	v.Mul(Vector2{
		X: scale,
		Y: scale,
	})
}

func (v *Vector2) Clear() {
	v.X = 0
	v.Y = 0
}

func (v Vector2) ToVector3() Vector3 {
	return Vector3{
		X: v.X,
		Y: v.Y,
		Z: 0,
	}
}
