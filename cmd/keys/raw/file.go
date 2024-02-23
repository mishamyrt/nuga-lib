// Package raw provides helpers for working with raw keys layout
package raw

type keysDump struct {
	Mac []uint32 `json:"mac"`
	Win []uint32 `json:"win"`
}
