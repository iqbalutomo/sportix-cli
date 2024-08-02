package utils

import (
	"fmt"
	"reflect"
	"sportix-cli/constants"
)

// ConvertStructSliceToStringSlice converts a slice of any struct type to a slice of strings
// based on the specified field name.
//
// Parameters:
//   - slice: The slice of structs to be converted.
//   - fieldName: The name of the field in the struct to be used for conversion.
//
// Returns:
//   - []string: A slice of strings containing the values of the specified field from each struct.
//   - error: An error if the field is not found, is not a string, or if the input is not a struct slice.
func ConvertStructSliceToStringSlice[T any](slice []T, fieldName string) ([]string, error) {
	result := make([]string, len(slice))

	for i, item := range slice {
		v := reflect.ValueOf(item)
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected struct, got %v", v.Kind())
		}

		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("field %s not found in struct", fieldName)
		}

		if field.Kind() != reflect.String {
			return nil, fmt.Errorf("field %s is not a string", fieldName)
		}

		result[i] = field.String()
	}

	return result, nil
}

// IsYes checks if the input string matches the "Yes" option defined in constants.YesNoOptions.
//
// Parameters:
//   - input: The string to be checked.
//
// Returns:
//   - bool: true if the input matches the "Yes" option, otherwise false.
func IsYes(input string) bool {
	return input == constants.YesNoOptions[0]
}
