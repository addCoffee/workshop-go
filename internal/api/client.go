package api

type Client interface {
	// GetJore returns one joke
	GetJoke() (*JokeResponse, error)
}
