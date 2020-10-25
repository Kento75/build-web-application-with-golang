package sec2

import (
	"file"
)

func ReadWrite() bool {
	file.Open("file")

	defer file.Close()

	if isHogeHoge {
		return false
	}

	if isFugaFuga {
		return false
	}

	return true
}

func lifo {
	// deferはLIFOモード。出力は4 3 2 1 0となる
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}