package helpers

import (
	"github.com/vlbarou/sampleproject/internal/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBlankOrEmpty(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"String is empty", args{value: constants.EmptyString}, true},
		{"String is blank", args{value: "  "}, true},
		{"String has value", args{value: "something"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlankOrEmpty(tt.args.value); got != tt.want {
				t.Errorf("IsBlankOrEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"String is empty", args{value: constants.EmptyString}, true},
		{"String is not empty", args{value: "some_value"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.value); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"String is empty", args{value: constants.EmptyString}, false},
		{"String is not empty", args{value: "some_value"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.args.value); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"String is not blank", args{value: "some_value"}, false},
		{"String is  blank", args{value: " "}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.value); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotBlank(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"String has value", args{value: "some_value"}, true},
		{"String is blank", args{value: " "}, false},
		{"String is empty NOT blank", args{value: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotBlank(tt.args.value); got != tt.want {
				t.Errorf("IsNotBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAt(t *testing.T) {

	// arrange
	input := "VNFType/SDL/VNFs/VNFName-sdl-ris-0/VNFCType/ops/VNFCs/sdl-ris-0-ops-vm0.sdl.com/Container//Processes/"

	// act & assert
	assert.Equal(t, "VNFType", GetAt(input, constants.Slash, 0))
	assert.Equal(t, "VNFs", GetAt(input, constants.Slash, 2))
	assert.Equal(t, "sdl-ris-0-ops-vm0.sdl.com", GetAt(input, constants.Slash, 7))
	assert.Equal(t, constants.EmptyString, GetAt(input, constants.Slash, 30))
	assert.Equal(t, constants.EmptyString, GetAt(input, constants.BackSlash, 30))

}

func TestTakeOne(t *testing.T) {
	assert.Equal(t, "test", TakeOne("test", "default"))
	assert.Equal(t, "default", TakeOne("", "default"))
}

func TestAreEqualIgnoreCaseTrim(t *testing.T) {
	testCases := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{
			str1:     " single  ",
			str2:     "SINGLE",
			expected: true,
		},
		{
			str1:     "  array   ",
			str2:     "ARRAY",
			expected: true,
		},
		{
			str1:     " arrayy   ",
			str2:     "ARRAY",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		result := AreEqualIgnoreCaseTrim(testCase.str1, testCase.str2)

		if result != testCase.expected {
			t.Errorf("Expected: %v, but got: %v", testCase.expected, result)
		}
	}
}

func TestIsNumberAndIsNotNumber(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{
			str:      "12345",
			expected: true,
		},
		{
			str:      "-987",
			expected: true,
		},
		{
			str:      "0",
			expected: true,
		},
		{
			str:      "123abc",
			expected: false,
		},
		{
			str:      "abc123",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		isNumber := IsNumber(testCase.str)
		isNotNumber := IsNotNumber(testCase.str)

		if isNumber != testCase.expected || isNotNumber != !testCase.expected {
			t.Errorf("Expected: %v,%v but got: %v,%v.", testCase.expected, !testCase.expected, isNumber, isNotNumber)
		}
	}
}

func TestRemoveEmptyStringOrDigit(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"", "CounterA", "123", "CounterB", ""},
			expected: []string{"CounterA", "CounterB"},
		},
		{
			input:    []string{"123", "", "456", "789", ""},
			expected: []string{},
		},
		{
			input:    []string{"CounterA", "CounterB", "CounterC", "CounterD", "CounterE"},
			expected: []string{"CounterA", "CounterB", "CounterC", "CounterD", "CounterE"},
		},
	}

	for _, testCase := range testCases {
		result := RemoveEmptyStringOrDigit(testCase.input)

		if len(result) != len(testCase.expected) {
			t.Errorf("Expected length: %d, but got: %d", len(testCase.expected), len(result))
		}

		for i := range result {
			if result[i] != testCase.expected[i] {
				t.Errorf("Mismatch at index %d. Expected: %s, but got: %s", i, testCase.expected[i], result[i])
			}
		}
	}
}
