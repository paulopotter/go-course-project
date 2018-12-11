package anime

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const URLBASE = "https://api.jikan.moe/v3/"

type Anime struct {
	Title string `json:"title"`
}

type Result struct {
	Results []*Anime `json:"results"`
}

func SearchAnime(animeName string) (*Result, error) {
	url := fmt.Sprintf("%ssearch/anime?q=%s", URLBASE, animeName)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result := &Result{}

	err = json.NewDecoder(res.Body).Decode(result)

	return result, nil

}
