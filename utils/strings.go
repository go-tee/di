package utils

func FirstRune(str string) (first rune) {
	for _, c := range str {
		first = c
		break
	}

	return
}
