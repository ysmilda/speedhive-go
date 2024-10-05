package eventresult

import (
	"net/http"

	"github.com/ysmilda/speedhive-go"
)

// Implementation is based on the Swagger API definition:
// https://eventresults-api.speedhive.com/swagger/docs/v1

type Client struct {
	c *speedhive.Client

	Accounts      *accountsService
	Championships *championshipsService
	Events        *eventsService
	Organizations *organizationsService
	Search        *searchService
	Sessions      *sessionsService
}

func NewClient(client *http.Client) *Client {
	c, err := speedhive.NewClient(client, "https://eventresults-api.speedhive.com")
	if err != nil {
		panic(err)
	}

	return &Client{
		c: c,

		Accounts:      &accountsService{c: c},
		Championships: &championshipsService{c: c},
		Events:        &eventsService{c: c},
		Organizations: &organizationsService{c: c},
		Search:        &searchService{c: c},
		Sessions:      &sessionsService{c: c},
	}
}

// Health returns the health status of the API.
// If the API is healthy, it will return nil.
func (s Client) Health() error {
	req, err := s.c.Request("GET", "/api/health", nil)
	if err != nil {
		return err
	}

	return s.c.Do(req, nil)
}
