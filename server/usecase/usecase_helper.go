package usecase

// ManageLimit は、Limitを制御する。
func ManageLimit(targetLimit, maxLimit, minLimit, defaultLimit int) int {
	if targetLimit < minLimit || maxLimit < targetLimit {
		return defaultLimit
	}
	return targetLimit
}
