package practice

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type chipsService struct {
	c *speedhive.Client
}

type ChipsActivitiesOptions struct {
	LocationID *int                  `url:"locationId,omitempty"`
	Year       *int                  `url:"year,omitempty"`
	Sport      *speedhive.SportValue `url:"sport,omitempty"`
	Order      *speedhive.OrderValue `url:"order,omitempty"`
	Count      *int                  `url:"count,omitempty"`
	Offset     *int                  `url:"offset,omitempty"`
}

// ActivitiesByChipNumber retrieves a collection of training activities by chip number.
// The list can be filtered by passing in the ChipsActivitiesOptions.
func (s chipsService) ActivitiesByChipNumber(chipNumber int, opt *ChipsActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/chips/%d/training/activities", chipNumber)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}

// ActivitiesByChipNumberAndCarID retrieves a collection of training activities by chip number and car ID.
func (s chipsService) ActivitiesByChipNumberAndCarID(chipNumber int, carID int, opt *ChipsActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/chips/%d/carId/%d/training/activities", chipNumber, carID)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}

// ActivitiesByChipCode retrieves a collection of training activities by chip code.
func (s chipsService) ActivitiesByChipCode(chipCode string, opt *ChipsActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/chips/code/%s/training/activities", chipCode)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}

// ActivitiesByChipCode retrieves a collection of training activities by chip code and car ID.
func (s chipsService) ActivitiesByChipCodeAndCarID(chipCode string, carID int, opt *ChipsActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/chips/code/%s/carId/%d/training/activities", chipCode, carID)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}
