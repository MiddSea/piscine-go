package main

import "testing"

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"make school", "makeSchool"},
		{"camel case", "camelCase"},
		{"", ""},
		{"single", "single"},
	}

	for _, test := range tests {
		result := camelCase(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, result)
		}
	}
}