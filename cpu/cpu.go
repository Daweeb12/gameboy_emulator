package cpu

import (
	"fmt"
	"gameboy_emulator/memory"
)

const (
	IORegisterSize = 0x7f
)

type CPU struct {
	A, F        byte
	B, C        byte
	D, E        byte
	H, L        byte
	SP          uint16
	PC          uint16
	IORegisters [IORegisterSize]byte
	Bus         memory.Bus
	Flags       *Flags
}
type Flags struct {
	Z bool
	N bool
	H bool
	C bool
}

func newFlags() *Flags {
	return &Flags{Z: false, N: false, H: false, C: false}
}

var OpcodeTable [256]func(*CPU)

func Init(cpu *CPU) {
	cpu.A = 0x01
	cpu.F = 0xB0
	cpu.B = 0x00
	cpu.C = 0x13
	cpu.D = 0x00
	cpu.E = 0xD8
	cpu.H = 0x01
	cpu.L = 0x4D
	cpu.SP = 0xFFFE
	cpu.PC = 0x0100
	cpu.Flags = newFlags()
}

func (cpu *CPU) checkFlags() (byte, byte, byte, byte) {
	z := cpu.F & (1 << 7)
	n := cpu.F & (1 << 6)
	halfCarry := cpu.F & (1 << 5)
	carry := cpu.F & (1 << 4)
	return z, n, halfCarry, carry
}
func (cpu *CPU) executeOPCode(opcode byte) {

}

func (cpu *CPU) getByteFromMemory(addr uint16) (byte, error) {
	if addr <= 0x7ff {
		return cpu.Bus.ROM[addr], nil
	} else if addr >= 0x8000 && addr <= 0x9ff {
		return cpu.Bus.VRAM[addr-0x8000], nil
	} else if addr >= 0xa000 && addr <= 0xbff {
		//eramu
		return cpu.Bus.ExtRAM[addr-0xa000], nil
	} else if addr >= 0xc000 && addr <= 0xdfff {
		//wram1
		return cpu.Bus.WRAM[addr-0xc000], nil
	} else if addr >= 0xe000 && addr <= 0xfdff {
		//echo ram
		return cpu.Bus.ERAM[addr-0xe000], nil
	} else if addr >= 0xfe00 && addr <= 0xfe9f {
		//object attr memory
		return cpu.Bus.OAM[addr-0xfe00], nil
	} else if addr >= 0xff00 && addr <= 0xff7f {
		//i/o registers
		return cpu.IORegisters[addr-0xff00], nil
	} else if addr >= 0xff80 && addr <= 0xfffe {
		// hram
		return cpu.Bus.HRAM[addr-0xff80], nil
	} else if addr == 0xffff {
		// interupt enable register
		return cpu.Bus.IE, nil
	}
	return 0xff, fmt.Errorf("invalid memory location")
}
func (cpu *CPU) NOP() {
}
