package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Instrument holds the schema definition for the Instrument entity.
type Instrument struct {
	ent.Schema
}

// Fields of the Instrument.
func (Instrument) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Instrument.
func (Instrument) Edges() []ent.Edge {
	return nil
}
