package practice

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type locationsService struct {
	c *speedhive.Client
}

type LocationListOptions struct {
	SportCategory *speedhive.SportCategoryValue `url:"sportCategory,omitempty"`
	Sport         *speedhive.SportValue         `url:"sport,omitempty"`
	Country       *speedhive.CountryValue       `url:"country,omitempty"`
	Status        *StatusValue                  `url:"status,omitempty"`
}

// List retrieves a collection of all (practice) locations.
// The list can be filtered by passing in the LocationListOptions.
func (s locationsService) List(opt *LocationListOptions) (*LocationList, error) {
	u := fmt.Sprintf("/api/v1/locations")
	return speedhive.Get[LocationList](s.c, u, opt)
}

// Get retrieves a location by ID.
func (s locationsService) Get(locationID int) (*Location, error) {
	u := fmt.Sprintf("/api/v1/locations/%d", locationID)
	return speedhive.Get[Location](s.c, u, nil)
}

type LocationActivitiesOptions struct {
	Year   *int                  `url:"year,omitempty"`
	Sport  *speedhive.SportValue `url:"sport,omitempty"`
	Order  *speedhive.OrderValue `url:"order,omitempty"`
	Count  *int                  `url:"count,omitempty"`
	Offset *int                  `url:"offset,omitempty"`
}

// Activities retrieves a collection of training activities by locationID.
// The list can be filtered by passing in the LocationActivitiesOptions.
func (s locationsService) Activities(locationID int, opt *LocationActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/locations/%d/activities", locationID)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}

type SessionsByLocationAndChipCodeOptions struct {
	Order *speedhive.OrderValue `url:"order,omitempty"`
}

// SessionsByLocationAndChipCode retrieves all practice sessions of a transponder of a specific track.
// The list can be filtered by passing in the SessionsByLocationAndChipCodeOptions.
func (s locationsService) SessionsByLocationAndChipCode(locationID int, chipCode string, opt *SessionsByLocationAndChipCodeOptions) (*PracticeSessions, error) {
	u := fmt.Sprintf("/api/v1/locations/%d/%s", locationID, chipCode)
	return speedhive.Get[PracticeSessions](s.c, u, opt)
}
