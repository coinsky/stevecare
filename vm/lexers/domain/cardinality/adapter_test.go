package cardinality

import (
	"reflect"
	"testing"
)

func TestAdapter_withSpecific_withRemaining_isSuccess(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	specific := uint8(8)
	cardinalityData := []byte{
		Open,
		specific,
		Close,
	}

	data := append(cardinalityData, expectedRemaining...)
	cardinality, retRemaining, err := NewAdapter().ToCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !cardinality.IsSpecific() {
		t.Errorf("the cardinality was expected to contain a specific amount")
		return
	}

	if cardinality.IsRange() {
		t.Errorf("the cardinality was expected to NOT contain a range")
		return
	}

	pSpecific := cardinality.Specific()
	if specific != *pSpecific {
		t.Errorf("the specific amount was expected to be %d, %d returned", specific, *pSpecific)
		return
	}

	if !reflect.DeepEqual(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, retRemaining)
		return
	}
}

func TestAdapter_withSpecific_withoutRemaining_isSuccess(t *testing.T) {
	expectedRemaining := []byte{}
	specific := uint8(8)
	cardinalityData := []byte{
		Open,
		specific,
		Close,
	}

	data := append(cardinalityData, expectedRemaining...)
	cardinality, retRemaining, err := NewAdapter().ToCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !cardinality.IsSpecific() {
		t.Errorf("the cardinality was expected to contain a specific amount")
		return
	}

	if cardinality.IsRange() {
		t.Errorf("the cardinality was expected to NOT contain a range")
		return
	}

	pSpecific := cardinality.Specific()
	if specific != *pSpecific {
		t.Errorf("the specific amount was expected to be %d, %d returned", specific, *pSpecific)
		return
	}

	if !reflect.DeepEqual(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, retRemaining)
		return
	}
}

func TestAdapter_withMinimum_withRemaining_isSuccess(t *testing.T) {
	minimum := uint8(8)
	expectedRemaining := []byte{0, 3, 4, 5}
	cardinalityData := []byte{
		Open,
		minimum,
		Separator,
		Close,
	}

	data := append(cardinalityData, expectedRemaining...)
	cardinality, retRemaining, err := NewAdapter().ToCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if cardinality.IsSpecific() {
		t.Errorf("the cardinality was expected to NOT contain a specific amount")
		return
	}

	if !cardinality.IsRange() {
		t.Errorf("the cardinality was expected to contain a range")
		return
	}

	rnge := cardinality.Range()
	if rnge.HasMax() {
		t.Errorf("the range was expected to NOT contain a maximum")
		return
	}

	min := rnge.Min()
	if minimum != min {
		t.Errorf("the minimum was expected to be %d, %d returned", minimum, min)
		return
	}

	if !reflect.DeepEqual(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, retRemaining)
		return
	}
}

func TestAdapter_withMinimum_withoutRemaining_isSuccess(t *testing.T) {
	minimum := uint8(8)
	expectedRemaining := []byte{}
	cardinalityData := []byte{
		Open,
		minimum,
		Separator,
		Close,
	}

	data := append(cardinalityData, expectedRemaining...)
	cardinality, retRemaining, err := NewAdapter().ToCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if cardinality.IsSpecific() {
		t.Errorf("the cardinality was expected to NOT contain a specific amount")
		return
	}

	if !cardinality.IsRange() {
		t.Errorf("the cardinality was expected to contain a range")
		return
	}

	rnge := cardinality.Range()
	if rnge.HasMax() {
		t.Errorf("the range was expected to NOT contain a maximum")
		return
	}

	min := rnge.Min()
	if minimum != min {
		t.Errorf("the minimum was expected to be %d, %d returned", minimum, min)
		return
	}

	if !reflect.DeepEqual(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, retRemaining)
		return
	}
}

func TestAdapter_withMinimum_withMaximum_withRemaining_isSuccess(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(10)
	expectedRemaining := []byte{0, 3, 4, 5}
	cardinalityData := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Close,
	}

	data := append(cardinalityData, expectedRemaining...)
	cardinality, retRemaining, err := NewAdapter().ToCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if cardinality.IsSpecific() {
		t.Errorf("the cardinality was expected to NOT contain a specific amount")
		return
	}

	if !cardinality.IsRange() {
		t.Errorf("the cardinality was expected to contain a range")
		return
	}

	rnge := cardinality.Range()
	if !rnge.HasMax() {
		t.Errorf("the range was expected to contain a maximum")
		return
	}

	pMax := rnge.Max()
	if maximum != *pMax {
		t.Errorf("the maximum was expected to be %d, %d returned", maximum, *pMax)
		return
	}

	min := rnge.Min()
	if minimum != min {
		t.Errorf("the minimum was expected to be %d, %d returned", minimum, min)
		return
	}

	if !reflect.DeepEqual(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, retRemaining)
		return
	}
}

func TestAdapter_withMinimum_withMaximum_isSuccess(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(10)
	expectedRemaining := []byte{0, 3, 4, 5}
	cardinalityData := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Close,
	}

	data := append(cardinalityData, expectedRemaining...)
	cardinality, retRemaining, err := NewAdapter().ToCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if cardinality.IsSpecific() {
		t.Errorf("the cardinality was expected to NOT contain a specific amount")
		return
	}

	if !cardinality.IsRange() {
		t.Errorf("the cardinality was expected to contain a range")
		return
	}

	rnge := cardinality.Range()
	if !rnge.HasMax() {
		t.Errorf("the range was expected to contain a maximum")
		return
	}

	pMax := rnge.Max()
	if maximum != *pMax {
		t.Errorf("the maximum was expected to be %d, %d returned", maximum, *pMax)
		return
	}

	min := rnge.Min()
	if minimum != min {
		t.Errorf("the minimum was expected to be %d, %d returned", minimum, min)
		return
	}

	if !reflect.DeepEqual(expectedRemaining, retRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, retRemaining)
		return
	}
}

func TestAdapter_withMinimum_withMaximum_maximumIsSmallerThanMinimum_isError(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(6)
	data := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Close,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withMinimum_withMaximum_maximumIsSameAsMinimum_isError(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(8)
	data := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Close,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withMinimum_withDataTooSmall_isError(t *testing.T) {
	data := []byte{
		Open,
		Close,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withMinimum_firstDataElementIsInvalid_isError(t *testing.T) {
	specific := uint8(8)
	data := []byte{
		specific,
		specific,
		Close,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withSpecific_closeElementIsInvalid_isError(t *testing.T) {
	specific := uint8(8)
	data := []byte{
		Open,
		specific,
		specific,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withMinimum_withInvalidSeparator_isError(t *testing.T) {
	minimum := uint8(8)
	data := []byte{
		Open,
		minimum,
		Open,
		Close,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withMinimum_withInvalidClose_isError(t *testing.T) {
	minimum := uint8(8)
	data := []byte{
		Open,
		minimum,
		Separator,
		Open,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withRange_withInvalidSeparator_isError(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(15)
	data := []byte{
		Open,
		minimum,
		Open,
		maximum,
		Close,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withRange_withInvalidClose_isError(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(15)
	data := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Separator,
	}

	_, _, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}
