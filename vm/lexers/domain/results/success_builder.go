package results

type successBuilder struct {
	pIndex *uint
}

func createSuccessBuilder() SuccessBuilder {
	out := successBuilder{
		pIndex: nil,
	}

	return &out
}

// Create initializes the builder
func (app *successBuilder) Create() SuccessBuilder {
	return createSuccessBuilder()
}

// WithIndex adds an index to the builder
func (app *successBuilder) WithIndex(index uint) SuccessBuilder {
	app.pIndex = &index
	return app
}

// Now builds a new Success instance
func (app *successBuilder) Now() (Success, error) {
	if app.pIndex != nil {
		return createSuccessWithIndex(app.pIndex), nil
	}

	return createSuccess(), nil
}
