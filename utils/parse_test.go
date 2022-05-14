package utils

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	cap := "256MB"
	parse, err := Parse(cap)
	fmt.Println("parse, err: ", parse, err)

	cap = "1mb"
	parse, err = Parse(cap)
	fmt.Println("parse, err: ", parse, err)

	cap = "1kb"
	parse, err = Parse(cap)
	fmt.Println("parse, err: ", parse, err)

	cap = "kb"
	parse, err = Parse(cap)
	fmt.Println("parse, err: ", parse, err)
}
