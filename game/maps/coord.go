package maps

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// CreateCoordinate
func CreateCoordinate(x, y int) *Coordinate {
	var co Coordinate
	co.X = float64(x)
	co.Y = float64(y)
	return &co
}

//GetCoordinate return X, Y Coordinate
func (co *Coordinate) GetCoordinate() (float64, float64) {
	return co.X, co.Y
}
