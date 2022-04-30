package tokens

import "errors"

type builder struct {
	pIndex *uint
	lines  Lines
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		lines:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithList adds a Lines to the builder
func (app *builder) WithList(lines Lines) Builder {
	app.lines = lines
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.pIndex != nil {
		return nil, errors.New("the index is mandatory in order to build a Token instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Token instance")
	}

	return createToken(*app.pIndex, app.lines), nil
}
