// demonstrate unit testing for json unmarshal
// test case format: json string, and expected struct value after unmarshal
// reference: unit testing basics: https://blog.alexellis.io/golang-writing-unit-tests/

package main

import (
	"testing"
)

func TestConvJson(t *testing.T) {

	testcases := []struct {
		name     string
		input    string
		expected JsonSt
	}{
		{"simple case 1", `{"First":"James","Last":"Bond"}`, JsonSt{"James", "Bond"}},
		{"simple case 2", `{"First":"Bruce","Last":"Wills"}`, JsonSt{"Bruce", "Wills"}},
		{"blank first name case ", `{"First":"","Last":"Wills"}`, JsonSt{"", "Wills"}},
		{"missing first name case ", `{"Last":"Wills"}`, JsonSt{"", "Wills"}},
	}

	for _, tc := range testcases {
		actual := ConvJson(tc.input)
		if actual != tc.expected {
			t.Errorf("test case %v failed: Expected %v , Actual %v", tc.name, actual, tc.expected)
		}
	}
}
