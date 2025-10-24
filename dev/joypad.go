package main

import (
	"fmt"
	"gameboy_emulator/cpu"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Joypad struct {
	Right, Left, Up, Down bool
	A, B, Select, Start   bool
	SelectBits            byte
	Mux                   sync.Mutex
}

func (j *Joypad) KeyDown(key string, c *cpu.CPU) {
	switch key {
	case "z":
		j.A = true
		fmt.Println("Button pressed: A")
	case "x":
		j.B = true
		fmt.Println("Button pressed: B")
	case "\n":
		j.Start = true
		fmt.Println("Button pressed: Start")
	case " ":
		j.Select = true
		fmt.Println("Button pressed: Select")
	case "w":
		j.Up = true
		fmt.Println("Button pressed: Up")
	case "s":
		j.Down = true
		fmt.Println("Button pressed: Down")
	case "a":
		j.Left = true
		fmt.Println("Button pressed: Left")
	case "d":
		j.Right = true
		fmt.Println("Button pressed: Right")
	}
	j.updateJoypadRegister(c)
}

func (j *Joypad) KeyUp(key string, c *cpu.CPU) {
	switch key {
	case "z":
		j.A = false
		fmt.Println("Button released: A")
	case "x":
		j.B = false
		fmt.Println("Button released: B")
	case "\n":
		j.Start = false
		fmt.Println("Button released: Start")
	case " ":
		j.Select = false
		fmt.Println("Button released: Select")
	case "w":
		j.Up = false
		fmt.Println("Button released: Up")
	case "s":
		j.Down = false
		fmt.Println("Button released: Down")
	case "a":
		j.Left = false
		fmt.Println("Button released: Left")
	case "d":
		j.Right = false
		fmt.Println("Button released: Right")
	}
	j.updateJoypadRegister(c)
}

func newJoypad() *Joypad {
	return &Joypad{Right: false, Left: false, Up: false, Down: false, A: false, B: false, Select: false, Start: false, SelectBits: 0x30}
}
func (j *Joypad) readKey(c *cpu.CPU) {
	ch := make(chan string)

	go func(ch chan string) {
		// disable input buffering

		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b []byte = make([]byte, 1)
		for {
			n, _ := os.Stdin.Read(b)
			if n != 1 {
				continue
			}
			fmt.Println("string being sent", string(b[:1]))
			ch <- string(b)
		}
	}(ch)
	var prevStr string = ""
	for {
		
		str, ok := <-ch
		if !ok {
			fmt.Println("channel was closed")
			return
		}
		if prevStr != "" && str != prevStr {
			j.Mux.Lock()
			j.KeyUp(prevStr, c)
			j.Mux.Unlock()
		}
		fmt.Println("key pressed ", str)
		j.Mux.Lock()
		j.KeyDown(str, c)
		j.Mux.Unlock()
		keyStatus(j)
		time.Sleep(time.Millisecond * 100)
		prevStr = str
	}
}

func (j *Joypad) updateJoypadRegister(c *cpu.CPU) error {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	result := j.SelectBits | 0xc0
	if j.SelectBits&byte(0x20) == 0 {
		if j.Right {
			result &^= 1 << 0
		} else if j.Left {
			result &^= 1 << 1
		} else if j.Up {
			result &^= 1 << 2
		} else if j.Down {
			result &^= 1 << 3
		} else {
			result |= 0xf
		}
	}
	if j.SelectBits&byte(0x10) == 0 {
		if j.Right {
			result &^= 1 << 0
		} else if j.Left {
			result &^= 1 << 1
		} else if j.Up {
			result &^= 1 << 2
		} else if j.Down {
			result &^= 1 << 3
		} else {
			result |= 0xf
		}
	}
	err := c.Bus.WriteByteToAddr(0xff00, result)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 100)
	return nil
}

func main() {
	var c cpu.CPU
	cpu.Init(&c)
	keyboard := newJoypad()
	keyboard.readKey(&c)
}

func keyStatus(joypad *Joypad) {
	fmt.Println("Key status: ")
	fmt.Println("========================")
	fmt.Println("A: ", joypad.A)
	fmt.Println("B: ", joypad.B)
	fmt.Println("Select: ", joypad.Select)
	fmt.Println("Start: ", joypad.Start)
	fmt.Println("Up: ", joypad.Up)
	fmt.Println("Down: ", joypad.Down)
	fmt.Println("Down: ", joypad.Down)
	fmt.Println("Right: ", joypad.Right)
	fmt.Println("Left: ", joypad.Left)

}
