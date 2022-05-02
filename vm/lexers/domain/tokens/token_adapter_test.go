package tokens

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestTokenAdapter_withHeightBitIndex_withRemaining_isSuccess(t *testing.T) {
	data := []byte{
		TokenPrefix,
	}

	indexValue := uint8(8)
	indexData := []byte{
		byte(8),
		indexValue,
	}

	linesData, expectedRemaining := NewLinesDataForTests(10, true)

	data = append(data, indexData...)
	data = append(data, linesData...)

	token, remaining, err := NewTokenAdapter().ToToken(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := token.Index()
	if index != 8 {
		t.Errorf("the index was expected to be %d,%d returned", 8, index)
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}

	lines := token.Lines().List()
	if len(lines) != 10 {
		t.Errorf("the amount of lines were expected to be %d,%d returned", 10, len(lines))
		return
	}
}

func TestTokenAdapter_withSixteenBitIndex_withRemaining_isSuccess(t *testing.T) {
	data := []byte{
		TokenPrefix,
	}

	indexValue := uint16(8344)
	indexValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint16(indexValueBytes, indexValue)

	indexData := []byte{
		byte(16),
	}

	indexData = append(indexData, indexValueBytes...)
	linesData, expectedRemaining := NewLinesDataForTests(10, true)

	data = append(data, indexData...)
	data = append(data, linesData...)

	token, remaining, err := NewTokenAdapter().ToToken(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := token.Index()
	if index != 8344 {
		t.Errorf("the index was expected to be %d,%d returned", 8344, index)
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}

	lines := token.Lines().List()
	if len(lines) != 10 {
		t.Errorf("the amount of lines were expected to be %d,%d returned", 10, len(lines))
		return
	}
}

func TestTokenAdapter_withThirtyTwoBitIndex_withRemaining_isSuccess(t *testing.T) {
	data := []byte{
		TokenPrefix,
	}

	indexValue := uint32(83343244)
	indexValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(indexValueBytes, indexValue)

	indexData := []byte{
		byte(32),
	}

	indexData = append(indexData, indexValueBytes...)
	linesData, expectedRemaining := NewLinesDataForTests(10, true)

	data = append(data, indexData...)
	data = append(data, linesData...)

	token, remaining, err := NewTokenAdapter().ToToken(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := token.Index()
	if index != 83343244 {
		t.Errorf("the index was expected to be %d,%d returned", 83343244, index)
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}

	lines := token.Lines().List()
	if len(lines) != 10 {
		t.Errorf("the amount of lines were expected to be %d,%d returned", 10, len(lines))
		return
	}
}

func TestTokenAdapter_withSixtyFourBitIndex_withRemaining_isSuccess(t *testing.T) {
	data := []byte{
		TokenPrefix,
	}

	indexValue := uint64(8334324534544)
	indexValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexValueBytes, indexValue)

	indexData := []byte{
		byte(64),
	}

	indexData = append(indexData, indexValueBytes...)
	linesData, expectedRemaining := NewLinesDataForTests(10, true)

	data = append(data, indexData...)
	data = append(data, linesData...)

	token, remaining, err := NewTokenAdapter().ToToken(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	index := token.Index()
	if index != 8334324534544 {
		t.Errorf("the index was expected to be %d,%d returned", 8334324534544, index)
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}

	lines := token.Lines().List()
	if len(lines) != 10 {
		t.Errorf("the amount of lines were expected to be %d,%d returned", 10, len(lines))
		return
	}
}

func TestTokenAdapter_withInvalidPrefix_isError(t *testing.T) {
	data := []byte{
		0,
	}

	indexValue := uint64(8334324534544)
	indexValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexValueBytes, indexValue)

	indexData := []byte{
		byte(64),
	}

	indexData = append(indexData, indexValueBytes...)
	linesData, _ := NewLinesDataForTests(10, false)

	data = append(data, indexData...)
	data = append(data, linesData...)

	_, _, err := NewTokenAdapter().ToToken(data)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestTokenAdapter_withNotEnoughDataElements_isError(t *testing.T) {
	data := []byte{
		TokenPrefix,
	}

	_, _, err := NewTokenAdapter().ToToken(data)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
