package cardinality

type cardinality struct {
	rnge      Range
	pSpecific *uint
}

func createCardinalityWithRange(
	rnge Range,
) Cardinality {
	return createCardinalityInternally(rnge, nil)
}

func createCardinalityWithSpecific(
	pSpecific *uint,
) Cardinality {
	return createCardinalityInternally(nil, pSpecific)
}

func createCardinalityInternally(
	rnge Range,
	pSpecific *uint,
) Cardinality {
	out := cardinality{
		rnge:      rnge,
		pSpecific: pSpecific,
	}

	return &out
}

// IsRange returns true if there is a range, false otherwise
func (obj *cardinality) IsRange() bool {
	return obj.rnge != nil
}

// Range returns the range, if any
func (obj *cardinality) Range() Range {
	return obj.rnge
}

// IsSpecific returns true if there is a specific value, false otherwise
func (obj *cardinality) IsSpecific() bool {
	return obj.pSpecific != nil
}

// Specific returns the specific value, if any
func (obj *cardinality) Specific() *uint {
	return obj.pSpecific
}
