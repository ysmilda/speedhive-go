package practice

import (
	"github.com/ysmilda/speedhive-go"
)

type ActivitiesInfoExclLocation struct {
	Activities    []Activity `json:"activities"`
	ActivityCount int        `json:"activityCount"`
}

type Activity struct {
	ActivityID          int              `json:"id"`
	Name                string           `json:"name"`
	Start               speedhive.Time   `json:"startTime"`
	End                 speedhive.Time   `json:"endTime"`
	FrameType           int              `json:"frameType"`
	Location            Location         `json:"location"`
	AccountID           int              `json:"accountId"`
	GaUId               string           `json:"gaUId"`
	UserID              string           `json:"userId"`
	ChipLabel           string           `json:"chipLabel"`
	ChipCode            string           `json:"chipCode"`
	ChipList            ChipList         `json:"chipListDto"`
	PracticeSessionsDto PracticeSessions `json:"practiceSessionsDto"`
}

type Location struct {
	Name        string                 `json:"name"`
	LocationID  int                    `json:"id"`
	Description string                 `json:"description"`
	URL         string                 `json:"url"`
	TrackLength int                    `json:"trackLength"`
	Status      string                 `json:"status"`
	Sport       speedhive.SportValue   `json:"sport"`
	Country     speedhive.CountryValue `json:"country"`
	X2ServiceID int                    `json:"x2ServiceId"`
	IPAddress   string                 `json:"ipAddress"`
	Sections    []Section              `json:"sections"`
}

type ChipList struct {
	Chips []Chip `json:"chips"`
}

type PracticeSessions struct {
	Locations []PracticeSession `json:"locations"`
}

type Section struct {
	Name      string `json:"name"`
	Length    int    `json:"length"`
	SpeedTrap bool   `json:"speedTrap"`
}

type Chip struct {
	Code       string `json:"code"`
	CodeNumber int    `json:"codeNr"`
	CarID      int    `json:"carId"`
	ID         int    `json:"id"`
}

type PracticeSession struct {
	ID         int            `json:"id"`
	LocationID int            `json:"locationId"`
	StartUTC   speedhive.Time `json:"starttimeutc"`
	EndUTC     speedhive.Time `json:"endtimeutc"`
	StartLocal speedhive.Time `json:"starttimelocal"`
	EndLocal   speedhive.Time `json:"endtimelocal"`
}

type ActivityYearCountList struct {
	YearCounts []ActivityYearCount `json:"yearCounts"`
	Total      int                 `json:"total"`
}

type ActivityYearCount struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

type LocationList struct {
	Locations []Location `json:"locations"`
}

type ActivitiesInfo struct {
	Location      Location               `json:"location"`
	Activities    []ActivityExclLocation `json:"activities"`
	ActivityCount int                    `json:"activityCount"`
}

type ActivityExclLocation struct {
	ActivityID int            `json:"id"`
	Name       string         `json:"name"`
	Start      speedhive.Time `json:"startTime"`
	End        speedhive.Time `json:"endTime"`
	FrameType  int            `json:"frameType"`
	AccountID  int            `json:"accountId"`
	GaUId      string         `json:"gaUId"`
	ChipLabel  string         `json:"chipLabel"`
	ChipCode   string         `json:"chipCode"`
}

type SportList struct {
	Sports []string `json:"sports"`
}

type ActivityStats struct {
	LapCount           int                `json:"lapCount"`
	FastestTime        speedhive.Duration `json:"fastestTime"`
	AverageTime        speedhive.Duration `json:"averageTime"`
	MedianTime         speedhive.Duration `json:"medianTime"`
	TotalTrainingTime  speedhive.Duration `json:"totalTrainingTime"`
	ActiveTrainingTime speedhive.Duration `json:"activeTrainingTime"`
	AverageSpeed       Speed              `json:"averageSpeed"`
	FastestSpeed       Speed              `json:"fastestSpeed"`
	Chip               Chip               `json:"chip"`
}

type Speed struct {
	Kph float64 `json:"kph"`
	Mps float64 `json:"mps"`
}

type SessionList struct {
	BestLap      OverallBestLap `json:"bestLap"`
	Stats        ActivityStats  `json:"stats"`
	Sessions     []Session      `json:"sessions"`
	Sections     []Section      `json:"sections"`
	OnlyNewSince string         `json:"onlyNewSince"`
}

type OverallBestLap struct {
	SessionID int                `json:"sessionId"`
	LapNumber int                `json:"lapNr"`
	Duration  speedhive.Duration `json:"duration"`
	Speed     Speed              `json:"speed"`
}

type Session struct {
	SessionID          int                `json:"id"`
	ChipID             int                `json:"chipId"`
	Start              speedhive.Time     `json:"dateTimeStart"`
	BestLap            BestLap            `json:"bestLap"`
	AverageLapDuration speedhive.Duration `json:"aveLapDuration"`
	MedianLapDuration  speedhive.Duration `json:"medianLapDuration"`
	Duration           speedhive.Duration `json:"duration"`
	Laps               []Lap              `json:"laps"`
}

type BestLap struct {
	Number   int    `json:"nr"`
	Duration string `json:"duration"`
	Speed    Speed  `json:"speed"`
}

type Lap struct {
	Number                int64              `json:"nr"`
	Start                 string             `json:"dateTimeStart"`
	Duration              string             `json:"duration"`
	Speed                 Speed              `json:"speed"`
	DifferencePreviousLap speedhive.Duration `json:"diffPrevLap"`
	SessionDuration       speedhive.Duration `json:"sessionDuration"`
	Status                LapStatusValue     `json:"status"`
	Sections              []LapSection       `json:"sections"`
	DataAttributes        []DataAttribute    `json:"dataAttributes"`
}

type LapSection struct {
	Name                  string             `json:"name"`
	Duration              speedhive.Duration `json:"duration"`
	DifferencePreviousLap speedhive.Duration `json:"diffPrevLap"`
	Speed                 Speed              `json:"speed"`
}

type DataAttribute struct {
	Type  DataAttributeTypeValue `json:"type"`
	Value interface{}            `json:"value"`
}
