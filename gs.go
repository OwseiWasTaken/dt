//TODO(1) fix all: change all this to fit as included


// flag index
const (
	Clean = iota
	Utracked = iota
	uncommited = iota
	unpushed = iota
	unpulled = iota
	restore = iota
	FlagLen = iota
)


// flag texts
var (
	FlagTexts = []string{
		"working tree clean",
		"Untracked files",
		"Changes to be committed",
		"Your branch is ahead",
		"behind",
		"Changes not staged for commit",
	}
	FlagIcons = []string{
		"✓",
		"+",
		"→",
		"↑",
		"↓",
		"*",
	}
	flags = make([]bool, FlagLen)
)

func GetGs ( dir string ) (string) {
	//TODO(3) colors for gs: make gs use cfg/colors
	var (
		GSOut string
		branch string
  )

	cmd := exec.Command("git", "status")
	cmd.Dir = dir
	Out, err := cmd.Output()

	// not git directory
	if err != nil{
		// TODO(5): cfg/show no .git
		PS(err)
		return "[no .git]"
	}

	GSOut = string(Out)
	branch = strings.Split(GSOut, "\n")[0]
	branch = strings.Join(strings.Split(branch, " ")[2:], " ")

	for i:=0;i<FlagLen;i++ {
		flags[i] = strings.Contains(GSOut, FlagTexts[i])
	}


	if branch != "master" && branch != "main" {
		GSOut = branch+" "
	} else {
		GSOut = ""
	}

	if flags[0] {
		GSOut += RGB(60, 255, 60)
	} else if flags[3] {
		GSOut += RGB(60, 60, 255)
	} else {
		GSOut += RGB(255, 60, 60)
	}

	for i:=0;i<FlagLen;i++ {
		if flags[i] {
			GSOut += FlagIcons[i]
		}
	}

	return GSOut
}
