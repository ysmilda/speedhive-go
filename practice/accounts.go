package practice

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type accountsService struct {
	c *speedhive.Client
}

type AccountsActivitiesOptions struct {
	LocationID *int                  `url:"locationId,omitempty"`
	Year       *int                  `url:"year,omitempty"`
	Sport      *speedhive.SportValue `url:"sport,omitempty"`
	Order      *speedhive.OrderValue `url:"order,omitempty"`
	Count      *int                  `url:"count,omitempty"`
	Offset     *int                  `url:"offset,omitempty"`
}

// Activities retrieves a collection of training activities by accountID.
// The list can be filtered by passing in the AccountsActivitiesOptions.
func (s accountsService) Activities(accountID string, opt *AccountsActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/accounts/%s/training/activities", accountID)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, nil)
}

// ActivityYears returns a list of years in which the account has training activities.
func (s accountsService) ActivityYears(accountID string) (*ActivityYearCountList, error) {
	u := fmt.Sprintf("/api/v1/accounts/%s/training/activities/years", accountID)
	return speedhive.Get[ActivityYearCountList](s.c, u, nil)
}
