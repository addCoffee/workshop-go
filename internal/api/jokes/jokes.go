package jokes

import (
	"fmt"
	"net/http"

	"gopkg.in/square/go-jose.v2/json"

	"workshop/internal/api"
)

const getJokePath = "/api?format=json"

type JokeClient struct {
	url string
}

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeCLient{
		url: baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (string, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %d", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
