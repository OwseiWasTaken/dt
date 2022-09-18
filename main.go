package main

Include "termin"
include "var"
include "filer"
include "abspath"
include "airline"

func main(){
	InitTermin()

	// set cursor type
	print("\033[2 q") // blink block

	InitVars()
	InitFiler()
	InitAirLine()

	Folder("file://home/ow/")

	StopTermin()
	exit(0)
}

