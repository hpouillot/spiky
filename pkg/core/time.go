package core

type Time float64

func (t *Time) Before(time Time) bool {
	return *t < time
}

func (t *Time) Add(duration Time) Time {
	return *t + duration
}

func (t *Time) ToInt() int {
	return int(*t)
}

func (t *Time) ToFloat() float64 {
	return float64(*t)
}
