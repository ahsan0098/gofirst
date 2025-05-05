package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ahsan0098/gofirst/internal/database"
)

func startScraping(db *database.Queries, limit int, timeBetween time.Duration) {

	ticker := time.NewTicker(timeBetween)

	for ; ; <-ticker.C {

		feeds, err := db.GetNextFeedsToFetched(context.Background(), int32(limit))

		if err != nil {
			log.Printf("error fetching feeds: %s", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()

		log.Printf("scrapping started on go routine %s duration", timeBetween)
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedASFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error updating feed as mark: %s", err)
		return
	}

	log.Printf("updated feed '%s' as last fetched", feed.Name)
}
