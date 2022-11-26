package models

type EventVote struct {
	EventID int
	Event Event

	EventUserID int
	EventUser EventUser

	TimeOption int
	TimeAvailability int
}
