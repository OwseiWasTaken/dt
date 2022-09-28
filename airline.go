// ==status=line==
// error line
//TODO(2) error line -> report line: (w/ error features)


var (
	airline string
	AirLineText string
	BadError string
	SimpleError string
	slap string
	cleanslap string
	bk string
	txt string
	airlinetxt string

	ErrorText []string

	ALW *Window
)

func InitAirLine () {
	// define colors
	airline = colors["AirLine"]
	BadError = colors["BadError"]
	SimpleError = colors["SimpleError"]
	txt = colors["TextBkGrey"]
	AirLineText = colors["AirLineText"]
	bk = colors["AirLineText"]
	slap = bk+sws+txt
	cleanslap = txt+sws

	// define errors
	ErrorText = []string{
		BadError+"No file from link",
		SimpleError+"Not a link",
		SimpleError+"Command Empty",
		BadError+"No Such Command \"%s\"",
	}

	// make airline window
	ALW = MakeWin(
		"AirLine Window",
		stdout, stdin,
		Win.LenY-2, Win.LenY, 0, Win.LenX,
	)
}

func ClearAllAirLine() {
	AirLine(slap)
	ReportLine(cleanslap)
}

// AirLine
func AirLine ( s string ) {
	// make bkground color
	wuprint(ALW, 0, 0, slap)
	// write
	wuprint(ALW, 0, 0, s)
}

func ClearAirLine() {
	AirLine(slap)
}

// warn
func Warn(warntype int) {
	ClearReport()
	ReportLine(ErrorText[warntype]+txt)
}

func AdvWarn(warntype int, inp ...string) {
	ClearReport()
	t:=ErrorText[warntype]
	for i:=0;i<len(inp);i++ {
		t = spf(t, inp[i])
	}
	ReportLine(t)
	wColor(txt)
}

func ReportInternalError( s string, ec int ) {
	ClearReport()
	ReportLine(s)
	wgtk(Win)
	if ec != 0 {
		exit(ec)
	}
}

func ClearReport() {
	ReportLine(cleanslap)
}

func ReportLine ( s string ) {
	wuprint(ALW, 1, 0, s)
}

