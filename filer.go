/// yel/ file:/(...)
// file://home/ow/
// https://google.com

//TODO
//enter file://file
//press 'G'

// file links
func fvalid (f string) (bool) {
	return strings.Index(f, "file:/")==0
}

func fexists (f string) (bool) {
	return exists(f[6:])
}

func fcan (f string) (bool) {
	return fvalid(f) && fexists(f)
}

func fload (f string) ([]string) {
	if fcan(f) {
		return strings.Split(ReadFile(f[6:]), "\n")
	} else {
		Warn(1) // no file from link
		return []string{"invalid file link"}
	}
}

func fopen (f string) (bool) {
	return Reader(fload(f), f[6:])
}

//READER

type Pair struct {
	y int
	x int
}

//init vars
var (
	ModeText [2]string
	mode = 0
)

func InitFiler() {
	// define colors
	ModeText = [...]string{
		colors["NormalMode"]+" NORMAL",
		colors["InsertMode"]+" INSERT",
	}
}

// use airline
func ReaderAirline (filename, k string, y, x int) {
	AirLine(
		spf("%s %s %s@%s%d:%d %s",
		ModeText[mode],
		bk, filename,
		airline, y+1, x,
		k,//9len = biggest wgtk ret
	)+txt)
}

func ClearAll () () {
	for i:=0;i<Win.LenY-2;i++{
		wprint(Win, i, 0, "\033[2K")
	}
	ClearAllAirLine()
}

func l( s string ) ( int ) {
	x:=len(s)
	if x == 0 {
		return 0
	}
	return x-1
}

func Reader (c []string, filename string) (bool) {
	// set cursor type
	print("\033[2 q") // block
	//TODO tab: tabs on the files kinda break
	var (
		// temp
		tstring string
		tint int
		//tbool bool
		// loop
		k string
		i int

		// read
		cl = len(c)
		off = cl-Win.LenY+1
		//ll []int

		y = 0
		x = 0
		w = 0//window shift
	)

	if off < 0 {
		off = 0
	}

	// div
	ClearAllAirLine()

	// clear temps
	tint = 0

	for k!="q" {
		//TODO command: use ';' to use a command
		// clear;draw text
		for i=0;(i+w)<cl&&(i<Win.LenY-2);i++{
			wprint(Win, i, 0, "\033[2K")
			wprint(Win, i, 0, c[i+w])
		}

		// print airline
		ReaderAirline(filename, k, y+w, x)

		// move cursor;get k
		wmove(Win, y, x)
		k = wgtk(Win)

		// use k
		switch (k) {
			case ("backspace"):
				ClearAll()
				return false
			case ("space"):
				ClearAllAirLine()
			case ("_"):
				x = 0
			case ("$"):
				//x = ll[y+w]
				x = l(c[y+w])
			case ("enter"):
				tstring = c[y+w][x:]
				tstring = strings.Split(tstring, " ")[0]
				if fvalid(tstring) {
					if fcan(tstring) {
							ClearAll()
							if (fopen(tstring)) {
								return true
							}
					} else {
						Warn(0) // warn: no file from link
					}
				} else {
					Warn(1) // warn: not a link
				}
			case ("g"):
				w = 0
				y = 0
				x = 0
			case ("G"):
				w = off
				y = cl-w-1 // -1 for the y
				if w != 0 {
					y-- //-1 for the w
				}
				x = 0
				tint = 0
			case ("z"):
				if off > w {
					w++
				}
			case ("x"):
				if w > 0 {
					w--
				}
			case ("j"):
				if y < Win.LenY-3 && (y+w+1) < cl {
					y++
					x = tint
					//ll
					if l(c[y+w]) < x {
						x = l(c[y+w])
					}
				} else {
					if off > w {
						w++
					}
				}
			case ("k"):
				if y > 0 {
					y--
					x = tint
					if l(c[y+w]) < x {
						x = l(c[y+w])
					}// else if ll[y+w] >= x {
					//	x = tint
					//}
				} else {
					if w > 0 {
						w--
					}
				}
			case ("l"):
				if x < l(c[y+w]) { // -1 no overhang
					x++
					tint=x
				}
			case ("h"):
				if x > 0 {
					x--
					tint=x
				}
			case "a", "i":
				if k == "i" {
					x--
				}
				if x > l(c[w+y]) {
					x = l(c[w+y])
				}
				mode = 1
				// change cursor type
				print("\033[6 q") // I-beam
				x++
				for k!="esc" {
					// move cursor;get k
					wmove(Win, y, x)
					k = wgtk(Win)

					// print airline
					ReaderAirline(filename, k, y+w, x)
					if len(k) == 1{
						c[y+w] = c[y+w][:x]+k+c[y+w][x:]
						x++
					} else {
						switch k {
							case ("space"):
								c[y+w] = c[y+w][:x]+" "+c[y+w][x:]
								x++
							case ("backspace"):
								if len(c[y+w])!=0 {
									c[y+w] = c[y+w][:x-1]+c[y+w][x:]
									x--
								}
							case ("left"):
								if x != 0 {
									x--
								}
							case ("down"):
							case ("up"):
							case ("right"):
								if x <= l(c[y+w]) {
									x++
								}
						}
					}
					// clear;draw text
					for i=0;(i+w)<cl&&(i<Win.LenY-2);i++{
						wprint(Win, i, 0, "\033[2K")
						wprint(Win, i, 0, c[i+w])
					}

				}
				x--
				if x == -1 {
					x = 0
				}

				mode = 0
				// change cursor type
				print("\033[1 q") // block
				//TODO insert mode
		}
	}
	clear()
	// reset cursor type
	print("\033[1 q") // blink block
	return true
}

// https links
func lvalid (f string) (bool) {
	return strings.Index(f, "https:/")==0
}
