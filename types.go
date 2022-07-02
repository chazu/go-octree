package octree

import (
	uuid "github.com/google/uuid"
)

type Entity struct {
	Id     uuid.UUID
	Radius int
	X      float64
	Y      float64
	Z      float64
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Octree struct {
	Root          bool
	Origin        Vec3
	children      [8]Octree
	halfDimension Vec3
	points        []Entity
	breakpoint    int
}
