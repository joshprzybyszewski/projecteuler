package easy

import (
	"fmt"
)

func SolveProblem19() string {
	/*
		How many Sundays fell on the first of the month
		during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
	*/
	ans := getSundaysOnFirstOfMonths(1901, 2000)
	return fmt.Sprintf("%v", ans)
}

func getSundaysOnFirstOfMonths(startYear, endYear int) int {
	numSundaysOnFirst := 0

	month := 1
	year := 1900

	firstDay := 1 // 1 Jan 1900 is a Monday

	for year <= endYear {
		if firstDay == 0 {
			if year >= startYear {
				numSundaysOnFirst++
			}
		}

		firstDay += daysInMonth(year, month)
		firstDay %= 7

		if month == 12 {
			month = 1
			year++
		} else {
			month++
		}
	}

	return numSundaysOnFirst
}

func daysInMonth(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if year%100 == 0 {
			if year%400 == 0 {
				return 29
			}
		} else if year%4 == 0 {
			return 29
		}
		return 28
	default:
		return 0
	}
}
