package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		zundoko := generate_zundoko()
		fmt.Println(zundoko)
		if zundoko == "ズンズンズンズンドコ" {
			fmt.Println("キヨシ!")
			break
		}
	}
}

func generate_zundoko() string {
	var result = ""

	for i := 0; i < 5; i++ {
		if rand.Intn(2) == 0 {
			result += "ズン"
		} else {
			result += "ドコ"
		}
	}

	return result
}

