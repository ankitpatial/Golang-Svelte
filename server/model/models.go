package model

import (
	"roofix/config"
	"roofix/ent"
	"roofix/pkg/util"
	"time"
)

type Page struct {
	After  string `json:"after"`
	First  uint16 `json:"first"`
	Before string `json:"before"`
	Last   uint16 `json:"last"`
}

type PageInput struct {
	After  *ent.Cursor `json:"after"`
	First  *int        `json:"first"`
	Before *ent.Cursor `json:"before"`
	Last   *int        `json:"last"`
}

type PartnerJob struct {
	ID              string     `json:"id"`
	OwnerFirstName  string     `json:"owner_first_name"`
	OwnerLastName   string     `json:"owner_last_name"`
	StreetNumber    string     `json:"street_number"`
	StreetName      string     `json:"street_name"`
	City            string     `json:"city"`
	State           string     `json:"state"`
	Region          string     `json:"region"`
	Zip             string     `json:"zip"`
	Latitude        float64    `json:"latitude"`
	Longitude       float64    `json:"longitude"`
	RepFirstName    string     `json:"rep_first_name"`
	RepLastName     string     `json:"rep_last_name"`
	RepEmail        string     `json:"rep_email"`
	RepMobile       string     `json:"rep_mobile"`
	CompanyName     string     `json:"company_name"`
	StatusID        uint8      `json:"status_id"`
	MaterialDate    *time.Time `json:"material_date"`
	RemoveDate      *time.Time `json:"remove_date"`
	InstallDate     *time.Time `json:"install_date"`
	CompletionDate  *time.Time `json:"completion_date"`
	Notes           *string    `json:"notes"`
	Price           *float64   `json:"price"`
	InviteID        string     `json:"invite_id"`
	InviteDate      *time.Time `json:"invited_at"`
	InviteReplyDate *time.Time `json:"invited_reply_at"`
	InviteAccepted  *bool      `json:"accepted"`
}

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetRange in app specific timezone
func (dr DateRange) GetRange() (time.Time, time.Time) {
	now := config.TimeNowAppTZ()
	switch dr {
	case DateRangeToday:
		return util.StartOfDay(now), util.EndOfDay(now)
	case DateRangeYesterday:
		to := util.StartOfDay(now)
		return util.SubtractDays(to, 1), to
	case DateRangeLastWeek:
		lw := util.LastWeek(now)
		return util.StartOfWeek(lw), util.EndOfWeek(lw)
	case DateRangeThisMonth:
		return util.StartOfMonth(now), util.EndOfDay(now)
	case DateRangeLastMonth:
		lm := util.LastMonth(now)
		return util.StartOfMonth(lm), util.EndOfMonth(lm)
	case DateRangeThisYear:
		return util.StartOfYear(now), util.EndOfDay(now)
	case DateRangeLastYear:
		ly := util.SubtractAYear(now)
		return util.StartOfYear(ly), util.EndOfYear(ly)

	default: // THIS_WEEK
		return util.StartOfWeek(now), util.EndOfDay(now)
	}
}
