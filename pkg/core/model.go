package core

type Model interface {
	Train(duration int)
	Run(duration int)
}
