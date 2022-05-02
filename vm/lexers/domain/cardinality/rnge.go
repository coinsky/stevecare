package cardinality

type rnge struct {
	min  uint8
	pMax *uint8
}

func createRange(
	min uint8,
) Range {
	return createRangeInternally(min, nil)
}

func createRangeWithMaximum(
	min uint8,
	pMax *uint8,
) Range {
	return createRangeInternally(min, pMax)
}

func createRangeInternally(
	min uint8,
	pMax *uint8,
) Range {
	out := rnge{
		min:  min,
		pMax: pMax,
	}

	return &out
}

// Bytes returns the []byte representation of the range
func (obj *rnge) Bytes() []byte {
	if obj.HasMax() {
		return []byte{
			Open,
			byte(obj.Min()),
			Separator,
			byte(*obj.pMax),
			Close,
		}
	}

	return []byte{
		Open,
		byte(obj.Min()),
		Separator,
		Close,
	}
}

// Min returns the minimum
func (obj *rnge) Min() uint8 {
	return obj.min
}

// HasMax returns true if there is a max, false otherwise
func (obj *rnge) HasMax() bool {
	return obj.pMax != nil
}

// Max returns the maximum, if any
func (obj *rnge) Max() *uint8 {
	return obj.pMax
}
