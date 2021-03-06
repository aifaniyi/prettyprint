package prettyprint

import (
	"fmt"
	"reflect"
	"strings"
)

const indt = "  "

// Struct :
func Struct(val interface{}) string {
	return structRec(reflect.ValueOf(val), indt)
}

// structRec :
func structRec(val reflect.Value, indent string) string {
	typ := val.Type()

	switch val.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", val.Interface())

	case reflect.String:
		return fmt.Sprintf("\"%v\"", val.Interface())

	case reflect.Ptr:
		if !val.IsNil() {
			return structRec(val.Elem(), indent+indt)
		}

	case reflect.Struct:
		arr := make([]string, 0)
		arr = append(arr, indent+"{")
		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i)

			// fmt.Printf("%d: %s %s = %v\n", i, typ.Field(i).Name, f.Type(), f.Interface())
			if f.IsValid() {
				arr = append(arr, fmt.Sprintf("%s\"%s\": %s", indent+indt,
					typ.Field(i).Name, structRec(f, indent+indt)))
			}
		}
		arr = append(arr, indent+"}")
		return strings.Join(arr, "\n")
	}

	return ""
}
