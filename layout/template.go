package layout

// Template represents keyboard template
type Template map[KeyName]uint32

// GetPosition returns key position
func (k Template) GetPosition(key KeyName) uint32 {
	return k[key]
}

// GetKey returns key name by position
func (k Template) GetKey(position uint32) KeyName {
	for k, v := range k {
		if v == position {
			return k
		}
	}
	return ""
}
