package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/trainingvideo"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/notification"
	"roofix/pkg/partner"
	"roofix/pkg/training"
	"roofix/pkg/util/log"
	"roofix/pkg/util/ptr"
	"roofix/pkg/util/storage"
	"roofix/server/generated"
	"roofix/server/model"
	"time"
)

// CreateTrainingCourse is the resolver for the createTrainingCourse field.
func (r *mutationResolver) CreateTrainingCourse(ctx context.Context, name string) (*ent.TrainingCourse, error) {
	return training.CreateCourse(ctx, account.CtxUserID(ctx), name)
}

// SaveTrainingVideo is the resolver for the saveTrainingVideo field.
func (r *mutationResolver) SaveTrainingVideo(ctx context.Context, inp *model.InputTrainingVideo) (bool, error) {
	if err := training.SaveVideo(ctx, account.CtxUserID(ctx), inp); err != nil {
		return false, err
	}

	m := notification.Message{
		Channel: enum.ChannelTrainingVideo,
		Topic:   enum.TopicChange,
		Title:   "Training Video",
		Message: "Saved training video " + inp.Title,
		Data: map[string]interface{}{
			"id":   inp.ID,
			"kind": inp.Kind,
		},
	}
	notification.Send(ctx, &m, notification.ToAdmins())
	return true, nil
}

// PartnerTrainingVideoAccess is the resolver for the partnerTrainingVideoAccess field.
func (r *mutationResolver) PartnerTrainingVideoAccess(ctx context.Context, partnerID string, videoID string, enabled bool) (bool, error) {
	if err := partner.TrainingVideoAccess(ctx, partnerID, videoID, enabled); err != nil {
		return false, err
	}

	db := ent.GetClient()
	defer db.CloseClient()
	tv, err := db.TrainingVideo.Query().Where(trainingvideo.ID(videoID)).Select(trainingvideo.FieldKind).Only(ctx)
	if tv != nil {
		log.Error(err)
	} else {
		m := notification.Message{
			Channel: enum.ChannelTrainingVideo,
			Topic:   enum.TopicAssigned,
			Title:   "Training Video",
			Message: "New training video assigned " + tv.Title,
			Data: map[string]interface{}{
				"id":   tv.ID,
				"kind": tv.Kind,
			},
		}
		notification.SendAndSave(ctx, &m, notification.ToAdmins(), notification.ToCompanyAdmins(partnerID))
	}
	return true, nil
}

// TrainingCourses is the resolver for the trainingCourses field.
func (r *queryResolver) TrainingCourses(ctx context.Context, search *string, page model.PageInput) (*ent.TrainingCourseConnection, error) {
	return training.Courses(ctx, search, &page)
}

// TrainingVideoKinds is the resolver for the trainingVideoKinds field.
func (r *queryResolver) TrainingVideoKinds(ctx context.Context) ([]*model.Entity, error) {
	var res []*model.Entity
	for _, k := range enum.AllTrainingTypes {
		res = append(res, &model.Entity{ID: k.String(), Name: k.Humanize()})
	}
	return res, nil
}

// TrainingVideoCourses is the resolver for the trainingVideoCourses field.
func (r *queryResolver) TrainingVideoCourses(ctx context.Context, kind enum.TrainingType, partnerID *string, pageSize *int) ([]*model.CourseVideos, error) {
	s := 10
	if pageSize != nil && *pageSize > 0 {
		s = *pageSize
	}
	return training.VideosCourses(ctx, kind, partnerID, false, s)
}

// TrainingVideos is the resolver for the trainingVideos field.
func (r *queryResolver) TrainingVideos(ctx context.Context, kind enum.TrainingType, courseID string, search *string, partnerID *string, page model.PageInput) (*ent.TrainingVideoConnection, error) {
	return training.Videos(ctx, kind, courseID, search, partnerID, false, &page)
}

// MyTrainingVideoKinds is the resolver for the myTrainingVideoKinds field.
func (r *queryResolver) MyTrainingVideoKinds(ctx context.Context) ([]*model.Entity, error) {
	u := account.CtxUser(ctx)
	if u.Partner == nil {
		return nil, msg.AsError(msg.FailedToFind, "Partner")
	}

	return training.VideoTypes(ctx, u.Partner.ID)
}

// MyTrainingVideoCourses is the resolver for the myTrainingVideoCourses field.
func (r *queryResolver) MyTrainingVideoCourses(ctx context.Context, kind enum.TrainingType, pageSize *int) ([]*model.CourseVideos, error) {
	u := account.CtxUser(ctx)
	if u.Partner == nil {
		return nil, msg.AsError(msg.FailedToFind, "Partner")
	}

	s := 10
	if pageSize != nil && *pageSize > 0 {
		s = *pageSize
	}
	return training.VideosCourses(ctx, kind, &u.Partner.ID, true, s)
}

// MyTrainingVideos is the resolver for the myTrainingVideos field.
func (r *queryResolver) MyTrainingVideos(ctx context.Context, kind enum.TrainingType, courseID string, search *string, page model.PageInput) (*ent.TrainingVideoConnection, error) {
	u := account.CtxUser(ctx)
	if u.Partner == nil {
		return nil, msg.AsError(msg.FailedToFind, "Partner")
	}

	return training.Videos(ctx, kind, courseID, search, &u.Partner.ID, true, &page)
}

// PosterURL is the resolver for the posterURL field.
func (r *trainingVideoResolver) PosterURL(ctx context.Context, obj *ent.TrainingVideo) (string, error) {
	if obj == nil || obj.Edges.Poster == nil {
		return "", nil
	}

	return fmt.Sprintf("%s/%s", config.Read().Website.AssetUrl, obj.Edges.Poster.Key), nil
}

// VideoURL is the resolver for the videoURL field.
func (r *trainingVideoResolver) VideoURL(ctx context.Context, obj *ent.TrainingVideo) (string, error) {
	if obj == nil || obj.Edges.Video == nil {
		return "", nil
	}

	e := obj.Edges.Video
	return storage.GetSignedUrl(ctx, e.Bucket, e.Key, time.Hour*24)
}

// Assigned is the resolver for the assigned field.
func (r *trainingVideoResolver) Assigned(ctx context.Context, obj *ent.TrainingVideo) (*bool, error) {
	if obj == nil {
		return nil, nil
	}

	if len(obj.Edges.TrainingVideos) > 0 {
		return &obj.Edges.TrainingVideos[0].Enabled, nil
	}

	return ptr.Bool(false), nil
}

// TrainingVideo returns generated.TrainingVideoResolver implementation.
func (r *Resolver) TrainingVideo() generated.TrainingVideoResolver { return &trainingVideoResolver{r} }

type trainingVideoResolver struct{ *Resolver }
