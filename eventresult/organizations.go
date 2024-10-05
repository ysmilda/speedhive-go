package eventresult

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type organizationsService struct {
	c *speedhive.Client
}

// Get retrieves an organization by ID.
func (s organizationsService) Get(organizationID int) (*Organization, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/organizations/%d", organizationID)
	return speedhive.Get[Organization](s.c, u, nil)
}

type OrganizationsGetEventsOptions struct {
	SportCategory *speedhive.SportCategoryValue `url:"sportCategory,omitempty"`
	Count         *int                          `url:"count,omitempty"`
	Offset        *int                          `url:"offset,omitempty"`
}

// Get returns a list of events for an organization.
// The list can be filtered by passing in the OrganizationsGetEventsOptions.
func (s organizationsService) GetEvents(organizationID int, opt *OrganizationsGetEventsOptions) ([]Event, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/organizations/%d/events", organizationID)
	res, err := speedhive.Get[[]Event](s.c, u, opt)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

type OrganizationsGetChampionshipsOptions struct {
	Count  *int `url:"count,omitempty"`
	Offset *int `url:"offset,omitempty"`
}

// GetChampionships returns a list of championships for an organization.
// The list can be filtered by passing in the OrganizationsGetChampionshipsOptions.
func (s organizationsService) GetChampionships(organizationID int, opt *OrganizationsGetChampionshipsOptions) ([]Championship, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/organizations/%d/championships", organizationID)
	res, err := speedhive.Get[ChampionshipList](s.c, u, opt)
	if err != nil {
		return nil, err
	}
	return res.Championships, nil
}

// GetChampionship retrieves a championship by organization ID and championship ID.
func (s organizationsService) GetChampionship(organizationID int, championshipID int) (*Championship, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/organizations/%d/championships/%d", organizationID, championshipID)
	return speedhive.Get[Championship](s.c, u, nil)
}
