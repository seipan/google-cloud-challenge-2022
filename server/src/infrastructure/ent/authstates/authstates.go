// Code generated by ent, DO NOT EDIT.

package authstates

const (
	// Label holds the string label denoting the authstates type in the database.
	Label = "auth_states"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldRedirectURL holds the string denoting the redirect_url field in the database.
	FieldRedirectURL = "redirect_url"
	// Table holds the table name of the authstates in the database.
	Table = "auth_states"
)

// Columns holds all SQL columns for authstates fields.
var Columns = []string{
	FieldID,
	FieldState,
	FieldRedirectURL,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
	// RedirectURLValidator is a validator for the "redirect_url" field. It is called by the builders before save.
	RedirectURLValidator func(string) error
)
