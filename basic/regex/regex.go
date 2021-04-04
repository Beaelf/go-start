package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	fmt.Println(match)
}
