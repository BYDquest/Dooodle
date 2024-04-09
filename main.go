package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/BYDquest/RRRDoodlers/utility")

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialize the global pseudo-random generator
	fmt.Println(utility.RandomFromInterval(1.0, 2.0))
}
