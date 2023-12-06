package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultPuppyInstanceCost is initially set the same as in wasmd
	DefaultPuppyInstanceCost uint64 = 60_000
	// DefaultPuppyCompileCost set to a large number for testing
	DefaultPuppyCompileCost uint64 = 3
)

// PuppyGasRegisterConfig is defaults plus a custom compile amount
func PuppyGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultPuppyInstanceCost
	gasConfig.CompileCost = DefaultPuppyCompileCost

	return gasConfig
}

func NewPuppyWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(PuppyGasRegisterConfig())
}
