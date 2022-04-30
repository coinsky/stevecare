package results

type result struct {
	mistake Mistake
	success Success
}

func createResultWithMistake(
	mistake Mistake,
) Result {
	return createResultInternally(mistake, nil)
}

func createResultWithSuccess(
	success Success,
) Result {
	return createResultInternally(nil, success)
}

func createResultInternally(
	mistake Mistake,
	success Success,
) Result {
	out := result{
		mistake: mistake,
		success: success,
	}

	return &out
}

// IsMistake returns true if there is a mistake, false otherwise
func (obj *result) IsMistake() bool {
	return obj.mistake != nil
}

// Mistake returns the mistake, if any
func (obj *result) Mistake() Mistake {
	return obj.mistake
}

// IsSuccess returns true if there is a success, false otherwise
func (obj *result) IsSuccess() bool {
	return obj.success != nil
}

// Success returns the success, if any
func (obj *result) Success() Success {
	return obj.success
}
