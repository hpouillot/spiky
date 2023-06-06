package main

import (
	"encoding/json"
	"math/rand"
	"spiky/pkg/core"
	"spiky/pkg/rl"
	"syscall/js"
)

type SpikyOptions struct {
	Layers      []int             `json:"layers"`
	Exploration float64           `json:"exploration"`
	Config      *core.ModelConfig `json:"config"`
}

func instantiate(this js.Value, args []js.Value) interface{} {
	env := rl.NewWasmEnvironment(args[0])
	optionsStr := args[1].String()

	options := SpikyOptions{}
	err := json.Unmarshal([]byte(optionsStr), &options)
	if err != nil {
		panic(err)
	}

	agent := core.BuildSequential(options.Layers, options.Config)
	agentId := env.Register()

	predict := func(state []float64) {
		agent.Encode(state)
		agent.Run()
		if rand.Float64() > options.Exploration {
			randomIdx := rand.Intn(agent.GetOutput().Size())
			agent.GetOutput().Visit(func(idx int, n *core.Neuron) {
				n.Reset()
				if randomIdx == idx {
					n.SetSpikeTime(agent.World, agent.World.Const.MaxTime)
				}
			})
		}
	}

	perform := func() {
		action := agent.DecodeClass()
		reward := env.Perform(agentId, action)
		agent.Stdp(reward)
		agent.Reset()
	}

	// Build agent
	return js.ValueOf(map[string]interface{}{
		"next": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			state := env.Observe(agentId)
			predict(state)
			perform()
			return true
		}),
	})
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("Spiky", js.FuncOf(instantiate))
	<-done
}
