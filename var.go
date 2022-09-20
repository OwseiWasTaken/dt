
const (
	dir = "./config/"
	file = "cfg"
	colordir = "./config/colors/"
)

//init vars
var (
	cfg = map[string]string{}
	colors = map[string]string{}
	FileColors = map[string]string{}
	sws string
	used = []
)

// run after Init
func debug () () {
	//vars
	for i:=0;i<Win.LenY-2;i++{
		wprint(Win, i, 0, "\033[2K")
	}
	i:=0
	for key, val := range cfg {
		wprint(Win, i, 0, spf("%s:%s", key, val))
		i++
	}
	wgtk(Win)
	for i:=0;i<Win.LenY-2;i++{
		wprint(Win, i, 0, "\033[2K")
	}
	//colors
	// add FileColors?
	i=0
	for key, val := range colors {
		wprint(Win, i, 0, spf("%s:%s████   %s", key, val, colors["nc"]))
		i++
	}
}

func InitVars () {
	cfg["flname"] = dir+file
	cfg["colordir"] = colordir
	sws = strings.Repeat(" ", Win.LenX)
	load(dir+file)
}

func load (f string) () {
	file := strings.Split(ReadFile(f), "\n")
	line := ""
	for i:=0;i<len(file);i++ {
		line = file[i]
		if len(line) < 2{continue}
		value := line[strings.Index(line, ":")+1:]
		name := line[:strings.Index(line, ":")]
		// oh my, after 750 lines, python would be useful for the first time!
		cfg[name] = value[1:len(value)-1]
	}
	LoadColors(cfg["colorfile"])
}

func InterpretColorLine ( line string ) ( string, string ) {
	line = strings.Replace(line, " ", "", -1)
	var ll = strings.Split(line, ":")
	var name = ""
	var code = ""
	name = ll[0]
	cd := strings.Split(ll[1], ",")

	if len(cd) == 1 {
		code = colors[cd[0]]
	} else if len(cd) == 3 {
		code = RGB(cd[0], cd[1], cd[2])
	} else if len(cd) == 6 {
		code = color(cd[0],cd[1],cd[2], cd[3],cd[4],cd[5])
	}

	return name, code
}

func LoadColors ( f string ) ( ) {
	var fl = strings.Split(ReadFile(cfg["colordir"]+f+".clrs"), "\n")
	var line string
	var cutoff int
	var name string
	var code string
	for i:=0;i<len(fl);i++ {
		line = fl[i]
		cutoff = strings.Index(line, " ")
		if cutoff != -1 {
			line = line[:cutoff]
		}
		if len(line) < 3 {continue}

		if line[0] == '.' {
			name, code = InterpretColorLine(line)
			if len(name) == 0 || len(code) == 0 {continue}
			FileColors[name] = code
		} else if line[0] != '#' {
			name, code = InterpretColorLine(line)
			if len(name) == 0 || len(code) == 0 {continue}
			colors[name] = code
		}
	}
}

func save(cfg map[string]string) (bool) {
	// TODO
	return true
}

func MakeDefaultFiles () {
	// TODO
}
