// Code generated by ent, DO NOT EDIT.

package trainingvideo

import (
	"fmt"
	"roofix/pkg/enum"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
)

const (
	// Label holds the string label denoting the trainingvideo type in the database.
	Label = "training_video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeTrainingVideos holds the string denoting the training_videos edge name in mutations.
	EdgeTrainingVideos = "training_videos"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// EdgePoster holds the string denoting the poster edge name in mutations.
	EdgePoster = "poster"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// Table holds the table name of the trainingvideo in the database.
	Table = "training_videos"
	// TrainingVideosTable is the table that holds the training_videos relation/edge.
	TrainingVideosTable = "partner_training_videos"
	// TrainingVideosInverseTable is the table name for the PartnerTrainingVideo entity.
	// It exists in this package in order to avoid circular dependency with the "partnertrainingvideo" package.
	TrainingVideosInverseTable = "partner_training_videos"
	// TrainingVideosColumn is the table column denoting the training_videos relation/edge.
	TrainingVideosColumn = "video_id"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "training_videos"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "creator_id"
	// CourseTable is the table that holds the course relation/edge.
	CourseTable = "training_videos"
	// CourseInverseTable is the table name for the TrainingCourse entity.
	// It exists in this package in order to avoid circular dependency with the "trainingcourse" package.
	CourseInverseTable = "training_courses"
	// CourseColumn is the table column denoting the course relation/edge.
	CourseColumn = "course_id"
	// PosterTable is the table that holds the poster relation/edge.
	PosterTable = "training_videos"
	// PosterInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	PosterInverseTable = "documents"
	// PosterColumn is the table column denoting the poster relation/edge.
	PosterColumn = "poster_id"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "training_videos"
	// VideoInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	VideoInverseTable = "documents"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_id"
)

// Columns holds all SQL columns for trainingvideo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldKind,
	FieldTitle,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "training_videos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"video_id",
	"poster_id",
	"course_id",
	"creator_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k enum.TrainingType) error {
	switch k.String() {
	case "ROOFING", "SOLAR", "SITE_SURVEY":
		return nil
	default:
		return fmt.Errorf("trainingvideo: invalid enum value for kind field: %q", k)
	}
}

// OrderOption defines the ordering options for the TrainingVideo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByKind orders the results by the kind field.
func ByKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKind, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByTrainingVideosCount orders the results by training_videos count.
func ByTrainingVideosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTrainingVideosStep(), opts...)
	}
}

// ByTrainingVideos orders the results by training_videos terms.
func ByTrainingVideos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTrainingVideosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCourseField orders the results by course field.
func ByCourseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCourseStep(), sql.OrderByField(field, opts...))
	}
}

// ByPosterField orders the results by poster field.
func ByPosterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPosterStep(), sql.OrderByField(field, opts...))
	}
}

// ByVideoField orders the results by video field.
func ByVideoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoStep(), sql.OrderByField(field, opts...))
	}
}
func newTrainingVideosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TrainingVideosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TrainingVideosTable, TrainingVideosColumn),
	)
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
	)
}
func newCourseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CourseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
	)
}
func newPosterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PosterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, PosterTable, PosterColumn),
	)
}
func newVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, VideoTable, VideoColumn),
	)
}

var (
	// enum.TrainingType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*enum.TrainingType)(nil)
	// enum.TrainingType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enum.TrainingType)(nil)
)