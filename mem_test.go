package brainfuck_vm

import (
	"fmt"
	"testing"
)

func TestMemSeg_Set(t *testing.T) {
	var mem1d MemSeg
	mem1d.Set([]byte{0}, 1)
	mem1d.Set([]byte{0, 1}, 1)
	mem1d.Set([]byte{0, 1, 1}, 255)

	fmt.Println(mem1d.Get([]byte{0}))
	fmt.Println(mem1d.Get([]byte{0, 1}))
	fmt.Println(mem1d.Get([]byte{0, 1, 1}))
}
