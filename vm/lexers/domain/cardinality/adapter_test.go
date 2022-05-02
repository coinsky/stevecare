package cardinality

import (
	"testing"
)

func TestAdapter_withSpecific_isSuccess(t *testing.T) {
	specific := uint8(8)
	data := []byte{
		Open,
		specific,
		Close,
	}

	cardinality, err := NewAdapter().ToCardinality(data)
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
}

func TestAdapter_withMinimum_isSuccess(t *testing.T) {
	minimum := uint8(8)
	data := []byte{
		Open,
		minimum,
		Separator,
		Close,
	}

	cardinality, err := NewAdapter().ToCardinality(data)
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
}

func TestAdapter_withMinimum_withMaximum_isSuccess(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(10)
	data := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Close,
	}

	cardinality, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}

func TestAdapter_withMinimum_withDataTooBig_isError(t *testing.T) {
	minimum := uint8(8)
	maximum := uint8(16)
	data := []byte{
		Open,
		minimum,
		Separator,
		maximum,
		Close,
		Close,
	}

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
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

	_, err := NewAdapter().ToCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to valid, nil returned")
		return
	}
}
