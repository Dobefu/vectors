package vectors

import (
	"math"
)

// IVector2 is the interface for a 2D vector.
// It defines all the operations that can be performed on a 2D vector.
type IVector2 interface {
	Add(vec Vector2)
	Sub(vec Vector2)
	Mul(vec Vector2)
	Div(vec Vector2)
	Bounce()
	Normalize()
	AngleRadians() float64
	AngleDegrees() float64
	IsZero() bool
	Magnitude() float64
	MagnitudeSquared() float64
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector3() Vector3
}

// Vector2 represents a 2D vector with X and Y coordinates.
// It implements the IVector2 interface and provides methods for vector operations.
type Vector2 struct {
	X float64 // X coordinate of the vector.
	Y float64 // Y coordinate of the vector.
}

// Add adds the values of another vector to this one.
func (v *Vector2) Add(vec Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

// Sub subtracts the values of another vector from this one.
func (v *Vector2) Sub(vec Vector2) {
	v.X -= vec.X
	v.Y -= vec.Y
}

// Mul multiplies this vector by another vector component-wise.
func (v *Vector2) Mul(vec Vector2) {
	v.X *= vec.X
	v.Y *= vec.Y
}

// Div divides this vector by another vector component-wise.
// Note: Division by zero will result in NaN or Inf values.
func (v *Vector2) Div(vec Vector2) {
	v.X /= vec.X
	v.Y /= vec.Y
}

// Bounce inverts the direction of the vector by negating all axes.
// This is equivalent to multiplying the vector by -1.
func (v *Vector2) Bounce() {
	v.X = -v.X
	v.Y = -v.Y
}

// Normalize scales the vector to have a magnitude of 1 while preserving its direction.
// If the vector is already zero, it remains unchanged.
// A normalized vector is also called a unit vector.
func (v *Vector2) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)

	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
	}
}

// AngleRadians returns the angle of the vector in radians relative to the positive X-axis.
// The angle is measured counterclockwise from the positive X-axis.
// Returns a value in the range [-π, π].
func (v Vector2) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleDegrees returns the angle of the vector in degrees relative to the positive X-axis.
// The angle is measured counterclockwise from the positive X-axis.
// Returns a value in the range [0, 360).
func (v Vector2) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	return math.Mod(angle+360, 360)
}

// IsZero returns true if all axes are zero.
// This indicates the vector has no magnitude and no direction.
func (v Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// Magnitude returns the length (magnitude) of the vector.
func (v Vector2) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

// MagnitudeSquared returns the squared magnitude of the vector.
// This is faster for magnitude comparisons, since it avoids the square root.
func (v Vector2) MagnitudeSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y)
}

// ClampMagnitude limits the magnitude of the vector to a maximum value.
// If the current magnitude exceeds maxValue, the vector is scaled down proportionally.
// If the vector is zero or already within the limit, no change is made.
// This preserves the direction while limiting the length.
func (v *Vector2) ClampMagnitude(maxValue float64) {
	magnitude := v.Magnitude()

	if magnitude == 0 || magnitude <= maxValue {
		return
	}

	scale := maxValue / magnitude

	v.X *= scale
	v.Y *= scale
}

// Clear sets the vector to zero on all axes.
// This is equivalent to setting all axes to 0.
func (v *Vector2) Clear() {
	v.X = 0
	v.Y = 0
}

// ToVector3 converts the 2D vector to a 3D vector by setting the Z component to 0.
// This is useful when working with 3D systems that need to represent 2D vectors.
func (v Vector2) ToVector3() Vector3 {
	return Vector3{
		X: v.X,
		Y: v.Y,
		Z: 0,
	}
}
