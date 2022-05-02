package tokens

type lineAdapter struct {
	builder        LineBuilder
	elementAdapter ElementWithCardinalityAdapter
}

func createLineAdapter(
	builder LineBuilder,
	elementAdapter ElementWithCardinalityAdapter,
) LineAdapter {
	out := lineAdapter{
		builder:        builder,
		elementAdapter: elementAdapter,
	}

	return &out
}

// ToLine converts data to a line instance
func (app *lineAdapter) ToLine(data []byte) (Line, []byte, error) {
	remaining := data
	list := []ElementWithCardinality{}
	for {
		if len(remaining) <= 0 {
			break
		}

		element, retRemaining, err := app.elementAdapter.ToElementWithCardinality(remaining)
		if err != nil {
			break
		}

		list = append(list, element)
		remaining = retRemaining
	}

	ins, err := app.builder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}
