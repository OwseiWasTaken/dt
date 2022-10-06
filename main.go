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

	var buffers = []Buffer{}
	//func MakeReader (win *Window, c []string, filename string) (R_Buffer) {
	bwin := MakeWin(
		"Filer/Editor Window",
		stdout, stdin,
		1, Win.LenY-Alw.LenY,
		0, Win.LenX,
	)
	buffers = append(buffers, MakeReader(bwin, []string{"test", "string"}, "test buff"))
	buffers[0].Run()


	ShowCursor()
	print("\033[2 q") // blink block
	clear()
	StopTermin()
	exit(0)
}

