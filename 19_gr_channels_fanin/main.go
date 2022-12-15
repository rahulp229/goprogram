package main

import (
	"errors"
	"fmt"
)

func main() {
	var errorss []error

	errorss = append(errorss, fmt.Errorf("err 1"))
	errorss = append(errorss, fmt.Errorf("err 2"))

	errorss = append(errorss, fmt.Errorf("err 3"))
	errorss = append(errorss, fmt.Errorf("err 4"))
	errorss = append(errorss, fmt.Errorf("err 5"))
	errorss = append(errorss, fmt.Errorf("err 6"))
	e := showerrors(errorss)
	fmt.Println("errr :", e)

}

func showerrors(e []error) error {
	var s string
	errCh := make(chan error)
	for _, val := range e {
		go func(errr <-chan error) {
			errStr := <-errr
			s = s + errStr.Error()
		}(errCh)
		errCh <- val
	}
	close(errCh)
	return errors.New(s)
}
func A() error {
	return errors.New("error A")
}
func B() error {
	return errors.New("error B")
}
