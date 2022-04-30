package results

type mistake struct {
	index uint
	path  []uint
}

func createMistake(
	index uint,
	path []uint,
) Mistake {
	out := mistake{
		index: index,
		path:  path,
	}

	return &out
}

// Index returns the index
func (obj *mistake) Index() uint {
	return obj.index
}

// Path returns the path
func (obj *mistake) Path() []uint {
	return obj.path
}
