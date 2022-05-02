package cardinality

import "errors"

type builder struct {
	rnge      Range
	pSpecific *uint8
}

func createBuilder() Builder {
	out := builder{
		rnge:      nil,
		pSpecific: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRange adds a range to the builder
func (app *builder) WithRange(rnge Range) Builder {
	app.rnge = rnge
	return app
}

// WithSpecific adds a specific value to the builder
func (app *builder) WithSpecific(specific uint8) Builder {
	app.pSpecific = &specific
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.rnge != nil {
		return createCardinalityWithRange(app.rnge), nil
	}

	if app.pSpecific != nil {
		return createCardinalityWithSpecific(app.pSpecific), nil
	}

	return nil, errors.New("the Cardinality is invalid")
}
