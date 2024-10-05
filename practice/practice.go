package practice

import (
	"net/http"

	"github.com/ysmilda/speedhive-go"
)

// Implementation is based on the Swagger API definition:
// https://practice-api.speedhive.com/swagger/docs/v1

type Client struct {
	c *speedhive.Client

	Accounts  *accountsService
	Admin     *adminService
	Training  *trainingService
	Chips     *chipsService
	Locations *locationsService
}

func NewClient(client *http.Client) *Client {
	c, err := speedhive.NewClient(client, "https://practice-api.speedhive.com")
	if err != nil {
		panic(err)
	}

	return &Client{
		c:         c,
		Accounts:  &accountsService{c: c},
		Admin:     &adminService{c: c},
		Training:  &trainingService{c: c},
		Chips:     &chipsService{c: c},
		Locations: &locationsService{c: c},
	}
}

// Health retrieves the health status of the API.
// If the API is healthy, it will return nil.
func (s Client) Health() error {
	req, err := s.c.Request(http.MethodGet, "/api/v1/health", nil)
	if err != nil {
		return err
	}

	return s.c.Do(req, nil)
}

// Sports retrieves a collection of all sports.
// The results may be used as a SportValue.
func (s Client) Sports() (*SportList, error) {
	return speedhive.Get[SportList](s.c, "/api/v1/sports", nil)
}
