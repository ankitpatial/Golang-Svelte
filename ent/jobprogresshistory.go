// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/apiuser"
	"roofix/ent/job"
	"roofix/ent/jobprogresshistory"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JobProgressHistory is the model entity for the JobProgressHistory schema.
type JobProgressHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// Status holds the value of the "status" field.
	Status enum.JobProgress `json:"status,omitempty"`
	// Complete holds the value of the "complete" field.
	Complete bool `json:"complete,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobProgressHistoryQuery when eager-loading is set.
	Edges        JobProgressHistoryEdges `json:"edges"`
	api_user_id  *string
	job_id       *string
	user_id      *string
	selectValues sql.SelectValues
}

// JobProgressHistoryEdges holds the relations/edges for other nodes in the graph.
type JobProgressHistoryEdges struct {
	// Job holds the value of the job edge.
	Job *Job `json:"job,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// CreatorAPIUser holds the value of the creator_api_user edge.
	CreatorAPIUser *ApiUser `json:"creator_api_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// JobOrErr returns the Job value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobProgressHistoryEdges) JobOrErr() (*Job, error) {
	if e.loadedTypes[0] {
		if e.Job == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: job.Label}
		}
		return e.Job, nil
	}
	return nil, &NotLoadedError{edge: "job"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobProgressHistoryEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// CreatorAPIUserOrErr returns the CreatorAPIUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobProgressHistoryEdges) CreatorAPIUserOrErr() (*ApiUser, error) {
	if e.loadedTypes[2] {
		if e.CreatorAPIUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: apiuser.Label}
		}
		return e.CreatorAPIUser, nil
	}
	return nil, &NotLoadedError{edge: "creator_api_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobProgressHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobprogresshistory.FieldComplete:
			values[i] = new(sql.NullBool)
		case jobprogresshistory.FieldID, jobprogresshistory.FieldStatus, jobprogresshistory.FieldNote:
			values[i] = new(sql.NullString)
		case jobprogresshistory.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case jobprogresshistory.ForeignKeys[0]: // api_user_id
			values[i] = new(sql.NullString)
		case jobprogresshistory.ForeignKeys[1]: // job_id
			values[i] = new(sql.NullString)
		case jobprogresshistory.ForeignKeys[2]: // user_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobProgressHistory fields.
func (jph *JobProgressHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jobprogresshistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				jph.ID = value.String
			}
		case jobprogresshistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				jph.CreatedAt = value.Time
			}
		case jobprogresshistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				jph.Status = enum.JobProgress(value.String)
			}
		case jobprogresshistory.FieldComplete:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field complete", values[i])
			} else if value.Valid {
				jph.Complete = value.Bool
			}
		case jobprogresshistory.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				jph.Note = value.String
			}
		case jobprogresshistory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_user_id", values[i])
			} else if value.Valid {
				jph.api_user_id = new(string)
				*jph.api_user_id = value.String
			}
		case jobprogresshistory.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_id", values[i])
			} else if value.Valid {
				jph.job_id = new(string)
				*jph.job_id = value.String
			}
		case jobprogresshistory.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				jph.user_id = new(string)
				*jph.user_id = value.String
			}
		default:
			jph.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobProgressHistory.
// This includes values selected through modifiers, order, etc.
func (jph *JobProgressHistory) Value(name string) (ent.Value, error) {
	return jph.selectValues.Get(name)
}

// QueryJob queries the "job" edge of the JobProgressHistory entity.
func (jph *JobProgressHistory) QueryJob() *JobQuery {
	return NewJobProgressHistoryClient(jph.config).QueryJob(jph)
}

// QueryCreator queries the "creator" edge of the JobProgressHistory entity.
func (jph *JobProgressHistory) QueryCreator() *UserQuery {
	return NewJobProgressHistoryClient(jph.config).QueryCreator(jph)
}

// QueryCreatorAPIUser queries the "creator_api_user" edge of the JobProgressHistory entity.
func (jph *JobProgressHistory) QueryCreatorAPIUser() *ApiUserQuery {
	return NewJobProgressHistoryClient(jph.config).QueryCreatorAPIUser(jph)
}

// Update returns a builder for updating this JobProgressHistory.
// Note that you need to call JobProgressHistory.Unwrap() before calling this method if this JobProgressHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (jph *JobProgressHistory) Update() *JobProgressHistoryUpdateOne {
	return NewJobProgressHistoryClient(jph.config).UpdateOne(jph)
}

// Unwrap unwraps the JobProgressHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jph *JobProgressHistory) Unwrap() *JobProgressHistory {
	_tx, ok := jph.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobProgressHistory is not a transactional entity")
	}
	jph.config.driver = _tx.drv
	return jph
}

// String implements the fmt.Stringer.
func (jph *JobProgressHistory) String() string {
	var builder strings.Builder
	builder.WriteString("JobProgressHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jph.ID))
	builder.WriteString("created_at=")
	builder.WriteString(jph.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", jph.Status))
	builder.WriteString(", ")
	builder.WriteString("complete=")
	builder.WriteString(fmt.Sprintf("%v", jph.Complete))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(jph.Note)
	builder.WriteByte(')')
	return builder.String()
}

// JobProgressHistories is a parsable slice of JobProgressHistory.
type JobProgressHistories []*JobProgressHistory
