package anime

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//URLBASE  Url utilizada para acessar a API
const URLBASE = "https://api.jikan.moe/v3/"

// Anime  Dados uteis recebidos pela a API
type Anime struct {
	ID          string `json:"mal_id"`
	Title       string `json:"title"`
	ImageURL    string `json:"image_url"`
	Description string `json:"synopsis"`
	Episodes    string `json:"episodes"`
}

// Result  Result do Search da API
type Result struct {
	Results []*Anime `json:"results"`
}

// SearchAnime  Busca o anime
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
