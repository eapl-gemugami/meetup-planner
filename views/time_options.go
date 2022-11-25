package views

import (
	//"fmt"
	"time"
)

func GetTimeOptions(startTimestamp int64, endTimestamp int64,
		intervalMins int, targetLoc time.Location)	[]string {
	timeStart := time.Unix(startTimestamp, 0)
	timeEnd := time.Unix(endTimestamp, 0)

	//fmt.Printf("%v - %v\n", timeStart, startTimestamp)
	//fmt.Printf("%v - %v\n", timeEnd, endTimestamp)

	var timeOptions []string

	// TODO: Check that the difference is less than a week
	// https://gosamples.dev/difference-between-dates/
	difference := timeEnd.Sub(timeStart)
	diff_weeks := int64(difference.Hours() / 24 / 7)
	if diff_weeks >= 1 {
		//fmt.Printf("Data range is too big - Only a week is allowed")
		return timeOptions
	}

	// Return an empty array
	if timeEnd.Before(timeStart) {
		//fmt.Printf("End time is before Start time")
		return timeOptions
	}

	// Init the first Date as the timeStart
	currentDate := time.Unix(startTimestamp, 0)

	currentIdx := 0
	currentTimeIsBeforetimeEnd := true
	for currentTimeIsBeforetimeEnd {
		// https://pkg.go.dev/fmt#hdr-Printing
		// For debugging purposes
		//fmt.Printf("%v - %s\n", currentIdx, currentDate.UTC())

		timeOptions = append(timeOptions, currentDate.In(&targetLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"))

		currentDate = currentDate.Add(time.Minute * time.Duration(intervalMins))
		currentTimeIsBeforetimeEnd = currentDate.Before(timeEnd)
		currentIdx += 1
	}

	// Add the last option as well
	timeOptions = append(timeOptions, currentDate.In(&targetLoc).Format("02/Jan/2006 - Monday @ 15:04 / 3:04pm"))

	return timeOptions
}