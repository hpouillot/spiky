package core

type Time int

func (t *Time) Before(time Time) bool {
	return *t < time
}

func (t *Time) Add(duration int) Time {
	return Time(int(*t) + duration)
}

func (t *Time) ToInt() int {
	return int(*t)
}

func (t *Time) ToFloat() float32 {
	return float32(*t)
}
