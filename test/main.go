package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p *People) String1() string {
	return fmt.Sprintf("print: %v", p)
}
func main() {
	p := &People{
		Name: "21",
	}
	_ = p.String1()
}
