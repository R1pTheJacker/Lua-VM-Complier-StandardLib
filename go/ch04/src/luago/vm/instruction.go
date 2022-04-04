package vm

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx>>1

type Instruction uint32

//extract opcode from instruction
func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}

//get params from iABC mode
func (self Instruction) ABC() (a, b, c int) {
	a = int(self>>6  &  0xFF)
	b = int(self>>14 & 0x1FF)
	c = int(self>>23 & 0x1FF)
	return
}

//get params from iABx mode
func (self Instruction) ABx() (a, bx int) {
	a  = int(self>>6  &  0xFF)
	bx = int(self>>14)
	return
}

//get params from iAsBx mode
func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

//get params from iAx mode
func (self Instruction) Ax() int {
	return int(self>>6)
}

//return opcode name
func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

//return opmode
func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

//return argBmode
func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

//return argCmode
func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}
