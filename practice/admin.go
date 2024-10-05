package practice

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type adminService struct {
	c *speedhive.Client
}

type AdminActivitiesOptions struct {
	LocationID *int                  `url:"locationId,omitempty"`
	Year       *int                  `url:"year,omitempty"`
	Sport      *speedhive.SportValue `url:"sport,omitempty"`
	Order      *speedhive.OrderValue `url:"order,omitempty"`
	Count      *int                  `url:"count,omitempty"`
	Offset     *int                  `url:"offset,omitempty"`
}

// Activities retrieves a collection of training activities by accountID.
// The list can be filtered by passing in the AdminActivitiesOptions.
func (s adminService) Activities(accountID int, opt *AdminActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/admin/accountId/%d", accountID)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}

// ActivitiesByChipNumber retrieves a collection of training activities by chip number.
func (s adminService) ActivitiesByChipNumber(chipNumber int, opt *AdminActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/admin/chipNumber/%d", chipNumber)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}

// ActivitiesByChipCode retrieves a collection of training activities by chip code.
func (s adminService) ActivitiesByChipCode(chipCode string, opt *AdminActivitiesOptions) (*ActivitiesInfoExclLocation, error) {
	u := fmt.Sprintf("/api/v1/admin/chipCode/%s", chipCode)
	return speedhive.Get[ActivitiesInfoExclLocation](s.c, u, opt)
}
