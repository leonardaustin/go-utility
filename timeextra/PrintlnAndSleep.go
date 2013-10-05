package timeextra

import (
	"log"
	"time"
)

func PrintlnAndSleep(userSeconds int64, errorCount *int64) int64 {
	log.Println("Sleep for", userSeconds, "seconds and try again. Count =",*errorCount)
	time.Sleep(time.Duration(userSeconds) * time.Second)
	*errorCount++

	return *errorCount
}