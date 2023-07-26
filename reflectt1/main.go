package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

const tagName = "base64"

type User struct {
	Id    int
	Name  string
	Email string `base64:"true"`
	Info  Info
}
type Info struct {
	Age    int
	Hobby  string
	Secret string `base64:"true"`
}

func SetField(source interface{}, fieldName string, fieldValue string) {
	v := reflect.ValueOf(source).Elem()

	fmt.Println(v.FieldByName(fieldName).CanSet())

	if v.FieldByName(fieldName).CanSet() {
		v.FieldByName(fieldName).SetString(fieldValue)
	}
}

func translate(i interface{}) error {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	// Get the type and kind of our user variable
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v (%v), tag:%v\n", i+1, field.Name, field.Type.Name(), tag)

		if field.Type.Name() == "string" && tag == "true" {
			structField := v.Field(i)
			original := structField.String()
			fmt.Printf("	value: %s \n", original)

			decoded, err := base64.StdEncoding.DecodeString(original)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("	decoded: %s \n", decoded)
			fmt.Printf("	CanSet: %v \n", v.FieldByName(field.Name).CanSet())
			// v.FieldByName(field.Name).SetString(string(decoded))

			// 	if v.FieldByName(field.Name).CanSet() {
			// 		v.FieldByName(field.Name).SetString(string(decoded))
			// 	}

			//SetField(&i, field.Name, string(decoded))
		}

	}
	return nil
}

func main() {
	info := Info{
		Age:    40,
		Hobby:  "Art",
		Secret: "VG9wU2VjcmV0",
	}
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "am9obkBleGFtcGxl",
		Info:  info,
	}
	fmt.Printf("Init value: %v\n", user)
	err := translate(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("After value: %v\n", user)
}
