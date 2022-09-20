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

func flist (f string) ([]string) {
	return ls(f[6:])
}

func fload (f string) ([]string) {
	if fcan(f) {
		return strings.Split(ReadFile(f[6:]), "\n")
	} else {
		Warn(1) // no file from link
		return []string{"invalid file link"}
	}
}

//TODO fopen check dir or file
func fopen (f string) (bool) {
	return Reader(fload(f), f[6:])
}

//READER

//init vars
var (
	ModeText [3]string
	bkgrey string
	FileColor string
	FolderColor string
	HiddenFileColor string

	// temp
	tstring string
	tint int
	tbool bool
	terror error

	mode = 0
)

// config file -> vars
func InitFiler() {
	// define colors
	bkgrey = colors["bkgrey"]
	FileColor = colors["FileColor"]
	FolderColor = colors["FolderColor"]
	HiddenFileColor = colors["HiddenFileColor"]
	ModeText = [...]string{
		colors["NormalMode"]+" NORMAL",
		colors["InsertMode"]+" INSERT",
		colors["NewTree"]+" NEWTREE ",
	}
}

// use airline
func ReaderAirline (filename, k string, y, x int) {
	ClearAirLine()
	AirLine(
		spf("%s %s %s@%s%d:%d %s%s",
		ModeText[mode],
		bk, filename,
		airline, y+1, x,
		k,
		txt,
	))
}

func WriterAirline (filename, k string, y, x, tint int) {
	ClearAirLine()
	AirLine(
		spf("%s %s %s@%s%d:%d::%d %s%s",
		ModeText[mode],
		bk, filename,
		airline, y+1, x, tint,
		k,
		txt,
	))
}

func ClearAll () () {
	for i:=0;i<Win.LenY-2;i++{
		wprint(Win, i, 0, "\033[2K")
	}
	ClearAllAirLine()
	ErrorLine(bkgrey+sws)
}

func l( s string ) ( int ) {
	x:=len(s)
	if x == 0 {
		return 0
	}
	return x-1
}

func retab ( l string ) ( string ) {
	for i:=0; len(l)-1 > i && (l[i]==' ') && (l[i+1] == ' ');i++ {
		l = strings.Replace(l, "  ", "\t", 1)
	}
	return l
}

func untab ( l string ) ( string ) {
	for i:=0; len(l) > i && (l[i]=='	');i+=2 {
		l = strings.Replace(l, "\t", "  ", 1)
	}
	return l
}

func Reader (c []string, filename string) (bool) {
	// normal mode
	mode = 0
	// set cursor type
	print("\033[2 q") // block
	var (
		// loop
		k string
		i int

		cmd string
		// read
		cl = len(c)
		off = cl-Win.LenY+1
		//ll []int

		y = 0
		x = 0
		w = 0//window shift
	)

	for i:=0;i<len(c);i++ {
		c[i] = untab(c[i])
	}

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
			wprint(Win, i, 0, "\033[2K"+bkgrey)
			wprint(Win, i, 0, c[i+w])
		}

		// print airline
		ReaderAirline(filename, k, y+w, x)

		// move cursor;get k
		wmove(Win, y, x)
		k = wgtk(Win)

		// use k
		switch (k) {
			//help
			case ("e"):
				tint = strings.Index(c[w+y][x:], " ")-1
				x = tint
			case ("w"):
				tint = strings.Index(c[w+y][x:], " ")+1
			//help
			case (":"):
				// change cursor type
				print("\033[6 q") // I-beam
				ClearWarn()
				wmove(Win, ALW.MinY+1, 0)
				tstring = ":"
				for {
					ErrorLine(tstring)
					k = wgtk(Win)
					if len(k) == 1 {
						tstring += k
					} else if k == "backspace" && len(tstring) != 0 {
						tstring = tstring[:len(tstring)-1]
					} else if k == "space" {
						tstring += " "
					} else if k == "enter" {
						break
					}
					// remove ':'
					if len(tstring) == 0 {
						break
					}
				}
				if len(tstring) != 0 {

					cl := strings.Split(tstring, " ")
					if len(cl) == 0 {
						Warn(2)// Empty
					}

					if len(cl) > 1 {
						cmd = cl[0]
						cl = cl[1:]
					} else {
						cmd = cl[0]
						cl = []string{}
					}

					tbool = false
					switch cmd {
						case (":w"):
							//save
							// retab
							WriteFile(filename, retab(strings.Join(c, "\n")))
						case (":q"):
							clear()
							// reset cursor type
							print("\033[1 q") // blink block
							return true
						default:
							tbool = true
					}
					if tbool {
						ErrorLine(ErrorText[3]+" '"+tstring+"'"+txt)
					} else {
						ErrorLine(tstring)
					}
					tstring = ""
				}
				// set cursor type
				print("\033[2 q") // block
			case ("backspace"):
				ClearAll()
				return false
			case ("space"):
				ClearAllAirLine()
			case ("_"):
				x = 0
				tint = x
			case ("$"):
				x = l(c[y+w])
				tint = x
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
				if k == "a" {
					x++
				}
				if x > len(c[w+y]) {
					x = len(c[w+y])
				} else if x < 0 {
					x = 0
				}
				// insert mode
				mode = 1
				// change cursor type
				print("\033[6 q") // I-beam
				for k!="esc" {
					// print airline
					WriterAirline(filename, k, y+w, x, tint)
					// move cursor;get k
					wmove(Win, y, x)
					//
					k = wgtk(Win)
					//
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
									tint = x
								}
							case ("right"):
								if x < len(c[y+w]) {
									x++
									tint = x
								}
							case ("down"):
								if y == Win.LenY-3 {
									if w < off {
										w++
									}
								} else {
									if y+1 < cl {
										y++
									}
								}
								if tint > x {
									x = tint
								}
								if l(c[y+w]) < x {
									x = len(c[y+w])
								}
							case ("up"):
								if y == 0 {
									if w != 0 {
										w--
									}
								} else {
									y--
								}
								if tint > x {
									x = tint
								}
								if l(c[y+w]) < x {
									x = len(c[y+w])
								}
							case ("tab"):
								c[y+w] = c[y+w][:x]+"  "+c[y+w][x:]
								x+=2
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

				// normal mode
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
//TODO (2): bkgrey into ErrorLine

// Reader out doesn't matter (when Folder().Reader())
//FOLDER
func Folder ( folder string ) () {
	// dir mode
	mode = 2
	// set cursor type
	print("\033[2 q") // block
	FolderAirline(folder, "no git yet")

	var (
		dirs []string
		ti = 0
		//y int
	)

	dirs = flist(folder)
	ErrorLine(bkgrey)
	for i:=0;i<Win.LenY-2;i++{
		wprint(Win, i, 0, "\033[2K"+bkgrey)
	}
	for i:=0;i<len(dirs);i++ {
		if dirs[i][0] == '.' {
			continue
			//wuprint(Win, i, 0, HiddenFileColor)
		}
		if ( dirs[i][len(dirs[i])-1] == '/' ) {
			wColor(FolderColor)
		} else {
			wColor(FileColor)
		}
		if tint = strings.Index(dirs[i], "."); tint != -1 {
			tstring = dirs[i][strings.Index(dirs[i], "."):]
			if tstring, tbool = FileColors[tstring]; tbool {
				wColor(tstring)
			}
		}
		wprint(Win, ti, 0, dirs[i])
		ti++
	}

	wmove(Win, 0,0)
	wgtk(Win)
}

//TODO(4) git: get gs.go's info
//TODO(1) debug: debug display colors FG/BK
func FolderAirline ( dir string, git string ) () {
	ClearAll()
	AirLine( spf(
		"%s%s %s%s",
		ModeText[mode], AirlineText, dir, git,
	))
}


//HTTP?
// https links
func lvalid (f string) (bool) {
	return strings.Index(f, "https:/")==0
}

