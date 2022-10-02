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

	Alw *Window
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
		"UNUSED ERROR",
		SimpleError+"Invalid Link",
		SimpleError+"Command Empty",
		BadError+"No Such Command \"%s\"",
		BadError+"Can't Write To File %s: %s",
		BadError+"Can't Create File %s",
	}

	// make airline window
	Alw = MakeWin(
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
	wuprint(Alw, 0, 0, slap)
	// write
	wuprint(Alw, 0, 0, s)
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
	wgtk(Alw)
	if ec != 0 {
		exit(ec)
	}
}

func ClearReport() {
	ReportLine(cleanslap)
}

func ReportLine ( s string ) {
	wuprint(Alw, 1, 0, s)
}

