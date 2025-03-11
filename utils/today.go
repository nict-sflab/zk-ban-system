package utils

import (
	"fmt"
	"time"
)

func Today() int64 {
	t := time.Now()
	day := time.Hour * 24

	today := t.UnixNano() / int64(day)
	fmt.Printf("today: %d\n", today)
	return today
}
