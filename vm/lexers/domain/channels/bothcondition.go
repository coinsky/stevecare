package channels

import "github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"

type bothCondition struct {
	prev tokens.Token
	next tokens.Token
}

func createBothCondition(
	prev tokens.Token,
	next tokens.Token,
) BothCondition {
	out := bothCondition{
		prev: prev,
		next: next,
	}

	return &out
}

// Previous returns the previous tokens
func (obj *bothCondition) Previous() tokens.Token {
	return obj.prev
}

// Next returns the next tokens
func (obj *bothCondition) Next() tokens.Token {
	return obj.next
}
