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
	AirlineText = colors["AirlineText"]
	bk = colors["AirlineText"]
	slap = bk+sws+txt
	cleanslap = txt+sws

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
	AirLine(slap)
	ErrorLine(slap)
}

// Airline
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
	ClearWarn()
	ErrorLine(ErrorText[warntype]+txt)
}

func ClearWarn () {
	ErrorLine(cleanslap)
}

func ErrorLine ( s string ) {
	wuprint(ALW, 1, 0, s)
}

