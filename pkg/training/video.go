package training

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/document"
	"roofix/ent/partnertrainingvideo"
	"roofix/ent/trainingcourse"
	"roofix/ent/trainingvideo"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"roofix/server/model"
	"strings"
	"sync"
)

func SaveVideo(ctx context.Context, ctxUserID string, inp *model.InputTrainingVideo) error {
	db := ent.GetClient()
	defer db.CloseClient()

	title := strings.TrimSpace(inp.Title)
	// check title already exists
	c, err := db.TrainingVideo.Query().
		Where(func(tv *sql.Selector) {
			tv.Where(sql.NEQ(tv.C(trainingvideo.FieldID), inp.ID))
			tv.Where(sql.EQ(tv.C(trainingvideo.CourseColumn), inp.CourseID))
			tv.Where(sql.EQ(tv.C(trainingvideo.FieldTitle), title))
			tv.Where(sql.EQ(tv.C(trainingvideo.FieldTitle), title))
		}).
		Count(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	}
	if c > 0 {
		return msg.AsError(msg.AlreadyExists, "Video with same title")
	}

	// upsert video now.
	return db.TrainingVideo.Create().
		SetID(inp.ID).
		SetKind(inp.Kind).
		SetTitle(title).
		SetDescription(inp.Description).
		SetCourseID(inp.CourseID).
		SetPosterID(inp.PosterID).
		SetVideoID(inp.VideoID).
		SetCreatorID(ctxUserID).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
}

func VideoTypes(ctx context.Context, partnerID string) ([]*model.Entity, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var kinds []enum.TrainingType
	err := db.PartnerTrainingVideo.Query().
		Where(func(ptv *sql.Selector) {
			tv := sql.Table(trainingvideo.Table)
			ptv.Join(tv).On(ptv.C(partnertrainingvideo.VideoColumn), tv.C(trainingvideo.FieldID))
			ptv.Where(sql.And(
				sql.EQ(partnertrainingvideo.PartnerColumn, partnerID),
				sql.EQ(partnertrainingvideo.FieldEnabled, true),
			))
			ptv.Select(tv.C(trainingvideo.FieldKind))
		}).
		Unique(true).
		Select().
		Scan(ctx, &kinds)
	if err != nil {
		return nil, err
	}

	var res []*model.Entity
	for _, e := range kinds {
		res = append(res, &model.Entity{
			ID:   e.String(),
			Name: e.Humanize(),
		})
	}

	return res, err
}

func VideosCourses(
	ctx context.Context, kind enum.TrainingType, partnerID *string, onlyAssigned bool, pageSize int,
) ([]*model.CourseVideos, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	courses, err := db.TrainingCourse.Query().
		Where(trainingcourse.HasTrainingVideosWith(trainingvideo.KindEQ(kind))).
		Order(trainingcourse.ByCreatedAt()).
		All(ctx)
	if err != nil {
		log.Error(err)
		return nil, msg.AsError(msg.ServerError)
	}

	var res []*model.CourseVideos
	var mut sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(courses))
	for _, c := range courses {
		go func(wg *sync.WaitGroup, course *ent.TrainingCourse, p model.PageInput) {
			defer wg.Done()
			// Note: sending courseVideos the `ctx` value is causing issue here, results in not triggering the query
			videos, er := courseVideos(context.Background(), db, kind, course.ID, nil, partnerID, onlyAssigned, &p)
			if er != nil {
				log.Error(er)
			} else if videos.TotalCount > 0 {
				mut.Lock()
				res = append(res, &model.CourseVideos{ID: course.ID, Name: course.Name, Videos: videos})
				mut.Unlock()
			}
		}(&wg, c, model.PageInput{First: &pageSize})
	}
	wg.Wait()
	return res, nil
}

func Videos(
	ctx context.Context, kind enum.TrainingType, courseID string, search, partnerID *string, onlyAssigned bool, page *model.PageInput,
) (*ent.TrainingVideoConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	return courseVideos(ctx, db, kind, courseID, search, partnerID, onlyAssigned, page)
}

// courseVideos
//
// @kind is training video type
// @courseID course under the previous kind
// @partnerID to load partner:video assigned status
// @onlyAssigned will force pull only assigned to give partnerID
// @page pager
func courseVideos(
	ctx context.Context, db *ent.Client, kind enum.TrainingType, courseID string, search, partnerID *string, onlyAssigned bool,
	page *model.PageInput,
) (*ent.TrainingVideoConnection, error) {
	qry := db.TrainingVideo.Query().
		WithCourse().
		WithPoster(func(p *ent.DocumentQuery) {
			p.Select(document.FieldBucket, document.FieldKey)
		}).
		WithVideo(func(v *ent.DocumentQuery) {
			v.Select(document.FieldBucket, document.FieldKey)
		}).
		Where(func(tv *sql.Selector) {
			tv.Where(sql.EQ(tv.C(trainingvideo.FieldKind), kind))
			tv.Where(sql.EQ(tv.C(trainingvideo.CourseColumn), courseID))
		})

	if partnerID != nil {
		qry.WithTrainingVideos(func(ptv *ent.PartnerTrainingVideoQuery) {
			ptv.Where(func(s *sql.Selector) {
				s.Where(sql.EQ(s.C(partnertrainingvideo.PartnerColumn), *partnerID))
			})
		})

		if onlyAssigned {
			qry.Where(func(tv *sql.Selector) {
				ptv := sql.Table(partnertrainingvideo.Table)
				tv.Join(ptv).On(tv.C(trainingvideo.FieldID), ptv.C(partnertrainingvideo.VideoColumn))
				tv.Where(sql.And(
					sql.EQ(ptv.C(partnertrainingvideo.PartnerColumn), *partnerID),
					sql.EQ(ptv.C(partnertrainingvideo.FieldEnabled), true),
				))
			})
		}
	}

	q := str.Val(search)
	if q != "" {
		qry.Where(trainingvideo.Title(q))
	}
	by := &ent.TrainingVideoOrder{
		Direction: entgql.OrderDirectionAsc,
		Field:     ent.TrainingVideoOrderFieldCreatedAt,
	}
	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithTrainingVideoOrder(by))
}
