package kratos

import (
	"fmt"
	"strings"
	"time"
)

type EpochTimeStamp time.Time

// UnmarshalJSON function for converting UNIX timestamps to time objects
func (t *EpochTimeStamp) UnmarshalJSON(timeStamp []byte) error {
	timeString := string(timeStamp)
	timeString = strings.ReplaceAll(timeString, `"`, "")
	parseTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return fmt.Errorf("time stamp conversion failed: %v", err)
	}
	*(*time.Time)(t) = parseTime

	return nil
}

// Return JSON epoch timestamp as Time.time
func (t EpochTimeStamp) Time() time.Time {
	return time.Time(t)
}
