package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type User struct {
	Username string `required:"true" unique:"true"`
	Password string `required:"true" min:"8"`
}

type Required func(any) bool

func ReadField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name:", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)

		isRequired := structField.Tag.Get("required")
		if isRequired == "" {
			fmt.Println(structField.Name, "has no tag `required`")
		} else {
			fmt.Println("required:", isRequired)
		}
	}
}

func RequiredField(value any) (result bool) {
	result = true

	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			if data == "" {
				result = false
			}
		}
	}

	return result
}

func LenMinEight(value any) bool {
	// Ensure the input is a struct
	v := reflect.ValueOf(value)
	fmt.Println(v.Field(0))
	if v.Kind() != reflect.Struct {
		fmt.Println("Error: Input is not a struct")
		return false
	}

	t := v.Type()

	// Iterate through struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get("min")

		// Check if the "min" tag is set to "8"
		if tagValue == "8" {
			fieldValue := v.Field(i)

			// Ensure the field is a string
			if fieldValue.Kind() == reflect.String {
				str := fieldValue.String()

				// Validate length
				if len(str) < 8 {
					return false
				}
			} else {
				fmt.Printf("Warning: Field %s is not a string\n", field.Name)
			}
		}
	}

	return true
}

func main() {
	// ReadField(Sample{"Fatih"})
	// ReadField(User{"Fatih", "password"})

	user := User{
		Username: "",
		Password: "passwor",
	}
	fmt.Println(RequiredField(user))
	fmt.Println(LenMinEight(user))
}
