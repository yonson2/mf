package search

import (
	"encoding/json"
	"github.com/yonson2/leetflix/config"
	"io/ioutil"
	"net/http"
)

type SearchItem struct {
	Name    string `json:"name"`
	Link    string `json:"link"`
	Date    string `json:"date"`
	Size    string `json:"size"`
	Seeders int    `json:"seeders"`
	Peers   int    `json:"peers"`
}

func Search(query string, resultsNo int) ([]SearchItem, error) {
	var result []SearchItem
	client := &http.Client{}

	req, _ := http.NewRequest("GET", config.ApiUrl, nil)
	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return result, err
	}
	return result[:getMaxSize(len(result), resultsNo)], nil
}

func getMaxSize(sliceSize, maxResults int) int {
	if sliceSize > maxResults {
		return maxResults
	}
	return sliceSize
}
