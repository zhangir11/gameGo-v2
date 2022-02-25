package game

const (
	a string = "floor_1"
	b string = "floor_2"
	c string = "floor_3"
	d string = "floor_4"
)

func LoadMap() [][]string {

	level := [][]string{
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, b, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, c, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, c, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, d, a, a, a, a, a, a, a},
		{a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a, a},
	}

	return level
}
