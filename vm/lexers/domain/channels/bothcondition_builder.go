package channels

import (
	"errors"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

type bothConditionBuilder struct {
	prev tokens.Token
	next tokens.Token
}

func createBothConditionBuilder() BothConditionBuilder {
	out := bothConditionBuilder{
		prev: nil,
		next: nil,
	}

	return &out
}

// Create initializes the builder
func (app *bothConditionBuilder) Create() BothConditionBuilder {
	return createBothConditionBuilder()
}

// WithPrevious adds a previous token
func (app *bothConditionBuilder) WithPrevious(prev tokens.Token) BothConditionBuilder {
	app.prev = prev
	return app
}

// WithNext adds a next token to the builder
func (app *bothConditionBuilder) WithNext(next tokens.Token) BothConditionBuilder {
	app.next = next
	return app
}

// Now builds a new BothCondition instance
func (app *bothConditionBuilder) Now() (BothCondition, error) {
	if app.next == nil {
		return nil, errors.New("the next token is mandatory in order to build a BothCondition instance")
	}

	if app.prev == nil {
		return nil, errors.New("the previous token is mandatory in order to build a BothCondition instance")
	}

	return createBothCondition(app.prev, app.next), nil
}
