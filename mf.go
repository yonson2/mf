package main

import (
	"flag"
	"strings"
)
func main() {
	resultsNo := flag.Int("n", 5, "Number of results to be displayed")
	isSearch := flag.Bool("s", false, "Search for results instead of auto selecting the best match")
	isGeneric := flag.Bool("g", false, "Search for all multimedia content (tv/movies) instead of just anime")
	preferredPlayer := flag.String("p", "", "preferred video player, defaults to mpv, mplayer and vlc")
	flag.Parse()

	if len(flag.Args()) == 0 {
		//Launch GUI
		launchGUI(*resultsNo, *isGeneric, *isSearch, *preferredPlayer)
	} else {
		//Launch CLI
		searchQuery := strings.Join(flag.Args(), " ")
		launchCLI(searchQuery, *resultsNo, *isGeneric, *isSearch, *preferredPlayer)
	}
}
