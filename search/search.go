package search

import (
	"github.com/mmcdole/gofeed"
)

const searchURL = "https://nyaa.si/?f=0&c=1_0&s=seeders&o=desc&page=rss&q="

type SearchItem struct {
	Name     string
	Link     string
	Date     string
	Size     string
	Seeders  string
	Leechers string
}

func Search(query string, resultsNo int) ([]SearchItem, error) {
	var result []SearchItem
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(searchURL + query)
	if err != nil {
		return result, err
	}

	for _, item := range feed.Items[:getMaxSize(len(feed.Items), resultsNo)] {
		result = append(result, SearchItem{Name: item.Title, Link: item.Link, Date: item.Published, Seeders: item.Extensions["nyaa"]["seeders"][0].Value, Leechers: item.Extensions["nyaa"]["leechers"][0].Value})
	}

	return result, nil
}

func getMaxSize(sliceSize, maxResults int) int {
	if sliceSize > maxResults {
		return maxResults
	}
	return sliceSize
}
