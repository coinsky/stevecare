package tokens

type token struct {
	index uint
	lines Lines
}

func createToken(
	index uint,
	lines Lines,
) Token {
	out := token{
		index: index,
		lines: lines,
	}

	return &out
}

// Index return the index
func (obj *token) Index() uint {
	return obj.index
}

// Lines return the lines
func (obj *token) Lines() Lines {
	return obj.lines
}
