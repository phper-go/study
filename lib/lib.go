package lib

import (
	"fmt"
)

type T struct {
	Name string
}

func (this *T) Output() {

	fmt.Println("name")
}
