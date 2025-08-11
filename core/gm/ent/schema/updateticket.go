package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UpdateTicket holds the schema definition for the UpdateTicket entity.
type UpdateTicket struct {
	ent.Schema
}

// Fields of the UpdateTicket.
func (UpdateTicket) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("ticket"),
	}
}

// Edges of the UpdateTicket.
func (UpdateTicket) Edges() []ent.Edge {
	return nil
}
