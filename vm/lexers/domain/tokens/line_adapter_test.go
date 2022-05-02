package tokens

import (
	"reflect"
	"testing"
)

func TestLineAdapter_withRemaining_isSuccess(t *testing.T) {
	first, _ := NewElementWithCardinalityDataForTests(false)
	second, _ := NewElementWithCardinalityDataForTests(false)
	third, _ := NewElementWithCardinalityDataForTests(false)
	fourth, _ := NewElementWithCardinalityDataForTests(false)
	fifth, expectedRemaining := NewElementWithCardinalityDataForTests(true)

	data := []byte{}
	data = append(data, first...)
	data = append(data, second...)
	data = append(data, third...)
	data = append(data, fourth...)
	data = append(data, fifth...)

	line, remaining, err := NewLineAdapter().ToLine(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := line.List()
	if len(list) != 5 {
		t.Errorf("%d element were expected, %d returned", 5, len(list))
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestLineAdapter_withoutRemaining_isSuccess(t *testing.T) {
	first, _ := NewElementWithCardinalityDataForTests(false)
	second, _ := NewElementWithCardinalityDataForTests(false)
	third, _ := NewElementWithCardinalityDataForTests(false)
	fourth, _ := NewElementWithCardinalityDataForTests(false)
	fifth, expectedRemaining := NewElementWithCardinalityDataForTests(false)

	data := []byte{}
	data = append(data, first...)
	data = append(data, second...)
	data = append(data, third...)
	data = append(data, fourth...)
	data = append(data, fifth...)

	line, remaining, err := NewLineAdapter().ToLine(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := line.List()
	if len(list) != 5 {
		t.Errorf("%d element were expected, %d returned", 5, len(list))
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}
