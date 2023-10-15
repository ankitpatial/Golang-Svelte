package partner

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/ent"
	entDoc "roofix/ent/document"
	"roofix/ent/partnertrainingvideo"
	"roofix/pkg/enum"
	"roofix/pkg/util/storage"
	"time"
)

func TrainingVideoAccess(ctx context.Context, partnerID, videoID string, enabled bool) error {
	db := ent.GetClient()
	defer db.CloseClient()

	// check exists
	first, err := db.PartnerTrainingVideo.
		Query().
		Where(func(ptv *sql.Selector) {
			ptv.Where(sql.And(
				sql.EQ(ptv.C(partnertrainingvideo.VideoColumn), videoID),
				sql.EQ(ptv.C(partnertrainingvideo.PartnerColumn), partnerID),
			))
		}).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	if first != nil {
		err = first.Update().SetEnabled(enabled).Exec(ctx)
	} else {
		err = db.PartnerTrainingVideo.Create().
			SetPartnerID(partnerID).
			SetVideoID(videoID).
			SetEnabled(enabled).
			Exec(ctx)
	}

	return err
}

func TrainingVideos(ctx context.Context, partnerID string) ([]*ent.TrainingVideo, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var res []*ent.TrainingVideo
	err := db.Document.
		Query().
		Where(func(d *sql.Selector) {
			ptv := sql.Table(partnertrainingvideo.Table)
			j := d.LeftJoin(ptv)
			cond := fmt.Sprintf("%s AND %s = '%s'", ptv.C(partnertrainingvideo.VideoColumn), ptv.C(partnertrainingvideo.PartnerColumn), partnerID)
			j.On(d.C(entDoc.FieldID), cond)

			d.Where(sql.And(
				sql.EQ(d.C(entDoc.FieldReady), true),
				sql.EQ(d.C(entDoc.FieldFolder), enum.FolderTrainingVideos),
			))
			d.Select(
				d.C(entDoc.FieldID),
				fmt.Sprintf("%s as name", d.C(entDoc.FieldFilename)),
				fmt.Sprintf("%s  as contentType", d.C(entDoc.FieldContentType)),
				fmt.Sprintf("ifnull(%s, 0) as enabled", ptv.C(partnertrainingvideo.FieldEnabled)),
			)
		}).
		Select().
		Scan(ctx, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func TrainingVideosEnabled(ctx context.Context, partnerID string) ([]*ent.TrainingVideo, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var res []*ent.TrainingVideo
	err := db.Document.
		Query().
		Where(func(d *sql.Selector) {
			ptv := sql.Table(partnertrainingvideo.Table)
			d.LeftJoin(ptv).On(d.C(entDoc.FieldID), ptv.C(partnertrainingvideo.VideoColumn))
			d.Where(sql.And(
				sql.EQ(d.C(entDoc.FieldReady), true),
				sql.EQ(d.C(entDoc.FieldFolder), enum.FolderTrainingVideos),
				sql.EQ(ptv.C(partnertrainingvideo.PartnerColumn), partnerID),
				sql.EQ(ptv.C(partnertrainingvideo.FieldEnabled), true),
			))
			d.Select(
				d.C(entDoc.FieldID),
				fmt.Sprintf("%s as name", d.C(entDoc.FieldFilename)),
				fmt.Sprintf("%s as contentType", d.C(entDoc.FieldContentType)),
				fmt.Sprintf("ifnull(%s, 0) as enabled", ptv.C(partnertrainingvideo.FieldEnabled)),
			)
		}).
		Select().
		Scan(ctx, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func TrainingVideoURL(ctx context.Context, partnerID, videoID string) (string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	doc, err := db.Document.
		Query().
		Where(func(d *sql.Selector) {
			ptv := sql.Table(partnertrainingvideo.Table)
			d.LeftJoin(ptv).On(d.C(entDoc.FieldID), ptv.C(partnertrainingvideo.VideoColumn))
			d.Where(sql.And(
				sql.EQ(d.C(entDoc.FieldReady), true),
				sql.EQ(d.C(entDoc.FieldFolder), enum.FolderTrainingVideos),
				sql.EQ(ptv.C(partnertrainingvideo.PartnerColumn), partnerID),
				sql.EQ(ptv.C(partnertrainingvideo.VideoColumn), videoID),
				sql.EQ(ptv.C(partnertrainingvideo.FieldEnabled), true),
			))
		}).
		Only(ctx)

	if err != nil {
		return "", err
	}

	return storage.GetSignedUrl(ctx, doc.Bucket, doc.Key, time.Hour)
}
