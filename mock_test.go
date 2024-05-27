package main

import (
	"context"
	"log"
	"testing"

	"github.com/wupyweb/realworld-kit/ent"
	"github.com/wupyweb/realworld-kit/ent/article"
	"github.com/wupyweb/realworld-kit/ent/tag"
	"github.com/wupyweb/realworld-kit/ent/user"

	//"github.com/wupyweb/realworld-kit/internal/apis/conduit"
	"github.com/wupyweb/realworld-kit/mock"
)

func TestCreate(t *testing.T) {
	client, err := ent.Open("sqlite3", "test.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	mock.Do(client)
}

func TestQueryArticles(t *testing.T) {
	client, err := ent.Open("sqlite3", "test.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// store := conduit.NewStorage(client)

	// Query articles.
	// articles, _ := store.GetArticles(&conduit.ArticleQuery{Tag: "aut"})
	// a := articles[0]
	// t.Log(a)
	// t.Log(a.Edges.Owner)
	a, _ := client.Article.Query().
		WithOwner().
		WithTags().
		Where(article.HasTagsWith(tag.Name("aut"))).
		All(context.Background())

	t.Log(a)
	t.Log(a[0].Edges.Owner)
	t.Log(a[0].Edges.Tags)
}

func TestDeleteArticle(t *testing.T) {
	client, err := ent.Open("sqlite3", "test.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// err = client.Article.DeleteOneID(2).Exec(context.Background())
	if err != nil {
		t.Error(err)
	}

	// err = client.Article.Update().
	// 	Where(article.ID(2)).
	// 	ClearComments().
	// 	ClearLikedUsers().
	// 	ClearOwner().
	// 	ClearTags().
	// 	Exec(context.Background())
	// if err != nil {
	// 	t.Error(err)
	// }

	err = client.Article.DeleteOneID(2).Exec(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteComment(t *testing.T) {
	client, err := ent.Open("sqlite3", "test.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	err = client.Comment.DeleteOneID(102).Exec(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestLikeArticle(t *testing.T) {
	client, err := ent.Open("sqlite3", "test.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	flag, err := client.Article.
		Query().Where(article.SlugEQ("assumenda"), article.HasLikedUsersWith(user.ID(1))).
		Exist(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(flag)

	err = client.Article.
		Update().
		AddLikedUserIDs(1).
		Where(
			article.SlugEQ("assumenda"),
		).
		Exec(context.Background())
	if err != nil {
		t.Error(err)
	}
}
