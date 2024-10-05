package eventresult

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type accountsService struct {
	c *speedhive.Client
}

type AccountsEventsOptions struct {
	Sport         *speedhive.SportValue         `url:"sport,omitempty"`
	SportCategory *speedhive.SportCategoryValue `url:"sportCategory,omitempty"`
	Count         *int                          `url:"count,omitempty"`
	Offset        *int                          `url:"offset,omitempty"`
}

// Events returns a list of events for a user.
func (s accountsService) Events(userID int, opt *AccountsEventsOptions) (*EventDto, error) {
	u := fmt.Sprintf("/api/v0.2.3/accounts/%d/events", userID)
	return speedhive.Get[EventDto](s.c, u, opt)
}

type AccountsEventSessionOptions struct {
	Count  *int `url:"count,omitempty"`
	Offset *int `url:"offset,omitempty"`
}

// EventSessions returns a list of sessions for an event of a user.
func (s accountsService) EventSessions(userID int, eventID int, opt *AccountsEventSessionOptions) ([]Session, error) {
	u := fmt.Sprintf("/api/v0.2.3/accounts/%d/events/%d/sessions", userID, eventID)
	res, err := speedhive.Get[[]Session](s.c, u, opt)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

type AccountsSessionsOptions struct {
	EventID *int `url:"eventId,omitempty"`
	Count   *int `url:"count,omitempty"`
	Offset  *int `url:"offset,omitempty"`
}

// Sessions return a list of sessions for a user.
func (s accountsService) Sessions(userID int, opt *AccountsSessionsOptions) ([]Session, error) {
	u := fmt.Sprintf("/api/v0.2.3/accounts/%d/sessions", userID)
	res, err := speedhive.Get[[]Session](s.c, u, opt)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

type AccountsSessionsClassificationOptions struct {
	Count  *int `url:"count,omitempty"`
	Offset *int `url:"offset,omitempty"`
}

// SessionsClassification returns a list of classifications for a session of a user.
func (s accountsService) SessionsClassification(userID int, sessionID int, opt *AccountsSessionsClassificationOptions) (*IRunClassification, error) {
	u := fmt.Sprintf("/api/v0.2.3/accounts/%d/sessions/%d/classification", userID, sessionID)
	return speedhive.Get[IRunClassification](s.c, u, opt)
}
