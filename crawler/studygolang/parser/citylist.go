package parser

import (
	"go-start/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(/go/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        "https://studygolang.com" + string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
