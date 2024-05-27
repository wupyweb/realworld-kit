package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("body"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).
			UpdateDefault(time.Now),

		field.Int("user_id").Optional(),
		field.Int("article_id").Optional(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("comments").Unique().Field("user_id"),
		edge.From("article", Article.Type).Ref("comments").Unique().Field("article_id"),
	}
}
