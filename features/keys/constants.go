package keys

var (
	cmdGetMacKeys = []byte{0x05, 0x84, 0xd8, 0x00, 0x00, 0x00}
	cmdGetWinKeys = []byte{0x05, 0x84, 0xd4, 0x00, 0x00, 0x00}
	cmdSetMacKeys = []byte{0x06, 0x04, 0xd8, 0x00, 0x40, 0x00, 0x00, 0x00}
	cmdSetWinKeys = []byte{0x06, 0x04, 0xd4, 0x00, 0x40, 0x00, 0x00, 0x00}
)
