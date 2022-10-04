package main
Include "termin"

include "var"
include "airline"
include "filer"
include "abspath"
include "gs"

func InitSec() {
	// init cfg file
	InitVars()

	// load cfg
	InitAirLine()
	InitFiler()
	InitGs()
}

func main(){
	// init screen
	InitTermin()
	// init internal-systems (+ colors)
	InitSec()


	// set cursor type
	print("\033[2 q") // blink block
	ClearAllAirLine()

	fopen("file://home/owsei/projs/dt/")
	//debug()

	ShowCursor()
	print("\033[2 q") // blink block
	clear()
	StopTermin()
	exit(0)
}

