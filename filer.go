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
			case ("a"):
				mode = 1
				for k!="esc" {
					// clear;draw text
					for i=0;(i+w)<cl&&(i<Win.LenY-2);i++{
						wprint(Win, i, 0, "\033[2K")
						wprint(Win, i, 0, c[i+w])
					}

					// print airline
					ReaderAirline(filename, k, y+w, x)
					if len(k) == 1{
						c[y+w] = c[y+w][:x]+k+c[y+w][x:]
					} else {

					}

					// move cursor;get k
					wmove(Win, y, x)
					k = wgtk(Win)
				}

				mode = 0
				//TODO insert mode
		}
	}
	clear()
	return true
}

// https links
func lvalid (f string) (bool) {
	return strings.Index(f, "https:/")==0
}
