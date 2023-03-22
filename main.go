package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "台湾九合abc一选举"
	ss := []rune(s)
	fmt.Println(len(ss))

	sss := utf8.RuneCountInString(s)
	fmt.Println(sss)

	char := "♥"
	ssss := []rune(char)
	fmt.Println(len(ssss))
	fmt.Println(len(char))
}
