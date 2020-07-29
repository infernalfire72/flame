package objects

type Target interface {
	GetName() string
	Write(...[]byte)
}
