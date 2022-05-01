package tokens

import "github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"

// NewBuilder creates a new tokens builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder instance
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewLinesBuilder creates a new lines builder
func NewLinesBuilder() LinesBuilder {
	return createLinesBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewElementWithCardinalityBuilder creates a new element with cardinality builder
func NewElementWithCardinalityBuilder() ElementWithCardinalityBuilder {
	return createElementWithCardinalityBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents a tokens builder
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents a list of token
type Tokens interface {
	List() []Token
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithIndex(index uint) TokenBuilder
	WithLines(lines Lines) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Index() uint
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
	WithCardinality(cardinality cardinality.Cardinality) ElementWithCardinalityBuilder
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
	WithReference(reference uint) ElementBuilder
	Now() (Element, error)
}

// Element represents a token element
type Element interface {
	IsByte() bool
	Byte() *byte
	IsToken() bool
	Token() Token
	IsReference() bool
	Reference() *uint
}
