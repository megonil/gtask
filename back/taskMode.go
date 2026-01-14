package back

type taskMode int

const (
	modePlanned taskMode = iota
	modeInProgress
)

func (m taskMode) String() string {
	switch m {
	case modePlanned:
		return "planned"
	case modeInProgress:
		return "in progress"
	default:
		return "unknown"
	}
}

func TaskModeFromString(str string) (taskMode, error) {
	switch {
	case checkString(str, "planned"):
		return modePlanned, nil
	case checkString(str, "in-progress"):
		return modeInProgress, nil
	}

	return 0, nil
}
