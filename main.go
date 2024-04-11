package main

import (
	"fmt"
	"github.com/BYDquest/Dooodle/doodle"
)

func main() {
	// Example of usage
	face := doodle.GenerateFaceContourPoints(100)
	fmt.Println(face.Center)


	leftEye, rightEye := doodle.GenerateBothEyes(50)
	fmt.Println("Left Eye Upper Points:", leftEye.Upper)
	fmt.Println("Right Eye Upper Points:", rightEye.Upper)
}
