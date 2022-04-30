package tokens

type element struct {
	pByte      *byte
	token      Token
	pReference *uint
}

func createElementWithByte(
	pByte *byte,
) Element {
	return createElementInternally(pByte, nil, nil)
}

func createElementWithToken(
	token Token,
) Element {
	return createElementInternally(nil, token, nil)
}

func createElementWithReference(
	pReference *uint,
) Element {
	return createElementInternally(nil, nil, pReference)
}

func createElementInternally(
	pByte *byte,
	token Token,
	pReference *uint,
) Element {
	out := element{
		pByte:      pByte,
		token:      token,
		pReference: pReference,
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

// IsReference returns true if there is a reference token, false otherwise
func (obj *element) IsReference() bool {
	return obj.pReference != nil
}

// Reference returns the reference, if any
func (obj *element) Reference() *uint {
	return obj.pReference
}
