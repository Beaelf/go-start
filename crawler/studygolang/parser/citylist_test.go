package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	 //content, err := fetcher.Fetch("https://studygolang.com/?tab=city")
	content, err := ioutil.ReadFile("studygolang.html")
	if(err != nil){
		 panic(err)
	 }
	 fmt.Printf("%s\n", content)

	//list := ParseCityList(content)





}
