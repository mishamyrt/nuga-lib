package keys

import (
	"testing"
)

var rawBytes = []byte{
	0x02, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x35, 0x00, 0x00, 0x00, 0x2b, 0x00, 0x00, 0x00, 0x39, 0x06, 0x00, 0x00, 0xe1, 0x06, 0x00, 0x00, 0xe0, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x1d, 0x06, 0x00, 0x00, 0xe2, 0x02, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x16, 0x02, 0x00, 0x00, 0x1d, 0x06, 0x00, 0x00, 0xe3, 0x02, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x07, 0x02, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x91, 0x02, 0x00, 0x00, 0x03, 0x02, 0x00, 0x00, 0x23, 0x00, 0x00, 0x00, 0x15, 0x02, 0x00, 0x00, 0x2c, 0x02, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x8b, 0x02, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x17, 0x02, 0x00, 0x00, 0x0d, 0x02, 0x00, 0x00, 0x25, 0x00, 0x00, 0x00, 0x2c, 0x02, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x23, 0x00, 0x00, 0x00, 0x1c, 0x02, 0x00, 0x00, 0x24, 0x02, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x8a, 0x02, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x90, 0x02, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x25, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x0e, 0x02, 0x00, 0x00, 0x27, 0x06, 0x00, 0x00, 0xe7, 0x02, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x0f, 0x02, 0x00, 0x00, 0x28, 0x20, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x27, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x33, 0x02, 0x00, 0x00, 0x29, 0x06, 0x00, 0x00, 0xe4, 0x02, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x2f, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x87, 0x00, 0x00, 0x00, 0x88, 0x02, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x2e, 0x02, 0x00, 0x00, 0x2f, 0x02, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x64, 0x06, 0x00, 0x00, 0xe7, 0x02, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x2a, 0x02, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x28, 0x06, 0x00, 0x00, 0xe5, 0x02, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x10, 0x02, 0x00, 0x00, 0x11, 0x00, 0x00, 0x08, 0x0c, 0x00, 0x00, 0x00, 0x4a, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x00, 0x00, 0x4e, 0x02, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x00, 0x00, 0x5a, 0x00, 0x00, 0x00, 0x62, 0x00, 0x00, 0x00, 0x4b, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x61, 0x00, 0x00, 0x00, 0x5e, 0x00, 0x00, 0x00, 0x5b, 0x00, 0x00, 0x00, 0x63, 0x00, 0x00, 0x00, 0x4e, 0x00, 0x00, 0x00, 0x56, 0x00, 0x00, 0x00, 0x57, 0x00, 0x00, 0x00, 0x85, 0x00, 0x00, 0x00, 0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x00, 0x00, 0x00, 0x61, 0x00, 0x00, 0x00, 0x5e, 0x00, 0x00, 0x00, 0x5b, 0x00, 0x00, 0x00, 0x63, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x56, 0x00, 0x00, 0x00, 0x57, 0x00, 0x00, 0x00, 0x85, 0x00, 0x00, 0x00, 0x58, 0x00, 0x00, 0x00, 0x00, 0x48, 0x61, 0x6c, 0x6f, 0x37, 0x35, 0x20, 0x33, 0x2e, 0x30, 0x4e, 0x75, 0x50, 0x68, 0x79, 0x20, 0x48, 0x61, 0x6c, 0x6f, 0x37, 0x35, 0x20, 0x00, 0x00, 0x00, 0x00, 0x3a, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x00, 0x3d, 0x04, 0x00, 0x00, 0xcf, 0x0e, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x43, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x45, 0x06, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x4a, 0x00, 0x00, 0x00, 0x29, 0x00, 0x00, 0x00, 0x52, 0x00, 0x00, 0x00, 0x51, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x4f, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x0a, 0x21, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x15, 0x00, 0x00, 0x00, 0x1b, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x00, 0x37, 0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x2e, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x31, 0x00, 0x00, 0x00, 0x2f, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x3a, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x3c, 0x00, 0x00, 0x00, 0x3d, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x43, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x45, 0x0e, 0x00, 0x00, 0x01, 0x0e, 0x00, 0x00, 0x10, 0x0e, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x35, 0x0c, 0x00, 0x01, 0x00, 0x0c, 0x00, 0x02, 0x00, 0x0b, 0x00, 0x03, 0x00, 0x12, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x16, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x0a, 0x20, 0x0e, 0x00, 0x00, 0x0c, 0x0e, 0x00, 0x00, 0x0d, 0x0e, 0x00, 0x00, 0x0e, 0x0e, 0x00, 0x00, 0x0f, 0x0e, 0x00, 0x00, 0x14, 0x0e, 0x00, 0x00, 0x1c, 0x0e, 0x00, 0x00, 0x1d, 0x0e, 0x00, 0x00, 0x0c, 0x0e, 0x00, 0x00, 0x0d, 0x0e, 0x00, 0x00, 0x0e, 0x0e, 0x00, 0x00, 0x0f, 0x0e, 0x00, 0x00, 0x18, 0x0e, 0x00, 0x00, 0x05, 0x0e, 0x00, 0x00, 0x1e, 0x0d, 0x00, 0x02, 0x00, 0x0d, 0x00, 0x01, 0x00, 0x0e, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x00, 0x2e, 0x0e, 0x00, 0x00, 0x1b, 0x0e, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x00, 0x2f, 0x0e, 0x00, 0x00, 0x19, 0x90, 0x0f, 0x07, 0x74, 0x11, 0xf0, 0xa3, 0x74, 0x08, 0xf0, 0x90, 0x0e, 0xfd, 0xe0, 0xff, 0x90, 0x0f, 0x07, 0xe0, 0xfc, 0xa3, 0xe0, 0xf5, 0x82, 0x8c, 0x83, 0xef, 0xf0, 0x53, 0x9b, 0xf0, 0x75, 0x9b, 0x01, 0x43, 0x97, 0x04, 0x22, 0x90, 0x0f, 0x06, 0x74, 0x02, 0xf0, 0x53, 0x9b, 0xf0, 0x75, 0x9b, 0x01, 0xa3, 0x74, 0x11, 0xf0, 0xa3, 0x74, 0x08, 0xf0, 0x90, 0x0f, 0x07, 0xe0, 0xfe, 0xa3, 0xe0, 0xf5, 0x82, 0x8e, 0x83, 0xe4, 0xf0, 0x43, 0x97, 0x04, 0x22, 0x12, 0xd0, 0xc6, 0x12, 0x00, 0x16, 0x7f, 0xf4, 0x7e, 0x01, 0x12, 0xb0, 0xc1, 0x12, 0xad, 0xf2, 0x12, 0xd2, 0x2b, 0x12, 0x7e, 0xd4, 0x12, 0xd2, 0xd1, 0x12, 0xd2, 0x73, 0x12, 0xd2, 0xd7, 0x02, 0xd1, 0x71, 0xd3, 0x90, 0x0e,
}

var keyCodes = []uint32{
	0xf000002, 0x35000000, 0x2b000000, 0x39000000, 0xe1000006, 0xe0000006, 0x02, 0x20000002, 0x14000000, 0x4000000, 0x1d000000, 0xe2000006, 0x1000002, 0x21000002, 0x1a000000, 0x16000000, 0x1d000002, 0xe3000006, 0x2000002, 0x22000002, 0x8000000, 0x7000000, 0x1e000002, 0x91000000, 0x3000002, 0x23000002, 0x15000000, 0x2c000002, 0x1f000002, 0x8b000000, 0x4000002, 0x22000000, 0x17000000, 0xd000002, 0x25000002, 0x2c000000, 0x5000002, 0x23000000, 0x1c000000, 0x24000002, 0x26000002, 0x8a000000, 0x6000002, 0x24000000, 0x18000000, 0xd000000, 0x10000000, 0x90000000, 0x7000002, 0x25000000, 0xc000000, 0xe000000, 0x27000002, 0xe7000006, 0x8000002, 0x26000000, 0x12000000, 0xf000000, 0x28000002, 0x20, 0x9000002, 0x27000000, 0x13000000, 0x33000000, 0x29000002, 0xe4000006, 0xa000002, 0x2d000000, 0x2f000000, 0x34000000, 0x87000000, 0x88000000, 0xb000002, 0x2e000000, 0x2f000002, 0x2d000002, 0x64000000, 0xe7000006, 0x18000002, 0x2a000000, 0x2d000002, 0x28000000, 0xe5000006, 0x12000002, 0x4c000000, 0x00, 0x00, 0x00, 0x10000002, 0x11000002, 0xc080000, 0x4a000000, 0x4d000000, 0x4b000000, 0x4e000000, 0x13000002, 0x4d000000, 0x54000000, 0x60000000, 0x5d000000, 0x5a000000, 0x62000000, 0x4b000000, 0x55000000, 0x61000000, 0x5e000000, 0x5b000000, 0x63000000, 0x4e000000, 0x56000000, 0x57000000, 0x85000000, 0x58000000, 0x00, 0x00, 0x55000000, 0x61000000, 0x5e000000, 0x5b000000, 0x63000000, 0x00, 0x56000000, 0x57000000, 0x85000000, 0x58000000, 0x00, 0x6f6c6148, 0x33203537, 0x754e302e, 0x20796850, 0x6f6c6148, 0x203537, 0x3a000000, 0x3b000000, 0x3c000000, 0x3d000000, 0xcf000004, 0x2000000e, 0x40000000, 0x41000000, 0x42000000, 0x43000000, 0x44000000, 0x45000000, 0xe3000006, 0xa000000, 0x4a000000, 0x29000000, 0x52000000, 0x51000000, 0x50000000, 0x4f000000, 0x1a000000, 0x16000000, 0x4000000, 0x7000000, 0x210a0000, 0x14000000, 0x1a000000, 0x8000000, 0x15000000, 0x1b000000, 0x6000000, 0x19000000, 0x1e000000, 0x1f000000, 0x20000000, 0x21000000, 0xb000000, 0x5000000, 0x11000000, 0x36000000, 0x37000000, 0x38000000, 0x2d000000, 0x2e000000, 0x9000000, 0x31000000, 0x2f000000, 0x30000000, 0x3a000000, 0x3b000000, 0x3c000000, 0x3d000000, 0x3e000000, 0x3f000000, 0x40000000, 0x41000000, 0x42000000, 0x43000000, 0x44000000, 0x45000000, 0x100000e, 0x1000000e, 0x400000e, 0x35000000, 0x1000c, 0x2000c, 0x3000b, 0x30012, 0x1a000000, 0x16000000, 0x1000001, 0x7000000, 0x200a0000, 0xc00000e, 0xd00000e, 0xe00000e, 0xf00000e, 0x1400000e, 0x1c00000e, 0x1d00000e, 0xc00000e, 0xd00000e, 0xe00000e, 0xf00000e, 0x1800000e, 0x500000e, 0x1e00000e, 0x2000d, 0x1000d, 0x1700000e, 0x2d000000, 0x2e000000, 0x1b00000e, 0x1a00000e, 0x2f000000, 0x1900000e, 0x74070f90, 0x74a3f011, 0xe90f008, 0x90ffe0fd, 0xfce0070f, 0x82f5e0a3, 0xf0ef838c, 0x75f09b53, 0x9743019b, 0xf902204, 0xf0027406, 0x75f09b53, 0x74a3019b, 0x74a3f011, 0xf90f008, 0xa3fee007, 0x8e82f5e0, 0x43f0e483, 0x12220497, 0x12c6d0, 0x7ef47f16, 0xc1b01201, 0x12f2ad12, 0x7e122bd2, 0xd1d212d4, 0x1273d212, 0xd102d7d2, 0xe90d371,
}

func TestPackKeyCodes(t *testing.T) {
	t.Parallel()
	result := PackKeyCodes(rawBytes)
	// Check some random values
	for i, v := range keyCodes {
		if result[i] != v {
			t.Errorf("result[%d] = %02x, want %02x", i, result[i], v)
		}
	}
}

func TestUnpackKeyCodes(t *testing.T) {
	t.Parallel()
	result := UnpackKeyCodes(keyCodes)
	for i, v := range rawBytes {
		if result[i] != v {
			t.Errorf("result[%d] = %02x, want %02x", i, result[i], v)
		}
	}
}
