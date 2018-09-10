package cpu

import (
	"fmt"
	"log"
)

// Memory is the 20 byte representation of RAM
// 14 bytes for instruction, 2 bytes for output, 2x2 bytes for input
type Memory [20]uint8

// Address will be an index in the Memory array
type Address uint8

// Register is used for both program counter and general purpose registers
type Register struct{ val uint16 }

type operation interface{}

// ISA maps from byte -> operation
type ISA map[uint8]operation

// Processor is the main type of this package
type Processor struct {
	registers map[uint8]*Register
	isa       ISA
	instrEnd  int
}

// New returns a new prcocessor
func New() *Processor {
	return &Processor{
		isa: ISA{
			0x01: LoadWord,
			0x02: StoreWord,
			0x03: Add,
			0x04: Sub,
			0x05: PrintStdOut,
			0xff: "HALT",
		},
		instrEnd: 0x0d,
		registers: map[uint8]*Register{
			0x00: &Register{val: 0}, // Program counter
			0x01: &Register{val: 0}, // General purpose
			0x02: &Register{val: 0}, // General purpose
		},
	}
}

// Process performs the fetch-decode-execute loop on a 20 byte "RAM" array
func (p *Processor) Process(memory *Memory) {
	currentInstr := p.registers[0x00].val
	complete := false
	for int(currentInstr) <= p.instrEnd && !complete {
		jump := uint16(3) // 1 byte for opcode, and 1 byte for each "param"
		instruction := memory[currentInstr]
		_, ok := p.isa[instruction]
		if !ok {
			log.Fatal("Invalid instruction:", instruction)
		}
		switch instruction {
		case 0xff: // HALT
			complete = true
		case 0x01:
			registerID := memory[currentInstr+1]
			address := Address(memory[currentInstr+2])
			LoadWord(p.registers[registerID], address, memory)
		case 0x02:
			registerID := memory[currentInstr+1]
			address := Address(memory[currentInstr+2])
			StoreWord(p.registers[registerID], address, memory)
		case 0x03:
			r1 := p.registers[memory[currentInstr+1]]
			r2 := p.registers[memory[currentInstr+2]]
			Add(r1, r2)
		case 0x04:
			r1 := p.registers[memory[currentInstr+1]]
			r2 := p.registers[memory[currentInstr+2]]
			Sub(r1, r2)
		case 0x05:
			PrintStdOut(*memory)
			jump = 1
		}
		p.registers[0x00].val += jump
		currentInstr = p.registers[0x00].val
	}
	p.registers[0x00].val = 0
}

// LoadWord loads value at given address from memory into register
func LoadWord(r *Register, addr Address, mem *Memory) {
	r.val = uint16(LittleEndianDecode(mem, addr))
}

// StoreWord stores the value from register to the given address
func StoreWord(r *Register, addr Address, mem *Memory) {
	littleEnd, bigEnd := LittleEndianEncode(int(r.val))
	mem[addr] = uint8(littleEnd)
	mem[addr+1] = uint8(bigEnd)
}

// Add sets r1 = r1 + r2
func Add(r1, r2 *Register) {
	r1.val = r1.val + r2.val
}

// Sub sets r1 = r1 - r2
func Sub(r1, r2 *Register) {
	r1.val = r1.val - r2.val
}

// PrintStdOut prints the output from memory (addr 0x0e and 0x0f)
func PrintStdOut(m Memory) {
	fmt.Println("Out: ", LittleEndianDecode(&m, Address(0x0e)))
}

// LittleEndianEncode returns 2 bytes in little endian
func LittleEndianEncode(n int) (int, int) {
	bigEnd := n / 256
	littleEnd := n % 256
	return littleEnd, bigEnd
}

// LittleEndianDecode returns 2 bytes in little endian
func LittleEndianDecode(mem *Memory, addr Address) (ret int) {
	littleEnd := uint16(mem[addr])
	bigEnd := uint16(mem[addr+1]) * 256
	return int(littleEnd + bigEnd)
}
