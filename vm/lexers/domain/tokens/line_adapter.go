package tokens

import (
	"errors"
	"fmt"
)

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
	if len(data) <= 0 {
		return nil, nil, errors.New("the data must contain at least 1 element in order be converted to a Line instance")
	}

	if data[0] != LinePrefix {
		str := fmt.Sprintf("the line prefix was expected to be %d, %d provided", LinePrefix, data[0])
		return nil, nil, errors.New(str)
	}

	remaining := data[1:]
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
