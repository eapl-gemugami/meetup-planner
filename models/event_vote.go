package models

type EventVote struct {
	EventUserID int
	EventUser EventUser

	TimeOption int
	TimeAvailability int
}
