package cpu

import (
	"fmt"
)

// 0x0C
func (cpu *CPU) INC_C() {
	cpu.Flags.Z = ((cpu.C + 1) == 0)
	cpu.Flags.H = ((0x0f & cpu.C) + 1) > 0x0f
	cpu.C++
	cpu.PC += 1
	cpu.Flags.N = false
}

// 0x0D
func (cpu *CPU) DEC_C() {
	cpu.Flags.N = true
	cpu.Flags.H = (cpu.C & 0x0F) == 0x00
	cpu.C--
	cpu.Flags.Z = (cpu.C == 0x0)
}

// 0x03
func (cpu *CPU) INC_BC() {
	var val int = (int(cpu.B)<<8+int(cpu.C))&0xffff + 1
	cpu.B = byte(val >> 4)
	cpu.C = byte(val & 0xff)
	cpu.PC += 1
	cpu.Flags.N = false
}

// 0x04
func (cpu *CPU) INC_B() {
	var val int = int(cpu.B + 1)
	cpu.Flags.H = ((0x0f & cpu.B) + 1) > 0x0f
	cpu.Flags.Z = (val == 0)
	cpu.PC += 1
	cpu.B = byte(0xff & val)
	cpu.Flags.N = false

}

// 0x0B
func (cpu *CPU) DEC_BC() {
	var bc int = int(cpu.B<<8) + int(cpu.C) - 1
	cpu.B = byte(bc >> 8)
	cpu.C = byte(bc & 0xff)
	cpu.Flags.H = (0x0f & cpu.C) == 0x0
	cpu.Flags.N = true
	cpu.Flags.Z = (bc == 0)
	cpu.PC++
}

// 0x09
func (cpu *CPU) ADD_HL_BC() {
	var hl int = int(cpu.H<<8) + int(cpu.L)
	var bc int = int(cpu.B<<8) + int(cpu.C)
	val := bc + hl
	cpu.Flags.H = (0x0f & val) > 0x0f
	cpu.Flags.C = (0xff & val) > 0xff
	high := byte(val >> 8)
	low := byte(val & 0xff)
	cpu.B = high
	cpu.C = low
	cpu.Flags.N = false
	cpu.PC += 1
}

// 0x8x
func (cpu *CPU) ADD(opcode byte) {
	var mask byte = 0x0f
	var halfCarry byte = 0
	if cpu.Flags.H {
		halfCarry = 1
	}
	low := mask & opcode
	var val int
	switch low {
	case 0:
		val = int(cpu.A) + int(cpu.B)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.B & 0xF)) > 0xF
	case 1:
		val = int(cpu.A) + int(cpu.C)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.C & 0xF)) > 0xF
	case 2:
		val = int(cpu.A) + int(cpu.D)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.D & 0xF)) > 0xF
	case 3:
		val = int(cpu.A) + int(cpu.E)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.E & 0xF)) > 0xF
	case 4:
		val = int(cpu.A) + int(cpu.H)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.H & 0xF)) > 0xF
	case 5:
		val = int(cpu.A) + int(cpu.L)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.L & 0xF)) > 0xF
	case 6:
		var addr uint16 = uint16(cpu.H)<<8 + uint16(cpu.L)
		b, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		val = int(cpu.A) + int(b)
		cpu.Flags.H = ((cpu.A & 0xF) + (b & 0xF)) > 0xF
	case 7:
		val = int(cpu.A) + int(cpu.A)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.A & 0xF)) > 0xF
	case 8:
		val = int(cpu.A) + int(cpu.B) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.B & 0xF) + (halfCarry & 0xF)) > 0xF
	case 9:
		val = int(cpu.A) + int(cpu.C) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.C & 0xF) + (halfCarry & 0xF)) > 0xF
	case 0xA:
		val = int(cpu.A) + int(cpu.D) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.D & 0xF) + (halfCarry & 0xF)) > 0xF
	case 0xB:
		val = int(cpu.A) + int(cpu.E) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.E & 0xF) + (halfCarry & 0xF)) > 0xF
	case 0xC:
		val = int(cpu.A) + int(cpu.H) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.H & 0xF) + (halfCarry & 0xF)) > 0xF
	case 0xD:
		val = int(cpu.A) + int(cpu.H) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.H & 0xF) + (halfCarry & 0xF)) > 0xF
	case 0xE:
		val = int(cpu.A) + int(cpu.H) + int(halfCarry)
		cpu.Flags.H = ((cpu.A & 0xF) + (cpu.H & 0xF) + (halfCarry & 0xF)) > 0xF
	default:
		fmt.Errorf("somehow we ended up here")
	}
	cpu.Flags.Z = (val == 0)
	cpu.Flags.C = val > 0xff
	cpu.A = byte(val & 0xff)
	cpu.PC++
	cpu.Flags.N = false
}

// 0x05
func (cpu *CPU) DEC_B() {
	cpu.Flags.N = true
	cpu.Flags.H = (cpu.B & 0x0F) == 0x00
	cpu.B--
	cpu.Flags.Z = (cpu.B == 0x0)
	cpu.PC++

}

// 0x07
func (cpu *CPU) RLCA() {
	cpu.Flags.Z = false
	cpu.Flags.H = false
	cpu.Flags.N = false
	cpu.Flags.C = 0x80&cpu.A > 0
	c := 0
	if cpu.Flags.C {
		c = 1
	}
	cpu.A = (cpu.A << 1) + byte(c)
	cpu.PC += 1
}

// 0x13
func (cpu *CPU) INC_DE() {
	val := uint16(cpu.D)<<8 + uint16(cpu.E)
	cpu.D = byte(val >> 8)
	cpu.E = byte(0xff & val)
	cpu.PC++
}

// 0x14
func (cpu *CPU) INC_D() {
	cpu.Flags.H = ((0xf & cpu.D) + 1) > 0xf
	cpu.Flags.N = false
	cpu.D++
	cpu.Flags.Z = (cpu.D == 0)
	cpu.PC++
}

// 0x15
func (cpu *CPU) DEC_D() {
	cpu.Flags.N = true
	cpu.Flags.H = (0xf & cpu.D) == 0x0
	cpu.D--
	cpu.Flags.Z = cpu.D == 0
	cpu.PC++
}

// 0x17
func (cpu *CPU) RLA() {
	var carry byte = 0x0
	if cpu.Flags.C {
		carry = 1
	}
	cpu.A = byte(0xff&uint16(cpu.A<<1)) + carry
	cpu.Flags.Z = false
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.Flags.C = (0x80 & cpu.A) > 0
	cpu.PC++
}

// 0x19
func (cpu *CPU) ADD_HL_DE() {
	hl := int(cpu.H)<<8 | int(cpu.L)
	de := int(cpu.D)<<8 | int(cpu.E)
	cpu.Flags.N = false
	cpu.Flags.H = int(hl&0xfff)+int(de&0xffff) > 0xf
	res := hl + de
	cpu.Flags.C = res > 0xffff
	cpu.H = byte(res >> 8)
	cpu.L = byte(res & 0xff)
	cpu.PC++
}

// 0x1B
func (cpu *CPU) DEC_DE() {
	val := uint16(cpu.D)<<8 + uint16(cpu.E) - 1
	cpu.D = byte(val >> 8)
	cpu.E = byte(0xff & val)
}

// 0x1C
func (cpu *CPU) INC_E() {
	cpu.Flags.H = (cpu.E&0xf)+1 > 0xf
	cpu.Flags.Z = (cpu.E + 1) == 0
	cpu.Flags.N = false
	cpu.E++
	cpu.PC++
}

// 0x1D
func (cpu *CPU) DEC_E() {
	cpu.Flags.H = (0xf & cpu.E) == 0
	cpu.Flags.Z = (cpu.E - 1) == 0
	cpu.Flags.N = true
	cpu.C--
	cpu.PC++
}

// 0x0f
func (cpu *CPU) RRCA() {
	cpu.A = cpu.A>>1 + (cpu.A&0x1)<<7
	cpu.Flags.C = (0x1 & cpu.A) == 1
	cpu.Flags.Z = false
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.PC++
}

func (cpu *CPU) RRA() {
	var carry byte = 0
	if cpu.Flags.C {
		carry = 1
	}
	cpu.A = cpu.A>>1 + carry
	cpu.Flags.C = (0x1 & cpu.A) == 1
	cpu.Flags.Z = false
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.PC++
}
