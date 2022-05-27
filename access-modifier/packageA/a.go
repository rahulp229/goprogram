package packageA

import (
	"access-modifier/packageB"
	"fmt"
)

func PA() {
	fmt.Println("My self PA")
	packageB.PB()
	return
}
