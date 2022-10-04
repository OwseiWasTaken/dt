
const (
	dir = "./config/"
	file = "cfg"
	colordir = "./config/colors/"
)

//init vars
var (
	cfg = map[string]string{}
	FileColors = map[string]string{}
	colors = map[string]string{}
	sws string
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
	for i=0;i<Win.LenY-2;i++{
		wprint(Win, i, 0, "\033[2K")
	}
	//colors
	// add FileColors?
	i=0
	for key, val := range colors {
		wprint(Win, i, 0, spf("%s:%spog████		%s", key, val, colors["nc"]))
		i++
	}
	wgtk(Win)
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
		cfg[name] = value
	}
	LoadColors(cfg["colorfile"])
}

func InterpretColorLine ( line string ) ( string, string ) {
	h := line[0]=='"'
	if h {
		line = line[1:]
	}
	line = strings.Replace(line, " ", "", -1)
	var ll = strings.Split(line, ":")
	if len(ll) == 1 {
		// can't use config colors lol
		wuprint(Win, 0, 0, RGB(255,0,0)+"invalid color line"+RGB(255,255,255))
		wuprint(Win, 1, 0, "\""+line+"\"")
		exit(1)
	}
	var name = ""
	var code = ""
	name = ll[0]
	cd := strings.Split(ll[1], ",")

	if len(cd) == 1 {
		code = colors[cd[0]]
	} else if len(cd) == 3 {
		if h{
			code = bkcolor(cd[0], cd[1], cd[2])
		} else {
			code = RGB(cd[0], cd[1], cd[2])
		}
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
		if line[0] == '"'  {
			name, code = InterpretColorLine(line)
			if len(name) == 0 || len(code) == 0 {continue}
			colors[name] = code
		}
		} else if line[0] != '#' {
			name, code = InterpretColorLine(line)
			if len(name) == 0 || len(code) == 0 {continue}
			colors[name] = code
		}
	}
}

const (
	T_bool = iota
	T_string = iota
	T_int = iota
)

func ReadCFG (name string, T int) (interface{}) {
	s := cfg[name]
	switch (T) {
		case (T_bool):
			if s == "false" || s == "true" {
				return s == "true"
			} else {
				//TODO: break
			}
	}
	return nil
}

func RCfgB(name string) (bool) {
	return ReadCFG(name, T_bool).(bool)
}

func RCfgS(name string) (string) {
	return ReadCFG(name, T_string).(string)
}
