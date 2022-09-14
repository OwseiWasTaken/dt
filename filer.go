/// yel/ file:/(...)
// file://home/ow/
// https://google.com

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
		ReaderWarn(1) // no file from link
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
	bk string
	airline string
	txt string
	slap string
	cleanslap string
	ModeText [2]string
	AirlineText string
	BadError string
	SimpleError string
	ErrorText []string
	mode = 0
)

func InitFiler() {
	AirlineText = colors["AirLineText"]
	bk = colors["BK"]
	airline = colors["Airline"]
	txt = colors["Text"]
	slap = bk+strings.Repeat(" ", Win.LenX)+txt
	cleanslap = txt+strings.Repeat(" ", Win.LenX)
	BadError = colors["BadError"]
	SimpleError = colors["SimpleError"]
	ModeText = [...]string{
		colors["NormalMode"]+" NORMAL",
		colors["InsertMode"]+" INSERT",
	}
	ErrorText = []string{
		BadError+"No file from link",
		SimpleError+"Not a link",
	}
}


func ReaderClear() {
	clear()
	wuprint(
		Win,
		Win.LenY-2, 0,
		slap,
	)
}

func ReaderWarn(warntype int) {
	ClearWarn()
	wuprint(
		Win, Win.LenY-1, 0, ErrorText[warntype]+txt,
	)
}

func ClearWarn() {
	wuprint(
		Win, Win.LenY-1, 0, cleanslap,
	)
}

func ReaderAirline (filename, k string, y, x int) {
	wuprint(Win, Win.LenY-2, 0,
		spf("%s %s %s@%s%d:%d %s",
		ModeText[mode],
		bk, filename,
		airline, y+1, x,
		k+strings.Repeat(" ", 9-len(k)),//9len = biggest wgtk ret
	)+txt)
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
		l = len(c)
		off = l-Win.LenY+1
		ll []int

		y = 0
		x = 0
		w = 0//window shift
	)

	for i=0;i<l;i++ {
		tint = len(c[i])-1
		if tint == -1 {
			tint = 0
		}
		ll = append(ll, tint)
	}

	// div
	ReaderClear()

	// clear temps
	tint = 0

	for k!="q" {
		// clear;draw text
		for i=0;(i+w)<l&&(i<Win.LenY-2);i++{
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
				return false
			case ("space"):
				ClearWarn()
			case ("_"):
				x = 0
			case ("$"):
				x = ll[y+w]
			case ("enter"):
				//TODO(1) link: get if link from line[x:]
				tstring = c[y+w][x:]
				tstring = strings.Split(tstring, " ")[0]
				if fvalid(tstring) {
					if fcan(tstring) {
							if (fopen(tstring)) {
								return true
							}
							ReaderClear()
					} else {
						ReaderWarn(0) // warn: no file from link
					}
				} else {
					ReaderWarn(1) // warn: not a link
				}
			case ("g"):
				w = 0
				y = 0
				x = 0
			case ("G"):
				w = off
				y = Win.LenY-3
				x = 0
			case ("z"):
				if off > w {
					w++
				}
			case ("x"):
				if w > 0 {
					w--
				}
			case ("j"):
				if y < Win.LenY-3 && (y+w+1) < l {
					y++
					x = tint
					if ll[y+w] < x {
						x = ll[y+w]
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
					if ll[y+w] < x {
						x = ll[y+w]
					}// else if ll[y+w] >= x {
					//	x = tint
					//}
				} else {
					if w > 0 {
						w--
					}
				}
			case ("l"):
				if x < ll[y+w] { // -1 no overhang
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
					for i=0;(i+w)<l&&(i<Win.LenY-2);i++{
						wprint(Win, i, 0, "\033[2K")
						wprint(Win, i, 0, c[i+w])
					}

					// print airline
					ReaderAirline(filename, k, y+w, x)

					// move cursor;get k
					wmove(Win, y, x)
					k = wgtk(Win)
				}
				mode = 0
				//TODO insert mode
		}
	}
	return true
}

// https links
func lvalid (f string) (bool) {
	return strings.Index(f, "https:/")==0
}
