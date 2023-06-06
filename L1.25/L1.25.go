package main

import (
	"time"
)

// After  - возвращает канал, в который по истечению времени отправляет текущее время
func MySleep(duration time.Duration) {
	timer := time.After(duration)
	<-timer
}

func main() {
	MySleep(time.Second)
}
