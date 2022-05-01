package tokens

import "errors"

type tokenBuilder struct {
	pIndex *uint
	lines  Lines
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		pIndex: nil,
		lines:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithIndex adds an index to the builder
func (app *tokenBuilder) WithIndex(index uint) TokenBuilder {
	app.pIndex = &index
	return app
}

// WithLines adds a Lines to the builder
func (app *tokenBuilder) WithLines(lines Lines) TokenBuilder {
	app.lines = lines
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Token instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Token instance")
	}

	return createToken(*app.pIndex, app.lines), nil
}
