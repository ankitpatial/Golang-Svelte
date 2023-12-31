// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/partnercontact"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PartnerContact is the model entity for the PartnerContact schema.
type PartnerContact struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// PartnerID holds the value of the "partner_id" field.
	PartnerID string `json:"partner_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Role holds the value of the "role" field.
	Role enum.PartnerContactRole `json:"role,omitempty"`
	// Type holds the value of the "type" field.
	Type enum.PartnerContact `json:"type,omitempty"`
	// Title holds the value of the "title" field.
	Title *string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// InvoicingEmail holds the value of the "invoicing_email" field.
	InvoicingEmail *string `json:"invoicing_email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartnerContactQuery when eager-loading is set.
	Edges        PartnerContactEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PartnerContactEdges holds the relations/edges for other nodes in the graph.
type PartnerContactEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Partner holds the value of the partner edge.
	Partner *Partner `json:"partner,omitempty"`
	// Sessions holds the value of the sessions edge.
	Sessions []*UserSession `json:"sessions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedSessions map[string][]*UserSession
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartnerContactEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// PartnerOrErr returns the Partner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartnerContactEdges) PartnerOrErr() (*Partner, error) {
	if e.loadedTypes[1] {
		if e.Partner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: partner.Label}
		}
		return e.Partner, nil
	}
	return nil, &NotLoadedError{edge: "partner"}
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e PartnerContactEdges) SessionsOrErr() ([]*UserSession, error) {
	if e.loadedTypes[2] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PartnerContact) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case partnercontact.FieldID, partnercontact.FieldPartnerID, partnercontact.FieldUserID, partnercontact.FieldRole, partnercontact.FieldType, partnercontact.FieldTitle, partnercontact.FieldDescription, partnercontact.FieldInvoicingEmail:
			values[i] = new(sql.NullString)
		case partnercontact.FieldCreatedAt, partnercontact.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartnerContact fields.
func (pc *PartnerContact) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partnercontact.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pc.ID = value.String
			}
		case partnercontact.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case partnercontact.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pc.UpdatedAt = value.Time
			}
		case partnercontact.FieldPartnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partner_id", values[i])
			} else if value.Valid {
				pc.PartnerID = value.String
			}
		case partnercontact.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pc.UserID = value.String
			}
		case partnercontact.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				pc.Role = enum.PartnerContactRole(value.String)
			}
		case partnercontact.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pc.Type = enum.PartnerContact(value.String)
			}
		case partnercontact.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pc.Title = new(string)
				*pc.Title = value.String
			}
		case partnercontact.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pc.Description = new(string)
				*pc.Description = value.String
			}
		case partnercontact.FieldInvoicingEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoicing_email", values[i])
			} else if value.Valid {
				pc.InvoicingEmail = new(string)
				*pc.InvoicingEmail = value.String
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PartnerContact.
// This includes values selected through modifiers, order, etc.
func (pc *PartnerContact) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the PartnerContact entity.
func (pc *PartnerContact) QueryUser() *UserQuery {
	return NewPartnerContactClient(pc.config).QueryUser(pc)
}

// QueryPartner queries the "partner" edge of the PartnerContact entity.
func (pc *PartnerContact) QueryPartner() *PartnerQuery {
	return NewPartnerContactClient(pc.config).QueryPartner(pc)
}

// QuerySessions queries the "sessions" edge of the PartnerContact entity.
func (pc *PartnerContact) QuerySessions() *UserSessionQuery {
	return NewPartnerContactClient(pc.config).QuerySessions(pc)
}

// Update returns a builder for updating this PartnerContact.
// Note that you need to call PartnerContact.Unwrap() before calling this method if this PartnerContact
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PartnerContact) Update() *PartnerContactUpdateOne {
	return NewPartnerContactClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PartnerContact entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PartnerContact) Unwrap() *PartnerContact {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartnerContact is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PartnerContact) String() string {
	var builder strings.Builder
	builder.WriteString("PartnerContact(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("partner_id=")
	builder.WriteString(pc.PartnerID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(pc.UserID)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", pc.Role))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pc.Type))
	builder.WriteString(", ")
	if v := pc.Title; v != nil {
		builder.WriteString("title=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pc.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pc.InvoicingEmail; v != nil {
		builder.WriteString("invoicing_email=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedSessions returns the Sessions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pc *PartnerContact) NamedSessions(name string) ([]*UserSession, error) {
	if pc.Edges.namedSessions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pc.Edges.namedSessions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pc *PartnerContact) appendNamedSessions(name string, edges ...*UserSession) {
	if pc.Edges.namedSessions == nil {
		pc.Edges.namedSessions = make(map[string][]*UserSession)
	}
	if len(edges) == 0 {
		pc.Edges.namedSessions[name] = []*UserSession{}
	} else {
		pc.Edges.namedSessions[name] = append(pc.Edges.namedSessions[name], edges...)
	}
}

// PartnerContacts is a parsable slice of PartnerContact.
type PartnerContacts []*PartnerContact
