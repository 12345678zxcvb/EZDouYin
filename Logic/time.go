package Logic

import "time"

func GetCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("200601021504")
}
