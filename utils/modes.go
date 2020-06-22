package utils

func DbMode(mode byte) string {
	switch mode {
	case 0:
		return "std"
	case 1:
		return "taiko"
	case 2:
		return "ctb"
	case 3:
		return "mania"
	}
	return "std"
}
