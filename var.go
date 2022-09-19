type _cfg struct {
	flname string
	colordir string
	colorfile string
}

const (
	dir = "./config/"
	file = "cfg"
	colordir = "./config/colors/"
)

//init vars
var (
	cfg = _cfg{file, colordir, ""}
	colors = map[string]string{}
)

func InitVars () {
	clear()
	load(dir+file)
}

func load (f string) () {
	file := strings.Split(ReadFile(f), "\n")
	cfg.flname = f
	line := ""
	for i:=0;i<len(file);i++ {
		line = file[i]
		if len(line) < 2{continue}
		value := line[strings.Index(line, ":")+1:]
		switch (i){
			case 0:
				cfg.colorfile=value[1:len(value)-1]
		}
	}
	LoadColors(cfg.colorfile)
}

func InterpretColorLine ( line string ) ( string, string ) {
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
	var fl = strings.Split(ReadFile(cfg.colordir+f+".clrs"), "\n")
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
		if line[0] != '#' {
			name, code = InterpretColorLine(line)
			if len(name) == 0 || len(code) == 0 {continue}
			colors[name] = code
		}
	}
}

func save(cfg _cfg) (bool) {
	// TODO
	return true
}

func MakeDefaultFiles () {
	// TODO
}
