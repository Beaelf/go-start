package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	content, _ := ioutil.ReadFile("profile.html")
	fmt.Println(content)
	profile := ParseProfile(content)
	fmt.Print("hhh: ",profile)

}
