package parser

import (
	"go-start/crawler/engine"
	"go-start/crawler/model"
	"regexp"
)

const nameReg = `<li>\n.*<label>名字:</label>\n.*<span>(.+)</span>\n.*</li>`
const emailReg = `<li>\n.*<label>Email:</label>\n.*<span><a href=".*">(.+)</a></span>\n.*</li>`

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = extractString(contents, nameReg)
	profile.Email = extractString(contents, emailReg)
	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(contents []byte, re string) string {
	reg := regexp.MustCompile(re)
	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
