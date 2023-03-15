// Code generated by ent, DO NOT EDIT.

package personal

const (
	// Label holds the string label denoting the personal type in the database.
	Label = "personal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the personal in the database.
	Table = "t_personal"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "t_personal"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "t_user"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
)

// Columns holds all SQL columns for personal fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhone,
	FieldUserID,
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