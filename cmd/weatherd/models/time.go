package models

type Time struct {
	ID                    string `json:"$id"`
	CurrentDateTime       string `json:"currentDateTime"`
	UTCOffset             string `json:"utcOffset"`
	IsDayLightSavingsTime bool   `json:"isDayLightSavingsTime"`
	DayOfTheWeek          string `json:"dayOfTheWeek"`
	TimeZoneName          string `json:"UTC"`
	CurrentFileTime       int64  `json:"currentFileTime"`
	OrdinalDate           string `json:"ordinalDate"`
}
