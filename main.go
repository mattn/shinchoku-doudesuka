package main

import (
	"fmt"
	"math/rand"
	"time"
)

type word int

const (
	進捗 word = iota
	どう
	です
	か
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0
	var w word
	for {
		n++
		s := word(rand.Intn(4))
		fmt.Print(s)
		if w != s {
			if s == 0 {
				w = 1
			} else {
				w = 0
			}
		} else {
			w++
			if w == 4 {
				break
			}
		}
	}
	fmt.Println(`
 ＿人人人人人人人＿
 ＞進捗どうですか＜
 ￣Y^Y^Y^Y^Y^Y^Y￣`)
	fmt.Printf("%d回で煽られました\n", n)
}
