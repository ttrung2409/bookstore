package utils

import "reflect"

func Nameof(myType any) string {
	return reflect.TypeOf(myType).Elem().Name()
}
