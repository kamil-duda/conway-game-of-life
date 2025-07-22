package conway

type IsLive = bool
type Grid = [][]IsLive

type Universe struct {
	grid Grid
}

func New() *Universe {
	return new(Universe)
}

// todo: parallel update
func (universe *Universe) Update() {

}
