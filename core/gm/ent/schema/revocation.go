package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Revocation holds the schema definition for the Revocation entity.
type Revocation struct {
	ent.Schema
}

// Fields of the Credential.
func (Revocation) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("nym"),
		field.Int("revoked_period"),
		field.Int("signed_period"),
	}
}

// Edges of the Credential.
func (Revocation) Edges() []ent.Edge {
	return nil
}
