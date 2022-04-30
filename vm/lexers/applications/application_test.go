package applications

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

func TestLexer_withReference_withSuccessIndex_Success(t *testing.T) {
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

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	success := result.Success()
	if !success.HasIndex() {
		t.Errorf("the success was expected to contain an index")
		return
	}

	pIndex := success.Index()
	if *pIndex != 7 {
		t.Errorf("the index was expected to be %d, %d returned", 7, *pIndex)
		return
	}
}

func TestLexer_withReference_isInfiniteRecursive_Mistake(t *testing.T) {
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

	if !result.IsMistake() {
		t.Errorf("the result was expected to be successful")
		return
	}
}

func TestLexer_withUndeclaredReference_Mistake(t *testing.T) {
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

	if !result.IsMistake() {
		t.Errorf("the result was expected to be successful")
		return
	}
}

func TestLexer_withOneLine_withSpecificCardinality_withSubTokens_withSuccessIndex_Success(t *testing.T) {
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

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	success := result.Success()
	if !success.HasIndex() {
		t.Errorf("the success was expected to contain an index")
		return
	}

	pIndex := success.Index()
	if *pIndex != 3 {
		t.Errorf("the index was expected to be %d, %d returned", 3, *pIndex)
		return
	}

}

func TestLexer_withOneLine_withSpecificCardinality_withByte_withoutSuccessIndex_Success(t *testing.T) {
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

	success := result.Success()
	if success.HasIndex() {
		t.Errorf("the success was expected to NOT contain an index")
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

	success := result.Success()
	if success.HasIndex() {
		t.Errorf("the success was expected to NOT contain an index")
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withMinimumPlusOccurences_Success(t *testing.T) {
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

	success := result.Success()
	if success.HasIndex() {
		t.Errorf("the success was expected to NOT contain an index")
		return
	}
}

func TestLexer_withOneLine_withMinimumCardinality_withByte_withLessThanMinimum_Mistake(t *testing.T) {
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

	index := mistake.Index()
	if index != 0 {
		t.Errorf("the mistake index was expected to be %d, %d returned", 0, index)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withMaximumExcceeded_Mistake(t *testing.T) {
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

	index := mistake.Index()
	if index != 0 {
		t.Errorf("the mistake index was expected to be %d, %d returned", 0, index)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withExactlyMaximumOccurences_Success(t *testing.T) {
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

	success := result.Success()
	if success.HasIndex() {
		t.Errorf("the success was expected to NOT contain an index")
		return
	}
}

func TestLexer_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_Success(t *testing.T) {
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

	success := result.Success()
	if success.HasIndex() {
		t.Errorf("the success was expected to NOT contain an index")
		return
	}
}
