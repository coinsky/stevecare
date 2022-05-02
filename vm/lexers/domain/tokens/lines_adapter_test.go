package tokens

import (
	"reflect"
	"testing"
)

func TestLinesAdapter_withRemaining_isSuccess(t *testing.T) {
	data, expectedRemaining := NewLinesDataForTests(5, true)
	lines, remaining, err := NewLinesAdapter().ToLines(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := lines.List()
	if len(list) != 5 {
		t.Errorf("%d element were expected, %d returned", 5, len(list))
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}

func TestLinesAdapter_withoutRemaining_isSuccess(t *testing.T) {
	data, expectedRemaining := NewLinesDataForTests(5, false)
	lines, remaining, err := NewLinesAdapter().ToLines(data)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := lines.List()
	if len(list) != 5 {
		t.Errorf("%d element were expected, %d returned", 5, len(list))
		return
	}

	if !reflect.DeepEqual(remaining, expectedRemaining) {
		t.Errorf("the expected remaining data was expected to be: %v, %v, returned", expectedRemaining, remaining)
		return
	}
}
