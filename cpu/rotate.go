package cpu

import (
	"fmt"
)

// 0xcb00-7
// 0xcb10-7
func (cpu *CPU) RLC(opcode byte) {
	carry := 0
	var prevCarry byte = 0x0
	if cpu.Flags.C {
		prevCarry = 0x1
	}
	switch opcode {
	case 0x00, 0x10:
		cpu.Flags.C = (0x80 & cpu.B) > 0
		cpu.B <<= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x01 {
			cpu.B += byte(carry)
		} else {
			cpu.B += prevCarry
		}
		cpu.Flags.Z = cpu.B == 0
	case 0x01, 0x11:
		cpu.Flags.C = (0x80 & cpu.C) > 0
		cpu.C <<= 1
		if cpu.Flags.C {
			carry = 1
		}

		if opcode < 0x01 {
			cpu.C += byte(carry)
		} else {
			cpu.C += prevCarry
		}
		cpu.Flags.Z = cpu.C == 0
	case 0x02, 0x12:
		cpu.Flags.C = (0x80 & cpu.D) > 0
		cpu.D <<= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x01 {
			cpu.D += byte(carry)
		} else {
			cpu.D += prevCarry
		}
		cpu.Flags.Z = cpu.D == 0
	case 0x03, 0x13:
		cpu.Flags.C = (0x80 & cpu.E) > 0
		cpu.E <<= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x01 {
			cpu.E += byte(carry)
		} else {
			cpu.E += prevCarry
		}
		cpu.Flags.Z = cpu.E == 0
	case 0x04, 0x14:
		cpu.Flags.C = (0x80 & cpu.H) > 0
		cpu.H <<= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x01 {
			cpu.H += byte(carry)
		} else {
			cpu.H += prevCarry
		}
		cpu.Flags.Z = cpu.H == 0
	case 0x05, 0x15:
		cpu.Flags.C = (0x80 & cpu.L) > 0
		cpu.L <<= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x01 {
			cpu.L += byte(carry)
		} else {
			cpu.L += prevCarry
		}
		cpu.Flags.Z = cpu.L == 0
	case 0x06, 0x16:
		break
	case 0x07, 0x17:
		cpu.Flags.C = (0x80 & cpu.A) > 0
		cpu.A <<= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x01 {
			cpu.A += byte(carry)
		} else {
			cpu.A += prevCarry
		}
		cpu.Flags.Z = cpu.A == 0
	default:
		fmt.Println("could not find opcode")
	}
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.PC += 2

}

// 0xcb08-f
func (cpu *CPU) RRC(opcode byte) {
	carry := 0
	var prevCarry byte = 0
	if cpu.Flags.C {
		prevCarry = 1
	}
	switch opcode {
	case 0x08, 0x18:
		cpu.Flags.C = (0x01 & cpu.B) > 0
		cpu.B >>= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x10 {
			cpu.B += byte(carry) << 7
		} else {
			cpu.B += prevCarry
		}
		cpu.Flags.Z = cpu.B == 0
	case 0x09, 0x19:
		cpu.Flags.C = (0x01 & cpu.C) > 0
		cpu.C >>= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x10 {
			cpu.C += byte(carry) << 7
		} else {
			cpu.C += prevCarry
		}
		cpu.Flags.Z = cpu.C == 0
	case 0x0A, 0x1A:
		cpu.Flags.C = (0x01 & cpu.D) > 0
		cpu.D >>= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x10 {
			cpu.E += byte(carry) << 7
		} else {
			cpu.E += prevCarry
		}

		cpu.Flags.Z = cpu.D == 0
	case 0x0B, 0x1B:
		cpu.Flags.C = (0x01 & cpu.E) > 0
		cpu.E >>= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x10 {
			cpu.E += byte(carry) << 7
		} else {
			cpu.E += prevCarry
		}
		cpu.Flags.Z = cpu.E == 0
	case 0x0C, 0x1C:
		cpu.Flags.C = (0x01 & cpu.H) > 0
		cpu.H >>= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x10 {
			cpu.H += byte(carry) << 7
		} else {
			cpu.H += prevCarry
		}
		cpu.Flags.Z = cpu.H == 0
	case 0x0D, 0x1D:
		cpu.Flags.C = (0x01 & cpu.L) > 0
		cpu.L >>= 1
		if cpu.Flags.C {
			carry = 1
		}
		if opcode < 0x10 {
			cpu.L += byte(carry) << 7
		} else {
			cpu.L += prevCarry
		}
		cpu.Flags.Z = cpu.L == 0
	case 0x0E, 0x1E:
		break
	case 0x0F, 0x1F:
		cpu.Flags.C = (0x01 & cpu.A) > 0
		cpu.A >>= 1
		if cpu.Flags.C {
			carry = 1
		}

		if opcode < 0x10 {
			cpu.A += byte(carry) << 7
		} else {
			cpu.A += prevCarry
		}
		cpu.Flags.Z = cpu.A == 0
	default:
		fmt.Println("could not find opcode")
	}
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.PC += 2

}

func (cpu *CPU) SLA(opcode byte) {
	var b byte = 0
	var val byte = 0xff
	switch opcode {
	case 0x20:
		b = (0x80 & cpu.B)
		cpu.B <<= 1
		val = cpu.B
	case 0x21:
		b = (0x80 & cpu.C)
		cpu.C <<= 1
		val = cpu.C
	case 0x22:
		b = (0x80 & cpu.D)
		cpu.D <<= 1
		val = cpu.D
	case 0x23:
		b = (0x80 & cpu.E)
		cpu.E <<= 1
		val = cpu.E
	case 0x24:
		b = (0x80 & cpu.H)
		cpu.H <<= 1
		val = cpu.H
	case 0x25:
		b = (0x80 & cpu.L)
		cpu.L <<= 1
		val = cpu.L
	case 0x26:
		break
	case 0x27:
		b = (0x80 & cpu.A)
		cpu.B <<= 1
		val = cpu.A
	default:
	}
	cpu.Flags.Z = val == 0
	cpu.Flags.C = b > 0
	cpu.Flags.H = false
	cpu.Flags.N = false
	cpu.PC += 2
}

// this is wrong
func (cpu *CPU) SRA(opcode byte) {
	var b byte = 0
	var val byte = 0xff
	switch opcode {
	case 0x28:
		b = (0x01 & cpu.B)
		cpu.B >>= 1
		val = cpu.B + b
	case 0x29:
		b = (0x01 & cpu.C)
		cpu.C >>= 1
		val = cpu.C + b
	case 0x2A:
		b = (0x01 & cpu.D)
		cpu.D >>= 1
		val = cpu.D + b
	case 0x2B:
		b = (0x01 & cpu.E)
		cpu.E >>= 1
		val = cpu.E + b
	case 0x2C:
		b = (0x01 & cpu.H)
		cpu.H >>= 1
		val = cpu.H + b
	case 0x2D:
		b = (0x01 & cpu.L)
		cpu.L >>= 1
		val = cpu.L + b
	case 0x2E:
		break
	case 0x2F:
		b = (0x01 & cpu.A)
		cpu.B >>= 1
		val = cpu.A + b
	default:
	}
	cpu.Flags.Z = val == 0
	cpu.Flags.C = b > 0
	cpu.Flags.H = false
	cpu.Flags.N = false
	cpu.PC += 2
}

func (cpu *CPU) SWAP(opcode byte) {
	var low byte
	var high byte
	var val byte
	switch opcode {
	case 0x30:
		low = cpu.B & 0xf
		high = (cpu.B & 0xf0) >> 4
		val = low<<4 + high
	case 0x31:
		low = cpu.C & 0xf
		high = (cpu.C & 0xf0) >> 4
		val = low<<4 + high
	case 0x32:
		low = cpu.D & 0xf
		high = (cpu.D & 0xf0) >> 4
		val = low<<4 + high
	case 0x33:
		low = cpu.E & 0xf
		high = (cpu.E & 0xf0) >> 4
		val = low<<4 + high
	case 0x34:
		low = cpu.H & 0xf
		high = (cpu.H & 0xf0) >> 4
		val = low<<4 + high
	case 0x35:
		low = cpu.L & 0xf
		high = (cpu.L & 0xf0) >> 4
		val = low<<4 + high
	case 0x36:
		break
	case 0x37:
		low = cpu.A & 0xf
		high = (cpu.A & 0xf0) >> 4
		val = low<<4 + high
	}
	cpu.Flags.Z = val == 0
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.Flags.C = false
	cpu.PC += 2
}

func (cpu *CPU) SRL(opcode byte) {
	var b byte = 0
	var val byte = 0xff
	switch opcode {
	case 0x38:
		b = (0x01 & cpu.B)
		cpu.B >>= 1
		val = cpu.B
	case 0x39:
		b = (0x01 & cpu.C)
		cpu.C >>= 1
		val = cpu.C
	case 0x3A:
		b = (0x01 & cpu.D)
		cpu.D >>= 1
		val = cpu.D
	case 0x3B:
		b = (0x01 & cpu.E)
		cpu.E >>= 1
		val = cpu.E
	case 0x3C:
		b = (0x01 & cpu.H)
		cpu.H >>= 1
		val = cpu.H
	case 0x3D:
		b = (0x01 & cpu.L)
		cpu.L >>= 1
		val = cpu.L
	case 0x3E:
		break
	case 0x3F:
		b = (0x01 & cpu.A)
		cpu.B >>= 1
		val = cpu.A
	default:
	}
	cpu.Flags.Z = val == 0
	cpu.Flags.C = b > 0
	cpu.Flags.H = false
	cpu.Flags.N = false
	cpu.PC += 2
}

// 0x5x
func (cpu *CPU) BIT(opcode byte) {
	bitNumber := -1
	if 0x40 <= opcode && opcode <= 0x47 {
		bitNumber = 0
	} else if 0x48 <= opcode && opcode <= 0x4f {
		bitNumber = 1
	} else if 0x50 <= opcode && opcode <= 0x57 {
		bitNumber = 2
	} else if 0x58 <= opcode && opcode <= 0x5f {
		bitNumber = 3
	} else if 0x60 <= opcode && opcode <= 0x67 {
		bitNumber = 4
	} else if 0x68 <= opcode && opcode <= 0x6f {
		bitNumber = 5
	} else if 0x70 <= opcode && opcode <= 0x77 {
		bitNumber = 6
	} else if 0x78 <= opcode && opcode <= 0x7f {

		bitNumber = 7
	} else {
		fmt.Println("invalid opcode ", opcode)
		return
	}

	switch opcode {
	case 0x50, 0x60, 0x70:

	case 0x41, 0x51, 0x61, 0x71:
	case 0x42, 0x52, 0x62, 0x72:
	case 0x43, 0x53, 0x63, 0x73:
	case 0x44, 0x54, 0x64, 0x74:
	case 0x45, 0x55, 0x65, 0x75:
	case 0x46, 0x56, 0x66, 0x76:
	case 0x47, 0x57, 0x67, 0x77:
	case 0x48, 0x58, 0x68, 0x78:
	case 0x49, 0x59, 0x69, 0x79:
	case 0x4A, 0x5A, 0x6A, 0x7A:
	case 0x4B, 0x5B, 0x6B, 0x7B:
	case 0x4C, 0x5C, 0x6C, 0x7C:
	case 0x4D, 0x5D, 0x6D, 0x7D:
	case 0x4E, 0x5E, 0x6E, 0x7E:
	case 0x5F, 0x6F, 0x7F:
	}
	cpu.Flags.N = false
	cpu.Flags.H = true
	cpu.PC += 2
}
