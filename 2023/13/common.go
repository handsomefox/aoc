package main

func parseCols(rows []string) []string {
	cols := []string{}
	if len(rows) == 0 {
		return cols
	}
	for chIndex := 0; chIndex < len(rows[0]); chIndex++ {
		col := ""
		for row := 0; row < len(rows); row++ {
			col += string(rows[row][chIndex])
		}
		cols = append(cols, col)
	}

	return cols
}

func SliceCopy[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}
