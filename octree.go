package octree

func CreateOctree(root bool, origin Vec3, breakpoint int, halfDimension Vec3) *Octree {
	tree := Octree()
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

func (o Octree) InitializeChildren() {
}
