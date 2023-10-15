package survey

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/config"
	"roofix/ent"
	entSurvey "roofix/ent/survey"
	"roofix/pkg/enum"
	"roofix/server/model"
	"time"
)

func AvailableSlots(ctx context.Context, ctxUserID, date string) ([]*model.SurveySlot, error) {
	_, err := config.TimeParseInAppTZ(time.DateOnly, date)
	if err != nil {
		return nil, ErrInvalidDate
	}

	db := ent.GetClient()
	defer db.CloseClient()

	var slots []*slotCount
	err = db.Survey.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entSurvey.FieldDate), date))
			s.Where(sql.And(
				sql.EQ(s.C(entSurvey.FieldStatus), enum.SurveyStatusRequesting),
				sql.GT(s.C(entSurvey.FieldUntil), time.Now().UTC()),
				sql.NEQ(s.C(entSurvey.CreatedByColumn), ctxUserID),
			))
		}).
		GroupBy(entSurvey.FieldSlotID).
		Aggregate(ent.Count()).
		Scan(ctx, &slots)
	if err != nil {
		return nil, err
	}

	var res []*model.SurveySlot
	now := config.TimeNowAppTZ()
	for _, s := range allSlots() {
		a := &model.SurveySlot{
			ID:        s.ID,
			Name:      s.String(),
			Available: true,
		}

		for _, b := range slots {
			if b.SlotID == s.ID {
				a.Available = b.Count < s.Capacity
				break
			}
		}

		// must not be in past
		st, _ := config.TimeParseInAppTZ(time.DateTime, fmt.Sprintf("%s %s", date, s.From))
		if st.Before(now) {
			a.Available = false
		}

		res = append(res, a)
	}

	return res, nil
}

func allSlots() []*Slot {
	return []*Slot{
		{"01GZ92CV2VTRRE1ATY6TD8698C", "08:00:00", "10:00:00", 1},
		{"01GZ92CV2VT94HF3PH4M43WRGY", "10:00:00", "12:00:00", 1},
		{"01GZ92CV2VKJQYSRF5H7XSEHX1", "12:00:00", "14:00:00", 1},
		{"01GZ92CV2V2Q7S99G0NJB4ADRM", "14:00:00", "16:00:00", 1},
		{"01GZ92CV2V1ZGW5QYH8CV4CV1Q", "16:00:00", "18:00:00", 1},
		{"01GZ92CV2VBJ9G6W2MK8DQV19T", "18:00:00", "20:00:00", 1},
	}
}

func SlotByID(id string) *Slot {
	if id == "" {
		return nil
	}

	for _, s := range allSlots() {
		if s.ID == id {
			return s
		}
	}

	return nil
}
