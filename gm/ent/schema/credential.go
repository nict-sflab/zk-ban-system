package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Credential holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

// Fields of the Credential.
func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("credential"),
		field.String("identifier").Optional(),
	}
}

// Edges of the Credential.
func (Credential) Edges() []ent.Edge {
	return nil
}
