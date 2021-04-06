package persist

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"go-start/crawler/engine"
	"log"
)

func ItemServer() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Server: got item %d: %v", itemCount, item)
			itemCount++

			err = save(client, item)
			if err != nil {
				log.Printf("Item Server: error saving item %v: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item) error {
	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
