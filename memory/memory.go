package memory

import (
	"fmt"
	"os"
	"sync"
)

const (
	ROMSize     = 0x8000
	VRAMSize    = 0x2000
	ERAMSize    = 0x2000
	WRAMSize    = 0x2000
	EchoRAMSize = 0x1dff
	OAMSize     = 0x9f
	HRAMSize    = 0x7e
)

type Bus struct {
	ROM         []byte
	VRAM        [VRAMSize]byte
	ERAM        [ERAMSize]byte
	ExtRAM      [ERAMSize]byte
	WRAM        [WRAMSize]byte
	OAM         [OAMSize]byte
	HRAM        [HRAMSize]byte
	IE          byte
	JoypadInput byte
	Mux         sync.Mutex
}

func NewBus() *Bus {
	return &Bus{}
}

func (bus *Bus) WriteByteToAddr(addr uint16, b byte) error {
	if addr == 0xff00 {
		bus.JoypadInput = b
		return nil
	}
	if addr <= 0x7ff {
		bus.ROM[addr] = b
		return nil
	} else if addr >= 0x8000 && addr <= 0x9ff {

		bus.VRAM[addr-0x8000] = b
		return nil
	} else if addr >= 0xa000 && addr <= 0xbff {
		//eramu

		bus.ExtRAM[addr-0xa000] = b
		return nil
	} else if addr >= 0xc000 && addr <= 0xdfff {
		bus.WRAM[addr-0xc000] = b
		//wram1
		return nil
	} else if addr >= 0xe000 && addr <= 0xfdff {
		//echo ram
		bus.ERAM[addr-0xe000] = b
		return nil
	} else if addr >= 0xfe00 && addr <= 0xfe9f {
		bus.OAM[addr-0xfe00] = b
		//object attr memory
		return nil
	} else if addr >= 0xff80 && addr <= 0xfffe {
		// hram
		bus.HRAM[addr-0xff80] = b
		return nil
	} else if addr == 0xffff {
		bus.IE = b
		// interupt enable register
		return nil
	}
	return fmt.Errorf("invalid memory location")
}

func (bus *Bus) LoadProgramIntoMemory(filename string) error {
	rom, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	bus.ROM = rom
	return nil
}
