package main

Include "termin"
include "var"
include "filer"
include "abspath"

func main(){
	InitTermin()
	InitVars()
	InitFiler()

	fopen("file://home/ow/70lines.txt")

	StopTermin()
	exit(0)
}

