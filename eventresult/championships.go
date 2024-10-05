package eventresult

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type championshipsService struct {
	c *speedhive.Client
}

// Get retrieves a championship by ID.
func (s championshipsService) Get(championshipID int) (*ChampionshipData, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/championships/%d", championshipID)
	return speedhive.Get[ChampionshipData](s.c, u, nil)
}

// GetCSV returns the championship data in CSV format.
func (s championshipsService) GetCSV(championshipID int) ([]byte, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/championships/%d/csv", championshipID)
	return s.c.GetBody(u, nil)
}

// GetOverview retrieves the championship overview by ID.
// This returns an html page that renders the championship overview.
func (s championshipsService) GetOverview(championshipID int) ([]byte, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/organizations/championships/%d/overview", championshipID)
	return s.c.GetBody(u, nil)
}
