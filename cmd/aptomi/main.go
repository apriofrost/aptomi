package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := aptomiCmd.Execute(); err != nil {
		panic(fmt.Errorf("error while executing command: %s", err))
	}
}
