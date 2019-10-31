package retry

import (
	"fmt"
	"log"
	"time"
)

func Retry(maxAttempts int, sleep time.Duration, function func() error) (err error) {
	for currentAttempt := 0; currentAttempt < maxAttempts; currentAttempt++ {
		err = function()
		if err == nil {
			return
		}
		for i := 0; i <= currentAttempt; i++ {
			time.Sleep(sleep)
		}
		log.Println("Retrying after error:", err)
	}
	return fmt.Errorf(fmt.Sprintf("After %d attempts, last error: %s", maxAttempts, err))
}
