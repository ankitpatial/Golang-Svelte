package survey

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrOnQuery          = errors.New("unexpected error on querying DB")
	ErrDoesNotBelong    = errors.New("this survey does not belong to you, so you cannot change it")
	ErrAlreadySubmitted = errors.New("request for this survey is already submitted")
	ErrRequestNotFound  = errors.New("unable to find survey request info")
	ErrSlotNotFound     = errors.New("slot for given ID does not exists")
	ErrInvalidDate      = errors.New("survey scheduled date must be in YYYY-MM-DD format")
	ErrDateIsInPast     = errors.New("survey can not be scheduled for past DateTime")
	ErrSlotNotAvailable = errors.New("all slots for give date and time are reserved, please try some other date or time")
)

type slotCount struct {
	SlotID string `json:"slot_id"`
	Count  int    `json:"count"`
}

type Slot struct {
	ID       string
	From     string
	To       string
	Capacity int
}

func (s *Slot) String() string {
	f, _ := time.Parse(time.TimeOnly, s.From)
	t, _ := time.Parse(time.TimeOnly, s.To)

	return fmt.Sprintf("%s-%s", f.Format("3"), t.Format("3 PM"))
}
