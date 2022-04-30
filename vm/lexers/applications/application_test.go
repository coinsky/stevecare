package applications

import (
	"reflect"
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

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withExactlyMinOccurences_Success(t *testing.T) {
	tokenIndex := uint(0)
	minimum := uint(2)
	byteVal := []byte("(")
	data := []byte("((")

	application := NewApplication()
	token := NewTokenWithMinimumCardinalityWithByteForTests(tokenIndex, minimum, byteVal[0])
	result, err := application.Execute(token, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withMinimumPlusOccurences_returnsSuccess(t *testing.T) {
	tokenIndex := uint(0)
	minimum := uint(2)
	byteVal := []byte("(")
	data := []byte("(((")

	application := NewApplication()
	token := NewTokenWithMinimumCardinalityWithByteForTests(tokenIndex, minimum, byteVal[0])
	result, err := application.Execute(token, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withLessThanMinimum_returnsMistake(t *testing.T) {
	tokenIndex := uint(0)
	minimum := uint(2)
	byteVal := []byte("(")
	data := []byte("(")

	application := NewApplication()
	token := NewTokenWithMinimumCardinalityWithByteForTests(tokenIndex, minimum, byteVal[0])
	result, err := application.Execute(token, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	mistake := result.Mistake()
	if mistake.Index() != 0 {
		t.Errorf("the mistake was expected to be at index %d, %d returned", 0, mistake.Index())
		return
	}

	expectedPath := []uint{
		token.Index(),
	}

	retPath := mistake.Path()
	if !reflect.DeepEqual(expectedPath, retPath) {
		t.Errorf("the mistake path was expected to be %v, %v returned", expectedPath, retPath)
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withMaximumExcceeded_returnsMistake(t *testing.T) {
	tokenIndex := uint(0)
	minimum := uint(2)
	maximum := uint(5)
	byteVal := []byte("(")
	data := []byte("((((((")

	application := NewApplication()
	token := NewTokenWithRangeCardinalityWithByteForTests(tokenIndex, minimum, maximum, byteVal[0])
	result, err := application.Execute(token, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to be NOT be successful")
		return
	}

	mistake := result.Mistake()
	if mistake.Index() != 0 {
		t.Errorf("the mistake was expected to be at index %d, %d returned", 0, mistake.Index())
		return
	}

	expectedPath := []uint{
		token.Index(),
	}

	retPath := mistake.Path()
	if !reflect.DeepEqual(expectedPath, retPath) {
		t.Errorf("the mistake path was expected to be %v, %v returned", expectedPath, retPath)
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withExactlyMaximumOccurences_returnsSuccess(t *testing.T) {
	tokenIndex := uint(0)
	minimum := uint(2)
	maximum := uint(5)
	byteVal := []byte("(")
	data := []byte("(((((")

	application := NewApplication()
	token := NewTokenWithRangeCardinalityWithByteForTests(tokenIndex, minimum, maximum, byteVal[0])
	result, err := application.Execute(token, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be be successful")
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_returnsSuccess(t *testing.T) {
	tokenIndex := uint(0)
	minimum := uint(2)
	maximum := uint(5)
	byteVal := []byte("(")
	data := []byte("((((")

	application := NewApplication()
	token := NewTokenWithRangeCardinalityWithByteForTests(tokenIndex, minimum, maximum, byteVal[0])
	result, err := application.Execute(token, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be be successful")
		return
	}
}
