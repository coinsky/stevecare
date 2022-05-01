package channels

import "github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"

type condition struct {
	and  BothCondition
	or   BothCondition
	xor  BothCondition
	prev tokens.Token
	next tokens.Token
}

func createConditionWithAnd(
	and BothCondition,
) Condition {
	return createConditionInternally(and, nil, nil, nil, nil)
}

func createConditionWithOr(
	or BothCondition,
) Condition {
	return createConditionInternally(nil, or, nil, nil, nil)
}

func createConditionWithXor(
	xor BothCondition,
) Condition {
	return createConditionInternally(nil, nil, xor, nil, nil)
}

func createConditionWithPrevious(
	prev tokens.Token,
) Condition {
	return createConditionInternally(nil, nil, nil, prev, nil)
}

func createConditionWithNext(
	next tokens.Token,
) Condition {
	return createConditionInternally(nil, nil, nil, nil, next)
}

func createConditionInternally(
	and BothCondition,
	or BothCondition,
	xor BothCondition,
	prev tokens.Token,
	next tokens.Token,
) Condition {
	out := condition{
		and:  and,
		or:   or,
		xor:  xor,
		prev: prev,
		next: next,
	}

	return &out
}

// IsAnd returns true if there is an and condition, false otherwise
func (obj *condition) IsAnd() bool {
	return obj.and != nil
}

// And returns the and condition, if any
func (obj *condition) And() BothCondition {
	return obj.and
}

// IsOr returns true if there is an or condition, false otherwise
func (obj *condition) IsOr() bool {
	return obj.or != nil
}

// Or returns the or condition, if any
func (obj *condition) Or() BothCondition {
	return obj.and
}

// IsXor returns true if there is a xor condition, false otherwise
func (obj *condition) IsXor() bool {
	return obj.xor != nil
}

// Xor returns the xor condition, if any
func (obj *condition) Xor() BothCondition {
	return obj.xor
}

// IsPrevious returns true if there is a previous condition, false otherwise
func (obj *condition) IsPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous condition, if any
func (obj *condition) Previous() tokens.Token {
	return obj.prev
}

// IsNext returns true if there is a next condition, false otherwise
func (obj *condition) IsNext() bool {
	return obj.next != nil
}

// Next returns the next condition, if any
func (obj *condition) Next() tokens.Token {
	return obj.next
}
