package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/dialog"
	"log"
	"github.com/yonson2/mf/config"
	"github.com/yonson2/mf/torrent"
	"github.com/yonson2/mf/search"
)

func launchGUI(maxResults int, generic bool, isSearch bool) {
	a := app.New()
	w := a.NewWindow(config.AppName)

	content := container.NewVBox()

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	inputColumn := container.New(layout.NewMaxLayout(), input)

	searchButton := widget.NewButton("Search", func() {})
	luckyButton := widget.NewButton("I'm feeling lucky", func () {})

	searchColumn := container.NewHBox(
		layout.NewSpacer(),
		searchButton,
		luckyButton,
	)

	findTorrent := func (query string, lucky bool) {
		pb := widget.NewProgressBarInfinite()
		content.Add(pb)
		results, err := search.Search(query, maxResults, true)
		if err != nil {
			log.Println("Error searching for results", err)
			a.Quit()
		}
		if len(results) == 0 {
			msg := "No results found for: " + query
			dialog.ShowInformation("No results", msg, w)
			content.Remove(pb)
			input.Text = ""
			input.Enable()
			searchButton.Enable()
			luckyButton.Enable()
			return
		}
		if lucky {
			err = torrent.StreamTorrent(results[0].Link)
			if err != nil {
				log.Println("Error streaming file", err)
				a.Quit()
			}
			a.Quit()
		}
		pb.Hide()
		//create a dict of key = name of result, value = link.
		searchResultsDict := make(map[string]string)
		var searchOptions []string
		for _, r := range results {
			searchResultsDict[r.Name] = r.Link
			searchOptions = append(searchOptions, r.Name)
		}
		options := widget.NewRadioGroup(searchOptions, func (string) {})
		options.OnChanged = func (value string) {
			options.Hide()
			pb.Show()
			err = torrent.StreamTorrent(searchResultsDict[value])
			if err != nil {
				log.Println("Error streaming file", err)
				a.Quit()
			}
			a.Quit()
		}
		content.Add(options)
	}

	searchButton.OnTapped = func () {
		input.Disable()
		searchButton.Disable()
		luckyButton.Disable()
		findTorrent(input.Text, false)
	}

	luckyButton.OnTapped = func () {
		input.Disable()
		searchButton.Disable()
		luckyButton.Disable()
		findTorrent(input.Text, true)
	}

	input.OnSubmitted = func(query string) {
		input.Disable()
		searchButton.Disable()
		luckyButton.Disable()
		findTorrent(query, true)
	}

	grid := container.New(layout.NewGridLayout(2), inputColumn, searchColumn)

	content.Add(grid)

	w.SetContent(content)
	w.Canvas().Focus(input)
	w.SetMaster()
	w.SetFullScreen(false)
	w.SetFixedSize(true)
	w.CenterOnScreen()

	w.ShowAndRun()
}
