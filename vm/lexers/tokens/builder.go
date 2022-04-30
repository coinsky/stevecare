package tokens

import "errors"

type builder struct {
	lines Lines
}

func createBuilder() Builder {
	out := builder{
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a Lines to the builder
func (app *builder) WithList(lines Lines) Builder {
	app.lines = lines
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Token instance")
	}

	return createToken(app.lines), nil
}
