package utils

import (
	"reflect"
)

func Nameof(myType any) string {
	return reflect.TypeOf(myType).Elem().Name()
}

func FieldsOfType(myType any) []string {
	metadata := reflect.TypeOf(myType).Elem()
	fields := make([]string, metadata.NumField())

	for i := 0; i < metadata.NumField(); i++ {
		fields = append(fields, metadata.Field(i).Name)
	}

	return fields
}
