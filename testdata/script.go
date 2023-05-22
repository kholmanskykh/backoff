package main

import (
	"log"
	"time"

	"github.com/kholmanskykh/backoff"
)

func main() {
	b := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: true,
	}

	for i := 0; i < 10; i++ {
		d := b.Duration()
		log.Println(d)
	}

	log.Println("reset")
	b.Reset()
}
