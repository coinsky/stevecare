package results

type result struct {
	index     uint
	cursor    uint
	path      []uint
	isSuccess bool
}

func createResult(
	index uint,
	cursor uint,
	path []uint,
	isSuccess bool,
) Result {
	out := result{
		index:     index,
		cursor:    cursor,
		path:      path,
		isSuccess: isSuccess,
	}

	return &out
}

// Index returns the index
func (obj *result) Index() uint {
	return obj.index
}

// Cursor returns the cursor
func (obj *result) Cursor() uint {
	return obj.cursor
}

// Path returns the path
func (obj *result) Path() []uint {
	return obj.path
}

// IsSuccess returns true if successful, falseotherwise
func (obj *result) IsSuccess() bool {
	return obj.isSuccess
}
