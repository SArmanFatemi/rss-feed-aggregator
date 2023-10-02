package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RssFeed, error) {
	httClient := http.Client{
		Timeout: 10 * time.Second,
	}
	rssFeed := RssFeed{}

	resp, err := httClient.Get(url)
	if err != nil {
		return rssFeed, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssFeed, err
	}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return rssFeed, err
	}

	return rssFeed, nil
}
