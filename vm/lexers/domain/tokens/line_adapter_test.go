package tokens

import (
	"reflect"
	"testing"
)

func TestLineAdapter_withRemaining_isSuccess(t *testing.T) {
	data, expectedRemaining := NewLineDataForTests(5, true)
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
	data, expectedRemaining := NewLineDataForTests(15, false)
	line, remaining, err := NewLineAdapter().ToLine(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := line.List()
	if len(list) != 15 {
		t.Errorf("%d element were expected, %d returned", 15, len(list))
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}
