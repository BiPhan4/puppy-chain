package app

import (
	"encoding/json"
)

// GenesisState: The genesis state of the blockchain is represented here as a map of raw json
// messages keyd by a identifier string.
// Within this application default genesis information is retrieved from the
// ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func newDefaultGenesisState() GenesisState {
	encodingConfig := MakeEncodingConfig()
	return ModuleBasics.DefaultGenesis(encodingConfig.Marshaler)

}
