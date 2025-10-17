package cpu

import (
	"fmt"

	"github.com/u2takey/go-utils/retry"
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

func (cpu *CPU) POP_BC() {
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
	cpu.C = low
	cpu.B = high
	cpu.SP++
	cpu.PC += 3
}
func (cpu *CPU) PUSH_BC() {
	low := cpu.C
	high := cpu.B
	err := cpu.Bus.WriteByteToAddr(cpu.SP+1, low)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.SP++
	err = cpu.Bus.WriteByteToAddr(cpu.SP+1, high)
	if err != nil {
		fmt.Println(err)
		return
	}
	cpu.SP++
	cpu.PC++
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
