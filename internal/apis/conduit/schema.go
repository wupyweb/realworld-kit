package conduit

import (
	"time"

	"github.com/wupyweb/realworld-kit/ent"
)

type (
	UserBase struct {
		Email    string `json:"email,omitempty"`
		Username string `json:"username,omitempty"`
		Bio      string `json:"bio,omitempty"`
		Image    string `json:"image,omitempty"`
	}

	LoginRequest struct {
		User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	LoginResponse struct {
		User struct {
			UserBase
			Token string `json:"token"`
		} `json:"user"`
	}

	RegisterRequest struct {
		User struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	UserUpdate struct {
		User struct {
			UserBase
			Password string `json:"password,omitempty"`
		} `json:"user"`
	}

	Profile struct {
		UserBase
		Following bool `json:"following"`
	}

	ProfileResponse struct {
		Profile `json:"profile"`
	}

	ArticleQuery struct {
		Tag       string `form:"tag,omitempty"`
		Author    string `form:"author,omitempty"`
		Favorited string `form:"favorited,omitempty"`
		Limit     int    `form:"limit,omitempty"`
		Offset    int    `form:"offset,omitempty"`
	}

	ArticleBase struct {
		Slug           string   `json:"slug"`
		Title          string   `json:"title"`
		Description    string   `json:"description"`
		Body           string   `json:"body"`
		TagList        []string `json:"tagList,omitempty"`
		CreatedAt      string   `json:"createdAt,omitempty"`
		UpdatedAt      string   `json:"updatedAt,omitempty"`
		Favorited      bool     `json:"favorited,omitempty"`
		FavoritesCount int      `json:"favoritesCount,omitempty"`
		Author         Profile  `json:"author"`
	}

	SingleArticle struct {
		Article ArticleBase `json:"article"`
	}

	MultipleArticles struct {
		Articles      []ArticleBase `json:"articles"`
		ArticlesCount int           `json:"articlesCount"`
	}

	ArticleRequest struct {
		Article struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			Body        string   `json:"body"`
			TagList     []string `json:"tagList,omitempty"`
		}
	}

	CommentRequest struct {
		Comment struct {
			Body string `json:"body"`
		}
	}

	CommentBase struct {
		Id        int     `json:"id"`
		Body      string  `json:"body"`
		CreatedAt string  `json:"createdAt"`
		UpdatedAt string  `json:"updatedAt"`
		Author    Profile `json:"author"`
	}

	SingleComment struct {
		Comment CommentBase `json:"comment"`
	}

	MultipleComments struct {
		Comments []CommentBase `json:"comments"`
	}

	TagsResponse struct {
		Tags []string `json:"tags"`
	}
)

func toProfile(u *ent.User) Profile {
	return Profile{
		UserBase: UserBase{
			Username: u.Username,
			Bio:      u.Bio,
			Image:    u.Image,
		},
		Following: false,
	}
}

func toTagList(tags []*ent.Tag) []string {
	var tagList []string
	for _, tag := range tags {
		tagList = append(tagList, tag.Name)
	}
	return tagList
}

func toArticle(a *ent.Article) ArticleBase {

	return ArticleBase{
		Slug:           a.Slug,
		Title:          a.Title,
		Description:    a.Description,
		Body:           a.Body,
		TagList:        toTagList(a.Edges.Tags),
		CreatedAt:      a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      a.UpdatedAt.Format(time.RFC3339),
		Favorited:      false,
		FavoritesCount: 0,
		Author:         toProfile(a.Edges.Owner),
	}
}

func toComment(m *ent.Comment) CommentBase {

	return CommentBase{
		Id:          m.ID,
		Body:        m.Body,
		CreatedAt:   m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   m.UpdatedAt.Format(time.RFC3339),
		Author:      toProfile(m.Edges.Owner),
	}
}
