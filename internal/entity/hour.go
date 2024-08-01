package entity

type FieldAvailableHour struct {
	FieldAvailableHourID uint
	FieldID              uint
	AvailableHourID      AvailableHour
	Status               string
}

type AvailableHour struct {
	AvailableHourID uint
	StartTime       string
	EndTime         string
}
