package hex

import "fmt"

func PrintBytes(b []byte) {
	if len(b) == 0 {
		return
	}
	result := FromByte(b[0])
	if len(b) > 1 {
		rest := b[1:]
		for _, v := range rest {
			result += ", " + FromByte(v)
		}
	}
	result = "[" + result + "]"
	fmt.Println(result)
}
