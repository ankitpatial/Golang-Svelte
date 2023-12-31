// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"roofix/ent/partner"
	"roofix/ent/partnerservicecity"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PartnerServiceCity is the model entity for the PartnerServiceCity schema.
type PartnerServiceCity struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// PostalID holds the value of the "postal_id" field.
	PostalID string `json:"postal_id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// NaicsCode holds the value of the "naics_code" field.
	NaicsCode uint `json:"naics_code,omitempty"`
	// LicenseNo holds the value of the "license_no" field.
	LicenseNo *string `json:"license_no,omitempty"`
	// ProofDocID holds the value of the "proof_doc_id" field.
	ProofDocID *string `json:"proof_doc_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartnerServiceCityQuery when eager-loading is set.
	Edges        PartnerServiceCityEdges `json:"edges"`
	partner_id   *string
	selectValues sql.SelectValues
}

// PartnerServiceCityEdges holds the relations/edges for other nodes in the graph.
type PartnerServiceCityEdges struct {
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
func (e PartnerServiceCityEdges) PartnerOrErr() (*Partner, error) {
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
func (*PartnerServiceCity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case partnerservicecity.FieldActive:
			values[i] = new(sql.NullBool)
		case partnerservicecity.FieldNaicsCode:
			values[i] = new(sql.NullInt64)
		case partnerservicecity.FieldID, partnerservicecity.FieldPostalID, partnerservicecity.FieldName, partnerservicecity.FieldLicenseNo, partnerservicecity.FieldProofDocID:
			values[i] = new(sql.NullString)
		case partnerservicecity.FieldCreatedAt, partnerservicecity.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case partnerservicecity.ForeignKeys[0]: // partner_id
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartnerServiceCity fields.
func (psc *PartnerServiceCity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partnerservicecity.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				psc.ID = value.String
			}
		case partnerservicecity.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				psc.CreatedAt = value.Time
			}
		case partnerservicecity.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				psc.UpdatedAt = value.Time
			}
		case partnerservicecity.FieldPostalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_id", values[i])
			} else if value.Valid {
				psc.PostalID = value.String
			}
		case partnerservicecity.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				psc.Active = value.Bool
			}
		case partnerservicecity.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				psc.Name = value.String
			}
		case partnerservicecity.FieldNaicsCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field naics_code", values[i])
			} else if value.Valid {
				psc.NaicsCode = uint(value.Int64)
			}
		case partnerservicecity.FieldLicenseNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field license_no", values[i])
			} else if value.Valid {
				psc.LicenseNo = new(string)
				*psc.LicenseNo = value.String
			}
		case partnerservicecity.FieldProofDocID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field proof_doc_id", values[i])
			} else if value.Valid {
				psc.ProofDocID = new(string)
				*psc.ProofDocID = value.String
			}
		case partnerservicecity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field partner_id", values[i])
			} else if value.Valid {
				psc.partner_id = new(string)
				*psc.partner_id = value.String
			}
		default:
			psc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PartnerServiceCity.
// This includes values selected through modifiers, order, etc.
func (psc *PartnerServiceCity) Value(name string) (ent.Value, error) {
	return psc.selectValues.Get(name)
}

// QueryPartner queries the "partner" edge of the PartnerServiceCity entity.
func (psc *PartnerServiceCity) QueryPartner() *PartnerQuery {
	return NewPartnerServiceCityClient(psc.config).QueryPartner(psc)
}

// Update returns a builder for updating this PartnerServiceCity.
// Note that you need to call PartnerServiceCity.Unwrap() before calling this method if this PartnerServiceCity
// was returned from a transaction, and the transaction was committed or rolled back.
func (psc *PartnerServiceCity) Update() *PartnerServiceCityUpdateOne {
	return NewPartnerServiceCityClient(psc.config).UpdateOne(psc)
}

// Unwrap unwraps the PartnerServiceCity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (psc *PartnerServiceCity) Unwrap() *PartnerServiceCity {
	_tx, ok := psc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartnerServiceCity is not a transactional entity")
	}
	psc.config.driver = _tx.drv
	return psc
}

// String implements the fmt.Stringer.
func (psc *PartnerServiceCity) String() string {
	var builder strings.Builder
	builder.WriteString("PartnerServiceCity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", psc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(psc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(psc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("postal_id=")
	builder.WriteString(psc.PostalID)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", psc.Active))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(psc.Name)
	builder.WriteString(", ")
	builder.WriteString("naics_code=")
	builder.WriteString(fmt.Sprintf("%v", psc.NaicsCode))
	builder.WriteString(", ")
	if v := psc.LicenseNo; v != nil {
		builder.WriteString("license_no=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := psc.ProofDocID; v != nil {
		builder.WriteString("proof_doc_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// PartnerServiceCities is a parsable slice of PartnerServiceCity.
type PartnerServiceCities []*PartnerServiceCity
