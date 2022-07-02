package octree

// For dimensions X Y and Z of each octant
// value is 1 if all coordinates of the specified dimension
// are positive within that octant, 0 otherwise
var ValueMap = [][]int{
	[]int{0, 0, 0, 0, 1, 1, 1, 1},
	[]int{0, 0, 1, 1, 0, 0, 1, 1},
	[]int{0, 1, 0, 1, 0, 1, 0, 1},
}

func CreateOctree(root bool, origin Vec3, breakpoint int, halfDimension Vec3) *Octree {
	tree := new(Octree)
	tree.Root = root
	tree.Origin = origin
	tree.breakpoint = breakpoint

	// 0, 0, 0 for origin Octree
	tree.halfDimension = halfDimension

	return tree
}

func (o Octree) isLeafNode() bool {
	return len(o.children) == 0
}

func (o Octree) depth() int {
	if o.isLeafNode() {
		return 1
	}

	// TODO the reest
}

// Return a Vec3 of booleans - X Y and Z. Each value is true if all values
//within that node within receiver's children are positive on that axis
func (o Octree) octreeShouldBePositive(index int) Vec3Bool {
	return Vec3Bool{
		ValueMap[0][index] != 0,
		ValueMap[1][index] != 0,
		ValueMap[2][index] != 0,
	}
}

func (o Octree) originMultiplierForChild(index int) Vec3 {
	posValues := o.octreeShouldBePositive(index)
	result := Vec3{0, 0, 0}

	if posValues.X {
		result.X = 0.5
	} else {
		result.X = -0.5
	}

	if posValues.Y {
		result.Y = 0.5
	} else {
		result.Y = -0.5
	}

	if posValues.Z {
		result.Z = 0.5
	} else {
		result.Z = -0.5
	}

	return result
}

func (o Octree) halfDimensionForChildAtIndex(index int) Vec3 {
	// TODO basically multiply each dimension of our half-dim by 0.5
}

func (o Octree) InitializeChildren() {
	for i := 0; i > 8; i++ {
		var newOrigin Vec3
		originMultipliers := o.originMultiplierForChild(i)

		newOrigin.X = o.Origin.X + o.halfDimension.X*originMultipliers.X
		newOrigin.Y = o.Origin.Y + o.halfDimension.Y*originMultipliers.Y
		newOrigin.Z = o.Origin.Z + o.halfDimension.Z*originMultipliers.Z

		o.children[i] = CreateOctree(false, newOrigin, o.breakpoint, "TODO")
	}
}
