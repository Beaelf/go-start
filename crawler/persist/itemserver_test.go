package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go-start/crawler/engine"
	"go-start/crawler/model"
	"testing"
)

func TestItemServer(t *testing.T) {
	//engine.Item{
	//	Url:
	//}
	profile := model.Profile{Name: "ming", Email: "123@qq.com"}
	err := save(engine.Item{})
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("studygolang").Id("101").Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("resp: %s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if profile == actual {
		fmt.Println("equaled ...")
	}
}
