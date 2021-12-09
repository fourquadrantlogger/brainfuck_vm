package brainfuck_vm

func Add(a, b []byte) (c []byte) {

	if len(b) > len(a) {
		a, b = b, a
	}
	//todo check a[0]==0
	idxa := len(a) - 1
	idxb := len(b) - 1
	adder := byte(0)
	for idxb >= 0 {
		var x int16 = int16(a[idxa]) + int16(b[idxb]) + int16(adder)
		if x >= (0xff) {
			adder = byte(x / 255)
			a[idxa] = byte(x % 255)
		} else {
			a[idxa] = byte(x)
			adder = 0
		}

		idxa--
		idxb--
	}
	if adder > 0 {
		a = append([]byte{adder}, a...)
	}
	return a
}
