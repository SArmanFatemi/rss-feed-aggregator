package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/sarmanfatemi/rssagg/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)

	/*
		If we are using a for loop without like:
		for range ticker.C {
		}
		It does not trigger immediately, it waits for the first tick to happen.
	*/
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error while fetching feeds:", err)
			continue
		}

		if len(feeds) == 0 {
			log.Println("No feeds to fetch")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkedFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error while marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error while fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		log.Println("Found post", item.Title, "on feed", feed.Name)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
