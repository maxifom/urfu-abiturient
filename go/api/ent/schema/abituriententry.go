package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AbiturientEntry holds the schema definition for the AbiturientEntry entity.
type AbiturientEntry struct {
	ent.Schema
}

// Fields of the AbiturientEntry.
func (AbiturientEntry) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(1000),
		field.Int64("number"),
		field.String("status"),
		field.String("type"),
		field.Bool("statement_given").Default(false).
			StructTag(`json:"statement_given"`),
		field.Bool("original_given").Default(false).
			StructTag(`json:"original_given"`),
		field.String("speciality"),
		field.String("program"),
		field.Enum("form").Values("och", "zaoch", "och-zaoch"),
		field.String("basis"),
		field.Int64("sum"),
	}
}
