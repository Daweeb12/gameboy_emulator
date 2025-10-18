package cpu

import (
	"fmt"
)

func (cpu *CPU) PrintRegisterState() {
	fmt.Println("Register contents:")
	fmt.Printf("A: %x\t", cpu.A)
	fmt.Printf("F: %x\n", cpu.F)
	fmt.Printf("B: %x\t", cpu.B)
	fmt.Printf("C: %x\n", cpu.C)
	fmt.Printf("D: %x\t", cpu.D)
	fmt.Printf("E: %x\n", cpu.E)
	fmt.Printf("H: %x\t", cpu.H)
	fmt.Printf("L: %x\n", cpu.L)
	fmt.Printf("SP %x\n", cpu.SP)
	fmt.Printf("PC: %x\n\n", cpu.PC)
}
func (cpu *CPU) PrintFlags() {
	fmt.Printf("zero flag: %b\n", cpu.Flags.Z)
	fmt.Printf("subtraction flag: %b\n", cpu.Flags.N)
	fmt.Printf("half carry flag: %b\n", cpu.Flags.H)
	fmt.Printf("carry flag: %b\n", cpu.Flags.C)
}

func (cpu *CPU) debugMode() {
	cpu.PrintRegisterState()
	cpu.PrintFlags()
}
