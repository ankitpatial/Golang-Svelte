package survey

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/config"
	"roofix/ent"
	entSurvey "roofix/ent/survey"
	"roofix/pkg/enum"
	"roofix/pkg/partner"
	"time"
)

func Request(ctx context.Context, creatorID string, date, slotID string) (string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// get Slot by its ID
	available, s, err := isSlotAvailable(ctx, db, creatorID, date, slotID)
	if err != nil {
		return "", err
	} else if !available {
		return "", ErrSlotNotAvailable
	}

	existing, err := db.Survey.
		Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.And(
				sql.EQ(s.C(entSurvey.CreatedByColumn), creatorID),
				sql.EQ(s.C(entSurvey.FieldStatus), enum.SurveyStatusRequesting),
				sql.EQ(s.C(entSurvey.FieldDate), date),
				sql.EQ(s.C(entSurvey.FieldSlotID), slotID),
			))
		}).
		Select(entSurvey.FieldID).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", err
	}

	span := time.Minute * 30 // time a REQUESTING request will live

	// existing record exist, just update only until date
	if existing != nil {
		err := db.Survey.UpdateOneID(existing.ID).
			SetUntil(time.Now().Add(span)).
			Exec(ctx)
		if err != nil {
			return "", err
		}

		return existing.ID, nil
	}

	// merge date and Slot times To get From & To dates
	f, _ := config.TimeParseInAppTZ(time.DateTime, fmt.Sprintf("%s %s", date, s.From))
	t, _ := config.TimeParseInAppTZ(time.DateTime, fmt.Sprintf("%s %s", date, s.To))

	// create a new survey  with the provided date, time, and a status of "REQUESTING"
	qry := db.Survey.Create().
		SetDate(date).
		SetSlot(s.String()).
		SetSlotID(s.ID).
		SetFrom(f.UTC()).
		SetTo(t.UTC()).
		SetUntil(time.Now().Add(span)). // request will be active until given time
		SetStatus(enum.SurveyStatusRequesting).
		SetCreatedByID(creatorID)

	// set related partner ID
	p, _ := partner.ByUserID(ctx, creatorID)
	if p != nil {
		qry.SetPartnerID(p.ID)
	}

	survey, err := qry.Save(ctx)
	if err != nil {
		return "", err
	}

	return survey.ID, nil
}

func isSlotAvailable(ctx context.Context, db *ent.Client, ctxUserID, date, slotID string) (bool, *Slot, error) {
	// parse date
	_, err := config.TimeParseInAppTZ(time.DateOnly, date)
	if err != nil {
		return false, nil, ErrInvalidDate
	}

	// find slot
	s := SlotByID(slotID)
	if s == nil {
		return false, nil, ErrSlotNotFound
	}

	// ensure not in past
	now := config.TimeNowAppTZ()
	st, _ := config.TimeParseInAppTZ(time.DateTime, fmt.Sprintf("%s %s", date, s.From))
	if st.Before(now) { // asked for past date-time
		return false, nil, ErrDateIsInPast
	}

	var reserved []*slotCount
	err = db.Survey.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entSurvey.FieldDate), date))
			s.Where(sql.EQ(s.C(entSurvey.FieldSlotID), slotID))
			s.Where(sql.And(
				sql.EQ(s.C(entSurvey.FieldStatus), enum.SurveyStatusRequesting),
				sql.GT(s.C(entSurvey.FieldUntil), time.Now().UTC()),
				sql.NEQ(s.C(entSurvey.CreatedByColumn), ctxUserID),
			))
		}).
		GroupBy(entSurvey.FieldSlotID).
		Aggregate(ent.Count()).
		Scan(ctx, &reserved)
	if err != nil {
		return false, nil, err
	}

	// not record in db, so good to go
	if len(reserved) == 0 {
		return true, s, nil
	}

	// check capacity
	if reserved[0].Count >= s.Capacity {
		return false, nil, nil
	}

	return true, s, err
}
