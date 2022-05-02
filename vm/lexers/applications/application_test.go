package applications

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/channels"
	"github.com/steve-care-software/stevecare/vm/lexers/domain/tokens"
)

/*
func TestCompile_oneToken_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		0@rootToken: $37[23,32] rootToken? subToken?;
		1@subToken: $37;
	`
	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	fmt.Printf("\n%s\n", grammar)
}*/

func TestLexer_withReference_withSuccessIndex_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		4@rootToken : .openParenthesis .rootToken .closeParenthesis
					| .five .minus .five
					;

		0@openParenthesis: $40;
		1@closeParenthesis: $41;
		2@five: $53;
		3@minus: $60;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	data := []byte("((5<5))567")
	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 7 {
		t.Errorf("the cursor was expected to be %d, %d returned", 7, cursor)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{4, 0, 4, 0, 4, 2, 3, 2, 1, 1}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withReference_withSuccessIndex_notEnoughData_cannotHavePrefix_isNotSuccess(t *testing.T) {
	script := `
		%rootToken;
		4@rootToken : .openParenthesis .rootToken .closeParenthesis
					| .five .minus .five
					;

		0@openParenthesis: $40;
		1@closeParenthesis: $41;
		2@five: $53;
		3@minus: $60;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	data := []byte("((5<5)")
	result, err := application.Execute(grammar, data, false)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d,%d returned", 0, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 0 {
		t.Errorf("the cursor was expected to be %d, %d returned", 0, cursor)
		return
	}

	if result.IsSuccess() {
		t.Errorf("the result was expected to NOT be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{4}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withReference_withSuccessIndex_notEnoughData_withPrefix_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		4@rootToken : .openParenthesis .rootToken .closeParenthesis
					| .five .minus .five
					;

		0@openParenthesis: $40;
		1@closeParenthesis: $41;
		2@five: $53;
		3@minus: $60;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	data := []byte("((5<5)")
	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 1 {
		t.Errorf("the index was expected to be %d,%d returned", 1, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 6 {
		t.Errorf("the cursor was expected to be %d, %d returned", 6, cursor)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{4, 0, 4, 2, 3, 2, 1}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withReference_isInfiniteRecursive_isNotSuccess(t *testing.T) {
	script := `
		%rootToken;
		5@rootToken : .rootToken;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	data := []byte("((5<5))")
	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 0 {
		t.Errorf("the cursor was expected to be %d, %d returned", 0, cursor)
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

// we don't use the script here because its impossible to convert a script to a grammar with an undeclared reference
func TestLexer_withUndeclaredReference_withPrefix_isSuccess(t *testing.T) {
	invalidReferenceIndex := uint(20)
	openTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(0), uint8(1), []byte("(")[0]), NewCardinalityWithSpecificForTests(1))
	closeTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(1), uint8(1), []byte(")")[0]), NewCardinalityWithSpecificForTests(1))
	fiveTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(2), uint8(1), []byte("5")[0]), NewCardinalityWithSpecificForTests(1))
	smallerThanTokenElWithCard := NewElementWithCardinalityWithTokenAndCardinalityForTests(NewTokenWithSpecificCardinalityWithByteForTests(uint(3), uint8(1), []byte("<")[0]), NewCardinalityWithSpecificForTests(1))

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

	grammar := NewGrammarForTests(NewTokenWithLinesForTests(uint(5), []tokens.Line{
		conditionFirstLine,
		conditionSecondLine,
	}))

	data := []byte("((5<5))")
	application := NewApplication()
	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 2 {
		t.Errorf("the index was expected to be %d, %d returned", 2, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
		return
	}

	if !result.IsSuccess() {
		t.Errorf("the result was expected to be successful")
		return
	}

	path := result.Path()
	expectedPath := []uint{5, 2, 3, 2}
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf("the path was expected to be %v, %v returned", expectedPath, path)
		return
	}
}

func TestLexer_withOneLine_withSpecificCardinality_withSubTokens_withSuccessIndex_isSuccess(t *testing.T) {
	script := `
		%rootToken;
		3@rootToken : .openParenthesis .hyphen .closeParenthesis;
		0@openParenthesis: $40;
		1@hyphen: $45;
		2@closeParenthesis: $41;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	data := []byte("(-)345")
	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 3 {
		t.Errorf("the cursor was expected to be %d, %d returned", 3, cursor)
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
	script := `
		%openParenthesis;
		0@openParenthesis : $40;
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 1 {
		t.Errorf("the cursor was expected to be %d, %d returned", 1, cursor)
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

	script := `
		%openParenthesis;
		0@openParenthesis : $40[2,];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("(("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 2 {
		t.Errorf("the cursor was expected to be %d, %d returned", 2, cursor)
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
	script := `
		%openParenthesis;
		0@openParenthesis : $40[2,];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 3 {
		t.Errorf("the cursor was expected to be %d, %d returned", 3, cursor)
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
	script := `
		%openParenthesis;
		0@openParenthesis : $40[2,];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 0 {
		t.Errorf("the cursor was expected to be %d, %d returned", 0, cursor)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withMaximumExcceeded_withPrefix_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		0@openParenthesis : $40[2,5];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("(((((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 0 {
		t.Errorf("the index was expected to be %d, %d returned", 0, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withExactlyMaximumOccurences_isSuccess(t *testing.T) {
	script := `
		%openParenthesis;
		0@openParenthesis : $40[2,5];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("((((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
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
	script := `
		%openParenthesis;
		0@openParenthesis : $40[2,5];
	`

	application := NewApplication()
	grammar, err := application.Compile(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	result, err := application.Execute(grammar, []byte("(((("), true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 4 {
		t.Errorf("the cursor was expected to be %d, %d returned", 4, cursor)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_withChannel_isSuccess(t *testing.T) {
	minimum := uint8(2)
	maximum := uint8(5)
	byteVal := []byte("(")
	data := []byte(" (((( ")

	application := NewApplication()
	grammar := NewGrammarWithChannelsForTests(
		NewTokenWithRangeCardinalityWithByteForTests(uint(0), minimum, maximum, byteVal[0]),
		[]channels.Channel{
			NewChannelForTests(NewTokenWithMinimumCardinalityWithByteForTests(uint(1), uint8(0), []byte(" ")[0])),
		},
	)

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 6 {
		t.Errorf("the cursor was expected to be %d, %d returned", 6, cursor)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_withChannel_withOpenParenthesisNextElement_isSuccess(t *testing.T) {
	minimum := uint8(2)
	maximum := uint8(5)
	byteVal := []byte("(")
	data := []byte(" (((( ")

	application := NewApplication()
	grammar := NewGrammarWithChannelsForTests(
		NewTokenWithRangeCardinalityWithByteForTests(uint(0), minimum, maximum, byteVal[0]),
		[]channels.Channel{
			NewChannelWithConditionsForTests(
				NewTokenWithMinimumCardinalityWithByteForTests(uint(1), uint8(0), []byte(" ")[0]),
				NewConditionWithNext(
					NewTokenWithRangeCardinalityWithByteForTests(uint(2), uint8(0), uint8(2), []byte("(")[0]),
				),
			),
		},
	)

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
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

func TestLexer_withOneLine_withRangeCardinality_withByte_withinRangeOccurences_withChannel_withOpenParenthesisPreviousElement_isSuccess(t *testing.T) {
	minimum := uint8(2)
	maximum := uint8(5)
	byteVal := []byte("(")
	data := []byte(" (((( ")

	application := NewApplication()
	grammar := NewGrammarWithChannelsForTests(
		NewTokenWithRangeCardinalityWithByteForTests(uint(0), minimum, maximum, byteVal[0]),
		[]channels.Channel{
			NewChannelWithConditionsForTests(
				NewTokenWithMinimumCardinalityWithByteForTests(uint(1), uint8(0), []byte(" ")[0]),
				NewConditionWithPrevious(
					NewTokenWithRangeCardinalityWithByteForTests(uint(2), uint8(0), uint8(2), []byte("(")[0]),
				),
			),
		},
	)

	result, err := application.Execute(grammar, data, true)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := result.Index()
	if index != 1 {
		t.Errorf("the index was expected to be %d, %d returned", 1, index)
		return
	}

	cursor := result.Cursor()
	if cursor != 5 {
		t.Errorf("the cursor was expected to be %d, %d returned", 5, cursor)
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
