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

	fopen("file://home/ow/70lines.txt")

	StopTermin()
	exit(0)
}

