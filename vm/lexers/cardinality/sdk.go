package cardinality

// Builder represents the cardinality builder
type Builder interface {
	Create() Builder
	WithRange(rnge Range) Builder
	WithSpecific(specific uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents the cardinality
type Cardinality interface {
	IsRange() bool
	Range() Range
	IsSpecific() bool
	Specific() uint
}

// RangeBuilder represents the range builder
type RangeBuilder interface {
	Create() RangeBuilder
	WithMinimum(min uint) RangeBuilder
	WithMaximum(max uint) RangeBuilder
	Now() (Range, error)
}

// Range represents the cardinality range
type Range interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
