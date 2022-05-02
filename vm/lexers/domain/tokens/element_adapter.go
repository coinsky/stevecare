package tokens

import (
	"errors"
	"fmt"
)

type elementAdapter struct {
	elementBuilder ElementBuilder
}

func createElementAdapter(
	elementBuilder ElementBuilder,
) ElementAdapter {
	out := elementAdapter{
		elementBuilder: elementBuilder,
	}

	return &out
}

// ToElement converts data to an element
func (app *elementAdapter) ToElement(data []byte) (Element, []byte, error) {
	if len(data) <= 1 {
		return nil, nil, errors.New("the data must contain at least 2 elements in order be converted to an Element instance")
	}

	if data[0] == TokenPrefix {

	}

	if data[0] == ReferencePrefix {
		pHeight, pSixteen, pThirtyTwo, pSixtyFour, remaining, err := parseUintData(data[1:])
		if err != nil {
			return nil, nil, err
		}

		builder := app.elementBuilder.Create()
		if pHeight != nil {
			builder.WithReference(uint(*pHeight))
		}

		if pSixteen != nil {
			builder.WithReference(uint(*pSixteen))
		}

		if pThirtyTwo != nil {
			builder.WithReference(uint(*pThirtyTwo))
		}

		if pSixtyFour != nil {
			builder.WithReference(uint(*pSixtyFour))
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, remaining, nil
	}

	if data[0] == BytePrefix {
		ins, err := app.elementBuilder.Create().WithByte(data[1]).Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, data[2:], nil
	}

	str := fmt.Sprintf("the data prefix (%d) is invalid and therefore cannot be converted to an Element instance", data[0])
	return nil, nil, errors.New(str)
}
