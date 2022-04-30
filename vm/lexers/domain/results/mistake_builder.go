package results

import "errors"

type mistakeBuilder struct {
	pIndex *uint
	path   []uint
}

func createMistakeBuilder() MistakeBuilder {
	out := mistakeBuilder{
		pIndex: nil,
		path:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *mistakeBuilder) Create() MistakeBuilder {
	return createMistakeBuilder()
}

// WithIndex adds an index to the builder
func (app *mistakeBuilder) WithIndex(index uint) MistakeBuilder {
	app.pIndex = &index
	return app
}

// WithPath adds a path to the builder
func (app *mistakeBuilder) WithPath(path []uint) MistakeBuilder {
	app.path = path
	return app
}

// Now builds a new Mistake instance
func (app *mistakeBuilder) Now() (Mistake, error) {
	if app.pIndex == nil {
		return nil, errors.New("the indexis mandatory in order to build a Mistake instance")
	}

	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("there must be at least 1 element in the Path in order to build a Mistake instance")
	}

	return createMistake(*app.pIndex, app.path), nil
}
