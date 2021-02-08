package main

import (
	"github.com/yonson2/mf/search"
	"github.com/briandowns/spinner"
	"time"
	"log"
	"fmt"
	"github.com/yonson2/mf/torrent"
	"os"
)

func launchCLI(query string, maxResults int, generic bool, isSearch bool) {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Suffix = " Searching..."
	s.Color("green", "bold")
	s.Start()
	results, err := search.Search(query, maxResults, generic)
	s.Stop()
	if err != nil {
		log.Println("Error searching for results", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("No results found for", query)
		os.Exit(1)
	}

	var result search.SearchItem
	if isSearch {
		prompt := genPrompt(results, maxResults)
		i, _, err := prompt.Run()

		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}
		result = results[i]
	} else {
		result = results[0]
	}
	err = torrent.StreamTorrent(result.Link)
	if err != nil {
		log.Println("There was an error:", err)
	}
}
