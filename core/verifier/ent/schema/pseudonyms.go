package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Pseudonyms holds the schema definition for the Pseudonyms entity.
type Pseudonyms struct {
	ent.Schema
}

// Fields of the Pseudonyms.
func (Pseudonyms) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("nym"),
		field.Int("count"),
		field.Int("period"),
	}
}

// Edges of the Pseudonyms.
func (Pseudonyms) Edges() []ent.Edge {
	return nil
}
