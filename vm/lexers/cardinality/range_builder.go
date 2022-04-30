package cardinality

import "errors"

type rangeBuilder struct {
	pMin *uint
	pMax *uint
}

func createRangeBuilder() RangeBuilder {
	out := rangeBuilder{
		pMin: nil,
		pMax: nil,
	}

	return &out
}

// Create initializes the builder
func (app *rangeBuilder) Create() RangeBuilder {
	return createRangeBuilder()
}

// WithMinimum adds a minimum to the builder
func (app *rangeBuilder) WithMinimum(min uint) RangeBuilder {
	app.pMin = &min
	return app
}

// WithMaximum adds a maximum to the builder
func (app *rangeBuilder) WithMaximum(max uint) RangeBuilder {
	app.pMax = &max
	return app
}

// Now builds a new Range instance
func (app *rangeBuilder) Now() (Range, error) {
	if app.pMin != nil {
		return nil, errors.New("the minimum is mandatory in order to build a Range instance")
	}

	if app.pMax != nil {
		return createRangeWithMaximum(*app.pMin, app.pMax), nil
	}

	return createRange(*app.pMin), nil
}
