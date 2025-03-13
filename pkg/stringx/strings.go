package helpers

import (
	"github.com/vlbarou/sampleproject/internal/constants"
	"regexp"
	"strconv"
	"strings"
)

func AreEqual(str1, str2 string) bool {
	return strings.Compare(str1, str2) == 0
}

func AreEqualTrim(str1, str2 string) bool {
	return strings.Compare(strings.TrimSpace(str1), strings.TrimSpace(str2)) == 0
}

func AreEqualIgnoreCaseTrim(str1, str2 string) bool {
	return AreEqualTrim(strings.ToLower(str1), strings.ToLower(str2))
}

func AreEqualIgnoreCase(str1, str2 string) bool {
	return AreEqual(strings.ToLower(str1), strings.ToLower(str2))
}

func IsBlankOrEmpty(value string) bool {

	return IsBlank(value) || IsEmpty(value)
}

func IsEmpty(value string) bool {
	if value == constants.EmptyString {
		return true
	}
	return false
}

func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

func IsBlank(value string) bool {
	if value == constants.EmptyString {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(value) {
		return true
	}
	return false
}

func IsNotBlank(value string) bool {
	return !IsBlank(value)
}

func GetAt(input string, delimeter string, position int) string {

	arrayString := strings.Split(input, delimeter)

	if len(arrayString) <= position {
		return constants.EmptyString
	}

	return arrayString[position]
}

// TakeOne returns valid string if not empty or later one.
func TakeOne(valid, or string) string {
	if len(valid) > constants.Zero {
		return valid
	}

	return or
}

func RemoveEmptyStringOrDigit(s []string) []string {

	var result []string

	for _, value := range s {
		if IsNotEmpty(value) && IsNotNumber(value) {
			result = append(result, value)
		}
	}
	return result
}

func IsNumber(str string) bool {
	_, err := strconv.ParseInt(str, 10, 32)
	return err == nil
}

func IsNotNumber(str string) bool {
	return !IsNumber(str)
}
