package vectors

import (
	"math"
)

// IVector3 is the interface for a 3D vector.
// It defines all the operations that can be performed on a 3D vector.
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
	MagnitudeSquared() float64
	Distance(vec Vector3) float64
	DistanceSquared(vec Vector3) float64
	ClampMagnitude(maxValue float64)
	Clear()
	ToVector2() Vector2
}

// Vector3 represents a 3D vector with X, Y, and Z coordinates.
// It implements the IVector3 interface and provides methods for 3D vector operations.
type Vector3 struct {
	X float64 // X coordinate of the vector.
	Y float64 // Y coordinate of the vector.
	Z float64 // Z coordinate of the vector.
}

// Add adds the values of another vector to this one.
func (v *Vector3) Add(vec Vector3) {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
}

// Sub subtracts the values of another vector from this one.
func (v *Vector3) Sub(vec Vector3) {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
}

// Mul multiplies this vector by another vector component-wise.
func (v *Vector3) Mul(vec Vector3) {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
}

// Div divides this vector by another vector component-wise.
// Note: Division by zero will result in NaN or Inf values.
func (v *Vector3) Div(vec Vector3) {
	v.X /= vec.X
	v.Y /= vec.Y
	v.Z /= vec.Z
}

// Bounce inverts the direction of the vector by negating all axes.
// This is equivalent to multiplying the vector by -1.
func (v *Vector3) Bounce() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Normalize scales the vector to have a magnitude of 1 while preserving its direction.
// If the vector is already zero, it remains unchanged.
// A normalized vector is also called a unit vector.
func (v *Vector3) Normalize() {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
		v.Z /= magnitude
	}
}

// AngleRadians returns the angle of the vector in radians relative to the positive X-axis.
// The angle is measured counterclockwise from the positive X-axis.
// Returns a value in the range [-π, π].
// Note: This ignores the Z component and projects the vector onto the XY plane.
func (v Vector3) AngleRadians() float64 {
	return math.Atan2(v.Y, v.X)
}

// AngleDegrees returns the angle of the vector in degrees relative to the positive X-axis.
// The angle is measured counterclockwise from the positive X-axis.
// Returns a value in the range [0, 360).
// Note: This ignores the Z component and projects the vector onto the XY plane.
func (v Vector3) AngleDegrees() float64 {
	angle := math.Atan2(v.Y, v.X) * 180 / math.Pi

	return math.Mod(angle+360, 360)
}

// IsZero returns true if all axes are zero.
// This indicates the vector has no magnitude and no direction.
func (v Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// Magnitude returns the length (magnitude) of the vector.
func (v Vector3) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// MagnitudeSquared returns the squared magnitude of the vector.
// This is faster for magnitude comparisons, since it avoids the square root.
func (v Vector3) MagnitudeSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Distance returns the distance between this vector and another vector.
// This is equivalent to the magnitude of the difference between the vectors.
func (v Vector3) Distance(vec Vector3) float64 {
	dx := v.X - vec.X
	dy := v.Y - vec.Y
	dz := v.Z - vec.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// DistanceSquared returns the squared distance between this vector and another vector.
// This is faster for distance comparisons, since it avoids the square root.
func (v Vector3) DistanceSquared(vec Vector3) float64 {
	dx := v.X - vec.X
	dy := v.Y - vec.Y
	dz := v.Z - vec.Z
	return dx*dx + dy*dy + dz*dz
}

// ClampMagnitude limits the magnitude of the vector to a maximum value.
// If the current magnitude exceeds maxValue, the vector is scaled down proportionally.
// If the vector is zero or already within the limit, no change is made.
// This preserves the direction while limiting the length.
func (v *Vector3) ClampMagnitude(maxValue float64) {
	magnitude := v.Magnitude()

	if magnitude == 0 || magnitude <= maxValue {
		return
	}

	scale := maxValue / magnitude
	v.X *= scale
	v.Y *= scale
	v.Z *= scale
}

// Clear sets the vector to zero on all axes.
// This is equivalent to setting all axes to 0.
func (v *Vector3) Clear() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}

// ToVector2 converts the 3D vector to a 2D vector by discarding the Z component.
// This is useful when working with 2D systems that need to represent 3D vectors.
func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}
