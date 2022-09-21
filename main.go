package main

Include "termin"

include "var"
include "airline"
include "filer"
include "abspath"
include "gs"

func main(){
	// init screen
	InitTermin()

	// init cfg file
	InitVars()

	// load cfg
	InitAirLine()
	InitFiler()


	// set cursor type
	print("\033[2 q") // blink block

	Folder("file://home/ow/code/golang/dt")

	StopTermin()
	exit(0)
}

