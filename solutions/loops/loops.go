package loops

func compute(begin, end int) int {
	for i := begin; i < end; i++ {
		if i%7 == 0 && i%5 == 0 {
			return i
		}
	}
	return -1
}
