package applications

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

func TestLexer_withReference_withSuccessIndex_isSuccess(t *testing.T) {
	openTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(0), uint(1), []byte("(")[0]), NewCardinalityWithSpecificForTests(1))
	closeTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(1), uint(1), []byte(")")[0]), NewCardinalityWithSpecificForTests(1))
	fiveTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(2), uint(1), []byte("5")[0]), NewCardinalityWithSpecificForTests(1))
	smallerThanTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(3), uint(1), []byte("<")[0]), NewCardinalityWithSpecificForTests(1))

	conditionFirstLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		openTokenElWithCard,
		NewElementWithCardinalityWithReferenceAndCardinalityForTests(uint(4), uint(5), NewCardinalityWithSpecificForTests(1)),
		closeTokenElWithCard,
	})

	conditionSecondLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		fiveTokenElWithCard,
		smallerThanTokenElWithCard,
		fiveTokenElWithCard,
	})

	rootToken := NewTokenWithLinesForTests(uint(5), []tokens.Line{
		conditionFirstLine,
		conditionSecondLine,
	})

	data := []byte("((5<5))567")
	application := NewApplication()
	result, err := application.Execute(rootToken, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 7 {
		t.Errorf("the index was expected to be %d, %d returned", 7, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{5, 0, 5, 0, 5, 2, 3, 2, 1, 1}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withReference_withSuccessIndex_notEnoughData_isNotSuccess(t *testing.T) {
	openTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(0), uint(1), []byte("(")[0]), NewCardinalityWithSpecificForTests(1))
	closeTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(1), uint(1), []byte(")")[0]), NewCardinalityWithSpecificForTests(1))
	fiveTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(2), uint(1), []byte("5")[0]), NewCardinalityWithSpecificForTests(1))
	smallerThanTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(3), uint(1), []byte("<")[0]), NewCardinalityWithSpecificForTests(1))

	conditionFirstLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		openTokenElWithCard,
		NewElementWithCardinalityWithReferenceAndCardinalityForTests(uint(4), uint(5), NewCardinalityWithSpecificForTests(1)),
		closeTokenElWithCard,
	})

	conditionSecondLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		fiveTokenElWithCard,
		smallerThanTokenElWithCard,
		fiveTokenElWithCard,
	})

	rootToken := NewTokenWithLinesForTests(uint(5), []tokens.Line{
		conditionFirstLine,
		conditionSecondLine,
	})

	data := []byte("((5<5)")
	application := NewApplication()
	result, err := application.Execute(rootToken, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{5}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withReference_isInfiniteRecursive_isNotSuccess(t *testing.T) {
	conditionFirstLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		NewElementWithCardinalityWithReferenceAndCardinalityForTests(uint(4), uint(5), NewCardinalityWithSpecificForTests(1)),
	})

	rootToken := NewTokenWithLinesForTests(uint(5), []tokens.Line{
		conditionFirstLine,
	})

	data := []byte("((5<5))")
	application := NewApplication()
	result, err := application.Execute(rootToken, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{5}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withUndeclaredReference_isNotSuccess(t *testing.T) {
	invalidReferenceIndex := uint(20)
	openTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(0), uint(1), []byte("(")[0]), NewCardinalityWithSpecificForTests(1))
	closeTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(1), uint(1), []byte(")")[0]), NewCardinalityWithSpecificForTests(1))
	fiveTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(2), uint(1), []byte("5")[0]), NewCardinalityWithSpecificForTests(1))
	smallerThanTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(3), uint(1), []byte("<")[0]), NewCardinalityWithSpecificForTests(1))

	conditionFirstLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		openTokenElWithCard,
		NewElementWithCardinalityWithReferenceAndCardinalityForTests(uint(4), invalidReferenceIndex, NewCardinalityWithSpecificForTests(1)),
		closeTokenElWithCard,
	})

	conditionSecondLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		fiveTokenElWithCard,
		smallerThanTokenElWithCard,
		fiveTokenElWithCard,
	})

	rootToken := NewTokenWithLinesForTests(uint(5), []tokens.Line{
		conditionFirstLine,
		conditionSecondLine,
	})

	data := []byte("((5<5))")
	application := NewApplication()
	result, err := application.Execute(rootToken, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{5}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withSpecificCardinality_withSubTokens_withSuccessIndex_isSuccess(t *testing.T) {
	openTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(0), uint(1), []byte("(")[0]), NewCardinalityWithSpecificForTests(1))
	hyphenTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(1), uint(1), []byte("-")[0]), NewCardinalityWithSpecificForTests(1))
	closeTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(2), uint(1), []byte(")")[0]), NewCardinalityWithSpecificForTests(1))
	tokenLine := NewLineWithElementWithCardinalityList([]tokens.ElementWithCardinality{
		openTokenElWithCard,
		hyphenTokenElWithCard,
		closeTokenElWithCard,
	})

	rootToken := NewTokenWithLinesForTests(uint(3), []tokens.Line{
		tokenLine,
	})

	data := []byte("(-)345")
	application := NewApplication()
	result, err := application.Execute(rootToken, data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 3 {
		t.Errorf("the index was expected to be %d, %d returned", 3, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{3, 0, 1, 2}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}

}

func TestLexer_withOneLine_withSpecificCardinality_withByte_withoutSuccessIndex_isSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 1 {
		t.Errorf("the index was expected to be %d, %d returned", 1, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withExactlyMinOccurences_isSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 2 {
		t.Errorf("the index was expected to be %d, %d returned", 2, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withMinimumPlusOccurences_isSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 3 {
		t.Errorf("the index was expected to be %d, %d returned", 3, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withLessThanMinimum_isNotSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withMaximumExcceeded_isNotSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withExactlyMaximumOccurences_isSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 5 {
		t.Errorf("the index was expected to be %d, %d returned", 5, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_isSuccess(t *testing.T) {
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

	index := result.Index()
	if index != 4 {
		t.Errorf("the index was expected to be %d, %d returned", 4, index)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{0}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}
