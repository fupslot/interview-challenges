package main

import (
	"fmt"
	"time"

	"github.com/fupslot/go-shift-report/pkg/common"
)

func main() {
	minute := int64(60)
	curr := time.Date(2022, time.January, 01, 12, 0, 0, 0, time.UTC).Unix()

	s := &common.Shifts{}
	s.AddTime(curr) // opening 1st shift
	curr += 55 * minute
	s.AddTime(curr)

	curr += 350 * minute // waiting

	s.AddTime(curr) // opening 2nd shift
	curr += 165 * minute
	s.AddTime(curr)

	curr += 1440 * minute

	s.AddTime(curr) // opening 3rd shift
	curr += 98 * minute
	s.AddTime(curr)

	s.Traverse()

	// next := s.Next()
	// for next != nil {
	// 	fmt.Printf("S%d E%d DIFF%d\n", next.Start, next.End, (next.End-next.Start)/60)
	// 	next = s.Next()
	// }

	beginingOfMonth := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	// endOfMonth := time.Date(2022, time.January, 31, 24, 59, 59, 0, time.UTC).Unix()

	// |0         24|25        48|49         72|
	// 1|.....ooooo|2|oooooooooo|3|o.oo..oo..|
	//        ^           $
	day := int64(24 * time.Hour / time.Second)
	tail := day

	var days []int64
	dayN := 1
	period := day * 31
	shift := s.Next()
	for tail <= period {
		if len(days) < dayN {
			days = append(days, 0)
		}

		for shift != nil {
			// head - the time in minutes from the beginig of the shift
			// end - the time in minutes when the shift ends, relative to the begining of the month
			head := shift.Start - beginingOfMonth
			end := shift.End - beginingOfMonth

			// when a shift is ahead of a current time period
			// move on to the next day
			if head < 0 || head > tail {
				break
			}

			// body - the duration in minutes of a current shift since head
			body := shift.End - shift.Start

			// if a shift duration expands beyond the current day
			// we add the time differense to the current period then
			// calculate the amount of hours left and slice them by days
			// head moves at the end of a calculatin period
			//
			// if a shift diration fits the current day we add the time
			// difference to it and continue
			if head+body < tail {
				days[dayN-1] += body
				shift = s.Next()
				continue
			}

			days[dayN-1] += tail - head
			for {
				head = tail
				tail += day
				dayN += 1

				if end >= tail {
					days = append(days, day)
					continue
				}

				days = append(days, end-head)
				shift = s.Next()
				break
			}
		}

		dayN += 1
		tail += day
	}

	// Printing results
	var total int64
	for idx, d := range days {
		total += d
		fmt.Printf("D %02d U %.1f\n", idx+1, float64(d)/3600)
	}

	fmt.Printf("Total %.1f hours\n", float64(total)/3600)
}
