package cardinality

type cardinality struct {
	rnge      Range
	pSpecific *uint8
}

func createCardinalityWithRange(
	rnge Range,
) Cardinality {
	return createCardinalityInternally(rnge, nil)
}

func createCardinalityWithSpecific(
	pSpecific *uint8,
) Cardinality {
	return createCardinalityInternally(nil, pSpecific)
}

func createCardinalityInternally(
	rnge Range,
	pSpecific *uint8,
) Cardinality {
	out := cardinality{
		rnge:      rnge,
		pSpecific: pSpecific,
	}

	return &out
}

// Bytes returns the []byte representation of the cardinality
func (obj *cardinality) Bytes() []byte {
	if obj.IsRange() {
		return obj.Range().Bytes()
	}

	return []byte{
		Open,
		byte(*obj.pSpecific),
		Close,
	}
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
func (obj *cardinality) Specific() *uint8 {
	return obj.pSpecific
}
