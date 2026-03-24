package main

import (
	"fmt"
	"net/http"
	"context"
	"io"
	"encoding/xml"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
    if err != nil {
        return nil, err
    }

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
    	return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rssfeed RSSFeed
	err = xml.Unmarshal([]byte(b), &rssfeed)
	if err != nil {
		fmt.Printf("error: %v", err)
		return &RSSFeed{}, err
	}

	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)

	for i := range rssfeed.Channel.Item {
    	rssfeed.Channel.Item[i].Title = html.UnescapeString(rssfeed.Channel.Item[i].Title)
		rssfeed.Channel.Item[i].Description = html.UnescapeString(rssfeed.Channel.Item[i].Description)
	}

	return &rssfeed, nil
}