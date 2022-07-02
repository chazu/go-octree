package octree

import (
	"testing"

	"github.com/google/uuid"
)

func TestOctreeMinMaxXYZ(t *testing.T) {
	point := &Entity{
		uuid.New(),
		50,
		0,
		0,
		0,
	}

	if int(point.MinX()) != -50 {
		t.Errorf("MinX is wrong - got %f, wanted -50", point.MinX())
	}

	if int(point.MinY()) != -50 {
		t.Errorf("MinY is wrong - got %f, wanted -50", point.MinY())
	}

	if int(point.MinZ()) != -50 {
		t.Errorf("MinZ is wrong - got %f, wanted -50", point.MinZ())
	}

	if int(point.MaxX()) != 50 {
		t.Errorf("MaxX is wrong - got %f, wanted 50", point.MaxX())
	}

	if int(point.MaxY()) != 50 {
		t.Errorf("MaxY is wrong - got %f, wanted 50", point.MaxY())
	}

	if int(point.MaxZ()) != 50 {
		t.Errorf("MaxZ is wrong - got %f, wanted 50", point.MaxZ())
	}

}

func TestOctreeEqualsPointIsTrue(t *testing.T) {
	point_one := &Entity{
		uuid.New(),
		50,
		100,
		100,
		100,
	}

	point_two := &Entity{
		uuid.New(),
		100,
		100,
		100,
		100,
	}
	result := point_one.EqualsPoint(*point_two)
	if result != true {
		t.Errorf("Equals point failed - expected true, got %v", result)
	}
}

func TestOctreeEqualsPointIsFalse(t *testing.T) {
	point_one := &Entity{
		uuid.New(),
		50,
		100,
		100,
		100,
	}

	point_two := &Entity{
		uuid.New(),
		100,
		-100,
		-100,
		-100,
	}
	result := point_one.EqualsPoint(*point_two)
	if result != false {
		t.Errorf("Equals point failed - expected false, got %v", result)
	}
}

// NOTE This test needs to be more thorough
func TestOctreeDistanceToSimpleCase(t *testing.T) {
	point_one := &Entity{
		uuid.New(),
		5,
		0,
		0,
		0,
	}

	point_two := &Entity{
		uuid.New(),
		5,
		100,
		0,
		0,
	}
	result := point_one.DistanceTo(*point_two)
	if int(result) != 90 {
		t.Errorf("Equals point failed - expected 90, got %v", int(result))
	}
}

func TestOctreeDistanceToOverlappingCase(t *testing.T) {
	point_one := &Entity{
		uuid.New(),
		10,
		10,
		0,
		0,
	}

	point_two := &Entity{
		uuid.New(),
		5,
		0,
		0,
		0,
	}
	result := point_one.DistanceTo(*point_two)
	expected := 0
	if int(result) != expected {
		t.Errorf("DistanceTo Failed - expected %v, got %v", expected, int(result))
	}
}

// NOTE This test needs to be more thorough
func TestOctreeDistanceToPoint(t *testing.T) {
	point_one := &Entity{
		uuid.New(),
		5,
		0,
		0,
		0,
	}

	point_two := &Entity{
		uuid.New(),
		5,
		100,
		0,
		0,
	}
	result := point_one.DistanceToPoint(*point_two)
	if int(result) != 100 {
		t.Errorf("DistanceToPoint failed - expected 100, got %v", int(result))
	}
}

func TestOctreeIntersectsFalse(t *testing.T) {
	point_one := &Entity{
		uuid.New(),
		5,
		0,
		0,
		0,
	}

	point_two := &Entity{
		uuid.New(),
		5,
		100,
		0,
		0,
	}

	result := point_one.Intersects(*point_two)
	expected := false

	if result != expected {
		t.Errorf("Fail - expected %v, got %v", expected, result)
	}
}
