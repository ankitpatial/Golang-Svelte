// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/trainingcourse"
	"roofix/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TrainingCourse is the model entity for the TrainingCourse schema.
type TrainingCourse struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TrainingCourseQuery when eager-loading is set.
	Edges        TrainingCourseEdges `json:"edges"`
	creator_id   *string
	selectValues sql.SelectValues
}

// TrainingCourseEdges holds the relations/edges for other nodes in the graph.
type TrainingCourseEdges struct {
	// TrainingVideos holds the value of the training_videos edge.
	TrainingVideos []*TrainingVideo `json:"training_videos,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTrainingVideos map[string][]*TrainingVideo
}

// TrainingVideosOrErr returns the TrainingVideos value or an error if the edge
// was not loaded in eager-loading.
func (e TrainingCourseEdges) TrainingVideosOrErr() ([]*TrainingVideo, error) {
	if e.loadedTypes[0] {
		return e.TrainingVideos, nil
	}
	return nil, &NotLoadedError{edge: "training_videos"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TrainingCourseEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TrainingCourse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case trainingcourse.FieldID, trainingcourse.FieldName:
			values[i] = new(sql.NullString)
		case trainingcourse.FieldCreatedAt, trainingcourse.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case trainingcourse.ForeignKeys[0]: // creator_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TrainingCourse fields.
func (tc *TrainingCourse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case trainingcourse.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				tc.ID = value.String
			}
		case trainingcourse.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tc.CreatedAt = value.Time
			}
		case trainingcourse.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tc.UpdatedAt = value.Time
			}
		case trainingcourse.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tc.Name = value.String
			}
		case trainingcourse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_id", values[i])
			} else if value.Valid {
				tc.creator_id = new(string)
				*tc.creator_id = value.String
			}
		default:
			tc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TrainingCourse.
// This includes values selected through modifiers, order, etc.
func (tc *TrainingCourse) Value(name string) (ent.Value, error) {
	return tc.selectValues.Get(name)
}

// QueryTrainingVideos queries the "training_videos" edge of the TrainingCourse entity.
func (tc *TrainingCourse) QueryTrainingVideos() *TrainingVideoQuery {
	return NewTrainingCourseClient(tc.config).QueryTrainingVideos(tc)
}

// QueryCreator queries the "creator" edge of the TrainingCourse entity.
func (tc *TrainingCourse) QueryCreator() *UserQuery {
	return NewTrainingCourseClient(tc.config).QueryCreator(tc)
}

// Update returns a builder for updating this TrainingCourse.
// Note that you need to call TrainingCourse.Unwrap() before calling this method if this TrainingCourse
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TrainingCourse) Update() *TrainingCourseUpdateOne {
	return NewTrainingCourseClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the TrainingCourse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TrainingCourse) Unwrap() *TrainingCourse {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TrainingCourse is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TrainingCourse) String() string {
	var builder strings.Builder
	builder.WriteString("TrainingCourse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(tc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(tc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedTrainingVideos returns the TrainingVideos named value or an error if the edge was not
// loaded in eager-loading with this name.
func (tc *TrainingCourse) NamedTrainingVideos(name string) ([]*TrainingVideo, error) {
	if tc.Edges.namedTrainingVideos == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := tc.Edges.namedTrainingVideos[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (tc *TrainingCourse) appendNamedTrainingVideos(name string, edges ...*TrainingVideo) {
	if tc.Edges.namedTrainingVideos == nil {
		tc.Edges.namedTrainingVideos = make(map[string][]*TrainingVideo)
	}
	if len(edges) == 0 {
		tc.Edges.namedTrainingVideos[name] = []*TrainingVideo{}
	} else {
		tc.Edges.namedTrainingVideos[name] = append(tc.Edges.namedTrainingVideos[name], edges...)
	}
}

// TrainingCourses is a parsable slice of TrainingCourse.
type TrainingCourses []*TrainingCourse
