package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

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

	return NewTokenWithSigleLineForTests(tokenIndex, element, cardinality)
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

	return NewTokenWithSigleLineForTests(tokenIndex, element, cardinality)
}

// NewTokenWithSpecificCardinalityWithByteForTests creates a new token with specific cardinality with byte for tests
func NewTokenWithSpecificCardinalityWithByteForTests(tokenIndex uint, specific uint, byteVal byte) tokens.Token {
	element, err := tokens.NewElementBuilder().Create().WithByte(byteVal).Now()
	if err != nil {
		panic(err)
	}

	cardinality, err := cardinality.NewBuilder().Create().WithSpecific(specific).Now()
	if err != nil {
		panic(err)
	}

	return NewTokenWithSigleLineForTests(tokenIndex, element, cardinality)
}

// NewTokenWithSigleLineForTests creates a new token with single line for tests
func NewTokenWithSigleLineForTests(tokenIndex uint, element tokens.Element, cardinality cardinality.Cardinality) tokens.Token {
	elementWithCardinality, err := tokens.NewElementWithCardinalityBuilder().
		Create().
		WithElement(element).
		WithCardinality(cardinality).
		Now()

	if err != nil {
		panic(err)
	}

	line, err := tokens.NewLineBuilder().Create().WithList([]tokens.ElementWithCardinality{
		elementWithCardinality,
	}).Now()

	if err != nil {
		panic(err)
	}

	lines, err := tokens.NewLinesBuilder().Create().WithList([]tokens.Line{
		line,
	}).Now()

	if err != nil {
		panic(err)
	}

	token, err := tokens.NewBuilder().Create().WithIndex(tokenIndex).WithLines(lines).Now()
	if err != nil {
		panic(err)
	}

	return token
}
