package cpu

import (
	"fmt"
)

func initOpcodes(opcode byte) map[byte]func(*CPU) {
	opcodes := make(map[byte]func(*CPU))

	opcodes[0x00] = (*CPU).NOP
	opcodes[0x01] = (*CPU).LD_BC
	opcodes[0x02] = (*CPU).LD_A_TO_BC_addr
	opcodes[0x03] = (*CPU).INC_BC
	opcodes[0x04] = (*CPU).INC_B
	opcodes[0x05] = (*CPU).DEC_B
	opcodes[0x06] = (*CPU).LD_B_D8
	opcodes[0x07] = (*CPU).RLCA
	opcodes[0x08] = (*CPU).LD_A16_SP
	opcodes[0x09] = (*CPU).ADD_HL_BC
	opcodes[0x0A] = (*CPU).LD_BC_INTO_A
	opcodes[0x0B] = (*CPU).DEC_BC
	opcodes[0x0C] = (*CPU).INC_C
	opcodes[0x0D] = (*CPU).DEC_C
	opcodes[0x0E] = (*CPU).LD_C_D8
	opcodes[0x0F] = (*CPU).RRCA
	for i := byte(0x80); i < 0x90; i++ {
		opcodes[i] = func(cpu *CPU) { cpu.ADD(i) }
	}
	for i := byte(0x90); i < 0xa0; i++ {
		opcodes[i] = func(cpu *CPU) { cpu.SUB(i) }
	}
	return opcodes
}

// find instruction not in the opcode table
func FindInstruction(opcode byte) (func(*CPU, byte), error) {
	high := 0xf0 & opcode
	low := 0xf & opcode
	if high >= 0x8 && high < 0x9 {
		return (*CPU).ADD, nil
	}
	if high >= 0xa && low <= 0x7 {
		return (*CPU).AND, nil
	}

	return nil, fmt.Errorf("instruction not found")

}

