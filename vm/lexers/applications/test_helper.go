package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

// NewCardinalityWithSpecificForTests creates cardinality with specific for tests
func NewCardinalityWithSpecificForTests(specific uint) cardinality.Cardinality {
	cardinality, err := cardinality.NewBuilder().Create().WithSpecific(specific).Now()
	if err != nil {
		panic(err)
	}

	return cardinality
}

// NewLineWithElementWithCardinalityList creates a new line with ElementWithCardinality list for tests
func NewLineWithElementWithCardinalityList(list []tokens.ElementWithCardinality) tokens.Line {
	line, err := tokens.NewLineBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return line
}

// NewElementWithCardinalityWithReferenceAndCardinalityForTests creates a new elementWithCardinality with token reference and cardinality for  tests
func NewElementWithCardinalityWithReferenceAndCardinalityForTests(tokenIndex uint, refIndex uint, cardinality cardinality.Cardinality) tokens.ElementWithCardinality {
	element, err := tokens.NewElementBuilder().Create().WithReference(refIndex).Now()
	if err != nil {
		panic(err)
	}

	return NewElementWithCardinalityWithElementAndCardinalityForTests(element, cardinality)
}

// NewTokenWithRangeCardinalityWithByteForTests creates a new token with range cardinality with byte for tests
func NewTokenWithRangeCardinalityWithByteForTests(tokenIndex uint, min uint, max uint, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	rnge, err := cardinality.NewRangeBuilder().WithMinimum(min).WithMaximum(max).Now()
	if err != nil {
		panic(err)
	}

	cardinality, err := cardinality.NewBuilder().Create().WithRange(rnge).Now()
	if err != nil {
		panic(err)
	}

	return NewTokenWithSingleElementInSingleLineForTests(tokenIndex, element, cardinality)
}

// NewTokenWithMinimumCardinalityWithByteForTests creates a new token with min cardinality with byte for tests
func NewTokenWithMinimumCardinalityWithByteForTests(tokenIndex uint, min uint, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	rnge, err := cardinality.NewRangeBuilder().WithMinimum(min).Now()
	if err != nil {
		panic(err)
	}

	cardinality, err := cardinality.NewBuilder().Create().WithRange(rnge).Now()
	if err != nil {
		panic(err)
	}

	return NewTokenWithSingleElementInSingleLineForTests(tokenIndex, element, cardinality)
}

// NewTokenWithSpecificCardinalityWithByteForTests creates a new token with specific cardinality with byte for tests
func NewTokenWithSpecificCardinalityWithByteForTests(tokenIndex uint, specific uint, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityWithSpecificForTests(specific)
	return NewTokenWithSingleElementInSingleLineForTests(tokenIndex, element, cardinality)
}

// NewTokenWithSingleElementInSingleLineForTests creates a new token with single element in a singleline for tests
func NewTokenWithSingleElementInSingleLineForTests(tokenIndex uint, element tokens.Element, cardinality cardinality.Cardinality) tokens.Token {
	elementWithCardinality := NewElementWithCardinalityWithElementAndCardinalityForTests(element, cardinality)
	line := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		elementWithCardinality,
	})

	return NewTokenWithLinesForTests(tokenIndex, []tokens.Line{
		line,
	})
}

// NewTokenWithLinesForTests creates a new token with lines for tests
func NewTokenWithLinesForTests(tokenIndex uint, linesList []tokens.Line) tokens.Token {
	lines, err := tokens.NewLinesBuilder().Create().WithList(linesList).Now()
	if err != nil {
		panic(err)
	}

	token, err := tokens.NewBuilder().Create().WithIndex(tokenIndex).WithLines(lines).Now()
	if err != nil {
		panic(err)
	}

	return token
}

// NewElementWithTokenForTests creates a new element with token for tests
func NewElementWithTokenForTests(token tokens.Token) tokens.Element {
	element, err := tokens.NewElementBuilder().Create().WithToken(token).Now()
	if err != nil {
		panic(err)
	}

	return element
}

// NewElementWithCardinalityWithElementAndCardinalityForTests creates a new elementWithCardinality with element and cardinality for tests
func NewElementWithCardinalityWithElementAndCardinalityForTests(element tokens.Element, cardinality cardinality.Cardinality) tokens.ElementWithCardinality {
	elementWithCardinality, err := tokens.NewElementWithCardinalityBuilder().
		Create().
		WithElement(element).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	return elementWithCardinality
}

// NewElementWithCardinalityWithTokenAndCardinalityForTests creates a new elementWithCardinality with token and cardinality for tests
func NewElementWithCardinalityWithTokenAndCardinalityForTests(token tokens.Token, cardinality cardinality.Cardinality) tokens.ElementWithCardinality {
	element := NewElementWithTokenForTests(token)
	return NewElementWithCardinalityWithElementAndCardinalityForTests(element, cardinality)
}
