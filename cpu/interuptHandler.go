package cpu

func (cpu *CPU) PendingInterrupts() byte {
	return cpu.IE & cpu.IF
}

func (cpu *CPU) InterruptHandler() {
	pending := cpu.PendingInterrupts()
	if pending == 0 {
		return
	}

	checkVBlank := (pending & 1) > 0
	if checkVBlank {
	}

	checkLCD := (pending & 2) > 0
	if checkLCD {
	}

	checkTimer := (pending & 1 << 2) > 0
	if checkTimer {

	}
	checkSerial := (pending & 1 << 3) > 0
	if checkSerial {
	}
	checkJoypad := (pending & 1 << 4) > 0
	if checkJoypad {
	}
}

func (cpu *CPU) serialHandler() {
}
func (cpu *CPU) timerHandler() {
}

func (cpu *CPU) vBlankHandler() {
}

func (cpu *CPU) lcdHandler() {
}

func (cpu *CPU) joypadHandler() {
}

func (cpu *CPU) requestInterrupt(bit byte) {
	cpu.Mux.Lock()
	defer cpu.Mux.Unlock()
	cpu.IF |= 1 << bit
}
