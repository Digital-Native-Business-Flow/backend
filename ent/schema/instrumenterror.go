package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// InstrumentError holds the schema definition for the InstrumentError entity.
type InstrumentError struct {
	ent.Schema
}

// Fields of the InstrumentError.
func (InstrumentError) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("instrument_id", uuid.UUID{}),
		field.Uint("code"),
		field.String("text").NotEmpty(),
		field.Time("occurred_at"),
	}
}

// Edges of the InstrumentError.
func (InstrumentError) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Instrument.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			StorageKey(
				edge.Symbol("fk__instrument_errors_instrument_id__instruments_id"),
			).
			Field("instrument_id").
			Unique().
			Required(),
	}
}
