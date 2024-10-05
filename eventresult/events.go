package eventresult

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type eventsService struct {
	c *speedhive.Client
}

type EventsListOptions struct {
	Sport         *speedhive.SportValue         `url:"sport,omitempty"`
	SportCategory *speedhive.SportCategoryValue `url:"sportCategory,omitempty"`
	StartDate     *speedhive.Time               `url:"startDate,omitempty"`
	EndDate       *speedhive.Time               `url:"endDate,omitempty"`
	Country       *speedhive.CountryValue       `url:"country,omitempty"`
	Count         *int                          `url:"count,omitempty"`
	Offset        *int                          `url:"offset,omitempty"`
}

// List returns a list of all events.
// The list can be filtered by passing in the EventsListOptions.
func (s *eventsService) List(opt *EventsListOptions) ([]EventDto, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/events")
	res, err := speedhive.Get[[]EventDto](s.c, u, opt)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

type EventsGetOptions struct {
	Sessions *bool   `url:"sessions,omitempty"`
	Guid     *string `url:"guid,omitempty"`
}

// Get returns a single event by its ID.
// The event can be expanded with sessions by passing in the EventsGetOptions.
func (s *eventsService) Get(eventID int, opt *EventsGetOptions) (*EventDto, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/events/%d", eventID)
	return speedhive.Get[EventDto](s.c, u, opt)
}

type EventsGetSessionsOptions struct {
	Guid *string `url:"guid,omitempty"`
}

// GetSessions returns a list of sessions for an event.
// The list can be filtered by passing in the EventsGetSessionsOptions.
func (s *eventsService) GetSessions(eventID int, opt *EventsGetSessionsOptions) (*SessionGrouping, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/events/%d/sessions", eventID)
	return speedhive.Get[SessionGrouping](s.c, u, opt)
}
