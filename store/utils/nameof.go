package utils

import "reflect"

func Nameof(myType interface{}) string {
	return reflect.TypeOf(myType).Elem().Name()
}
