package players

import "fmt"

type Cricketer interface {
}
type cricketer struct {
	Name, Age                string
	Runs, Centuries, Fifties int32
}

func NewCrickter() Cricketer {
	return cricketer{
		Name:      "Sachin Tendulkar",
		Age:       "40",
		Runs:      18000,
		Centuries: 51,
		Fifties:   95,
	}
}

func (c cricketer) get() {
	fmt.Println("crickter : ", c)
}
