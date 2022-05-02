package tokens

import (
	"encoding/binary"
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
	if len(data) <= 0 {
		return nil, nil, errors.New("the data must contain at least 1 element in order be converted to an Element instance")
	}

	if data[0] == TokenPrefix {

	}

	if data[0] == ReferencePrefix {
		data = data[1:]
		remaining := []byte{}
		builder := app.elementBuilder.Create()
		switch data[0] {
		case 8:
			builder.WithReference(uint(data[1]))
			remaining = data[2:]
			break
		case 16:
			value := binary.BigEndian.Uint16(data[1:9])
			builder.WithReference(uint(value))
			remaining = data[9:]
			break
		case 32:
			value := binary.BigEndian.Uint32(data[1:9])
			builder.WithReference(uint(value))
			remaining = data[9:]
			break
		case 64:
			value := binary.BigEndian.Uint64(data[1:9])
			builder.WithReference(uint(value))
			remaining = data[9:]
			break
		default:
			str := fmt.Sprintf("the referenced element was expected to contain one of these: [8, 16, 32, 64] in its data at index %d, %d provided", 1, data[1])
			return nil, nil, errors.New(str)
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, nil, err
		}

		return ins, remaining, nil
	}

	ins, err := app.elementBuilder.Create().WithByte(data[0]).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, data[1:], nil
}
