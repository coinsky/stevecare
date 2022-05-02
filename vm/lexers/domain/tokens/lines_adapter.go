package tokens

type linesAdapter struct {
	builder     LinesBuilder
	lineAdapter LineAdapter
}

func createLinesAdapter(
	builder LinesBuilder,
	lineAdapter LineAdapter,
) LinesAdapter {
	out := linesAdapter{
		builder:     builder,
		lineAdapter: lineAdapter,
	}

	return &out
}

// ToLines converts data to lines
func (app *linesAdapter) ToLines(data []byte) (Lines, []byte, error) {
	remaining := data
	list := []Line{}
	for {
		if len(remaining) <= 0 {
			break
		}

		element, retRemaining, err := app.lineAdapter.ToLine(remaining)
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
