package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AbiturientEntry holds the schema definition for the AbiturientEntry entity.
type LastUpdated struct {
	ent.Schema
}

// Fields of the AbiturientEntry.
func (LastUpdated) Fields() []ent.Field {
	return []ent.Field{
		field.Time("last_updated"),
	}
}
