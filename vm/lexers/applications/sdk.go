package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/grammars"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
)

// NewApplication creates a new application
func NewApplication() Application {
	resultBuilder := results.NewBuilder()
	return createApplication(resultBuilder)
}

// Application represents a lexer application
type Application interface {
	Execute(grammar grammars.Grammar, data []byte, canHavePrefix bool) (results.Result, error)
}
