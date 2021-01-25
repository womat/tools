package tools

import (
	"strings"
	"testing"
)

func TestIn(t *testing.T) {

	if In(1, 2, 3) {
		t.Errorf("value: %v shouldn't be found", 1)
	}

	if !In("0", "0") {
		t.Errorf("value: %v should be found", "0")
	}

	fields := strings.FieldsFunc("csv,influx", func(c rune) bool { return c == ',' })
	is := make([]interface{}, len(fields))
	for i, v := range fields {
		is[i] = v
	}

	if !In("csv", is...) {
		t.Errorf("value: %v should be found", "csv")
	}
}
