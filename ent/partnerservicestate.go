// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/partnerservicestate"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PartnerServiceState is the model entity for the PartnerServiceState schema.
type PartnerServiceState struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// LicenseNo holds the value of the "license_no" field.
	LicenseNo *string `json:"license_no,omitempty"`
	// LicenseExpDate holds the value of the "license_exp_date" field.
	LicenseExpDate *time.Time `json:"license_exp_date,omitempty"`
	// ProofDocID holds the value of the "proof_doc_id" field.
	ProofDocID *string `json:"proof_doc_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartnerServiceStateQuery when eager-loading is set.
	Edges        PartnerServiceStateEdges `json:"edges"`
	partner_id   *string
	selectValues sql.SelectValues
}

// PartnerServiceStateEdges holds the relations/edges for other nodes in the graph.
type PartnerServiceStateEdges struct {
	// Partner holds the value of the partner edge.
	Partner *Partner `json:"partner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// PartnerOrErr returns the Partner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartnerServiceStateEdges) PartnerOrErr() (*Partner, error) {
	if e.loadedTypes[0] {
		if e.Partner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: partner.Label}
		}
		return e.Partner, nil
	}
	return nil, &NotLoadedError{edge: "partner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PartnerServiceState) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case partnerservicestate.FieldID, partnerservicestate.FieldCountry, partnerservicestate.FieldState, partnerservicestate.FieldLicenseNo, partnerservicestate.FieldProofDocID:
			values[i] = new(sql.NullString)
		case partnerservicestate.FieldCreatedAt, partnerservicestate.FieldUpdatedAt, partnerservicestate.FieldLicenseExpDate:
			values[i] = new(sql.NullTime)
		case partnerservicestate.ForeignKeys[0]: // partner_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartnerServiceState fields.
func (pss *PartnerServiceState) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partnerservicestate.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pss.ID = value.String
			}
		case partnerservicestate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pss.CreatedAt = value.Time
			}
		case partnerservicestate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pss.UpdatedAt = value.Time
			}
		case partnerservicestate.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				pss.Country = value.String
			}
		case partnerservicestate.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				pss.State = value.String
			}
		case partnerservicestate.FieldLicenseNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field license_no", values[i])
			} else if value.Valid {
				pss.LicenseNo = new(string)
				*pss.LicenseNo = value.String
			}
		case partnerservicestate.FieldLicenseExpDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field license_exp_date", values[i])
			} else if value.Valid {
				pss.LicenseExpDate = new(time.Time)
				*pss.LicenseExpDate = value.Time
			}
		case partnerservicestate.FieldProofDocID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field proof_doc_id", values[i])
			} else if value.Valid {
				pss.ProofDocID = new(string)
				*pss.ProofDocID = value.String
			}
		case partnerservicestate.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partner_id", values[i])
			} else if value.Valid {
				pss.partner_id = new(string)
				*pss.partner_id = value.String
			}
		default:
			pss.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PartnerServiceState.
// This includes values selected through modifiers, order, etc.
func (pss *PartnerServiceState) Value(name string) (ent.Value, error) {
	return pss.selectValues.Get(name)
}

// QueryPartner queries the "partner" edge of the PartnerServiceState entity.
func (pss *PartnerServiceState) QueryPartner() *PartnerQuery {
	return NewPartnerServiceStateClient(pss.config).QueryPartner(pss)
}

// Update returns a builder for updating this PartnerServiceState.
// Note that you need to call PartnerServiceState.Unwrap() before calling this method if this PartnerServiceState
// was returned from a transaction, and the transaction was committed or rolled back.
func (pss *PartnerServiceState) Update() *PartnerServiceStateUpdateOne {
	return NewPartnerServiceStateClient(pss.config).UpdateOne(pss)
}

// Unwrap unwraps the PartnerServiceState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pss *PartnerServiceState) Unwrap() *PartnerServiceState {
	_tx, ok := pss.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartnerServiceState is not a transactional entity")
	}
	pss.config.driver = _tx.drv
	return pss
}

// String implements the fmt.Stringer.
func (pss *PartnerServiceState) String() string {
	var builder strings.Builder
	builder.WriteString("PartnerServiceState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pss.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pss.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pss.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(pss.Country)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(pss.State)
	builder.WriteString(", ")
	if v := pss.LicenseNo; v != nil {
		builder.WriteString("license_no=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pss.LicenseExpDate; v != nil {
		builder.WriteString("license_exp_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pss.ProofDocID; v != nil {
		builder.WriteString("proof_doc_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// PartnerServiceStates is a parsable slice of PartnerServiceState.
type PartnerServiceStates []*PartnerServiceState