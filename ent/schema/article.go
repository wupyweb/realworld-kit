package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(100),
		field.String("description").MaxLen(100),
		field.Text("body"),
		field.String("slug").MaxLen(100).Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),

		field.Int("user_id").Optional(),	// 自定义外键
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
	    edge.To("comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("tags", Tag.Type).
			// 设置级联删除，文章删除，会自动删除关联关系
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("owner", User.Type).
			Ref("articles").
			Unique().
			Field("user_id").
			Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.From("liked_users", User.Type).
			Ref("liked_articles").
			Through("favorites", Favorite.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
