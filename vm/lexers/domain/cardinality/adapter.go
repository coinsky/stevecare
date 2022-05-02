package cardinality

import (
	"errors"
	"fmt"
)

type adapter struct {
	builder      Builder
	rangeBuilder RangeBuilder
}

func createAdapter(
	builder Builder,
	rangeBuilder RangeBuilder,
) Adapter {
	out := adapter{
		builder:      builder,
		rangeBuilder: rangeBuilder,
	}

	return &out
}

// ToCardinality converts data to cardinality instance
func (app *adapter) ToCardinality(data []byte) (Cardinality, []byte, error) {
	length := len(data)
	if length < 3 {
		str := fmt.Sprintf("%s the data was expected to contain atleast 3 elements, %d provided", prefixErrLabel, length)
		return nil, nil, errors.New(str)
	}

	if data[0] != Open {
		str := fmt.Sprintf("%s the first (index: 0) element of the data was expected to contain %d, %d provided", prefixErrLabel, Open, data[0])
		return nil, nil, errors.New(str)
	}

	builder := app.builder.Create()
	if data[2] == Close {
		ins, err := builder.WithSpecific(data[1]).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, data[3:], nil
	}

	if length >= 4 && data[3] == Close {
		if data[2] != Separator {
			str := fmt.Sprintf("%s the data at index: 2 was expected to contain %d, %d provided, since the data length was 4", prefixErrLabel, Separator, data[2])
			return nil, nil, errors.New(str)
		}

		rnge, err := app.rangeBuilder.Create().WithMinimum(data[1]).Now()
		if err != nil {
			return nil, nil, err
		}

		ins, err := builder.WithRange(rnge).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, data[4:], nil
	}

	if length >= 5 && data[4] == Close {
		if data[2] != Separator {
			str := fmt.Sprintf("%s the data at index: 2 was expected to contain %d, %d provided, since the data length was 5", prefixErrLabel, Separator, data[2])
			return nil, nil, errors.New(str)
		}

		rnge, err := app.rangeBuilder.Create().WithMinimum(data[1]).WithMaximum(data[3]).Now()
		if err != nil {
			return nil, nil, err
		}

		ins, err := builder.WithRange(rnge).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, data[5:], nil
	}

	return nil, nil, errors.New("the data could not be converted to a Cardinality instance")
}
