package results

import "errors"

type builder struct {
	pIndex    *uint
	path      []uint
	isSuccess bool
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		path:   nil,
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

// WithPath adds a path to the builder
func (app *builder) WithPath(path []uint) Builder {
	app.path = path
	return app
}

// IsSuccess flags the builder as a success
func (app *builder) IsSuccess() Builder {
	app.isSuccess = true
	return app
}

// Now builds a new Result instance
func (app *builder) Now() (Result, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Result instance")
	}

	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("there must be at least 1 element in the Path in order to build a Result instance")
	}

	return createResult(*app.pIndex, app.path, app.isSuccess), nil
}
