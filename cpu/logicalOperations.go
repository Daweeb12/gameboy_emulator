package cpu

import ()

func (cpu *CPU) AND(opcode byte) error {
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
			return err
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
	return nil
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
