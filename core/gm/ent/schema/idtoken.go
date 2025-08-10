package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// IDToken holds the schema definition for the UpdateTicket entity.
type IDToken struct {
	ent.Schema
}

// Fields of the IDToken.
func (IDToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("idToken"),
	}
}

// Edges of the IDToken.
func (IDToken) Edges() []ent.Edge {
	return nil
}
