package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

// NewApplication creates a new application
func NewApplication() Application {
	resultBuilder := results.NewBuilder()
	return createApplication(resultBuilder)
}

// Application represents a lexer application
type Application interface {
	Execute(token tokens.Token, data []byte, canHavePrefix bool) (results.Result, error)
}
