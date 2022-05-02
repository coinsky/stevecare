package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/grammars"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

// NewApplication creates a new application
func NewApplication() Application {
	grammarBuilder := grammars.NewBuilder()
	tokenBuilder := tokens.NewTokenBuilder()
	linesBuilder := tokens.NewLinesBuilder()
	lineBuilder := tokens.NewLineBuilder()
	elementWithCardinalityBuilder := tokens.NewElementWithCardinalityBuilder()
	elementBuilder := tokens.NewElementBuilder()
	cardinalityBuilder := cardinality.NewBuilder()
	cardinalityRangeBuilder := cardinality.NewRangeBuilder()
	resultBuilder := results.NewBuilder()
	rootPrefix := []byte("%")[0]
	rootSuffix := []byte(";")[0]
	tokenNamePrefix := []byte(".")[0]
	bytePrefix := []byte("$")[0]
	indexTokenNameSeparator := []byte("@")[0]
	linesPrefix := []byte(":")[0]
	linesSuffix := []byte(";")[0]
	lineDelimiter := []byte("|")[0]
	cardinalityNonZeroMultiple := []byte("+")[0]
	cardinalityZeroMultiple := []byte("*")[0]
	cardinalityOptional := []byte("?")[0]
	cardinalityRangePrefix := []byte("[")[0]
	cardinalityRangeSuffix := []byte("]")[0]
	cardinalityRangeSeparator := []byte(",")[0]
	commentPrefix := []byte(";")[0]
	commentSuffix := []byte(";")[0]
	numbersCharacters := []byte("0123456789")
	tokenNameCharacters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	channelCharacters := []byte{
		[]byte("\t")[0],
		[]byte("\n")[0],
		[]byte("\r")[0],
		[]byte(" ")[0],
	}

	return createApplication(
		grammarBuilder,
		tokenBuilder,
		linesBuilder,
		lineBuilder,
		elementWithCardinalityBuilder,
		elementBuilder,
		cardinalityBuilder,
		cardinalityRangeBuilder,
		resultBuilder,
		rootPrefix,
		rootSuffix,
		tokenNamePrefix,
		bytePrefix,
		indexTokenNameSeparator,
		linesPrefix,
		linesSuffix,
		lineDelimiter,
		cardinalityNonZeroMultiple,
		cardinalityZeroMultiple,
		cardinalityOptional,
		cardinalityRangePrefix,
		cardinalityRangeSuffix,
		cardinalityRangeSeparator,
		commentPrefix,
		commentSuffix,
		numbersCharacters,
		tokenNameCharacters,
		channelCharacters,
	)
}

// Application represents a lexer application
type Application interface {
	Compile(script string) (grammars.Grammar, error)
	Execute(grammar grammars.Grammar, data []byte, canHavePrefix bool) (results.Result, error)
}
