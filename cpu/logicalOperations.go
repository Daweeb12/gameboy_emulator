package cpu

import (
	"fmt"
)

func (cpu *CPU) AND(opcode byte) {
	var low byte = 0xf & opcode
	switch low {
	case 0x0:
		//AND B
		cpu.A &= cpu.B
	case 0x1:
		// AND C
		cpu.A &= cpu.C
	case 0x2:
		// AND D
		cpu.A &= cpu.D
	case 0x3:
		// AND E
		cpu.A &= cpu.E
	case 0x4:
		//AND H
		cpu.A &= cpu.H
	case 0x5:
		// AND L
		cpu.A &= cpu.L
	case 0x6:
		// AND (HL) , A
		var addr uint16 = uint16(cpu.H)<<8 + uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.A &= val

	case 0x7:
		cpu.A &= cpu.A
	}
	cpu.Flags.Z = (cpu.A == 0)
	cpu.Flags.N = false
	cpu.Flags.H = true
	cpu.Flags.C = false
	cpu.PC += 1
}

func (cpu *CPU) XOR(opcode byte) {
	low := 0xf & opcode
	switch low {
	case 0x8:
		//xor B
		cpu.A ^= cpu.B
	case 0x9:
		// xor C
		cpu.A ^= cpu.C
	case 0xA:
		// XOR d
		cpu.A ^= cpu.D
	case 0xB:
		// XOR E
		cpu.A ^= cpu.E
	case 0xC:
		// xor H
		cpu.A ^= cpu.H
	case 0xD:
		// XOR L
		cpu.A ^= cpu.L
	}
	cpu.Flags.Z = (cpu.A == 0)
	cpu.Flags.N = false
	cpu.Flags.H = false
	cpu.Flags.C = false
	cpu.PC += 1
}

// 0xb0-7
func (cpu *CPU) OR(opcode byte) {
	low := 0xf & opcode
	switch low {
	case 0x0:
		cpu.A |= cpu.B
	case 0x1:
		cpu.A |= cpu.C
	case 0x2:
		cpu.A |= cpu.D
	case 0x3:
		cpu.A |= cpu.E
	case 0x4:
		cpu.A |= cpu.H
	case 0x5:
		cpu.A |= cpu.L
	case 0x6:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.A |= val
	case 0x7:
		cpu.A |= cpu.A
	default:
		fmt.Println("unknown opcode")
	}
	cpu.Flags.Z = (cpu.A == 0)
	cpu.Flags.H = false
	cpu.Flags.N = false
	cpu.Flags.C = false
	cpu.PC++
}

// 0xb8-f
func (cpu *CPU) CP(opcode byte) {
	low := opcode & 0xf
	switch low {
	case 0x8:
		cpu.Flags.Z = (cpu.A - cpu.B) == 0
		cpu.Flags.N = true
		cpu.Flags.H = (0xf & cpu.A) < (0xf & cpu.B)
		cpu.Flags.C = (cpu.A < cpu.B)
	case 0x9:
		cpu.Flags.Z = (cpu.A - cpu.C) == 0
		cpu.Flags.N = true
		cpu.Flags.H = (0xf & cpu.A) < (0xf & cpu.C)
		cpu.Flags.C = (cpu.A < cpu.C)
	case 0xA:
		cpu.Flags.Z = (cpu.A - cpu.D) == 0
		cpu.Flags.N = true
		cpu.Flags.H = (0xf & cpu.A) < (0xf & cpu.D)
		cpu.Flags.C = (cpu.A < cpu.D)
	case 0xb:
		cpu.Flags.Z = (cpu.A - cpu.E) == 0
		cpu.Flags.N = true
		cpu.Flags.H = (0xf & cpu.A) < (0xf & cpu.E)
		cpu.Flags.C = (cpu.A < cpu.E)
	case 0xc:
		cpu.Flags.Z = (cpu.A - cpu.H) == 0
		cpu.Flags.N = true
		cpu.Flags.H = (0xf & cpu.A) < (0xf & cpu.H)
		cpu.Flags.C = (cpu.A < cpu.H)
	case 0xd:
		cpu.Flags.Z = (cpu.A - cpu.L) == 0
		cpu.Flags.N = true
		cpu.Flags.H = (0xf & cpu.A) < (0xf & cpu.L)
		cpu.Flags.C = (cpu.A < cpu.L)
	case 0xe:
		addr := uint16(cpu.H)<<8 | uint16(cpu.L)
		val, err := cpu.getByteFromMemory(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.Flags.Z = (cpu.A - val) == 0x0
		cpu.Flags.H = (0xf & cpu.A) < (val & 0xf)
		cpu.Flags.N = true
		cpu.Flags.C = (cpu.A < val)
	case 0xf:
		cpu.Flags.Z = true
		cpu.Flags.N = true
		cpu.Flags.H = false
		cpu.Flags.C = false
	default:
		fmt.Println("no such opcode")
		return
	}

}
