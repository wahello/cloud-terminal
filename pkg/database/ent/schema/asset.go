package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Annotations of the User.
func (Asset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "assets"},
	}
}

// Fields of the Assets.
func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name"),
		field.String("ip"),
		field.String("protocol"),
		field.Int("port"),
		field.String("accountType"),
		field.String("username"),
		field.String("password"),
		field.String("credentialId"),
		field.String("privateKey"),
		field.String("passphrase"),
		field.String("description"),
		field.Bool("active"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).
			UpdateDefault(time.Now),
		field.String("tags"),
	}
}

// Edges of the Assets.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sessions", Session.Type).
			Ref("assets").
			Unique(),
	}
}