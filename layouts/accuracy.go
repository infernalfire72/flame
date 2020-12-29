package layouts

type Accuracy struct {
	N300 int16
	N100 int16
	N50  int16
	NGeki int16
	NKatu int16
	NMiss int16
}

func (a *Accuracy) Value() float32 {
	return 0
}