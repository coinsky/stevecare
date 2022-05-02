package tokens

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestElementAdapter_withByte_isSuccess(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	byteValue := uint8(8)
	elementData := []byte{
		BytePrefix,
		byteValue,
	}

	data := append(elementData, expectedRemaining...)
	element, remaining, err := NewElementAdapter().ToElement(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if element.IsReference() {
		t.Errorf("the element was expected to NOT be a reference")
		return
	}

	if element.IsToken() {
		t.Errorf("the element was expected to NOT be a token")
		return
	}

	if !element.IsByte() {
		t.Errorf("the element was expected to be a byte")
		return
	}

	pByte := element.Byte()
	if byteValue != *pByte {
		t.Errorf("the byte was expected to be %d, %d returned", byteValue, *pByte)
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestElementAdapter_withReference_withoutAdditionalData_isSuccess(t *testing.T) {
	elementData := []byte{
		ReferencePrefix,
	}

	_, _, err := NewElementAdapter().ToElement(elementData)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElementAdapter_withReference_withHeightBitSize_withoutAdditionalData_isSuccess(t *testing.T) {
	elementData := []byte{
		ReferencePrefix,
		uint8(8),
	}

	_, _, err := NewElementAdapter().ToElement(elementData)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElementAdapter_withReference_withSixteenBitSize_withoutEnoughAdditionalData_isSuccess(t *testing.T) {
	elementData := []byte{
		ReferencePrefix,
		uint8(16),
		byte(2),
	}

	_, _, err := NewElementAdapter().ToElement(elementData)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElementAdapter_withReference_withThirtyTwoBitSize_withoutEnoughAdditionalData_isSuccess(t *testing.T) {
	elementData := []byte{
		ReferencePrefix,
		uint8(32),
		byte(2),
	}

	_, _, err := NewElementAdapter().ToElement(elementData)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElementAdapter_withReference_withSixtyFourBitSize_withoutEnoughAdditionalData_isSuccess(t *testing.T) {
	elementData := []byte{
		ReferencePrefix,
		uint8(64),
		byte(2),
	}

	_, _, err := NewElementAdapter().ToElement(elementData)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElementAdapter_withReference_isHeightBit_isSuccess(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	refValue := uint8(8)
	elementData := []byte{
		ReferencePrefix,
		byte(8),
		refValue,
	}

	data := append(elementData, expectedRemaining...)
	element, remaining, err := NewElementAdapter().ToElement(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !element.IsReference() {
		t.Errorf("the element was expected to be a reference")
		return
	}

	pRef := element.Reference()
	if uint(refValue) != *pRef {
		t.Errorf("the reference was expected to be %d, %d returned", refValue, *pRef)
		return
	}

	if element.IsToken() {
		t.Errorf("the element was expected to NOT be a token")
		return
	}

	if element.IsByte() {
		t.Errorf("the element was expected to NOT be a byte")
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestElementAdapter_withReference_isSixteenBit_isSuccess(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	refValue := uint16(34354)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint16(refValueBytes, refValue)

	elementData := []byte{ReferencePrefix, uint8(16)}
	elementData = append(elementData, refValueBytes...)

	data := append(elementData, expectedRemaining...)
	element, remaining, err := NewElementAdapter().ToElement(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !element.IsReference() {
		t.Errorf("the element was expected to be a reference")
		return
	}

	pRef := element.Reference()
	if uint(refValue) != *pRef {
		t.Errorf("the reference was expected to be %d, %d returned", refValue, *pRef)
		return
	}

	if element.IsToken() {
		t.Errorf("the element was expected to NOT be a token")
		return
	}

	if element.IsByte() {
		t.Errorf("the element was expected to NOT be a byte")
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestElementAdapter_withReference_isThirtyTwoBit_isSuccess(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	refValue := uint32(3423235354)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(refValueBytes, refValue)

	elementData := []byte{ReferencePrefix, uint8(32)}
	elementData = append(elementData, refValueBytes...)

	data := append(elementData, expectedRemaining...)
	element, remaining, err := NewElementAdapter().ToElement(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !element.IsReference() {
		t.Errorf("the element was expected to be a reference")
		return
	}

	pRef := element.Reference()
	if uint(refValue) != *pRef {
		t.Errorf("the reference was expected to be %d, %d returned", refValue, *pRef)
		return
	}

	if element.IsToken() {
		t.Errorf("the element was expected to NOT be a token")
		return
	}

	if element.IsByte() {
		t.Errorf("the element was expected to NOT be a byte")
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestElementAdapter_withReference_isSixtyFourBit_isSuccess(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	refValue := uint64(34232353234554)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(refValueBytes, refValue)

	elementData := []byte{ReferencePrefix, uint8(64)}
	elementData = append(elementData, refValueBytes...)

	data := append(elementData, expectedRemaining...)
	element, remaining, err := NewElementAdapter().ToElement(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !element.IsReference() {
		t.Errorf("the element was expected to be a reference")
		return
	}

	pRef := element.Reference()
	if uint(refValue) != *pRef {
		t.Errorf("the reference was expected to be %d, %d returned", refValue, *pRef)
		return
	}

	if element.IsToken() {
		t.Errorf("the element was expected to NOT be a token")
		return
	}

	if element.IsByte() {
		t.Errorf("the element was expected to NOT be a byte")
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestElementAdapter_withReference_isInvalidBit_isError(t *testing.T) {
	expectedRemaining := []byte{0, 3, 4, 5}
	refValue := uint64(34232353234554)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(refValueBytes, refValue)

	elementData := []byte{ReferencePrefix, uint8(128)}
	elementData = append(elementData, refValueBytes...)

	data := append(elementData, expectedRemaining...)
	_, _, err := NewElementAdapter().ToElement(data)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
