package octree

import (
	"github.com/google/uuid"
)

type OctreeEntity struct {
	Id     uuid.UUID
	Radius int
	X      float64
	Y      float64
	Z      float64
}
