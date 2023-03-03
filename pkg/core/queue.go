package core

type Queue interface {
	Add(time Time, node Node)
	Count() int
	Pop() (Time, Node)
}
