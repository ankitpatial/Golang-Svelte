package training

import (
	"context"
	"entgo.io/contrib/entgql"
	"roofix/ent"
	"roofix/ent/trainingcourse"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"roofix/server/model"
	"strings"
)

// CreateCourse name
func CreateCourse(ctx context.Context, creatorID, name string) (*ent.TrainingCourse, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	n := strings.TrimSpace(name)

	// check, already exist
	count, err := db.TrainingCourse.Query().Where(trainingcourse.NameEQ(n)).Count(ctx)
	if err != nil {
		log.Error(err)
		return nil, msg.AsError(msg.ServerError)
	}
	if count > 0 {
		return nil, msg.AsError(msg.AlreadyExists, "Course Name")
	}

	// save
	return db.TrainingCourse.Create().SetCreatorID(creatorID).SetName(n).Save(ctx)
}

func Courses(ctx context.Context, search *string, page *model.PageInput) (*ent.TrainingCourseConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.TrainingCourse.Query()
	q := str.Val(search)
	if q != "" {
		qry.Where(trainingcourse.NameContainsFold(q))
	}

	orderBy := &ent.TrainingCourseOrder{
		Direction: entgql.OrderDirectionAsc,
		Field:     ent.TrainingCourseOrderFieldCreatedAt,
	}
	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithTrainingCourseOrder(orderBy))
}
