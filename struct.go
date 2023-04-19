package utils

import (
	"errors"
	"reflect"
	"strings"
)

// getTagName returns the tag name of a struct field.
// If customTag is not empty, it will return the tag value of the custom tag.
// Otherwise, it will return the field name.
func getTagName(field reflect.StructField, customTag string) string {
	tagName := field.Name

	if customTag != "" {
		if tagValue := field.Tag.Get(customTag); tagValue != "" {
			tagName = strings.Split(tagValue, ",")[0]
		}
	} else {
		tagName = field.Name
	}

	return tagName
}

// MapToStruct maps a map to a struct.
// The map key must be the same as the struct field name or the custom tag name.
// The map value must be the same type as the struct field type.
// The target must be a non-nil pointer to a struct.
func MapToStruct(source map[string]interface{}, target interface{}, customTag string) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return errors.New("target must be a non-nil pointer to a struct")
	}

	targetType := targetValue.Elem().Type()
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldName := getTagName(field, customTag)
		if fieldValue, ok := source[fieldName]; ok {
			targetValue.Elem().Field(i).Set(reflect.ValueOf(fieldValue))
		}
	}

	return nil
}

// StructToMap converts a struct to a map.
// The map key will be the struct field name or the custom tag name.
// The map value will be the struct field value.
// The source must be a struct or a pointer to a struct.
func StructToMap(source interface{}, customTag string) (map[string]interface{}, error) {
	sourceValue := reflect.ValueOf(source)

	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}

	if sourceValue.Kind() != reflect.Struct {
		return nil, errors.New("source must be a struct or a pointer to a struct")
	}

	sourceType := sourceValue.Type()
	result := make(map[string]interface{})
	for i := 0; i < sourceType.NumField(); i++ {
		field := sourceType.Field(i)
		fieldName := getTagName(field, customTag)
		result[fieldName] = sourceValue.Field(i).Interface()
	}

	return result, nil
}
