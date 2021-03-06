package tokens

import "errors"

type builder struct {
	list []Token
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list of tokens to the builder
func (app *builder) WithList(list []Token) Builder {
	return createBuilder()
}

// Now builds a new Tokens instance
func (app *builder) Now() (Tokens, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Tokens instance")
	}

	return createTokens(app.list), nil
}
