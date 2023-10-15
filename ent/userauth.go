// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/user"
	"roofix/ent/userauth"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserAuth is the model entity for the UserAuth schema.
type UserAuth struct {
	config `json:"-"`
	// ID of the ent.
	// user id from provider side
	ID string `json:"id,omitempty"`
	// ProviderID holds the value of the "provider_id" field.
	ProviderID uint8 `json:"provider_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserAuthQuery when eager-loading is set.
	Edges        UserAuthEdges `json:"edges"`
	user_id      *string
	selectValues sql.SelectValues
}

// UserAuthEdges holds the relations/edges for other nodes in the graph.
type UserAuthEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserAuthEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserAuth) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userauth.FieldProviderID:
			values[i] = new(sql.NullInt64)
		case userauth.FieldID:
			values[i] = new(sql.NullString)
		case userauth.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case userauth.ForeignKeys[0]: // user_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserAuth fields.
func (ua *UserAuth) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userauth.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ua.ID = value.String
			}
		case userauth.FieldProviderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field provider_id", values[i])
			} else if value.Valid {
				ua.ProviderID = uint8(value.Int64)
			}
		case userauth.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ua.CreatedAt = value.Time
			}
		case userauth.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ua.user_id = new(string)
				*ua.user_id = value.String
			}
		default:
			ua.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserAuth.
// This includes values selected through modifiers, order, etc.
func (ua *UserAuth) Value(name string) (ent.Value, error) {
	return ua.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserAuth entity.
func (ua *UserAuth) QueryUser() *UserQuery {
	return NewUserAuthClient(ua.config).QueryUser(ua)
}

// Update returns a builder for updating this UserAuth.
// Note that you need to call UserAuth.Unwrap() before calling this method if this UserAuth
// was returned from a transaction, and the transaction was committed or rolled back.
func (ua *UserAuth) Update() *UserAuthUpdateOne {
	return NewUserAuthClient(ua.config).UpdateOne(ua)
}

// Unwrap unwraps the UserAuth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ua *UserAuth) Unwrap() *UserAuth {
	_tx, ok := ua.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserAuth is not a transactional entity")
	}
	ua.config.driver = _tx.drv
	return ua
}

// String implements the fmt.Stringer.
func (ua *UserAuth) String() string {
	var builder strings.Builder
	builder.WriteString("UserAuth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ua.ID))
	builder.WriteString("provider_id=")
	builder.WriteString(fmt.Sprintf("%v", ua.ProviderID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ua.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserAuths is a parsable slice of UserAuth.
type UserAuths []*UserAuth