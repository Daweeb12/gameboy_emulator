package cpu

import (
	"fmt"
)

func  initOpcodes(opcode byte) map[byte]func(*CPU) {
	return map[byte]func(*CPU){
		0x00: (*CPU).NOP,
		0x01: (*CPU).LD_BC , 
		0x02: (*CPU).LD_A_TO_BC_addr,
		0x03: (*CPU).INC_BC,
		0x04: (*CPU).INC_B, 
		0x05:(*CPU).DEC_B,
		0x06: (*CPU).LD_B_D8,
		0x07:(*CPU).RLCA , 
		0x08: (*CPU).LD_A16_SP,
		0x09:(*CPU).ADD_HL_BC,
		0x0A:(*CPU).LD_BC_INTO_A,
		0x0B:(*CPU).DEC_BC,
		0x0C:(*CPU).INC_C,
		0x0D:(*CPU).DEC_C,
		0x0E:(*CPU).LD_C_D8,
		0x0F:(*CPU).RRCA,
	}
}

