package device

// LightsState represents raw lights state.
type LightsState struct {
	Colors       []byte `json:"colors"`
	Params       []byte `json:"effects"`
	CustomEffect []byte `json:"custom_effect"`
}

// KeysState represents raw keys state.
type KeysState struct {
	Mac    []byte `json:"mac"`
	Win    []byte `json:"win"`
	Macros []byte `json:"macros"`
}

// StateData represents raw state data.
type StateData struct {
	Lights *LightsState `json:"lights"`
	Keys   *KeysState   `json:"keys"`
}

// State represents raw device state. It contains data of all supported features
type State struct {
	Model Model     `json:"model"`
	Data  StateData `json:"state"`
}
