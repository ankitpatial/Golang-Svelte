// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/partnertrainingvideo"
	"roofix/ent/trainingvideo"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PartnerTrainingVideo is the model entity for the PartnerTrainingVideo schema.
type PartnerTrainingVideo struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartnerTrainingVideoQuery when eager-loading is set.
	Edges        PartnerTrainingVideoEdges `json:"edges"`
	partner_id   *string
	video_id     *string
	selectValues sql.SelectValues
}

// PartnerTrainingVideoEdges holds the relations/edges for other nodes in the graph.
type PartnerTrainingVideoEdges struct {
	// Video holds the value of the video edge.
	Video *TrainingVideo `json:"video,omitempty"`
	// Partner holds the value of the partner edge.
	Partner *Partner `json:"partner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// VideoOrErr returns the Video value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartnerTrainingVideoEdges) VideoOrErr() (*TrainingVideo, error) {
	if e.loadedTypes[0] {
		if e.Video == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: trainingvideo.Label}
		}
		return e.Video, nil
	}
	return nil, &NotLoadedError{edge: "video"}
}

// PartnerOrErr returns the Partner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartnerTrainingVideoEdges) PartnerOrErr() (*Partner, error) {
	if e.loadedTypes[1] {
		if e.Partner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: partner.Label}
		}
		return e.Partner, nil
	}
	return nil, &NotLoadedError{edge: "partner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PartnerTrainingVideo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case partnertrainingvideo.FieldEnabled:
			values[i] = new(sql.NullBool)
		case partnertrainingvideo.FieldID:
			values[i] = new(sql.NullString)
		case partnertrainingvideo.FieldCreatedAt, partnertrainingvideo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case partnertrainingvideo.ForeignKeys[0]: // partner_id
			values[i] = new(sql.NullString)
		case partnertrainingvideo.ForeignKeys[1]: // video_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartnerTrainingVideo fields.
func (ptv *PartnerTrainingVideo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partnertrainingvideo.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ptv.ID = value.String
			}
		case partnertrainingvideo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ptv.CreatedAt = value.Time
			}
		case partnertrainingvideo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ptv.UpdatedAt = value.Time
			}
		case partnertrainingvideo.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				ptv.Enabled = value.Bool
			}
		case partnertrainingvideo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partner_id", values[i])
			} else if value.Valid {
				ptv.partner_id = new(string)
				*ptv.partner_id = value.String
			}
		case partnertrainingvideo.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				ptv.video_id = new(string)
				*ptv.video_id = value.String
			}
		default:
			ptv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PartnerTrainingVideo.
// This includes values selected through modifiers, order, etc.
func (ptv *PartnerTrainingVideo) Value(name string) (ent.Value, error) {
	return ptv.selectValues.Get(name)
}

// QueryVideo queries the "video" edge of the PartnerTrainingVideo entity.
func (ptv *PartnerTrainingVideo) QueryVideo() *TrainingVideoQuery {
	return NewPartnerTrainingVideoClient(ptv.config).QueryVideo(ptv)
}

// QueryPartner queries the "partner" edge of the PartnerTrainingVideo entity.
func (ptv *PartnerTrainingVideo) QueryPartner() *PartnerQuery {
	return NewPartnerTrainingVideoClient(ptv.config).QueryPartner(ptv)
}

// Update returns a builder for updating this PartnerTrainingVideo.
// Note that you need to call PartnerTrainingVideo.Unwrap() before calling this method if this PartnerTrainingVideo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ptv *PartnerTrainingVideo) Update() *PartnerTrainingVideoUpdateOne {
	return NewPartnerTrainingVideoClient(ptv.config).UpdateOne(ptv)
}

// Unwrap unwraps the PartnerTrainingVideo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ptv *PartnerTrainingVideo) Unwrap() *PartnerTrainingVideo {
	_tx, ok := ptv.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartnerTrainingVideo is not a transactional entity")
	}
	ptv.config.driver = _tx.drv
	return ptv
}

// String implements the fmt.Stringer.
func (ptv *PartnerTrainingVideo) String() string {
	var builder strings.Builder
	builder.WriteString("PartnerTrainingVideo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ptv.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ptv.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ptv.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", ptv.Enabled))
	builder.WriteByte(')')
	return builder.String()
}

// PartnerTrainingVideos is a parsable slice of PartnerTrainingVideo.
type PartnerTrainingVideos []*PartnerTrainingVideo