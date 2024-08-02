package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// FormatRupiah formats a given amount in float64 to Indonesian Rupiah currency format.
//
// Parameters:
//   - amount: The amount to be formatted.
//
// Returns:
//   - string: The formatted string in Indonesian Rupiah currency format (e.g., "Rp 1,234.56").
func FormatRupiah(amount float64) string {
	amountStr := fmt.Sprintf("%.2f", amount)

	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	n := len(integerPart)
	var formattedInt string
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			formattedInt += ","
		}
		formattedInt += string(integerPart[i])
	}

	return fmt.Sprintf("Rp %s.%s", formattedInt, decimalPart)
}

// BoolToYesNo converts a boolean value to a string representation of "Yes" or "No".
//
// Parameters:
//   - b: The boolean value to be converted.
//
// Returns:
//   - string: "Yes" if the input is true, otherwise "No".
func BoolToYesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

// ParseDataType parses a string value into a specified type T.
// It supports float64, int, bool, and time.Time types. If the type is unsupported, it returns an error.
//
// Parameters:
//   - value: The string value to be parsed.
//
// Returns:
//   - T: The parsed value of type T.
//   - error: An error if the parsing fails or if the type is unsupported.
func ParseDataType[T any](value string) (T, error) {
	var result T
	switch any(result).(type) {
	case float64:
		v, err := strconv.ParseFloat(value, 64)
		return any(v).(T), err
	case int:
		v, err := strconv.Atoi(value)
		return any(v).(T), err
	case bool:
		v, err := strconv.ParseBool(value)
		return any(v).(T), err
	case time.Time:
		v, err := time.Parse(time.RFC3339, value)
		return any(v).(T), err
	default:
		return result, fmt.Errorf("unsupported type")
	}
}

// CheckNonNegative checks if a given value is non-negative.
// It supports int and float64 types. If the value is negative, it returns an error
// with a message indicating that the specified field cannot be negative.
// If the type is unsupported, it returns an error indicating the unsupported type.
//
// Parameters:
//   - value: The value to be checked. It can be of any type, but only int and float64 are supported.
//   - fieldName: The name of the field being checked. This is used in the error message.
//
// Returns:
//   - error: An error if the value is negative or if the type is unsupported. Otherwise, it returns nil.
func CheckNonNegative[T any](value T, fieldName string) error {
	switch v := any(value).(type) {
	case int:
		if v < 0 {
			return fmt.Errorf("%s cannot be negative", fieldName)
		}
	case float64:
		if v < 0 {
			return fmt.Errorf("%s cannot be negative", fieldName)
		}
	default:
		return fmt.Errorf("unsupported type")
	}
	return nil
}

// CheckEmptyFields is a function to check if there is an empty field in a struct
func CheckEmptyFields(data interface{}) bool {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			return true
		}
	}
	return false
}
