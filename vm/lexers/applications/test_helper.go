package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

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
