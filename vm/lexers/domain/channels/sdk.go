package channels

import "github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"

// Builder represents a channels builder
type Builder interface {
	Create() Builder
	WithList(list []Channel) Builder
	Now() (Channels, error)
}

// Channels represents channels
type Channels interface {
	List() []Channel
}

// ChannelBuilder represents a channel builder
type ChannelBuilder interface {
	Create() ChannelBuilder
	WithToken(token tokens.Token) ChannelBuilder
	WithCondition(condition Condition) ChannelBuilder
	Now() (Channel, error)
}

// Channel represents a channel
type Channel interface {
	Token() tokens.Token
	HasCondition() bool
	Condition() Condition
}

// ConditionBuilder represents the condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithAnd(and BothCondition) ConditionBuilder
	WithOr(or BothCondition) ConditionBuilder
	WithPrevious(prev tokens.Tokens) ConditionBuilder
	WithNext(next tokens.Tokens) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a channel condition
type Condition interface {
	IsAnd() bool
	And() BothCondition
	IsOr() bool
	Or() BothCondition
	IsPrevious() bool
	Previous() tokens.Tokens
	IsNext() bool
	Next() tokens.Tokens
}

// BothConditionBuilder represents a both condition builder
type BothConditionBuilder interface {
	Create() BothConditionBuilder
	WithPrevious(prev tokens.Tokens) BothConditionBuilder
	WithNext(next tokens.Tokens) BothConditionBuilder
	Now() (BothCondition, error)
}

// BothCondition represents an and or or condition
type BothCondition interface {
	Previous() tokens.Tokens
	Next() tokens.Tokens
}
