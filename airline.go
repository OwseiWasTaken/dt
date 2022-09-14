// ==status=line==
// error line

var (
	airline string
	AirlineText string
	BadError string
	SimpleError string
	slap string
	cleanslap string
	bk string
	txt string

	ErrorText []string

	ALW *Window
)

func InitAirLine () {
	// define errors
	ErrorText = []string{
		BadError+"No file from link",
		SimpleError+"Not a link",
	}
	// define colors
	airline = colors["Airline"]
	AirlineText = colors["AirLineText"]
	BadError = colors["BadError"]
	SimpleError = colors["SimpleError"]
	slap = bk+strings.Repeat(" ", Win.LenX)+txt
	cleanslap = txt+strings.Repeat(" ", Win.LenX)
	bk = colors["BK"]
	txt = colors["Text"]

	// make airline window
	ALW = MakeWin(
		"AirLine Window",
		stdout, stdin,
		Win.LenY-2, Win.LenY, 0, Win.LenX,
	)
}

func ClearAllAirLine() {
	AirLine(cleanslap)
	ErrorLine(slap)
}

// Airline
func AirLine ( s string ) {
	// make bkground color
	wuprint(ALW, 0, 0, bk)
	// write
	wuprint(ALW, 0, 0, s+txt)
}

func ClearAirLine() {
	wuprint(
		ALW, 1, 0, cleanslap,
	)
}

// warn
func Warn(warntype int) {
	ClearWarn()
	ErrorLine(ErrorText[warntype]+txt)
}

func ClearWarn () {
	AirLine(cleanslap)
	ErrorLine(slap)
}

func ErrorLine ( s string ) {
	wprint(ALW, 1, 0, s)
}

