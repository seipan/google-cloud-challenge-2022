// Code generated by ent, DO NOT EDIT.

package event

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// EdgeState holds the string denoting the state edge name in mutations.
	EdgeState = "state"
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the event in the database.
	Table = "events"
	// StateTable is the table that holds the state relation/edge.
	StateTable = "estates"
	// StateInverseTable is the table name for the EState entity.
	// It exists in this package in order to avoid circular dependency with the "estate" package.
	StateInverseTable = "estates"
	// StateColumn is the table column denoting the state relation/edge.
	StateColumn = "event_state"
	// TypeTable is the table that holds the type relation/edge.
	TypeTable = "etypes"
	// TypeInverseTable is the table name for the EType entity.
	// It exists in this package in order to avoid circular dependency with the "etype" package.
	TypeInverseTable = "etypes"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "event_type"
	// AdminTable is the table that holds the admin relation/edge.
	AdminTable = "events"
	// AdminInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminInverseTable = "users"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "event_admin"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_events"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDetail,
	FieldLocation,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_admin",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "event_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DetailValidator is a validator for the "detail" field. It is called by the builders before save.
	DetailValidator func(string) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
