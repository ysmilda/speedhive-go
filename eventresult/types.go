package eventresult

type StatusValue string

const (
	StatusValue_RED    StatusValue = "RED"
	StatusValue_GREEN  StatusValue = "GREEN"
	StatusValue_YELLOW StatusValue = "YELLOW"
)

type ClassificationTypeValue string

const (
	ClassificationTypeValue_PracticeAndQualification ClassificationTypeValue = "PracticeAndQualification"
	ClassificationTypeValue_Race                     ClassificationTypeValue = "Race"
	ClassificationTypeValue_RaceMerge                ClassificationTypeValue = "RaceMerge"
	ClassificationTypeValue_QualificationMerge       ClassificationTypeValue = "QualificationMerge"
	ClassificationTypeValue_PointMerge               ClassificationTypeValue = "PointMerge"
	ClassificationTypeValue_Null                     ClassificationTypeValue = "Null"
	ClassificationTypeValue_Timetrial                ClassificationTypeValue = "Timetrial"
	ClassificationTypeValue_Handicap                 ClassificationTypeValue = "Handicap"
)

type IRunRowStatusValue string

const (
	IRunRowStatusValue_Normal      IRunRowStatusValue = "Normal"
	IRunRowStatusValue_DNS         IRunRowStatusValue = "DNS"
	IRunRowStatusValue_DNF         IRunRowStatusValue = "DNF"
	IRunRowStatusValue_DQ          IRunRowStatusValue = "DQ"
	IRunRowStatusValue_Provisional IRunRowStatusValue = "Provisional"
)

type ProductCategoryValue string

const (
	ProductCategoryValue_Other  ProductCategoryValue = "Other"
	ProductCategoryValue_Flex   ProductCategoryValue = "Flex"
	ProductCategoryValue_X2     ProductCategoryValue = "X2"
	ProductCategoryValue_X2Link ProductCategoryValue = "X2Link"
	ProductCategoryValue_TR2    ProductCategoryValue = "TR2"
)

type SearchFilterValue string

const (
	SearchFilterValue_Events       SearchFilterValue = "Events"
	SearchFilterValue_Locations    SearchFilterValue = "Locations"
	SearchFilterValue_Racers       SearchFilterValue = "Racers"
	SearchFilterValue_Transponders SearchFilterValue = "Transponders"
	SearchFilterValue_AccountID    SearchFilterValue = "AccountId"
)

type SearchModeValue string

const (
	SearchModeValue_Any SearchModeValue = "Any"
	SearchModeValue_All SearchModeValue = "All"
)

type SearchTypeValue string

const (
	SearchTypeValue_Event   SearchTypeValue = "Event"
	SearchTypeValue_Racer   SearchTypeValue = "Racer"
	SearchTypeValue_Profile SearchTypeValue = "Profile"
	SearchTypeValue_All     SearchTypeValue = "All"
)
