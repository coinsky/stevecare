package grammars

import (
	"github.com/steve-care-software/stevecare/vm/lexers/domain/channels"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithRoot(root tokens.Token) Builder
	WithChannels(channels channels.Channels) Builder
	Now() (Grammar, error)
}

// Grammar represents a lexer grammar
type Grammar interface {
	Root() tokens.Token
	HasChannels() bool
	Channels() channels.Channels
}
