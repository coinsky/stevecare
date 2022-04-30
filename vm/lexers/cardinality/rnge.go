package cardinality

type rnge struct {
	min  uint
	pMax *uint
}

func createRange(
	min uint,
) Range {
	return createRangeInternally(min, nil)
}

func createRangeWithMaximum(
	min uint,
	pMax *uint,
) Range {
	return createRangeInternally(min, pMax)
}

func createRangeInternally(
	min uint,
	pMax *uint,
) Range {
	out := rnge{
		min:  min,
		pMax: pMax,
	}

	return &out
}

// Min returns the minimum
func (obj *rnge) Min() uint {
	return obj.min
}

// HasMax returns true if there is a max, false otherwise
func (obj *rnge) HasMax() bool {
	return obj.pMax != nil
}

// Max returns the maximum, if any
func (obj *rnge) Max() *uint {
	return obj.pMax
}
