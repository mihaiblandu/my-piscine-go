package piscine

func Sqrt(nb int) int {
	for i := nb / 2; i >= 1; i-- {
		tmp := i * i
		if tmp == nb {
			return i
		}
	}
	return 0
}
