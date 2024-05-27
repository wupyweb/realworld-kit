package mock

import (
	"context"
	"math/rand"

	"github.com/wupyweb/realworld-kit/ent"
	"golang.org/x/crypto/bcrypt"

	"github.com/bxcodec/faker/v3"
)

func CreateUsers(client *ent.Client) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	for i := 0; i < 10; i++ {
		client.User.Create().
			SetUsername(faker.Name()).
			SetEmail(faker.Email()).
			SetPassword(string(hashPassword)).
			Save(context.Background())
	}
}

func CreateTags(client *ent.Client) {

	for i := 0; i < 5; i++ {
		client.Tag.Create().
			SetName(faker.Word()).
			Save(context.Background())
	}
}

func CreateArticles(client *ent.Client) {
	users, _ := client.User.Query().All(context.Background())
	tags, _ := client.Tag.Query().All(context.Background())

	randomTags := func(tags []*ent.Tag, n int) []*ent.Tag {
		rand.Shuffle(len(tags), func(i, j int) {
			tags[i], tags[j] = tags[j], tags[i]
		})
		return tags[:n]
	}

	randomUser := func() *ent.User {
		return users[rand.Intn(len(users))]
	}

	for i := 0; i < 50; i++ {
		title := faker.Word()

		client.Article.Create().
			SetTitle(title).
			SetDescription(faker.Sentence()).
			SetBody(faker.Paragraph()).
			SetSlug(title).
			SetOwner(randomUser()).
			AddTags(randomTags(tags, 3)...).
			Save(context.Background())
	}
}

func CreateComments(client *ent.Client) {
	articles, _ := client.Article.Query().All(context.Background())
	users, _ := client.User.Query().All(context.Background())

	randomUser := func() *ent.User {
		return users[rand.Intn(len(users))]
    }

	randomArticle := func() *ent.Article {
		return articles[rand.Intn(len(articles))]
	}

   for i := 0; i < 100; i++ {
		client.Comment.Create().
			SetBody(faker.Sentence()).
			SetOwner(randomUser()).
			SetArticle(randomArticle()).
			Save(context.Background())
    }
}

func LikeArticles(client *ent.Client) {
	articles, _ := client.Article.Query().All(context.Background())
	users, _ := client.User.Query().All(context.Background())

	randomArticles := func(n int) []*ent.Article {
		rand.Shuffle(len(articles), func(i, j int) {
			articles[i], articles[j] = articles[j], articles[i]
		})
		return articles[:n]
    }

	for _, u := range users {
		client.User.UpdateOne(u).
			AddLikedArticles(randomArticles(rand.Intn(len(articles)))...).
			Save(context.Background())
	}
}

func FollowUsers(client *ent.Client) {
	users, _ := client.User.Query().All(context.Background())
	randomUser := func() *ent.User {
		return users[rand.Intn(len(users))]
	}

	for _, u := range users {
		client.User.UpdateOne(u).
			AddFollowing([]*ent.User{randomUser(), randomUser(), randomUser()}...).
			Save(context.Background())
	}
}

func Do(client *ent.Client) {
	CreateUsers(client)
	CreateTags(client)
	CreateArticles(client)
	CreateComments(client)
	LikeArticles(client)
	FollowUsers(client)
}