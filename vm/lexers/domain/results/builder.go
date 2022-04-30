package results

import "errors"

type builder struct {
	mistake Mistake
	success Success
}

func createBuilder() Builder {
	out := builder{
		mistake: nil,
		success: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMistake adds a mistake to the builder
func (app *builder) WithMistake(mistake Mistake) Builder {
	app.mistake = mistake
	return app
}

// WithSuccess adds a success to the builder
func (app *builder) WithSuccess(success Success) Builder {
	app.success = success
	return app
}

// Now builds a new Result instance
func (app *builder) Now() (Result, error) {
	if app.mistake != nil {
		return createResultWithMistake(app.mistake), nil
	}

	if app.success != nil {
		return createResultWithSuccess(app.success), nil
	}

	return nil, errors.New("the Result is invalid")
}
