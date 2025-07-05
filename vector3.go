package vectors

import (
	"math"
)

type IVector3 interface {
	Add(vec Vector3)
	Sub(vec Vector3)
	Mul(vec Vector3)
	Div(vec Vector3)
	Bounce()
	Normalize()
	AngleRadians() float64
	AngleDegrees() float64
	IsZero() bool
	Magnitude() float64
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector2() Vector2
}

type Vector3 struct {
	IVector3

	X float64
	Y float64
	Z float64
}

func (v *Vector3) Add(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

func (v *Vector3) Sub(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

func (v *Vector3) Mul(vec Vector3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

func (v *Vector3) Div(vec Vector3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}

func (v *Vector3) Bounce() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vector3) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
		v.Z /= magnitude
	}
}

func (v Vector3) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vector3) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	return math.Mod(angle+360, 360)
}

func (v *Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

func (v *Vector3) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v *Vector3) ClampMagnitude(maxValue float64) {
	magnitude := v.Magnitude()

	if magnitude == 0 || magnitude <= maxValue {
		return
	}

	scale := maxValue / magnitude

	v.Mul(Vector3{
		X: scale,
		Y: scale,
		Z: scale,
	})
}

func (v *Vector3) Clear() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}
