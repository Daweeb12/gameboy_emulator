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
