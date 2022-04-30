package applications

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/results"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

// Application represents a lexer application
type Application interface {
	Execute(token tokens.Token, data []byte) (results.Result, error)
}
