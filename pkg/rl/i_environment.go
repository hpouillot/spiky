package rl

type IEnvironment interface {
	Register() int
	Observe() []float64
	Perform(agentId int, action int) float64
}
