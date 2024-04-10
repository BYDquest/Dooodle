package main

import (
	"fmt"
	"github.com/BYDquest/Dooodle/doodle"
)

func main() {
	// Example of usage
	face := doodle.GenerateFaceContourPoints(100)

	fmt.Println(face.Center)
}
