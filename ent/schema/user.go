package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
	    field.String("username"),
		field.String("password"),
		field.String("email").Unique(),
		field.String("bio").Optional(),
		field.String("image").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
	    edge.To("articles", Article.Type),
		edge.To("comments", Comment.Type),
		edge.To("following", User.Type).From("followers"),
		edge.To("liked_articles", Article.Type).
			Through("favorites", Favorite.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
