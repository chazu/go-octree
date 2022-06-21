package octree

import "math"

func (e OctreeEntity) MinX() float64 {
	return e.X - float64(e.Radius)
}

func (e OctreeEntity) MinY() float64 {
	return e.Y - float64(e.Radius)
}

func (e OctreeEntity) MinZ() float64 {
	return e.Z - float64(e.Radius)
}

func (e OctreeEntity) MaxX() float64 {
	return e.X + float64(e.Radius)
}

func (e OctreeEntity) MaxY() float64 {
	return e.Y + float64(e.Radius)
}

func (e OctreeEntity) MaxZ() float64 {
	return e.Z + float64(e.Radius)
}

func (e OctreeEntity) EqualsPoint(o OctreeEntity) bool {
	return e.X == o.X &&
		e.Y == o.Y &&
		e.Z == o.Z
}

func (e OctreeEntity) Location() [3]float64 {
	return [3]float64{e.X, e.Y, e.Z}
}

func (e OctreeEntity) DistanceToPoint(o OctreeEntity) float64 {
	sum := math.Pow(o.X-e.X, 2) + math.Pow(o.Y-e.Y, 2) + math.Pow(o.Z-e.Z, 2)
	return math.Sqrt(sum)
}

func (e OctreeEntity) DistanceTo(o OctreeEntity) float64 {
	return e.DistanceTo(o) - (float64(e.Radius) + float64(o.Radius))
}

func (e OctreeEntity) Intersects(o OctreeEntity) bool {
	return e.DistanceTo(o) <= 0
}
