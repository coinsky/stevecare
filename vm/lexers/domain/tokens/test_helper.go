package tokens

import (
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
)

// NewLinesDataForTests returns the line's data for tests
func NewLinesDataForTests(amountOfLines uint, withRemaining bool) ([]byte, []byte) {
	data := []byte{}
	lastRemaining := []byte{}
	for i := 0; i < int(amountOfLines); i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		amountOfElements := r1.Intn(15) + 1

		hasRemaining := (i+1 >= int(amountOfLines)) && withRemaining
		elementData, remainingData := NewLineDataForTests(uint(amountOfElements), hasRemaining)
		lastRemaining = remainingData
		data = append(data, elementData...)
	}

	return data, lastRemaining
}

// NewLineDataForTests returns the line data for tests
func NewLineDataForTests(amountOfElements uint, withRemaining bool) ([]byte, []byte) {
	data := []byte{
		LinePrefix,
	}

	lastRemaining := []byte{}
	for i := 0; i < int(amountOfElements); i++ {
		hasRemaining := (i+1 >= int(amountOfElements)) && withRemaining
		elementData, remainingData := NewElementWithCardinalityDataForTests(hasRemaining)
		lastRemaining = remainingData
		data = append(data, elementData...)
	}

	return data, lastRemaining
}

// NewElementWithCardinalityDataForTests returns ElementWithCardinality's data for tests
func NewElementWithCardinalityDataForTests(withRemaining bool) ([]byte, []byte) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	expectedRemaining := []byte{}
	if withRemaining {
		expectedRemaining = []byte{
			uint8(r1.Intn(255)),
			uint8(r1.Intn(255)),
			uint8(r1.Intn(255)),
			uint8(r1.Intn(255)),
		}
	}

	refValue := uint32(r1.Intn(65000))

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
	return data, expectedRemaining
}
