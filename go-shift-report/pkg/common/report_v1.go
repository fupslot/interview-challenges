package common

import "time"

func ReportV1(beginingOfMonth int64, s *Shifts) []int64 {
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
	return days
}
