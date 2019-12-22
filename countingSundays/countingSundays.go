package main

/*
You are given the following information, but you may prefer to do some research for yourself.

    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
	A leap year occurs on any year evenly divisible by 4,
		but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century
(1 Jan 1901 to 31 Dec 2000)?
*/

import (
	"fmt"
	"time"
)

func main() {
	//How many sundays in the 20th century
	countSundays("Jan-01-1901", "Dec-31-2000")
}

func countSundays(startDate string, endDate string) (sundays int) {
	//reference date for layout Mon Jan 2 15:04:05 -0700 MST 2006
	fmt.Println(fmt.Sprintf("approximate should be 52/year %v", 52*100))

	start, end := getDateTimes(startDate, endDate)
	fmt.Println(fmt.Sprintf("counting Sundays from %s to %s", start.UTC(), end.UTC()))

	//find out what day of week we are starting on
	dayOfWeekFormat := "Monday"
	day := start.Format(dayOfWeekFormat)

	//add days to adjust the start date to first Sunday (makes the math easy)
	newStart := start.Add(time.Hour * 24 * time.Duration(daysToSunday(day)))
	fmt.Println(fmt.Sprintf("adjusting to first Sunday which is: %s %s", newStart.Format(dayOfWeekFormat), newStart.UTC()))

	//get the duration in Days
	durationDays := end.Sub(newStart).Hours() / 24
	fmt.Println(fmt.Sprintf("duration in days: %v", durationDays))

	//get how many sundays
	sundays = int(durationDays / 7)
	fmt.Println(fmt.Sprintf("%v Sundays between %s and %s", sundays, startDate, endDate))

	return
}

func getDateTimes(startDate string, endDate string) (time.Time, time.Time) {
	timeFormat := "Jan-02-2006"
	//get the date for the start
	start, errStart := time.Parse(timeFormat, startDate)
	if errStart != nil {
		fmt.Println(fmt.Sprintf("error %s converting from %s", errStart, startDate))
	}

	//get the date for the end
	end, errEnd := time.Parse(timeFormat, endDate)
	if errEnd != nil {
		fmt.Println(fmt.Sprintf("error %s converting from %s", errEnd, endDate))
	}
	return start, end
}

func daysToSunday(day string) int {
	if day == "Sunday" {
		return 0
	}

	weekDays := map[string]int{"Sunday": 0, "Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4, "Friday": 5, "Saturday": 6}

	return 7 - weekDays[day]
}
