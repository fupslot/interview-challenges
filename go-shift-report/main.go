package main

import (
	"fmt"
	"time"

	"github.com/fupslot/go-shift-report/pkg/common"
)

type Shift struct {
	checkIn  int64
	checkOut int64
}

var shifts []Shift

func lastEndTime(s *common.Shifts) int64 {
	return s.Last().Data.End
}

func main() {
	hours := int64(time.Hour / time.Second)
	curr := time.Date(2022, time.January, 01, 12, 0, 0, 0, time.UTC).Unix()

	s := &common.Shifts{}
	s.AddTime(curr) // opening 1st shift

	curr += 8 * hours
	s.AddTime(curr)

	curr += 18 * hours // waiting

	s.AddTime(curr) // opening 2nd shift
	curr += 13 * hours
	s.AddTime(curr)

	curr += 48 * hours
	s.AddTime(curr) // opening 3rd shift
	curr += 36 * hours
	s.AddTime(curr)

	s.Traverse()
	fmt.Printf("Size() %d", s.Size())

}

// "shift" respresented by two field "start" & "end" both are in "unix time" type
//
func main_A() {
	hours := int64(time.Hour / time.Second)
	curr := time.Date(2022, time.January, 01, 12, 0, 0, 0, time.UTC).Unix()

	s := &common.Shifts{}
	s.AddTime(curr) // opening 1st shift

	curr += 8 * hours
	s.AddTime(curr)

	curr += 18 * hours // waiting

	s.AddTime(curr) // opening 2nd shift
	curr += 13 * hours
	s.AddTime(curr)

	curr += 48 * hours
	s.AddTime(curr) // opening 3rd shift
	curr += 36 * hours
	s.AddTime(curr)

	s.Traverse()
	fmt.Printf("Size() %d", s.Size())

	beginingOfMonth := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	// endOfMonth := time.Date(2022, time.January, 31, 24, 59, 59, 0, time.UTC).Unix()

	// |0         24|25        48|49         72|
	// 1|.....ooooo|2|oooooooooo|3|o.oo..oo..|
	//        ^           $
	day := int64(24 * time.Hour / time.Minute)
	// head := 0
	tail := day

	var days []int64
	dayN := 1
	shiftN := 0
	period := day * 31
	for tail <= period {
		if len(days) < dayN {
			days = append(days, 0)
		}

		for i := shiftN; i < len(shifts); i++ {
			shift := shifts[i]

			// head - the time in minutes from the beginig of the shift
			// end - the time in minutes when the shift ends, relative to the begining of the month
			head := (shift.checkIn - beginingOfMonth) / 60
			end := (shift.checkOut - beginingOfMonth) / 60
			// this shift is ahead of a current time period
			// move on to the next day
			if head < 0 && head > tail {
				break
			}

			// body - the duration in minutes of a current shift since head
			body := (shift.checkOut - shift.checkIn) / 60

			// if a shift duration expands beyond the current day
			// we add the time differense to the current period then
			// calculate the amount of hours left and slice them by days
			// head moves at the end of a calculatin period
			//
			// if a shift diration fits the current day we add the time
			// difference to it and continue
			if head+body < tail {
				days[dayN-1] += body
				shiftN += i
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
				shiftN += 1
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
		fmt.Printf("D %02d U %.1f\n", idx+1, float64(d/60))
	}

	fmt.Printf("Total %.1f hours\n", float64(total/60))
}
