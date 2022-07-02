package octree

import "math"

func (e Entity) MinX() float64 {
	return e.X - float64(e.Radius)
}

func (e Entity) MinY() float64 {
	return e.Y - float64(e.Radius)
}

func (e Entity) MinZ() float64 {
	return e.Z - float64(e.Radius)
}

func (e Entity) MaxX() float64 {
	return e.X + float64(e.Radius)
}

func (e Entity) MaxY() float64 {
	return e.Y + float64(e.Radius)
}

func (e Entity) MaxZ() float64 {
	return e.Z + float64(e.Radius)
}

func (e Entity) EqualsPoint(o Entity) bool {
	return e.X == o.X &&
		e.Y == o.Y &&
		e.Z == o.Z
}

func (e Entity) Location() [3]float64 {
	return [3]float64{e.X, e.Y, e.Z}
}

func (e Entity) DistanceToPoint(o Entity) float64 {
	sum := math.Pow(o.X-e.X, 2) + math.Pow(o.Y-e.Y, 2) + math.Pow(o.Z-e.Z, 2)
	return math.Sqrt(sum)
}

func (e Entity) DistanceTo(o Entity) float64 {
	return math.Max(e.DistanceToPoint(o)-(float64(e.Radius)+float64(o.Radius)), 0)
}

func (e Entity) Intersects(o Entity) bool {
	return e.DistanceTo(o) <= 0
}
