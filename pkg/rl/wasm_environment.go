package rl

import "syscall/js"

// Wrapper around js environment
type WasmEnvironment struct {
	jsEnv js.Value
}

func (env *WasmEnvironment) Register() int {
	return env.jsEnv.Call("register").Int()
}

func (env *WasmEnvironment) Observe(agentId int) []float64 {
	jsState := env.jsEnv.Call("observe", agentId)
	goState := make([]uint8, jsState.Get("byteLength").Int())
	js.CopyBytesToGo(goState, jsState)
	float64State := make([]float64, len(goState))
	for i, v := range goState {
		float64State[i] = float64(v)
	}
	return float64State
}

func (env *WasmEnvironment) Perform(agentId int, action int) float64 {
	reward := env.jsEnv.Call("perform", agentId, action).Float()
	return reward
}

func NewWasmEnvironment(jsEnv js.Value) *WasmEnvironment {
	env := &WasmEnvironment{
		jsEnv: jsEnv,
	}
	return env
}
