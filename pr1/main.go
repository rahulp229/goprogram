package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// MyStruct ..
type MyStruct struct {
	String string `validate:"is-awesome"`
	A      string `validate:"answer"`
	B      int
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterValidation("is-awesome", ValidateMyVal)
	validate.RegisterStructValidation(Validate1, MyStruct{})
	s := MyStruct{String: "awesome", A: "gt", B: 3}
	err := validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.String = "awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}

func Validate1(fl validator.StructLevel) {
	val := fmt.Sprintln(fl.Current().FieldByName("a"))
	fmt.Println(">>>>", val)
}
