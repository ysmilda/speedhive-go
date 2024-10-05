package eventresult

import (
	"fmt"

	"github.com/ysmilda/speedhive-go"
)

type sessionsService struct {
	c *speedhive.Client
}

// Get retrieves a session by ID.
func (s sessionsService) Get(sessionID int) (*Session, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d", sessionID)
	return speedhive.Get[Session](s.c, u, nil)
}

// GetCSV returns the session data in CSV format.
func (s sessionsService) GetCSV(sessionID int) ([]byte, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d/csv", sessionID)
	return s.c.GetBody(u, nil)
}

// GetClassification retrieves the classification for a session.
func (s sessionsService) GetClassification(sessionID int) (*IRunClassification, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d/classification", sessionID)
	return speedhive.Get[IRunClassification](s.c, u, nil)
}

// GetLapChart retrieves the lap chart for a session.
func (s sessionsService) GetLapChart(sessionID int) (*LapChart, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d/lapchart", sessionID)
	return speedhive.Get[LapChart](s.c, u, nil)
}

// GetAnnouncements retrieves the announcements for a session.
func (s sessionsService) GetAnnouncements(sessionID int) (*RunAnnouncements, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d/announcements", sessionID)
	return speedhive.Get[RunAnnouncements](s.c, u, nil)
}

type SessionGetLapdataOptions struct {
	Count  *int `url:"count,omitempty"`
	Offset *int `url:"offset,omitempty"`
}

// GetLapdata retrieves the lap data for a session.
func (s sessionsService) GetLapdata(sessionID int, finishPosition int, opt *SessionGetLapdataOptions) (*LapDataPage, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d/lapdata/%d/laps", sessionID, finishPosition)
	return speedhive.Get[LapDataPage](s.c, u, opt)
}

// GetLapdataCSV retrieves the lap data for a session in CSV format.
func (s sessionsService) GetLapdataCSV(sessionID int, finishPosition int) ([]byte, error) {
	u := fmt.Sprintf("/api/v0.2.3/eventresults/sessions/%d/lapdata/%d/laps/csv", sessionID, finishPosition)
	return s.c.GetBody(u, nil)
}
