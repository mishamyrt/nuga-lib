// Package bit contains utils for bit masks
package bit

// Mask represents bit mask entity
type Mask uint32

type predicate func(f Mask) bool

func (m *Mask) assert(features []Mask, assert predicate, defaultExit bool) bool {
	for i := range features {
		if assert(features[i]) {
			return !defaultExit
		}
	}
	return defaultExit
}

// IsSet checks if bits is set.
func (m Mask) IsSet(bits ...Mask) bool {
	return m.assert(bits, func(feature Mask) bool {
		return m&feature == 0
	}, true)
}

// IsAnySet checks if any of bits is set.
func (m Mask) IsAnySet(bits ...Mask) bool {
	return m.assert(bits, func(feature Mask) bool {
		return m&feature != 0
	}, false)
}
