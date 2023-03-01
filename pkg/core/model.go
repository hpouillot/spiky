package core

type Model struct {
	Inputs  []Node
	Outputs []Node
}

func (m *Model) Run(duration int) {
	queue := NewQueue()
	time := Time(0)
	end_time := Time(duration)
	for _, input := range m.Inputs {
		queue.Add(time, input)
	}
	for queue.GetCount() != 0 && time < end_time {
		time, node := queue.PopMin()
		node.Compute(time, queue)
	}
}
