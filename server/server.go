package main

import (
	"fmt"
	"strings"
)

const (
	text1 = `
a,b,c
1,2,3`
	text2 = `
a,b,c
1,2,3
4,5,6
asdasd
asdasd
asdasdas
asdasd
`
)

func main() {
	fmt.Println(diff(text1, text2))
}

func diff(text1, text2 string)[]string{
	a := strings.Split(text1, "\n")
	amap := make(map[string]bool)
	b := strings.Split(text2, "\n")
	bmap := make(map[string]bool)
	for _, a2 := range a{
		amap[a2] = true
	}
	for _, b2 := range b{
		bmap[b2] = true
	}
	final := []string{}
	for key, _ := range bmap{
		if amap[key] == false{
			final = append(final, key)
		}
	}
	return final
}