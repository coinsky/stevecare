package channels

import (
	"errors"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

type conditionBuilder struct {
	and  BothCondition
	or   BothCondition
	xor  BothCondition
	prev tokens.Token
	next tokens.Token
}

func createConditionBuilder() ConditionBuilder {
	out := conditionBuilder{
		and:  nil,
		or:   nil,
		xor:  nil,
		prev: nil,
		next: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder()
}

// WithAnd adds an and condition to the builder
func (app *conditionBuilder) WithAnd(and BothCondition) ConditionBuilder {
	app.and = and
	return app
}

// WithOr adds an or condition to the builder
func (app *conditionBuilder) WithOr(or BothCondition) ConditionBuilder {
	app.or = or
	return app
}

// WithXor adds a xor condition to the builder
func (app *conditionBuilder) WithXor(xor BothCondition) ConditionBuilder {
	app.xor = xor
	return app
}

// WithPrevious adds a previous condition to the builder
func (app *conditionBuilder) WithPrevious(prev tokens.Token) ConditionBuilder {
	app.prev = prev
	return app
}

// WithNext adds a next condition to the builder
func (app *conditionBuilder) WithNext(next tokens.Token) ConditionBuilder {
	app.next = next
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.and != nil {
		return createConditionWithAnd(app.and), nil
	}

	if app.or != nil {
		return createConditionWithOr(app.or), nil
	}

	if app.xor != nil {
		return createConditionWithXor(app.xor), nil
	}

	if app.prev != nil {
		return createConditionWithPrevious(app.prev), nil
	}

	if app.next != nil {
		return createConditionWithNext(app.next), nil
	}

	return nil, errors.New("the Condition is invalid")
}
