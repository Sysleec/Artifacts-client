package utils

import (
	"fmt"
	"reflect"
)

// CharacterPrettyPrinter Prints fields of a struct line by line, skipping empty ones.
func CharacterPrettyPrinter(s interface{}) {
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typeOfS.Field(i)

		if !fieldType.IsExported() || field.IsZero() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			fmt.Printf("%s = %q\n", fieldType.Name, field.Interface())
		case reflect.Slice:
			if field.Len() == 0 {
				continue
			} else {
				for j := 0; j < field.Len(); j++ {
					item := field.Index(j)
					if item.FieldByName("Quantity").Interface().(int) != 0 {
						fmt.Printf("%s[%d] = %+v\n", fieldType.Name, j, item.Interface())
						if item.Kind() == reflect.Struct {
							CharacterPrettyPrinter(item.Interface())
						}
					}
				}
			}
		default:
			fmt.Printf("%s = %+v\n", fieldType.Name, field.Interface())
			if field.Kind() == reflect.Struct {
				CharacterPrettyPrinter(field.Interface())
			}
		}
	}
}
