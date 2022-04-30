package tokens

import "github.com/steve-care-software/stevecare/vm/lexers/cardinality"

// Builder represents the token builder
type Builder interface {
	Create() Builder
	WithLList(lines Lines) Builder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Lines() Lines
}

// LinesBuilder represents the lines builder
type LinesBuilder interface {
	Create() LinesBuilder
	WithList(lines []Line) LinesBuilder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	List() []Line
}

// LineBuilder represents the line builder
type LineBuilder interface {
	Create() LineBuilder
	WithList(elements []ElementWithCardinality) LineBuilder
	Now() (Line, error)
}

// Line represents token lines
type Line interface {
	List() []ElementWithCardinality
}

// ElementWithCardinalityBuilder represents the element with cardinality builder
type ElementWithCardinalityBuilder interface {
	Create() ElementWithCardinalityBuilder
	WithElement(element Element) ElementWithCardinalityBuilder
	WithCardinality(cardinality ElementWithCardinality) ElementWithCardinalityBuilder
	Now() (ElementWithCardinality, error)
}

// ElementWithCardinality represents element with cardinality
type ElementWithCardinality interface {
	Element() Element
	Cardinality() cardinality.Cardinality
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithByte(byteValue byte) ElementBuilder
	WithToken(token Token) ElementBuilder
	Now() (Element, error)
}

// Element represents a token element
type Element interface {
	IsByte() bool
	Byte() byte
	IsToken() bool
	Token() Token
}