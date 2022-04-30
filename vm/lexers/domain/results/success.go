package results

type success struct {
	pIndex *uint
}

func createSuccess() Success {
	return createSuccessInternally(nil)
}

func createSuccessWithIndex(
	pIndex *uint,
) Success {
	return createSuccessInternally(pIndex)
}

func createSuccessInternally(
	pIndex *uint,
) Success {
	out := success{
		pIndex: pIndex,
	}

	return &out
}

// HasIndex returns true if there is an index, false otherwise
func (obj *success) HasIndex() bool {
	return obj.pIndex != nil
}

// Index returns the index, if any
func (obj *success) Index() *uint {
	return obj.pIndex
}
