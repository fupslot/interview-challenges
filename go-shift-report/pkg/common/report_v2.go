package common

import (
	"fmt"
	"time"
)

func ReportV2(beginingOfPeriod int64, daysOfPeriod int, s *Shifts) []int64 {
	var days []int64 = make([]int64, daysOfPeriod)
	var day int64 = 24 * int64(time.Hour/time.Second)

	fmt.Printf("day is %d secs\n", day)

	var shiftCount int = 0
	shift := s.Next()
	for shift != nil {
		shiftCount += 1
		shiftDayStart := (day + (shift.Start - beginingOfPeriod)) / day
		shiftDayEnd := (day + (shift.End - beginingOfPeriod)) / day

		// exceeding days period? break and exit
		if shiftDayEnd > int64(daysOfPeriod) {
			return days
		}

		if shiftDayStart == shiftDayEnd {
			days[shiftDayStart-1] = shift.End - shift.Start
			shift = s.Next()
			continue
		}

		nextDayStart := (beginingOfPeriod + (shiftDayStart * day))
		currDayPeriod := shift.Start
		nextDay := shiftDayStart
		days[nextDay-1] += nextDayStart - currDayPeriod
		currDayPeriod = nextDayStart

		for nextDayStart < shift.End {
			if (nextDayStart + day) <= shift.End {
				nextDayStart += day
			} else {
				nextDayStart = shift.End
			}

			nextDay += 1

			days[nextDay-1] += nextDayStart - currDayPeriod
			currDayPeriod = nextDayStart
		}

		fmt.Printf("shift# %d day start %d day end %d\n", shiftCount, shiftDayStart, shiftDayEnd)
		shift = s.Next()
	}

	return days
}
