package tokens

import (
	"encoding/binary"
	"math/rand"
	"time"

	"github.com/steve-care-software/stevecare/vm/lexers/domain/cardinality"
)

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
