package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Favorite holds the schema definition for the Favorite entity.
type Favorite struct {
	ent.Schema
}

// Fields of the Favorite.
func (Favorite) Fields() []ent.Field {
	return []ent.Field{
		field.Time("favorited_at").Default(time.Now),

		field.Int("user_id"),
		field.Int("article_id"),
	}
}

// Edges of the Favorite.
func (Favorite) Edges() []ent.Edge {
	return []ent.Edge{
	    edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	    edge.To("article", Article.Type).
			Unique().
			Required().
			Field("article_id"),
	}
}
