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

func taskModeFromString(str string) (taskMode, error) {
	return 0, nil
}
