package TimeUtil

import (
	"time"
	"errors"
)

func ParseTimeToUnixTimeStamp(t string) (time.Time, error) {
	layouts := []string{time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339,
		time.RFC3339Nano, time.RubyDate, time.ANSIC, time.UnixDate, time.Stamp, time.StampMicro,
		time.StampMilli, time.StampNano, "1/2/06", "01/02/2006", "06/01/02", "2006/01/02", "1-2-06", "01-02-2006",
		"06-01-02", "2006-01-02"}

	tm, err := parseToTime(layouts, t)

	return tm, err
}

func parseToTime(layouts []string, t string) (time.Time, error) {

	for _, val := range layouts {
		tm, err := time.Parse(val, t)
		if err == nil {
			return tm, nil
		}
	}

	return time.Now(), errors.New("Failed to parse Time String")
}
