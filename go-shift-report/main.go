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

	curr += 120 * minute

	s.AddTime(curr) // opening 4rd shift
	curr += 9809 * minute
	s.AddTime(curr)

	s.Traverse()

	beginingOfMonth := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()

	days := common.ReportV2(beginingOfMonth, 31, s)

	// Printing results
	var total int64
	for idx, d := range days {
		total += d
		fmt.Printf("D %02d U %.1f\n", idx+1, float64(d)/3600)
	}

	fmt.Printf("Total %.1f hours\n", float64(total)/3600)
}
