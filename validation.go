package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(data).Field(i).Interface()
			if data == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	sample := Sample{"Yoga"}

	fmt.Println(IsValid(sample))
}
