package debug

import (
	"github.com/temnok/esp32/dmi"
)

func Session(block func(*Conn)) {
	dmi.Session(func(dmi *dmi.Conn) {
		block(&Conn{dmi})
	})
}
