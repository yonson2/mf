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
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Missing name to search. try: \nleetflix <content to stream>")
		os.Exit(1)
	}
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Suffix = " Searching..."
	s.Color("green", "bold")
	s.Start()
	results, err := search.Search(strings.Join(flag.Args(), " "), *resultsNo)
	s.Stop()
	if err != nil {
		log.Println("Error searching for results", err)
		return
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
