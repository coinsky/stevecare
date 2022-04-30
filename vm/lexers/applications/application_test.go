package applications

import (
	"fmt"
	"testing"
)

func TestLexer_withOneLine_withSpecificCardinality_withByte_Success(t *testing.T) {
	tokenIndex := uint(0)
	specific := uint(1)
	byteVal := []byte("(")

	application := NewApplication()
	token := NewTokenWithSpecificCardinalityWithByteForTests(tokenIndex, specific, byteVal[0])
	result, err := application.Execute(token, byteVal)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%v\n", result)
}
