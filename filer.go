// file://home/ow/
// https://google.com

//TODO: G error
//enter file://file
//press 'G'
//crash?

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
		ReportInternalError("faild to prevent opening an invalid link", 1)
		return []string{"nawh man"}
	}
}

//true:quit
func fopen (f string) (bool) {
	// try to read as dir
	_, err := ioutil.ReadDir(f)
	if err == nil {
		return Folder(f)
	} else {
		return Reader(fload(f), f[6:])
	}
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
	Few *Window

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
		colors["NormalMode"]+" NORMAL "+AirLineText+" ",
		colors["InsertMode"]+" INSERT "+AirLineText+" ",
		colors["NewTree"]+" NEWTREE "+AirLineText+" ",
	}
	Few = MakeWin(
		"Filer/Editor Window",
		stdout, stdin,
		0, Win.LenY-3,
		0, Win.LenX,
	)
}

func ShortenName (f string) (string) {
	var r = "/"
	var index int
	var gone int // chars already gone over
	for {
		index = strings.Index(f[gone:], "/")
		gone+=index+1
		if strings.Index(f[gone:], "/") == -1 {
			r+=f[gone:]
			break
		}
		r+=string(f[gone:gone+index+1][0])+"/"
	}
	return r
}

func MakeAirLine (s string) {
	ClearAirLine()
	AirLine(
		ModeText[mode]+s,
	)
}

func ReaderAirLine (filename, k string, y, x int) {
	MakeAirLine( spf(
		"%s%s@%d:%d %s%s",
		AirLineText,
		filename,
		y+1, x,
		k,
		txt,
	))
}

func WriterAirLine (filename, k string, y, x, tint int) {
	MakeAirLine( spf(
		"%s%s@%s%d:%d::%d %s%s",
		bk, filename,
		airline, y+1, x, tint,
		k,
		txt,
	))
}

func FolderAirLine ( dir string, git string ) () {
	MakeAirLine( spf(
		"%s %s %s%s",
		AirLineText, dir, git, txt,
	))
}

func ClearAll () () {
	for i:=0;i<Few.MaxY;i++{
		wprint(Few, i, 0, "\033[2K")
	}
	ClearAllAirLine()
	ReportLine(bkgrey+sws)
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
		l = strings.Replace(l, " "+" ", "\t", 1)
	}
	return l
}

func untab ( l string ) ( string ) {
	return strings.Replace(l, "\t", " "+" ", 1)
}

//TODO(1): make reader win
func Reader (c []string, filename string) (bool) {
	// normal mode
	mode = 0
	// set cursor type
	ShowCursor()
	wuprint(Few, 0, 0, "\033[2 q") // block
	var (
		// loop
		k string
		i int

		cmd string
		args []string
		shortname string
		// read
		cl = len(c)
		off = cl-Few.LenY-1
		//ll []int

		y = 0
		x = 0
		w = 0//window shift
	)

	shortname = ShortenName(filename[6:])

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
		// clear;draw text
		for i=0;(i+w)<cl&&(i<Few.LenY);i++{
			wprint(Few, i, 0, "\033[2K"+bkgrey)
			wprint(Few, i, 0, c[i+w])
		}

		// print airline
		ReaderAirLine(shortname, k, y+w, x)

		// move cursor;get k
		wmove(Few, y, x)
		k = wgtk(Few)

		// use k
		switch (k) {
			case (":"):
				// change cursor type
				print("\033[6 q") // I-beam
				ClearReport()
				wmove(Alw, Alw.LenY-2, 0)// move to report line
				tstring = ":"
				for {
					ReportLine(tstring)
					k = wgtk(Alw)
					if len(k) == 1 {
						tstring += k
					} else if k == "backspace" && len(tstring) != 0 {
						tstring = tstring[:len(tstring)-1]
					} else if k == "space" {
						tstring += " "
					} else if k == "enter" {
						break
					}
					if len(tstring) == 0 {
						break
					}
				}
				if len(tstring) != 0 {

					args = strings.Split(tstring, " ")
					if len(args) == 0 {
						Warn(2)// Empty
					}

					if len(args) > 1 {
						cmd = args[0]
						args = args[1:]
					} else {
						cmd = args[0]
						args = []string{}
					}

					// report command
					ReportLine(tstring)
					switch cmd {
						case (":w"):
							//save
							// retab
							if len(args) == 1 {
								shortname = ShortenName(args[0])
								filename = "file://"+args[0]
								_, terror = os.OpenFile(args[0], os.O_CREATE|os.O_WRONLY, 0644)
								if terror != nil {
									AdvWarn(5, spf("%v", terror))
								}
							}
							terror = os.WriteFile(
								filename,
								[]byte(retab(strings.Join(c, "\n"))),
								0644,//TODO: change file 0otype
							)
							if !exists(args[0]) {
								AdvWarn(4, filename + spf("%v", terror), "d")
							}
						case (":q"):
							clear()
							// reset cursor type
							print("\033[1 q") // blink block
							return true
						case (":wq"):
							clear()
							print("\033[1 q")
							WriteFile(filename, retab(strings.Join(c, "\n")))
							return true
						default:
							// overwrite report with error
							AdvWarn(3, tstring)
							//ReportLine(ErrorText[3]+" '"+tstring+"'"+txt)
					}
					tstring = ""
				}
				// set cursor type
				print("\033[2 q") // block
			case "backspace", "^H":
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
				if len(c[y+w])!=0 {
					c[y+w] = c[y+w][:x]+c[y+w][x+1:]
				}
				if len(c[y+w]) < x {
					x--
				}
			case ("c"):
				if w > 0 {
					w--
				}
			case ("j"):
				if y < Few.LenY-1 && (y+w+1) < cl {
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
					// clear;draw text
					for i=0;(i+w)<cl&&(i<Few.LenY);i++{
						wprint(Few, i, 0, "\033[2K")
						wprint(Few, i, 0, c[i+w])
					}

					// print airline
					WriterAirLine(filename, k, y+w, x, tint)
					// move cursor;get k
					wmove(Few, y, x)
					//
					k = wgtk(Few)
					//
					if len(k) == 1{
						c[y+w] = c[y+w][:x]+k+c[y+w][x:]
						x++
						continue
					}
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
							if y == Few.LenY-1 {
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
							c[y+w] = c[y+w][:x]+" "+" "+c[y+w][x:]
							x+=2
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
		}
	}
	clear()
	// reset cursor type
	print("\033[1 q") // blink block
	return true
}

// Reader out doesn't matter (when Folder().Reader())
//FOLDER
func Folder ( folder string ) (bool) {
	// dir mode
	mode = 2
	// set cursor type
	HideCursor()
	FolderAirLine(folder, "no git yet")

	var (
		dir []string
		Ldir []string
		fl = flist(folder)
		git string
		ShowHiddenFiles bool
		ShowFiles bool
		ShowDirs bool
		k string
		ld int
		i int
		y = 0
		mark string
	)
	//fl = append(fl, "../")
	mark = colors["red"]+"*"+colors["white"]

	git = GetGs(folder[6:])
	ShowHiddenFiles = RCfgB("ShowHiddenFiles")
	ShowFiles = RCfgB("ShowFiles")
	ShowDirs = RCfgB("ShowDirs")

	dir = FilterFolder(fl,
		ShowHiddenFiles, ShowFiles, ShowDirs, true,
	)
	Ldir = FilterFolder(fl,
		ShowHiddenFiles, ShowFiles, ShowDirs, false,
	)

	ld = len(dir)
	if len(Ldir) != ld {
		wuprint(Few, 0, 0, "fuck")
		wgtk(Few)
	}

	// clear screen
	ClearReport()
	wColor(bkgrey)
	for i:=0;i<Few.LenY;i++ {
		wprint(Few, i, 0, "\033[2K")
	}

	for k!="backspace"&&k!="^H"{
		FolderAirLine(folder, git)
		for i=0;i<ld;i++ {
			if i < ld {
				wprint(Few, i, 0, "\033[2K")
				if i == y {
					wprint(Few, i, 0, dir[i]+mark)
				} else {
					wprint(Few, i, 0, dir[i])
				}
			}
		}
		wmove(Few, y, 0)
		k = wgtk(Few)
		switch (k) {
			case ("j"):
				if ld != y+1{
					y++
				}
			case ("k"):
				if y != 0 {
					y--
				}
			case ("enter"):
				if Ldir[y] == "../" {
					//TODO: up dir
				} else if fopen(folder+Ldir[y]) {
					return true
				}
				HideCursor()
			case ("q"):
				return true
		}
		ReportLine(spf("%v", y))
	}

	ShowCursor()
	return false
}

func RemoveIndex ( s []string, i int ) ( []string ) {
	if len(s) <= i+2 {
		s = s[:i]
	} else {
		s = append(s[:i], s[i+2:]...)
	}
	return s
}

//Show [Hidden] File/Dir
//S[H]F, SD
//Use Colors
func FilterFolder ( dir []string, SHF, SF, SD, UC bool) ( []string ) {
	for i:=0;i<len(dir);i++ {
		if dir[i][0] == '.' {
			if SHF {
				if UC {
					dir[i] = HiddenFileColor+dir[i]
				}
			} else {
				dir = RemoveIndex(dir, i)
				i--
			}
		} else if ( dir[i][len(dir[i])-1] == '/' ) {
			if SD {
				if UC {
					dir[i] = FolderColor+dir[i]
				}
			} else {
				dir = RemoveIndex(dir, i)
				i--
			}
		} else {
			if SF {
				tint = strings.Index(dir[i], ".")
				if tint != -1 && len(dir[i]) != 1 {
					tstring = dir[i][tint:]
					if tstring, tbool = FileColors[tstring]; tbool {
						if UC {
							dir[i] = tstring+dir[i]
						}
					} else {
						if UC {
							dir[i] = txt+dir[i]
						}
					}
				} else {
					if UC {
						dir[i] = txt+dir[i]
					}
				}
			} else {
				dir = RemoveIndex(dir, i)
				i--
			}
		}
	}
	return dir
}

//HTTP?
// https links
func lvalid (f string) (bool) {
	return strings.Index(f, "https:/")==0
}

