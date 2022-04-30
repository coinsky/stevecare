package tokens

type token struct {
	lines Lines
}

func createToken(
	lines Lines,
) Token {
	out := token{
		lines: lines,
	}

	return &out
}

// Lines return the lines
func (obj *token) Lines() Lines {
	return obj.lines
}
