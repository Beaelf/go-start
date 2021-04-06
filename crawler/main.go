package main

import (
	"go-start/crawler/engine"
	"go-start/crawler/persist"
	"go-start/crawler/scheduler"
	"go-start/crawler/studygolang/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "https://studygolang.com/?tab=city",
	//	ParserFunc: parser.ParseCityList,
	//})
	saver, err := persist.ItemServer()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 10,
		ItemChan:  saver,
	}

	e.Run(engine.Request{
		Url:        "https://studygolang.com/?tab=city",
		ParserFunc: parser.ParseCityList,
	})
}
