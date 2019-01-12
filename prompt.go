package main

import (
	"github.com/manifoldco/promptui"
	"github.com/yonson2/mf/search"
)

func genPrompt(results []search.SearchItem, size int) promptui.Select {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Name | cyan | bold }} ({{ .Seeders | red | bold }} seeders)",
		Inactive: "{{ .Name | cyan }}",
		Selected: "{{ .Name | magenta | bold}}",
	}

	return promptui.Select{
		Label:     "Select title to stream:",
		Items:     results,
		Templates: templates,
		Size:      size,
	}
}
