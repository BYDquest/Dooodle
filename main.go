package main

import (
	"fmt"
	"github.com/BYDquest/Dooodle/doodle"
	"github.com/BYDquest/Dooodle/lib"
)

func main() {

	face := doodle.GenerateFaceContourPoints(100)
	fmt.Println(face.Center)

	leftEye, rightEye := doodle.GenerateBothEyes(50)
	fmt.Println("Left Eye Upper Points:", leftEye.Upper)
	fmt.Println("Right Eye Upper Points:", rightEye.Upper)
	lib.Svg2png()
}
