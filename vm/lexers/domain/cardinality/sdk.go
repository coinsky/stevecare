package cardinality

// Open represents the open byte of the cardinality (ascii of '[')
const Open = 91

// Close represents the close byte of the cardinality (ascii of ']')
const Close = 93

// Separator represents the separator byte of the cardinality (ascii of ',')
const Separator = 44

const prefixErrLabel = "in order to convert data to a Cardinality instance,"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	rangeBuilder := NewRangeBuilder()
	return createAdapter(builder, rangeBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewRangeBuilder creates a new range builder
func NewRangeBuilder() RangeBuilder {
	return createRangeBuilder()
}

// Adapter represents a cardinality adapter
type Adapter interface {
	ToCardinality(data []byte) (Cardinality, []byte, error)
}

// Builder represents the cardinality builder
type Builder interface {
	Create() Builder
	WithRange(rnge Range) Builder
	WithSpecific(specific uint8) Builder
	Now() (Cardinality, error)
}

// Cardinality represents the cardinality
type Cardinality interface {
	Bytes() []byte
	IsRange() bool
	Range() Range
	IsSpecific() bool
	Specific() *uint8
}

// RangeBuilder represents the range builder
type RangeBuilder interface {
	Create() RangeBuilder
	WithMinimum(min uint8) RangeBuilder
	WithMaximum(max uint8) RangeBuilder
	Now() (Range, error)
}

// Range represents the cardinality range
type Range interface {
	Bytes() []byte
	Min() uint8
	HasMax() bool
	Max() *uint8
}
