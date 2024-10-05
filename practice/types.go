package practice

func Ptr[T any](v T) *T {
	return &v
}

type StatusValue string

const (
	StatusValue_Unknown StatusValue = "UNKNOWN"
	StatusValue_Offline StatusValue = "OFFLINE"
	StatusValue_Online  StatusValue = "ONLINE"
	StatusValue_Active  StatusValue = "ACTIVE"
)

type LapStatusValue string

const (
	LapStatusValue_None        LapStatusValue = "NONE"
	LapStatusValue_Slower      LapStatusValue = "SLOWER"
	LapStatusValue_Equal       LapStatusValue = "EQUAL"
	LapStatusValue_Faster      LapStatusValue = "FASTER"
	LapStatusValue_SessionBest LapStatusValue = "SESSIONBEST"
	LapStatusValue_ReportBest  LapStatusValue = "REPORTBEST"
)

type DataAttributeTypeValue string

const (
	DataAttributeTypeValue_Voltage     DataAttributeTypeValue = "VOLTAGE"
	DataAttributeTypeValue_Temperature DataAttributeTypeValue = "TEMPERATUREs"
)
