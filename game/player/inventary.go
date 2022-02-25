package player

type Inventary struct {
	Slots []*Slot
}

type Slot struct {
	Name string `json:"name"`
	Sum  int    `json:"sum"`
}
