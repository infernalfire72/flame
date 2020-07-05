package achievements

type Achievement struct {
	ID          int    `db:"id"`
	Icon        string `db:"icon"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

// String concat is ugly but fast
func (a Achievement) String() string {
	return a.Icon + "+" + a.Name + "+" + a.Description
}
