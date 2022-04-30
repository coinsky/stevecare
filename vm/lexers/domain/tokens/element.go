package tokens

type element struct {
	pByte *byte
	token Token
}

func createElementWithByte(
	pByte *byte,
) Element {
	return createElementInternally(pByte, nil)
}

func createElementWithToken(
	token Token,
) Element {
	return createElementInternally(nil, token)
}

func createElementInternally(
	pByte *byte,
	token Token,
) Element {
	out := element{
		pByte: pByte,
		token: token,
	}

	return &out
}

// IsByte returns true if there is a byte, false otherwise
func (obj *element) IsByte() bool {
	return obj.pByte != nil
}

// Byte returns the byte, if any
func (obj *element) Byte() *byte {
	return obj.pByte
}

// IsToken returns true if there is a token, false otherwise
func (obj *element) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *element) Token() Token {
	return obj.token
}
