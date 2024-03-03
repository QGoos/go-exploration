package reflection

import (
	"reflect"
)

// Walk function for learning reflection
func Walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {

	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		for i := range val.NumField() {
			walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}

	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}

	}
}

func getValue(x interface{}) reflect.Value {

	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer { // can check for pointer to derefence?
		val = val.Elem()
	}

	return val
}
