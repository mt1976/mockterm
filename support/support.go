package support

import "time"

// GetTimeStamp returns the current date in the format "20060102".
func GetTimeStamp() string {
	return time.Now().Format("20060102")
}
