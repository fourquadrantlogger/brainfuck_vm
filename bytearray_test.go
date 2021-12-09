package brainfuck_vm

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	a := []byte{1, 1}
	c := Add(a, []byte{1})
	fmt.Println(c)

	a = []byte{255}
	c = Add(a, []byte{255})
	fmt.Println(c)
}
