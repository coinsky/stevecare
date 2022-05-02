package tokens

import (
	"errors"
	"fmt"
)

type tokenAdapter struct {
	builder      TokenBuilder
	linesAdapter LinesAdapter
}

func createTokenAdapter(
	builder TokenBuilder,
	linesAdapter LinesAdapter,
) TokenAdapter {
	out := tokenAdapter{
		builder:      builder,
		linesAdapter: linesAdapter,
	}

	return &out
}

// ToToken converts data to a Token instance
func (app *tokenAdapter) ToToken(data []byte) (Token, []byte, error) {
	if len(data) <= 2 {
		return nil, nil, errors.New("the data must contain at least 2 elements in order be converted to a Token instance")
	}

	if data[0] != TokenPrefix {
		str := fmt.Sprintf("the data prefix (%d) is invalid and therefore cannot be converted to a Token instance", data[0])
		return nil, nil, errors.New(str)
	}

	pHeight, pSixteen, pThirtyTwo, pSixtyFour, remaining, err := parseUintData(data[1:])
	if err != nil {
		return nil, nil, err
	}

	lines, remainingAfterLines, err := app.linesAdapter.ToLines(remaining)
	if err != nil {
		return nil, nil, err
	}

	builder := app.builder.Create().WithLines(lines)
	if pHeight != nil {
		index := uint(*pHeight)
		builder.WithIndex(index)
	}

	if pSixteen != nil {
		index := uint(*pSixteen)
		builder.WithIndex(index)
	}

	if pThirtyTwo != nil {
		index := uint(*pThirtyTwo)
		builder.WithIndex(index)
	}

	if pSixtyFour != nil {
		index := uint(*pSixtyFour)
		builder.WithIndex(index)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remainingAfterLines, nil
}
