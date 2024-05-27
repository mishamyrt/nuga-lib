package cli

// Args represents CLI arguments
type Args []*string

// UnpackArgs unpacks CLI arguments into provided Args
func UnpackArgs(args []string, targets Args) {
	for i, value := range targets {
		*value = args[i]
	}
}
