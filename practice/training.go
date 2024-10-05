package practice

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type trainingService struct {
	c *speedhive.Client
}

type TrainingActivitiesOptions struct {
	LocationID int                   `url:"locationId,omitempty"`
	AccountID  *string               `url:"accountId,omitempty"`
	Year       *int                  `url:"year,omitempty"`
	Sport      *speedhive.SportValue `url:"sport,omitempty"`
	Order      *speedhive.OrderValue `url:"order,omitempty"`
	Count      *int                  `url:"count,omitempty"`
	Offset     *int                  `url:"offset,omitempty"`
}

// TrainingActivities retrieves a collection of training activities by location ID.
// The list can be filtered by passing in the TrainingActivitiesOptions.
func (s trainingService) Activities(locationID int, opt *TrainingActivitiesOptions) (*ActivitiesInfo, error) {
	if opt == nil {
		opt = &TrainingActivitiesOptions{}
	}

	opt.LocationID = locationID

	return speedhive.Get[ActivitiesInfo](s.c, "/api/v1/training/activities", opt)
}

// Activity retrieves a training activity by ID.
func (s trainingService) Activity(activityID int) (*Activity, error) {
	u := fmt.Sprintf("/api/v1/training/activities/%d", activityID)
	return speedhive.Get[Activity](s.c, u, nil)
}

// Statistics retrieves statistics for a training activity by ID.
func (s trainingService) Statistics(activityID int) (*ActivityStats, error) {
	u := fmt.Sprintf("/api/v1/training/activities/%d/statistics", activityID)
	return speedhive.Get[ActivityStats](s.c, u, nil)
}

type TraininsSessionsOptions struct {
	Format       *string `url:"format,omitempty"`
	NumberOfLaps *int    `url:"numberOfLaps,omitempty"`
}

// Sessions retrieves the results of a training activity grouped by session.
// The list can be filtered by passing in the TraininsSessionsOptions.
func (s trainingService) Sessions(activityID int) (*SessionList, error) {
	u := fmt.Sprintf("/api/v1/training/activities/%d/sessions", activityID)
	return speedhive.Get[SessionList](s.c, u, nil)
}

func (s trainingService) Chips(activityID int) (*ChipList, error) {
	u := fmt.Sprintf("/api/v1/training/activities/%d/chips", activityID)
	return speedhive.Get[ChipList](s.c, u, nil)
}
