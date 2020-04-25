package api

// Client
type Client struct {
	// GetJore returns one joke
	GetJoke() (*JokeResponse, error)
}