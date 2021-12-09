package brainfuck_vm

func (mem *MemSeg) Get(addr []byte) (v byte) {
	if len(addr) > 0 {
		firstbyte := addr[0]
		if mem.childs == nil {
			return
		} else {
			return mem.childs[firstbyte].Get(addr[1:])
		}

	} else {
		return mem.value
	}
}
func (mem *MemSeg) Set(addr []byte, value byte) {
	if len(addr) == 0 {
		mem.value = value
	} else {
		if mem.childs == nil {
			mem.childs = &[256]MemSeg{}
		}
		mem.childs[addr[0]].Set(addr[1:], value)
	}
}

type MemSeg struct {
	value  byte
	childs *[256]MemSeg
}
