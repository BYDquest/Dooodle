package doodle


type Point struct {
	X float64
	Y float64
}


type ColorScheme struct {
	FaceColor        string
	BackgroundColor  string
	HairColor        string
	MouthColor       string
}


type Eye struct {
	UpperContour   [][]float64
	LowerContour   [][]float64
	WhiteContour   [][]float64
	Center         [2]float64
	OffsetX        float64
	OffsetY        float64
	PupilShiftX    float64
	PupilShiftY    float64
}

type Nose struct {
	RightX float64
	RightY float64
	LeftX  float64
	LeftY  float64
}

type Facial struct {
	FaceScale           float64
	FacePoints  [][]float64
	FaceHeight          float64
	FaceWidth           float64
	Center              [2]float64
	DistanceBetweenEyes float64
	EyeHeightOffset     float64
	LeftEye             Eye
	RightEye            Eye
	Nose                Nose
	Hairs               [][]float64
	MouthPoints         [][]float64
}

type Face struct {
	FacePoints  [][]float64
	FaceHeight          float64
	FaceWidth           float64
	Center              [2]float64
}

