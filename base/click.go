package base

import (
	"github.com/nsf/termbox-go"
)

func ClickHandle() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventMouse:
			if ev.Key == termbox.MouseLeft {

			} else if ev.Key == termbox.MouseRight {

			} else {

			}
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowDown:
				break
			case termbox.KeyArrowLeft:
				break
			case termbox.KeyArrowRight:
				break
			case termbox.KeyArrowUp:
				break
			default:
			}
		default:
		}
	}
}
