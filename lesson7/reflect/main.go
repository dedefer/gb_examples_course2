package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Passport struct {
	Serial string
	Number string
}

type User struct {
	Passport
	Username string `json:"user"`
	Password string
}

func main() {
	u := &User{
		Username: "Danila",
		Password: "123456789987654321",
	}

	PrintStruct(u)

	iVal := reflect.ValueOf(u)
	iElem := iVal.Elem()

	iElem.Set(reflect.ValueOf(User{
		Username: "12312312",
	}))

	fmt.Printf("i = %+v\n", u)
}

func PrintStruct(s interface{}) error {
	if s == nil {
		return errors.New("arg is nil")
	}

	sVal := reflect.ValueOf(s)

	if sVal.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
	}

	return printStruct(sVal.Type(), "")
}

func printStruct(sType reflect.Type, offset string) error {
	if sType.Kind() != reflect.Struct {
		return fmt.Errorf("arg is not struct; kind=%s", sType.Kind())
	}

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fmt.Printf(offset+"field: %s, type: %s, tag: %s\n", field.Name, field.Type, field.Tag)
		if field.Type.Kind() == reflect.Struct {
			printStruct(field.Type, offset+"  ")
		}
	}

	return nil
}
