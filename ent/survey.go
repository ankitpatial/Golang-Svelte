// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/survey"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Survey is the model entity for the Survey schema.
type Survey struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Date holds the value of the "date" field.
	Date string `json:"date,omitempty"`
	// Slot holds the value of the "slot" field.
	Slot string `json:"slot,omitempty"`
	// SlotID holds the value of the "slot_id" field.
	SlotID string `json:"slot_id,omitempty"`
	// From holds the value of the "from" field.
	From time.Time `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To time.Time `json:"to,omitempty"`
	// will be used with status REQUESTING and will denote expiry of request
	Until *time.Time `json:"until,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address *string `json:"address,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone *string `json:"phone,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes *string `json:"notes,omitempty"`
	// Status holds the value of the "status" field.
	Status enum.SurveyStatus `json:"status,omitempty"`
	// Progress holds the value of the "progress" field.
	Progress *enum.SurveyProgress `json:"progress,omitempty"`
	// ProgressAt holds the value of the "progress_at" field.
	ProgressAt *time.Time `json:"progress_at,omitempty"`
	// ProgressFlagAt holds the value of the "progress_flag_at" field.
	ProgressFlagAt *time.Time `json:"progress_flag_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SurveyQuery when eager-loading is set.
	Edges        SurveyEdges `json:"edges"`
	partner_id   *string
	user_id      *string
	selectValues sql.SelectValues
}

// SurveyEdges holds the relations/edges for other nodes in the graph.
type SurveyEdges struct {
	// ProgressHistory holds the value of the progress_history edge.
	ProgressHistory []*SurveyProgress `json:"progress_history,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy *User `json:"created_by,omitempty"`
	// Partner holds the value of the partner edge.
	Partner *Partner `json:"partner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedProgressHistory map[string][]*SurveyProgress
}

// ProgressHistoryOrErr returns the ProgressHistory value or an error if the edge
// was not loaded in eager-loading.
func (e SurveyEdges) ProgressHistoryOrErr() ([]*SurveyProgress, error) {
	if e.loadedTypes[0] {
		return e.ProgressHistory, nil
	}
	return nil, &NotLoadedError{edge: "progress_history"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SurveyEdges) CreatedByOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.CreatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.CreatedBy, nil
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// PartnerOrErr returns the Partner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SurveyEdges) PartnerOrErr() (*Partner, error) {
	if e.loadedTypes[2] {
		if e.Partner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: partner.Label}
		}
		return e.Partner, nil
	}
	return nil, &NotLoadedError{edge: "partner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Survey) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case survey.FieldID, survey.FieldDate, survey.FieldSlot, survey.FieldSlotID, survey.FieldName, survey.FieldAddress, survey.FieldPhone, survey.FieldNotes, survey.FieldStatus, survey.FieldProgress:
			values[i] = new(sql.NullString)
		case survey.FieldCreatedAt, survey.FieldUpdatedAt, survey.FieldFrom, survey.FieldTo, survey.FieldUntil, survey.FieldProgressAt, survey.FieldProgressFlagAt:
			values[i] = new(sql.NullTime)
		case survey.ForeignKeys[0]: // partner_id
			values[i] = new(sql.NullString)
		case survey.ForeignKeys[1]: // user_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Survey fields.
func (s *Survey) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case survey.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case survey.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case survey.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case survey.FieldDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				s.Date = value.String
			}
		case survey.FieldSlot:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slot", values[i])
			} else if value.Valid {
				s.Slot = value.String
			}
		case survey.FieldSlotID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slot_id", values[i])
			} else if value.Valid {
				s.SlotID = value.String
			}
		case survey.FieldFrom:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				s.From = value.Time
			}
		case survey.FieldTo:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				s.To = value.Time
			}
		case survey.FieldUntil:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field until", values[i])
			} else if value.Valid {
				s.Until = new(time.Time)
				*s.Until = value.Time
			}
		case survey.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = new(string)
				*s.Name = value.String
			}
		case survey.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				s.Address = new(string)
				*s.Address = value.String
			}
		case survey.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				s.Phone = new(string)
				*s.Phone = value.String
			}
		case survey.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				s.Notes = new(string)
				*s.Notes = value.String
			}
		case survey.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = enum.SurveyStatus(value.String)
			}
		case survey.FieldProgress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field progress", values[i])
			} else if value.Valid {
				s.Progress = new(enum.SurveyProgress)
				*s.Progress = enum.SurveyProgress(value.String)
			}
		case survey.FieldProgressAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field progress_at", values[i])
			} else if value.Valid {
				s.ProgressAt = new(time.Time)
				*s.ProgressAt = value.Time
			}
		case survey.FieldProgressFlagAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field progress_flag_at", values[i])
			} else if value.Valid {
				s.ProgressFlagAt = new(time.Time)
				*s.ProgressFlagAt = value.Time
			}
		case survey.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partner_id", values[i])
			} else if value.Valid {
				s.partner_id = new(string)
				*s.partner_id = value.String
			}
		case survey.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.user_id = new(string)
				*s.user_id = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Survey.
// This includes values selected through modifiers, order, etc.
func (s *Survey) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryProgressHistory queries the "progress_history" edge of the Survey entity.
func (s *Survey) QueryProgressHistory() *SurveyProgressQuery {
	return NewSurveyClient(s.config).QueryProgressHistory(s)
}

// QueryCreatedBy queries the "created_by" edge of the Survey entity.
func (s *Survey) QueryCreatedBy() *UserQuery {
	return NewSurveyClient(s.config).QueryCreatedBy(s)
}

// QueryPartner queries the "partner" edge of the Survey entity.
func (s *Survey) QueryPartner() *PartnerQuery {
	return NewSurveyClient(s.config).QueryPartner(s)
}

// Update returns a builder for updating this Survey.
// Note that you need to call Survey.Unwrap() before calling this method if this Survey
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Survey) Update() *SurveyUpdateOne {
	return NewSurveyClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Survey entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Survey) Unwrap() *Survey {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Survey is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Survey) String() string {
	var builder strings.Builder
	builder.WriteString("Survey(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(s.Date)
	builder.WriteString(", ")
	builder.WriteString("slot=")
	builder.WriteString(s.Slot)
	builder.WriteString(", ")
	builder.WriteString("slot_id=")
	builder.WriteString(s.SlotID)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(s.From.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(s.To.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := s.Until; v != nil {
		builder.WriteString("until=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.Address; v != nil {
		builder.WriteString("address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := s.Notes; v != nil {
		builder.WriteString("notes=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	if v := s.Progress; v != nil {
		builder.WriteString("progress=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := s.ProgressAt; v != nil {
		builder.WriteString("progress_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.ProgressFlagAt; v != nil {
		builder.WriteString("progress_flag_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedProgressHistory returns the ProgressHistory named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Survey) NamedProgressHistory(name string) ([]*SurveyProgress, error) {
	if s.Edges.namedProgressHistory == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedProgressHistory[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Survey) appendNamedProgressHistory(name string, edges ...*SurveyProgress) {
	if s.Edges.namedProgressHistory == nil {
		s.Edges.namedProgressHistory = make(map[string][]*SurveyProgress)
	}
	if len(edges) == 0 {
		s.Edges.namedProgressHistory[name] = []*SurveyProgress{}
	} else {
		s.Edges.namedProgressHistory[name] = append(s.Edges.namedProgressHistory[name], edges...)
	}
}

// Surveys is a parsable slice of Survey.
type Surveys []*Survey
