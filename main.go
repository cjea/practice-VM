package main

import (
	"bradfield-vm/cpu"
)

func main() {
	discs := getDiscs()
	processor := cpu.New()
	for _, disc := range discs {
		processor.Process(disc)
	}
	discs[len(discs)-1].Flush()
}

func getDiscs() []*cpu.Memory {
	init := &cpu.Memory{
		0x01, 0x01, 0x10, // LOAD R1 0x10
		0x01, 0x02, 0x12, // LOAD R2 0x12

		0x03, 0x01, 0x02, // R1 = R1 + R2
		0x03, 0x02, 0x01, // R2 = R2 + R1

		0xff, // HALT
		0x00, // END OF INSTRUCTIONS
		0x00, // OUTPUT
		0x00, // OUTPUT
		0x01, // INPUT 1 (LITTLE)
		0x00, // INPUT 1 (BIG)
		0x02, // INPUT 2 (LITTLE)
		0x00, // INPUT 2 (BIG)
	}
	add := &cpu.Memory{
		0x03, 0x01, 0x02, // R1 = R1 + R2
		0x03, 0x02, 0x01, // R2 = R2 + R1
		0x03, 0x01, 0x02, // R1 = R1 + R2
		0x03, 0x02, 0x01, // R2 = R2 + R1

		0xff, // HALT
	}
	flusher := &cpu.Memory{
		0x02, 0x01, 0x0e, // STORE R1 0x0e

		0xff, // HALT
	}
	return []*cpu.Memory{init, add, flusher}
}
