package main

import "fmt"

type house struct {
	windowType string
	doorType   string
	floor      int
}

func (h *house) Done() string {
	return fmt.Sprintf("House is completed!\n" +
		"The window is made up %s,\n" +
		"the door is made up %s, \n" +
		"and it has %d floors! \n", h.windowType, h.doorType, h.floor)
}