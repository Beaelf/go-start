package parser

import (
	"go-start/crawler/engine"
	"regexp"
)

const userRe = `<a href="(/user/[a-zA-Z0-9]+)".*title="([^>]*)">`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(userRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := ""
		//result.Items = append(result.Items, m[2])
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: "https://studygolang.com" + string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, url)
				},
			})
	}

	return result
}
