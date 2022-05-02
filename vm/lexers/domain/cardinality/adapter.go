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
func (app *adapter) ToCardinality(data []byte) (Cardinality, error) {
	length := len(data)
	if length < 3 || length > 5 {
		str := fmt.Sprintf("%s the data was expected to contain between 3 and 5 elements, %d provided", prefixErrLabel, length)
		return nil, errors.New(str)
	}

	if data[0] != Open {
		str := fmt.Sprintf("%s the first (index: 0) element of the data was expected to contain %d, %d provided", prefixErrLabel, Open, data[0])
		return nil, errors.New(str)
	}

	builder := app.builder.Create()
	if length == 3 {
		if data[2] != Close {
			str := fmt.Sprintf("%s the data at index: 2 was expected to contain %d, %d provided, since the data length was 3", prefixErrLabel, Close, data[2])
			return nil, errors.New(str)
		}

		builder.WithSpecific(data[1])
	}

	if length == 4 {
		if data[2] != Separator {
			str := fmt.Sprintf("%s the data at index: 2 was expected to contain %d, %d provided, since the data length was 4", prefixErrLabel, Separator, data[2])
			return nil, errors.New(str)
		}

		if data[3] != Close {
			str := fmt.Sprintf("%s the data at index: 3 was expected to contain %d, %d provided, since the data length was 4", prefixErrLabel, Close, data[3])
			return nil, errors.New(str)
		}

		rnge, err := app.rangeBuilder.Create().WithMinimum(data[1]).Now()
		if err != nil {
			return nil, err
		}

		builder.WithRange(rnge)
	}

	if length == 5 {
		if data[2] != Separator {
			str := fmt.Sprintf("%s the data at index: 2 was expected to contain %d, %d provided, since the data length was 5", prefixErrLabel, Separator, data[2])
			return nil, errors.New(str)
		}

		if data[4] != Close {
			str := fmt.Sprintf("%s the data at index: 4 was expected to contain %d, %d provided, since the data length was 5", prefixErrLabel, Close, data[4])
			return nil, errors.New(str)
		}

		rnge, err := app.rangeBuilder.Create().WithMinimum(data[1]).WithMaximum(data[3]).Now()
		if err != nil {
			return nil, err
		}

		builder.WithRange(rnge)
	}

	return builder.Now()
}
