package cpu

import (
	"fmt"
)

// 0x18
func (cpu *CPU) JR_s8() {
	offset, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC += uint16(offset)
}

// 0x20
func (cpu *CPU) JR_NZ_s8() {
	if !cpu.Flags.Z {
		offset, err := cpu.getByteFromMemory(cpu.PC + 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.PC += uint16(offset)
		return
	}
	cpu.PC += 2
}

// 0x28
func (cpu *CPU) JR_Z_s8() {
	if cpu.Flags.Z {
		offset, err := cpu.getByteFromMemory(cpu.PC + 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.PC += uint16(offset)
		return
	}
	cpu.PC += 2
}

//0x30

func (cpu *CPU) JR_NC_s8() {
	if !cpu.Flags.C {
		offset, err := cpu.getByteFromMemory(cpu.PC + 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.PC += uint16(offset)
		return
	}
	cpu.PC += 2
}

// 0x38
func (cpu *CPU) JR_C_s8() {
	if cpu.Flags.C {
		offset, err := cpu.getByteFromMemory(cpu.PC + 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		cpu.PC += uint16(offset)
		return
	}
	cpu.PC += 2
}

func (cpu *CPU) JR_NZ_a16() {
	if !cpu.Flags.Z {
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
		val := uint16(high)<<8 | uint16(low)
		cpu.PC = val
		return
	}
	cpu.PC += 3
}

func (cpu *CPU) JR_a16() {
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
	val := uint16(high)<<8 | uint16(low)
	cpu.PC = val
}

func (cpu *CPU) JR_Z_a16() {
	if cpu.Flags.Z {
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
		val := uint16(high)<<8 | uint16(low)
		cpu.PC = val
		return
	}
	cpu.PC += 3
}

func (cpu *CPU) JR_NC_a16() {
	if !cpu.Flags.C {
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
		val := uint16(high)<<8 | uint16(low)
		cpu.PC = val
		return
	}
	cpu.PC += 3
}

func (cpu *CPU) JR_C_a16() {
	if cpu.Flags.C {
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
		val := uint16(high)<<8 | uint16(low)
		cpu.PC = val
		return
	}
	cpu.PC += 3
}

func (cpu *CPU) JR_HL() {
	addr := uint16(cpu.H)<<8 | uint16(cpu.L)
	cpu.PC = addr
}
func (cpu *CPU) RET_NZ() {
	addr := uint16(cpu.PC) | uint16(cpu.PC+1)<<8
	cpu.SP += 2
	cpu.PC = addr
}

func pushAddr(cpu *CPU, addr uint16) error {
	cpu.SP--
	err := cpu.Bus.WriteByteToAddr(cpu.SP, byte(0xff&addr))
	if err != nil {
		fmt.Println(err)
		return err
	}
	cpu.SP--
	err = cpu.Bus.WriteByteToAddr(cpu.SP, byte(addr>>8))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (cpu *CPU) CALL_NZ_a16() {
	if !cpu.Flags.Z {
		err := pushAddr(cpu, cpu.PC)
		if err != nil {
			fmt.Println(err)
			return
		}
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
		addr := uint16(high)<<8 + uint16(low)
		cpu.PC = addr
	}
	cpu.PC += 3
}

func (cpu *CPU) PUSH(opcode byte) {
	var high byte
	var low byte
	switch opcode {
	case 0xc5:
		low = cpu.E
		high = cpu.D
	case 0xd5:
		low = cpu.C
		high = cpu.B
	case 0xe5:
		low = cpu.L
		high = cpu.H
	case 0xf5:
		low = cpu.F
		high = cpu.A
	}
	cpu.SP--
	err := cpu.Bus.WriteByteToAddr(cpu.SP, low)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.SP--
	err = cpu.Bus.WriteByteToAddr(cpu.SP, high)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (cpu *CPU) POP(opcode byte) {
	low, err := cpu.getByteFromMemory(cpu.SP)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.SP++
	high, err := cpu.getByteFromMemory(cpu.SP)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.SP++
	switch opcode {
	case 0xc1:
		cpu.C = low
		cpu.B = high
	case 0xd1:
		cpu.E = low
		cpu.D = high
	case 0xe1:
		cpu.L = low
		cpu.H = high
	case 0xf1:
		cpu.F = low
		cpu.A = high
	default:
		fmt.Println("invalid opcode")
		return
	}
	cpu.PC += 3
}

func makeCall(cpu *CPU) error {
	err := pushAddr(cpu, cpu.PC)
	if err != nil {
		fmt.Println(err)
		return err
	}
	low, err := cpu.getByteFromMemory(cpu.PC + 1)
	if err != nil {
		fmt.Println(err)
		return err
	}
	high, err := cpu.getByteFromMemory(cpu.PC + 2)
	if err != nil {
		fmt.Println(err)
		return err
	}
	addr := uint16(high)<<8 + uint16(low)
	cpu.PC = addr
	return nil
}
func (cpu *CPU) CALL(opcode byte) {
	if opcode == 0xc4 && !cpu.Flags.Z {
		err := makeCall(cpu)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if opcode == 0xcc && cpu.Flags.Z {
		err := makeCall(cpu)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if opcode == 0xd4 && !cpu.Flags.C {
		err := makeCall(cpu)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if opcode == 0xdc && cpu.Flags.C {
		err := makeCall(cpu)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println(fmt.Errorf("could not find opcode"))
		return
	}
	cpu.PC += 3
}

func (cpu *CPU) RST(addr uint16) {
	high := byte((cpu.PC >> 8) & 0xff)
	low := byte(cpu.PC & 0xff)
	cpu.SP--
	err := cpu.Bus.WriteByteToAddr(cpu.SP, low)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.SP--
	err = cpu.Bus.WriteByteToAddr(cpu.SP, high)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.PC = addr
}

func ret(cpu *CPU) error {
	low, err := cpu.getByteFromMemory(cpu.SP)
	if err != nil {
		return err
	}
	cpu.SP++
	high, err := cpu.getByteFromMemory(cpu.SP)
	if err != nil {
		return err
	}
	cpu.SP++
	cpu.PC = uint16(high)<<8 | uint16(low)
	return nil
}
func (cpu *CPU) RET(opcode uint16) {
	switch opcode {
	case 0xc0:
		if !cpu.Flags.Z {
			err := ret(cpu)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
	case 0xc8:
		if cpu.Flags.Z {
			err := ret(cpu)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
	case 0xc9:
		err := ret(cpu)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	case 0xd0:
		if cpu.Flags.C {
			err := ret(cpu)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
	case 0xd8:
		if !cpu.Flags.C {
			err := ret(cpu)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
	case 0xd9:
		if !cpu.Flags.Z {
			err := ret(cpu)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
	default:
		err := fmt.Errorf("could not find opcode")
		fmt.Println(err)
		return
	}
}
