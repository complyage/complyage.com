package abstract

//||------------------------------------------------------------------------------------------------||
//|| Helpers
//||------------------------------------------------------------------------------------------------||

func KeyLevelToWordCount(level int) int {
	switch level {
	case 2:
		return 6
	case 3:
		return 12
	case 4:
		return 18
	case 5:
		return 24
	}
	return 0
}
