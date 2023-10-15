package survey

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"roofix/ent"
	entSurvey "roofix/ent/survey"
	"roofix/ent/user"
	entUser "roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/util"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"roofix/server/model"
)

func ByID(ctx context.Context, creatorID string, id string) (*ent.Survey, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	if id == "" {
		return nil, errors.New("invalid survey ID")
	}

	if creatorID == "" {
		return nil, errors.New("invalid creator ID")
	}

	survey, err := db.Survey.Query().
		Where(entSurvey.ID(id)).
		Where(entSurvey.HasCreatedByWith(entUser.ID(creatorID))).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed To get survey: %v", err)
	}

	return survey, nil
}

func List(
	ctx context.Context, creatorID *string, progress *enum.SurveyProgress, search *string, betweenDates []string,
	page model.PageInput, orderBy *ent.SurveyOrder,
) (*ent.SurveyConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Survey.Query().Where(entSurvey.StatusNEQ(enum.SurveyStatusRequesting))

	// creator filter
	if creatorID != nil {
		qry.Where(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(entSurvey.CreatedByColumn), *creatorID))
		})
	}

	// progress filter
	if progress != nil {
		qry.Where(entSurvey.ProgressEQ(*progress))
	}

	// search filter
	q := str.Val(search)
	if q != "" {
		qry.Where(entSurvey.Or(
			entSurvey.NameContainsFold(q),
			entSurvey.AddressContainsFold(q),
			entSurvey.PhoneContainsFold(q),
		))
	}

	// date filters
	if len(betweenDates) == 2 {
		f := util.ParseISODate(betweenDates[0])
		t := util.ParseISODate(betweenDates[1])
		qry.Where(entSurvey.ProgressAtGTE(f), entSurvey.ProgressAtLTE(t))
	}

	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithSurveyOrder(orderBy))
}

func ByIdAndCreatorID(ctx context.Context, db *ent.Client, surveyID, creatorID string) (*ent.Survey, error) {
	s, err := db.Survey.Query().
		Where(entSurvey.ID(surveyID)).
		WithCreatedBy(func(c *ent.UserQuery) {
			c.Select(user.FieldID)
		}).
		Select(entSurvey.FieldStatus).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrRequestNotFound
		}

		log.Error(err)
		return nil, ErrOnQuery
	}

	// check creator relation
	if s.Edges.CreatedBy.ID != creatorID {
		return nil, ErrDoesNotBelong
	}

	return s, nil
}
