package devices

import (
	"fmt"
	"github.com/carlmango11/schmarlbox/backend/box/log"
	
)

const (
	maxRows = 25
	maxCols = 80
)

type Display struct {
	output *os.File

	row, col int
	screen   [maxRows][maxCols]byte
}

func NewDisplay() *Display {
	//f, err := os.OpenFile("/tmp/display.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	//if err != nil {
	//	panic(err)
	//}

	return &Display{
		//output: f,
	}
}

func (d *Display) Read(addr uint16) byte {
	return 0
}

func (d *Display) Write(addr uint16, val byte) {
	d.handle(val)
}

func (d *Display) handle(ch byte) {
	fmt.Print(string(ch))
	log.Printf("input char: 0x%x (%d)", ch, ch)
	return
	fmt.Print("\033[2J\033[H")
	if ch == '\n' {
		d.row++
		if d.row > maxRows {
			d.row = 0
		}

		return
	}

	d.screen[d.row][d.col] = ch

	d.col++
	if d.col == maxCols {
		d.col = 0
	}

	//d.printFile()
}

func (d *Display) printFile() {
	d.output.WriteString("\n=================\n")
	for r := range d.screen {
		for c := range d.screen[r] {
			ch := string(d.screen[r][c])

			if d.screen[r][c] == 0 {
				ch = " "
			}

			log.Println(ch)
			d.output.Write([]byte(fmt.Sprintf("%v", ch)))
		}
		fmt.Printf("\n")
	}
}

func (d *Display) State() [maxRows][maxCols]string {
	display := [maxRows][maxCols]string{}

	for x := range d.screen {
		for y := range d.screen[x] {
			display[x][y] = string(d.screen[x][y])
		}
	}

	return display
}
