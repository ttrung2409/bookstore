package utils

import (
	"reflect"
)

func FieldsOfObject(obj interface{}) []string {
	value := reflect.ValueOf(obj)  
	typeOfObj := value.Type()
	
	fields := make([]string, value.NumField())

	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).IsZero() == false {
			fields = append(fields, typeOfObj.Field(i).Name)			
		}
	}

	return fields
}

func FieldsOfType(obj interface{}) []string {
	typeOfObj := reflect.TypeOf(obj).Elem()

	fields := make([]string, typeOfObj.NumField())

	for i := 0; i < typeOfObj.NumField(); i++ {
		fields = append(fields, typeOfObj.Field(i).Name)
	}

	return fields	
}
