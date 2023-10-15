package job

import (
	"roofix/config"
	"roofix/pkg/enum"
	"roofix/pkg/util"
	"roofix/pkg/util/ptr"
	"time"
)

func FlagAtUTC(current enum.JobProgress) *time.Time {
	ns := nextStep(current)
	if ns == nil {
		return nil
	}

	var days uint
	for _, l := range progressLimits {
		if l.step == *ns {
			days = l.days
			break
		}
	}

	if days == 0 {
		return nil
	}

	// add business days to business time.now and also make until EOD of that day.
	dt := addBusinessDays(config.TimeNowAppTZ(), days)
	return ptr.Time(dt.UTC())
}

// addBusinessDays to give time
func addBusinessDays(d time.Time, days uint) time.Time {
	var n time.Time
	count := uint(0)
	for count < days {
		n = d.AddDate(0, 0, 1)
		weekday := n.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			count++
		}

		d = n
	}
	return util.EndOfDay(d)
}
