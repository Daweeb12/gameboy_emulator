package cpu

import (
	"fmt"
)

func (cpu *CPU) LD_C_D8() {
	imm, err := cpu.getByteFromMemory(cpu.PC)
	if err != nil {
		fmt.Println(err)
	}
	cpu.C = imm
}

func (cpu *CPU) LD_BC() {

	low, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	high, err := cpu.getByteFromMemory(cpu.PC + 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC += 3
	cpu.B = high
	cpu.C = low
	cpu.PC += 3
	cpu.Flags.N = false
}

func (cpu *CPU) LD_BC_INTO_A() {
	var addr uint16 = uint16(cpu.B)<<8 + uint16(cpu.C)
	err := cpu.Bus.WriteByteToAddr(addr, cpu.A)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC++
	cpu.Flags.N = false

}
func (cpu *CPU) LD_A_TO_BC_addr() {
	var addr uint16 = uint16(cpu.B)<<8 + uint16(cpu.C)
	err := cpu.Bus.WriteByteToAddr(addr, cpu.A)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC++

}

func (cpu *CPU) LD_A16_SP() {
	low, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	high, err := cpu.getByteFromMemory(cpu.PC + 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	var addr uint16 = uint16(high<<8) + uint16(low)
	err = cpu.Bus.WriteByteToAddr(addr, byte(cpu.PC&0xff))
	if err != nil {
		return
	}
	cpu.PC += 3
}
func (cpu *CPU) LD_B_D8() {
	low, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.B = low
	cpu.PC += 2
}
func (cpu *CPU) LD_DE_D16() {
	low, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	high, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.B = high
	cpu.C = low
	cpu.PC += 3
}

func (cpu *CPU) LD_from_A_into_DE() {
	var addr uint16 = uint16(cpu.D)<<8 + uint16(cpu.E)
	err := cpu.Bus.WriteByteToAddr(addr, cpu.A)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC++

}

// 0x16
func (cpu *CPU) LD_D8_to_D() {
	imm, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.D = imm
	cpu.PC += 2
}

// 0x1A
func (cpu *CPU) LD_fromDE_to_A() {
	addr := uint16(cpu.D)<<8 + uint16(cpu.E)
	imm, err := cpu.getByteFromMemory(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.A = imm
	cpu.PC++
}

// 0x40<=x<0x70 || 0x78<=x<0x80
func (cpu *CPU) LD_register(opcode byte) {
	switch opcode {
	case 0x40:
		cpu.B = cpu.B
	case 0x41:
		cpu.B = cpu.C
	case 0x42:
		cpu.B = cpu.D
	case 0x43:
		cpu.B = cpu.E
	case 0x44:
		cpu.B = cpu.H
	case 0x45:
		cpu.B = cpu.L
	case 0x46:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.B = val
	case 0x47:
		cpu.B = cpu.A
	case 0x48:
		cpu.C = cpu.B
	case 0x49:
		cpu.C = cpu.C
	case 0x4a:
		cpu.C = cpu.D
	case 0x4b:
		cpu.C = cpu.E
	case 0x4c:
		cpu.C = cpu.H
	case 0x4d:
		cpu.C = cpu.L
	case 0x4e:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.C = val
	case 0x4f:
		cpu.C = cpu.A
	case 0x50:
		cpu.D = cpu.B
	case 0x51:
		cpu.D = cpu.C
	case 0x52:
		cpu.D = cpu.D
	case 0x53:
		cpu.D = cpu.E
	case 0x54:
		cpu.D = cpu.H
	case 0x55:
		cpu.D = cpu.L
	case 0x56:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.D = val
	case 0x57:
		cpu.D = cpu.A
	case 0x58:
		cpu.E = cpu.B
	case 0x59:
		cpu.E = cpu.C
	case 0x5a:
		cpu.E = cpu.D
	case 0x5b:
		cpu.E = cpu.E
	case 0x5c:
		cpu.E = cpu.H
	case 0x5d:
		cpu.E = cpu.L
	case 0x5e:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.E = val
	case 0x5f:
		cpu.E = cpu.A
	case 0x60:
		cpu.H = cpu.B
	case 0x61:
		cpu.H = cpu.C
	case 0x62:
		cpu.H = cpu.D
	case 0x63:
		cpu.H = cpu.E
	case 0x64:
		cpu.H = cpu.H
	case 0x65:
		cpu.H = cpu.L
	case 0x66:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.H = val
	case 0x67:
		cpu.H = cpu.A
	case 0x68:
		cpu.L = cpu.B
	case 0x69:
		cpu.L = cpu.C
	case 0x6a:
		cpu.L = cpu.D
	case 0x6b:
		cpu.L = cpu.E
	case 0x6c:
		cpu.L = cpu.H
	case 0x6d:
		cpu.L = cpu.L
	case 0x6e:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.L = val
	case 0x6f:
		cpu.L = cpu.A
	case 0x78:
		cpu.A = cpu.B
	case 0x79:
		cpu.A = cpu.C
	case 0x7a:
		cpu.A = cpu.D
	case 0x7b:
		cpu.A = cpu.E
	case 0x7c:
		cpu.A = cpu.H
	case 0x7d:
		cpu.A = cpu.L
	case 0x7e:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.A = val
	case 0x7f:
		cpu.A = cpu.A
	default:
		fmt.Println("opcode not found")
		return
	}
	cpu.PC++
}

// 0x70 <= x <= 0x77 without halt instruction
func (cpu *CPU) LD_into_HL(opcode byte) {
	addr := uint16(cpu.H)<<8 | uint16(cpu.L)
	var val byte
	switch opcode {
	case 0x70:
		val = cpu.B
	case 0x71:
		val = cpu.C
	case 0x72:
		val = cpu.D
	case 0x73:
		val = cpu.E
	case 0x74:
		val = cpu.H
	case 0x75:
		val = cpu.L
	case 0x77:
		val = cpu.A
	default:
		fmt.Println("could not find instruction")
		return
	}
	err := cpu.Bus.WriteByteToAddr(addr, val)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC++
}
