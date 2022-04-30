package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

// NewApplication creates a new application
func NewApplication() Application {
	resultBuilder := results.NewBuilder()
	successBuilder := results.NewSuccessBuilder()
	mistakeBuilder := results.NewMistakeBuilder()
	return createApplication(resultBuilder, successBuilder, mistakeBuilder)
}

// Application represents a lexer application
type Application interface {
	Execute(token tokens.Token, data []byte) (results.Result, error)
}
