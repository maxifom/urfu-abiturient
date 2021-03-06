// Code generated by entc, DO NOT EDIT.

package lastupdated

const (
	// Label holds the string label denoting the lastupdated type in the database.
	Label = "last_updated"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// Table holds the table name of the lastupdated in the database.
	Table = "last_updateds"
)

// Columns holds all SQL columns for lastupdated fields.
var Columns = []string{
	FieldID,
	FieldLastUpdated,
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
