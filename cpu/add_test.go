package cpu

import (
	"testing"
)

func TestInit(b *testing.T) {
	var cpu CPU
	Init(&cpu)
	cpu.printRegisterState()
	cpu.printFlags()
}
func TestAdd(t *testing.T) {
	var cpu CPU
	Init(&cpu)

}
