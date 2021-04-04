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
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        "https://studygolang.com" + string(m[1]),
				ParserFunc: ParseProfile,
			})
	}

	return result
}
