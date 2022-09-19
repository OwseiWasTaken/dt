package main

Include "termin"
include "var"
include "filer"
include "abspath"
include "airline"

func main(){
	InitTermin()

	InitVars()
	InitFiler()
	InitAirLine()

	// set cursor type
	print("\033[2 q") // blink block

	Folder("file://home/ow/code/py")
	Folder("file://home/ow/code/golang")

	StopTermin()
	exit(0)
}

