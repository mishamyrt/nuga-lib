// Package device provides supported device models
package device

// Model represents supported keyboard model
type Model string

const (
	// Halo75 represents NuPhy Halo 75 keyboard name
	Halo75 Model = "Halo75"
	// Halo65 represents NuPhy Halo 65 keyboard name
	Halo65 Model = "Halo65"
	// Halo96 represents NuPhy Halo 96 keyboard name
	Halo96 Model = "Halo96"
	// Unknown represents unknown keyboard name
	Unknown Model = ""
)

// IsSupported returns true if model is supported by Nuga library
func IsSupported(m Model) bool {
	return m == Halo75 || m == Halo65 || m == Halo96
}
