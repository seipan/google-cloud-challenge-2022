// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/loginsessions"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

// LoginSessions is the model entity for the LoginSessions schema.
type LoginSessions struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LoginSessionsQuery when eager-loading is set.
	Edges LoginSessionsEdges `json:"edges"`
}

// LoginSessionsEdges holds the relations/edges for other nodes in the graph.
type LoginSessionsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LoginSessionsEdges) UserOrErr() (*User, error) {
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
func (*LoginSessions) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case loginsessions.FieldID:
			values[i] = new(sql.NullString)
		case loginsessions.FieldExpiry:
			values[i] = new(sql.NullTime)
		case loginsessions.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LoginSessions", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LoginSessions fields.
func (ls *LoginSessions) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loginsessions.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ls.ID = value.String
			}
		case loginsessions.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ls.UserID = *value
			}
		case loginsessions.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				ls.Expiry = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the LoginSessions entity.
func (ls *LoginSessions) QueryUser() *UserQuery {
	return (&LoginSessionsClient{config: ls.config}).QueryUser(ls)
}

// Update returns a builder for updating this LoginSessions.
// Note that you need to call LoginSessions.Unwrap() before calling this method if this LoginSessions
// was returned from a transaction, and the transaction was committed or rolled back.
func (ls *LoginSessions) Update() *LoginSessionsUpdateOne {
	return (&LoginSessionsClient{config: ls.config}).UpdateOne(ls)
}

// Unwrap unwraps the LoginSessions entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ls *LoginSessions) Unwrap() *LoginSessions {
	_tx, ok := ls.config.driver.(*txDriver)
	if !ok {
		panic("ent: LoginSessions is not a transactional entity")
	}
	ls.config.driver = _tx.drv
	return ls
}

// String implements the fmt.Stringer.
func (ls *LoginSessions) String() string {
	var builder strings.Builder
	builder.WriteString("LoginSessions(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ls.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ls.UserID))
	builder.WriteString(", ")
	builder.WriteString("expiry=")
	builder.WriteString(ls.Expiry.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LoginSessionsSlice is a parsable slice of LoginSessions.
type LoginSessionsSlice []*LoginSessions

func (ls LoginSessionsSlice) config(cfg config) {
	for _i := range ls {
		ls[_i].config = cfg
	}
}
