package cpu

import ()

func (cpu *CPU) EI() {
	cpu.imeDelay = 2
}

func (cpu *CPU) DI() {
	cpu.IME = false
}

func (cpu *CPU) HALT() {
	if cpu.IME {

	}
}
