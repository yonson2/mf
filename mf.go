package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/yonson2/mf/search"
	"github.com/yonson2/mf/torrent"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	resultsNo := flag.Int("n", 5, "Number of results to be displayed")
	isSearch := flag.Bool("s", false, "Search for results instead of auto selecting the best match")
	isGeneric := flag.Bool("g", false, "Search for all multimedia content (tv/movies) instead of just anime")
	flag.Parse()
	searchQuery := strings.Join(flag.Args(), " ")
	if len(flag.Args()) == 0 {
		fmt.Println("Missing name to search. try: \nmf <content to stream>")
		os.Exit(1)
	}
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Suffix = " Searching..."
	s.Color("green", "bold")
	s.Start()
	results, err := search.Search(searchQuery, *resultsNo, *isGeneric)
	s.Stop()
	if err != nil {
		log.Println("Error searching for results", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("No results found for", searchQuery)
		os.Exit(1)
	}

	var result search.SearchItem
	if *isSearch {
		prompt := genPrompt(results, *resultsNo)
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
