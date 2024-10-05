package eventresult

import (
	"github.com/ysmilda/speedhive-go"
)

type EventDto struct {
	EventID        int             `json:"id"`
	Name           string          `json:"name"`
	Sport          string          `json:"sport"`
	StartDate      speedhive.Time  `json:"startDate"`
	UpdatedAt      speedhive.Time  `json:"updatedAt"`
	UploadSoftware UploadSoftware  `json:"uploadSoftware"`
	HasChampion    bool            `json:"hasChamp"`
	Location       Location        `json:"location"`
	Organization   Organization    `json:"organization"`
	Sessions       SessionGrouping `json:"sessions"`
}

type Organization struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Logo    string  `json:"logo"`
	URL     string  `json:"url"`
	City    string  `json:"city"`
	Ref     *string `json:"ref,omitempty"`
	Country Country `json:"country"`
}

type SessionGrouping struct {
	Sessions []Session      `json:"sessions"`
	Groups   []SessionGroup `json:"groups"`
}

type Session struct {
	SessionID     int              `json:"id"`
	Name          string           `json:"name"`
	Comment       string           `json:"comment"`
	EventID       int              `json:"eventId"`
	Type          string           `json:"type"`
	StartTime     speedhive.Time   `json:"startTime"`
	GroupName     string           `json:"groupName"`
	IsMerge       bool             `json:"isMerge"`
	Participated  int              `json:"participated"`
	Announcements RunAnnouncements `json:"announcements"`
}

type SessionGroup struct {
	SessionGroupID int            `json:"id"`
	Name           string         `json:"name"`
	Date           speedhive.Time `json:"date"`
	SubGroups      []SessionGroup `json:"subGroups"`
	Sessions       []Session      `json:"sessions"`
}

type RunAnnouncements struct {
	Rows []AnnouncementRow `json:"rows"`
}

type AnnouncementRow struct {
	Timestamp speedhive.Time `json:"timestamp"`
	Text      string         `json:"text"`
}

type IRunClassification struct {
	Rows    []IRunRow               `json:"rows"`
	Classes []string                `json:"classes"`
	Type    ClassificationTypeValue `json:"type"`
}

type IRunRow struct {
	Position         int                `json:"position"`
	StartNumber      string             `json:"startNumber"`
	Name             string             `json:"name"`
	ResultClass      string             `json:"resultClass"`
	PositionInClass  int                `json:"positionInClass"`
	AdditionalFields []string           `json:"additionalFields"`
	IsQualified      bool               `json:"isQualified"`
	Status           IRunRowStatusValue `json:"status"`
	User             UserInfo           `json:"user"`
}

type UserInfo struct {
	MstUserID       string               `json:"mstUserId"`
	ProductCategory ProductCategoryValue `json:"productCategory"`
	Chip            Chip                 `json:"chip"`
}

type Chip struct {
	Code string `json:"code"`
}

type ChampionshipData struct {
	ID          int64                    `json:"id"`
	Name        string                   `json:"name"`
	Season      string                   `json:"season"`
	Class       string                   `json:"class"`
	Layout      string                   `json:"layout"`
	Events      []ChampionshipEvent      `json:"events"`
	Competitors []ChampionshipCompetitor `json:"competitors"`
}

type ChampionshipEvent struct {
	ID     int                      `json:"id"`
	Name   string                   `json:"name"`
	Date   string                   `json:"date"`
	Rounds []ChampionshipEventRound `json:"rounds"`
}

type ChampionshipCompetitor struct {
	No         string                        `json:"no"`
	FirstName  string                        `json:"firstName"`
	LastName   string                        `json:"lastName"`
	Class      string                        `json:"class"`
	Position   string                        `json:"position"`
	Total      float64                       `json:"total"`
	Difference float64                       `json:"diff"`
	Gap        float64                       `json:"gap"`
	Events     []ChampionshipCompetitorEvent `json:"events"`
}

type ChampionshipEventRound struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type ChampionshipCompetitorEvent struct {
	EventID  int                                `json:"eventId"`
	Points   float64                            `json:"points"`
	Position string                             `json:"position"`
	Rounds   []ChampionshipCompetitorEventRound `json:"rounds"`
}

type ChampionshipCompetitorEventRound struct {
	RoundID  int     `json:"roundId"`
	Points   float64 `json:"points"`
	Position string  `json:"position"`
}

type Event struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Sport          string         `json:"sport"`
	StartDate      speedhive.Time `json:"startDate"`
	UpdatedAt      speedhive.Time `json:"updatedAt"`
	UploadSoftware UploadSoftware `json:"uploadSoftware"`
	Location       Location       `json:"location"`
	Organization   Organization   `json:"organization"`
}

type UploadSoftware struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Location struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Length      float64 `json:"length"`
	LengthUnit  string  `json:"lengthUnit"`
	LengthLabel string  `json:"lengthLabel"`
	Country     Country `json:"country"`
	Lon         float64 `json:"lon"`
	Lat         float64 `json:"lat"`
}

type Country struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
}

type ChampionshipList struct {
	Championships []Championship `json:"championships"`
}

type Championship struct {
	ID             int    `json:"id"`
	Organization   int    `json:"organization"`
	ChampionshipID int64  `json:"championshipid"`
	Name           string `json:"name"`
	Season         string `json:"season"`
	HTML           string `json:"html"`
	XML            string `json:"xml"`
	BackupXML      string `json:"backupxml"`
}

type SearchResults struct {
	Results []Search `json:"results"`
}

type Search struct {
	EID   int                          `json:"eid"`
	AID   string                       `json:"aid"`
	Type  SearchTypeValue              `json:"type"`
	Sport speedhive.SportCategoryValue `json:"sport"`
	Title string                       `json:"title"`
	Desc  string                       `json:"desc"`
	CC    string                       `json:"cc"`
	Date  speedhive.Time               `json:"date"`
}

type LapChart struct {
	ID             int             `json:"id"`
	PositionRows   [][]PositionRow `json:"positionRows"`
	StartPositions []StartPosition `json:"startPositions"`
	HasStartGrid   bool            `json:"hasStartGrid"`
	NumberOfLaps   int             `json:"numberOfLaps"`
	EventID        int             `json:"eventId"`
}

type PositionRow struct {
	Position    int      `json:"position"`
	StartNumber string   `json:"startNumber"`
	InLeaderLap bool     `json:"inLeaderLap"`
	InPit       bool     `json:"inPit"`
	Status      []string `json:"status"`
}

type StartPosition struct {
	Name        string `json:"name"`
	StartNumber string `json:"startNumber"`
}

type LapDataPage struct {
	LapDataInfo LapDataInfo `json:"lapDataInfo"`
	Laps        []Lap       `json:"laps"`
}

type LapDataInfo struct {
	ParticipantInfo          ParticipantInfo         `json:"participantInfo"`
	LapCount                 int                     `json:"lapCount"`
	AllLapsHaveFieldPosition bool                    `json:"allLapsHaveFieldPos"`
	FirstLapNumber           int                     `json:"firstLapNr"`
	LapsDriven               int                     `json:"lapsDriven"`
	ClassificationType       ClassificationTypeValue `json:"classificationType"`
	SessionID                int                     `json:"sessionId"`
}

type Lap struct {
	LapNumber       int                `json:"lapNr"`
	TimeOfDay       speedhive.Time     `json:"timeOfDay"`
	LapTime         speedhive.Duration `json:"lapTime"`
	DiffWithLastLap speedhive.Duration `json:"diffWithLastLap"`
	DiffWithBestLap speedhive.Duration `json:"diffWithBestLap"`
	Speed           float64            `json:"speed"`
	SectionTimes    []interface{}      `json:"sectionTimes"`
	InPit           bool               `json:"inPit"`
	Status          []StatusValue      `json:"status"`
	FieldComparison LapComparison      `json:"fieldComparison"`
}

type ParticipantInfo struct {
	Name                string `json:"name"`
	Class               string `json:"class"`
	Transponder         string `json:"transponder"`
	UserID              string `json:"userId"`
	StartNumber         string `json:"startNr"`
	StartPosition       int    `json:"startPos"`
	FieldFinishPosition int    `json:"fieldFinishPos"`
	ClassFinishPosition int    `json:"classFinishPos"`
}

type LapComparison struct {
	Position   int           `json:"position"`
	LeaderLap  int           `json:"leaderLap"`
	Difference LapDifference `json:"diff"`
	GapAhead   LapDifference `json:"gapAhead"`
	GapBehind  LapDifference `json:"gapBehind"`
}

type LapDifference struct {
	Time speedhive.Duration `json:"time"`
	Laps int                `json:"laps"`
}
