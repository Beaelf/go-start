package main

import (
	"go-start/crawler/engine"
	"go-start/crawler/scheduler"
	"go-start/crawler/studygolang/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "https://studygolang.com/?tab=city",
	//	ParserFunc: parser.ParseCityList,
	//})
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 10,
	}

	e.Run(engine.Request{
		Url:        "https://studygolang.com/?tab=city",
		ParserFunc: parser.ParseCityList,
	})
}
