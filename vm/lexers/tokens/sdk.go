package tokens

import "github.com/steve-care-software/stevecare/vm/lexers/cardinality"

// Token represents a token
type Token interface {
	Lines() Lines
}

// Lines represents lines
type Lines interface {
	List() []Line
}

// Line represents token lines
type Line interface {
	List() []ElementWithCardinality
}

// ElementWithCardinality represents element with cardinality
type ElementWithCardinality interface {
	Element() Element
	Cardinality() cardinality.Cardinality
}

// Element represents a token element
type Element interface {
	IsByte() bool
	Byte() byte
	IsToken() bool
	Token() Token
}
