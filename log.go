package main

import (
	"fmt"
	"github.com/postelniql/logger-output/hash"
	"time"
)

func log() string {
	dt := time.Now()
	hash := hash.NewSHA1Hash()
	return dt.Format("01-02-2006T15:04:05.000Z") + ": " + hash
}

func main() {
	fmt.Println(log())

	tick := time.Tick(5000 * time.Millisecond)
	for range tick {
		fmt.Println(log())
	}
}
