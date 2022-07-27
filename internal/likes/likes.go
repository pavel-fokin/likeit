package likes

type Likes struct {}

func New() *Likes {
	return &Likes{}
}

func (l *Likes) Count() {}
func (l *Likes) Increment() {}
