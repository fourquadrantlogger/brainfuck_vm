package brainfuck_vm

import "io"

type cpu struct {
	addrpoint []byte
}

func (c *cpu) Run(vm *MemSeg, io io.ReadWriter) {
	for {
		opcode := vm.Get(c.addrpoint)

		switch opcode {
		case '>':
			c.addrpoint = Add(c.addrpoint, []byte{1})
		case '<':
			c.addrpoint = Add(c.addrpoint, []byte{-1})
		case '+':
			vm.Set(c.addrpoint, Add([]byte{vm.Get(c.addrpoint)}, []byte{1}))
		case '-':
			vm.Set(c.addrpoint, Add(vm.Get(c.addrpoint), []byte{-1}))
		}

	}
}
func (c *cpu) control(vm *MemSeg, io io.ReadWriter) {

}
func (c *cpu) branch(vm *MemSeg, io io.ReadWriter) {

}
func (c *cpu) compute(vm *MemSeg, io io.ReadWriter) {

}
