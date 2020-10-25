package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func panicTest() {
	if user != "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	// panicが発生した場合、復元を行う
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()

	// 関数fを実行する
	f()

	return
}

func main() {
	result := throwsPanic(panicTest)
	fmt.Println(result)
}
