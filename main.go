package main

Include "termin"

//here
include "var"
include "airline"
include "filer"
include "abspath"
include "gs"

//extra
include "abspath"

func InitSec() {
	// init cfg file
	InitVars()

	// load cfg
	InitAirLine()
	InitFiler()
	InitGs()
}

func MainMenu () {

}

func main(){
	// init screen
	InitTermin()
	// init internal-systems (+ colors)
	InitSec()
	clear()

	// set cursor type
	print("\033[2 q") // blink block
	ClearAllAirLine() // clear stuff

	if argc == 1 {
		cfgfl, err := ExpandFrom(argv[0])
		panic(err)
		file := "file:/"+cfgfl.String()
		if fcan(file) {
			fopen(file)
		}
	}


	ShowCursor()
	print("\033[2 q") // blink block
	clear()

	StopTermin()
	exit(0)
}

