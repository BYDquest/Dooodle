package main

import (
	"fmt"
	"github.com/BYDquest/RRRDoodlers/doodle"

	
)

func main() {
    // Example of usage
    face := doodle.GenerateFaceContourPoints(100)
	
	fmt.Println(face.FacePoints,face.FaceHeight, face.FaceWidth,face.Center )
}

