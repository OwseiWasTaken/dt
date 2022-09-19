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
	// define colors
	airline = colors["AirLine"]
	BadError = colors["BadError"]
	SimpleError = colors["SimpleError"]
	slap = bk+sws+txt
	cleanslap = txt+sws
	bk = colors["BK"]
	txt = colors["Text"]

	// define errors
	ErrorText = []string{
		BadError+"No file from link",
		SimpleError+"Not a link",
		SimpleError+"Command Empty",
		BadError+"No Such Command",
	}

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
	wuprint(ALW, 0, 0, bk+cleanslap)
	// write
	wuprint(ALW, 0, 0, s+txt)
}

func ClearAirLine() {
	AirLine(cleanslap)
}

// warn
func Warn(warntype int) {
	ClearWarn()
	ErrorLine(ErrorText[warntype]+txt)
}

func ClearWarn () {
	ErrorLine(slap)
}

func ErrorLine ( s string ) {
	wprint(ALW, 1, 0, s)
}

