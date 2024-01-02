package hid

// Handler represents hid request handler
type Handler interface {
	Send(payload []byte) error
	Read(count int) ([]byte, error)
	Request(payload []byte, count int) ([]byte, error)
	GetInfo() *DeviceInfo
	Close() error
}
