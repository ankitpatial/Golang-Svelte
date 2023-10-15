// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/job"
	"roofix/ent/jobassignmenthistory"
	"roofix/ent/partner"
	"roofix/pkg/enum"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JobAssignmentHistory is the model entity for the JobAssignmentHistory schema.
type JobAssignmentHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// Status holds the value of the "status" field.
	Status enum.JobAssignmentStatus `json:"status,omitempty"`
	// Note holds the value of the "Note" field.
	Note string `json:"Note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobAssignmentHistoryQuery when eager-loading is set.
	Edges        JobAssignmentHistoryEdges `json:"edges"`
	job_id       *string
	partner_id   *string
	selectValues sql.SelectValues
}

// JobAssignmentHistoryEdges holds the relations/edges for other nodes in the graph.
type JobAssignmentHistoryEdges struct {
	// Job holds the value of the job edge.
	Job *Job `json:"job,omitempty"`
	// Partner holds the value of the partner edge.
	Partner *Partner `json:"partner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// JobOrErr returns the Job value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobAssignmentHistoryEdges) JobOrErr() (*Job, error) {
	if e.loadedTypes[0] {
		if e.Job == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: job.Label}
		}
		return e.Job, nil
	}
	return nil, &NotLoadedError{edge: "job"}
}

// PartnerOrErr returns the Partner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobAssignmentHistoryEdges) PartnerOrErr() (*Partner, error) {
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
func (*JobAssignmentHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobassignmenthistory.FieldID, jobassignmenthistory.FieldStatus, jobassignmenthistory.FieldNote:
			values[i] = new(sql.NullString)
		case jobassignmenthistory.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case jobassignmenthistory.ForeignKeys[0]: // job_id
			values[i] = new(sql.NullString)
		case jobassignmenthistory.ForeignKeys[1]: // partner_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobAssignmentHistory fields.
func (jah *JobAssignmentHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jobassignmenthistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				jah.ID = value.String
			}
		case jobassignmenthistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				jah.CreatedAt = value.Time
			}
		case jobassignmenthistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				jah.Status = enum.JobAssignmentStatus(value.String)
			}
		case jobassignmenthistory.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Note", values[i])
			} else if value.Valid {
				jah.Note = value.String
			}
		case jobassignmenthistory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_id", values[i])
			} else if value.Valid {
				jah.job_id = new(string)
				*jah.job_id = value.String
			}
		case jobassignmenthistory.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partner_id", values[i])
			} else if value.Valid {
				jah.partner_id = new(string)
				*jah.partner_id = value.String
			}
		default:
			jah.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobAssignmentHistory.
// This includes values selected through modifiers, order, etc.
func (jah *JobAssignmentHistory) Value(name string) (ent.Value, error) {
	return jah.selectValues.Get(name)
}

// QueryJob queries the "job" edge of the JobAssignmentHistory entity.
func (jah *JobAssignmentHistory) QueryJob() *JobQuery {
	return NewJobAssignmentHistoryClient(jah.config).QueryJob(jah)
}

// QueryPartner queries the "partner" edge of the JobAssignmentHistory entity.
func (jah *JobAssignmentHistory) QueryPartner() *PartnerQuery {
	return NewJobAssignmentHistoryClient(jah.config).QueryPartner(jah)
}

// Update returns a builder for updating this JobAssignmentHistory.
// Note that you need to call JobAssignmentHistory.Unwrap() before calling this method if this JobAssignmentHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (jah *JobAssignmentHistory) Update() *JobAssignmentHistoryUpdateOne {
	return NewJobAssignmentHistoryClient(jah.config).UpdateOne(jah)
}

// Unwrap unwraps the JobAssignmentHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jah *JobAssignmentHistory) Unwrap() *JobAssignmentHistory {
	_tx, ok := jah.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobAssignmentHistory is not a transactional entity")
	}
	jah.config.driver = _tx.drv
	return jah
}

// String implements the fmt.Stringer.
func (jah *JobAssignmentHistory) String() string {
	var builder strings.Builder
	builder.WriteString("JobAssignmentHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jah.ID))
	builder.WriteString("created_at=")
	builder.WriteString(jah.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", jah.Status))
	builder.WriteString(", ")
	builder.WriteString("Note=")
	builder.WriteString(jah.Note)
	builder.WriteByte(')')
	return builder.String()
}

// JobAssignmentHistories is a parsable slice of JobAssignmentHistory.
type JobAssignmentHistories []*JobAssignmentHistory
