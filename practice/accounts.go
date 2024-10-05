package practice

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type accountsService struct {
	c *speedhive.Client
}

// ActivityYears retrieves a collection of years that contain training activities.
func (s accountsService) ActivityYears(accountID int) (*ActivityYearCountList, error) {
	u := fmt.Sprintf("/api/v1/accounts/%d/activities/years", accountID)
	return speedhive.Get[ActivityYearCountList](s.c, u, nil)
}
