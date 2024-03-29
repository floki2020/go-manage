// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-manage/pkg/ent/personal"
	"go-manage/pkg/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Personal is the model entity for the Personal schema.
type Personal struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 个人姓名
	Name string `json:"name,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty"`
	// 关联ID
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonalQuery when eager-loading is set.
	Edges PersonalEdges `json:"edges"`
}

// PersonalEdges holds the relations/edges for other nodes in the graph.
type PersonalEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonalEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Personal) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case personal.FieldID, personal.FieldUserID:
			values[i] = new(sql.NullInt64)
		case personal.FieldName, personal.FieldPhone:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Personal", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Personal fields.
func (pe *Personal) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case personal.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case personal.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case personal.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				pe.Phone = value.String
			}
		case personal.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pe.UserID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Personal entity.
func (pe *Personal) QueryUsers() *UserQuery {
	return NewPersonalClient(pe.config).QueryUsers(pe)
}

// Update returns a builder for updating this Personal.
// Note that you need to call Personal.Unwrap() before calling this method if this Personal
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Personal) Update() *PersonalUpdateOne {
	return NewPersonalClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Personal entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Personal) Unwrap() *Personal {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Personal is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Personal) String() string {
	var builder strings.Builder
	builder.WriteString("Personal(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(pe.Phone)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Personals is a parsable slice of Personal.
type Personals []*Personal
