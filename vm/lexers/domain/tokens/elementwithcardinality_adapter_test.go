package tokens

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
)

func TestElementWithCardinalityAdapter_withRemainingData_isSuccess(t *testing.T) {
	expectedRemaining := []byte{53, 34, 54, 56}
	refValue := uint32(523534)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(refValueBytes, refValue)

	elementData := []byte{
		ReferencePrefix,
		byte(32),
	}

	elementData = append(elementData, refValueBytes...)

	specific := uint8(12)
	cardinalityData := []byte{
		cardinality.Open,
		specific,
		cardinality.Close,
	}

	data := []byte{}
	data = append(data, elementData...)
	data = append(data, cardinalityData...)
	data = append(data, expectedRemaining...)

	elementWithCardinality, remaining, err := NewElementWithCardinalityAdapter().ToElementWithCardinality(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}

	pReference := elementWithCardinality.Element().Reference()
	if uint(refValue) != *pReference {
		t.Errorf("the element's reference was expected to be %d, %d returned", uint(refValue), *pReference)
		return
	}

	pSpecific := elementWithCardinality.Cardinality().Specific()
	if specific != *pSpecific {
		t.Errorf("the cardinality's specific was expected to be %d, %d returned", specific, *pSpecific)
		return
	}
}

func TestElementWithCardinalityAdapter_withInvalidElement_isSuccess(t *testing.T) {
	expectedRemaining := []byte{53, 34, 54, 56}
	refValue := uint32(523534)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(refValueBytes, refValue)

	elementData := []byte{
		byte(0),
		byte(32),
	}

	elementData = append(elementData, refValueBytes...)

	specific := uint8(12)
	cardinalityData := []byte{
		cardinality.Open,
		specific,
		cardinality.Close,
	}

	data := []byte{}
	data = append(data, elementData...)
	data = append(data, cardinalityData...)
	data = append(data, expectedRemaining...)

	_, _, err := NewElementWithCardinalityAdapter().ToElementWithCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestElementWithCardinalityAdapter_withInvalidCardinality_isSuccess(t *testing.T) {
	expectedRemaining := []byte{53, 34, 54, 56}
	refValue := uint32(523534)

	refValueBytes := make([]byte, 8)
	binary.BigEndian.PutUint32(refValueBytes, refValue)

	elementData := []byte{
		ReferencePrefix,
		byte(32),
	}

	elementData = append(elementData, refValueBytes...)

	specific := uint8(1)
	cardinalityData := []byte{
		specific,
		cardinality.Close,
	}

	data := []byte{}
	data = append(data, elementData...)
	data = append(data, cardinalityData...)
	data = append(data, expectedRemaining...)

	_, _, err := NewElementWithCardinalityAdapter().ToElementWithCardinality(data)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
